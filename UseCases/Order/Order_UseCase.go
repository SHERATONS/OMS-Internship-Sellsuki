package Order

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
	"time"
)

type OrderUseCase struct {
	Repo              Order.IOrderRepo
	RepoStock         Stock.IStockRepo
	RepoTransactionID Transaction.ITransactionIDRepo
}

func (o OrderUseCase) GetOrderById(ctx context.Context, orderID string) (Order2.Order, error) {
	ctx, span := tracer.Start(ctx, "GetOrderById_UseCase")
	defer span.End()

	return o.Repo.GetOrderByID(ctx, orderID)
}

func (o OrderUseCase) CreateOrder(ctx context.Context, TransactionID string) (Order2.Order, error) {
	ctx, span := tracer.Start(ctx, "CreateOrder_UseCase")
	defer span.End()

	TempOrder, err := o.RepoTransactionID.GetOrderByTransactionID(ctx, TransactionID)
	if err != nil {
		return Order2.Order{}, err
	}

	productList := strings.Split(TempOrder.TProductList, ", ")
	for _, product := range productList {
		parts := strings.Split(product, ":")
		PID := strings.TrimSpace(parts[0])
		PQuantity, _ := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)

		if _, err := o.RepoStock.GetStockByID(ctx, PID); err != nil {
			return Order2.Order{}, err
		}

		StockQuantity, _ := o.RepoStock.GetStockByID(ctx, PID)
		NewQuantity := StockQuantity.SQuantity - PQuantity

		TempStock := Stock2.Stock{
			SID:       PID,
			SQuantity: NewQuantity,
			SUpdated:  time.Now(),
		}

		_, err := o.RepoStock.UpdateStock(ctx, TempStock, PID)
		if err != nil {
			return Order2.Order{}, err
		}
	}

	var createOrder Order2.Order

	createOrder.OID = uuid.New()
	createOrder.OTranID = TempOrder.TID
	createOrder.OPaid = false
	createOrder.ODestination = TempOrder.TDestination
	createOrder.OPrice = TempOrder.TPrice
	createOrder.OStatus = "New"
	createOrder.OCreated = time.Now()

	return o.Repo.CreateOrder(ctx, createOrder)
}

func (o OrderUseCase) ChangeOrderStatus(ctx context.Context, orderID string, orderStatus string) (Order2.Order, error) {
	ctx, span := tracer.Start(ctx, "ChangeOrderStatus_UseCase")
	defer span.End()

	TempOrder, err := o.GetOrderById(ctx, orderID)
	if err != nil {
		return TempOrder, errors.New("order ID Not Found")
	}

	TempOrder, err = TempOrder.ChangeStatus(TempOrder, orderStatus)
	if err != nil {
		return TempOrder, err
	}

	return o.Repo.ChangeOrderStatus(ctx, TempOrder, orderID)
}

func NewOrderUseCase(Repo Order.IOrderRepo, RepoStock Stock.IStockRepo, RepoTransactionID Transaction.ITransactionIDRepo) IOrderUseCase {
	return &OrderUseCase{Repo: Repo, RepoStock: RepoStock, RepoTransactionID: RepoTransactionID}
}
