package UseCases

import (
	"context"
	"errors"
	Order2 "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Order"
	Stock2 "github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Stock"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository/Order"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository/Stock"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository/Transaction"
	"github.com/google/uuid"
	"strconv"
	"strings"
)

type OrderUseCase struct {
	Repo              Order.IOrderRepo
	RepoStock         Stock.IStockRepo
	RepoTransactionID Transaction.ITransactionIDRepo
}

func (o OrderUseCase) GetOrderById(ctx context.Context, orderID string) (Order2.Order, error) {
	ctx, span := tracerOrder.Start(ctx, "GetOrderById_UseCase")
	defer span.End()

	return o.Repo.GetOrderByID(ctx, orderID)
}

func (o OrderUseCase) CreateOrder(ctx context.Context, transactionID string) (Order2.Order, error) {
	ctx, span := tracerOrder.Start(ctx, "CreateOrder_UseCase")
	defer span.End()

	tempOrder, err := o.RepoTransactionID.GetOrderByTransactionID(ctx, transactionID)
	if err != nil {
		return Order2.Order{}, err
	}

	productList := strings.Split(tempOrder.TProductList, ", ")
	for _, product := range productList {
		parts := strings.Split(product, ":")
		pID := strings.TrimSpace(parts[0])
		pQuantity, _ := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)

		if _, err := o.RepoStock.GetStockByID(ctx, pID); err != nil {
			return Order2.Order{}, err
		}

		stockQuantity, _ := o.RepoStock.GetStockByID(ctx, pID)
		newQuantity := stockQuantity.SQuantity - pQuantity

		tempStock := Stock2.Stock{
			SID:       pID,
			SQuantity: newQuantity,
			//SUpdated:  time.Now(),
		}

		_, err := o.RepoStock.UpdateStock(ctx, tempStock, pID)
		if err != nil {
			return Order2.Order{}, err
		}
	}

	var createOrder Order2.Order

	createOrder.OID = uuid.New()
	createOrder.OTranID = tempOrder.TID
	createOrder.OPaid = false
	createOrder.ODestination = tempOrder.TDestination
	createOrder.OPrice = tempOrder.TPrice
	createOrder.OStatus = "New"
	//createOrder.OCreated = time.Now()

	return o.Repo.CreateOrder(ctx, createOrder)
}

func (o OrderUseCase) ChangeOrderStatus(ctx context.Context, orderID string, orderStatus string) (Order2.Order, error) {
	ctx, span := tracerOrder.Start(ctx, "ChangeOrderStatus_UseCase")
	defer span.End()

	tempOrder, err := o.GetOrderById(ctx, orderID)
	if err != nil {
		return tempOrder, errors.New("order ID Not Found")
	}

	tempOrder, err = tempOrder.ChangeStatus(tempOrder, orderStatus)
	if err != nil {
		return tempOrder, err
	}

	return o.Repo.ChangeOrderStatus(ctx, tempOrder, orderID)
}

func NewOrderUseCase(Repo Order.IOrderRepo, RepoStock Stock.IStockRepo, RepoTransactionID Transaction.ITransactionIDRepo) IOrderUseCase {
	return &OrderUseCase{Repo: Repo, RepoStock: RepoStock, RepoTransactionID: RepoTransactionID}
}
