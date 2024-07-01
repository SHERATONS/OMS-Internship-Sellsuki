package Order

import (
	"errors"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
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

func (o OrderUseCase) GetOrderById(orderID string) (Entities.Order, error) {
	return o.Repo.GetOrderByID(orderID)
}

func (o OrderUseCase) CreateOrder(TransactionID string) (Entities.Order, error) {
	TempOrder, err := o.RepoTransactionID.GetOrderByTransactionID(TransactionID)
	if err != nil {
		return Entities.Order{}, err
	}

	productList := strings.Split(TempOrder.TProductList, ", ")
	for _, product := range productList {
		parts := strings.Split(product, ":")
		PID := strings.TrimSpace(parts[0])
		PQuantity, _ := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)

		if _, err := o.RepoStock.GetStockByID(PID); err != nil {
			return Entities.Order{}, err
		}

		StockQuantity, _ := o.RepoStock.GetStockByID(PID)
		NewQuantity := StockQuantity.SQuantity - PQuantity

		TempStock := Entities.Stock{
			SID:       PID,
			SQuantity: NewQuantity,
			SUpdated:  time.Now(),
		}

		_, err := o.RepoStock.UpdateStock(TempStock, PID)
		if err != nil {
			return Entities.Order{}, err
		}
	}

	var createOrder Entities.Order

	createOrder.OID = uuid.New()
	createOrder.OTranID = TempOrder.TID
	createOrder.OPaid = false
	createOrder.ODestination = TempOrder.TDestination
	createOrder.OPrice = TempOrder.TPrice
	createOrder.OStatus = "New"
	createOrder.OCreated = time.Now()

	return o.Repo.CreateOrder(createOrder)
}

func (o OrderUseCase) ChangeOrderStatus(orderID string, orderStatus string) (Entities.Order, error) {
	TempOrder, err := o.GetOrderById(orderID)
	if err != nil {
		return TempOrder, errors.New("order ID Not Found")
	}

	switch orderStatus {
	case "Paid":
		if TempOrder.OStatus == "New" {
			TempOrder.OStatus = "Paid"
			TempOrder.OPaid = true
		} else {
			return TempOrder, errors.New("invalid Order Status")
		}

	case "Processing":
		if TempOrder.OStatus == "Paid" {
			if TempOrder.ODestination != "Branch" {
				TempOrder.OStatus = "Processing"
				TempOrder.OPaid = true
			} else {
				return TempOrder, errors.New("please Come Pick Up your Product at the Branch")
			}
		} else {
			return TempOrder, errors.New("invalid Order Status")
		}

	case "Done":
		if TempOrder.OStatus == "Processing" || (TempOrder.OStatus == "Paid" && TempOrder.ODestination == "Branch") {
			TempOrder.OStatus = "Done"
			TempOrder.OPaid = true
		} else {
			return TempOrder, errors.New("invalid Order Status")
		}

	default:
		return TempOrder, errors.New("invalid Order Status")
	}

	return o.Repo.ChangeOrderStatus(TempOrder, orderID)
}

func NewOrderUseCase(Repo Order.IOrderRepo, RepoStock Stock.IStockRepo, RepoTransactionID Transaction.ITransactionIDRepo) IOrderUseCase {
	return &OrderUseCase{Repo: Repo, RepoStock: RepoStock, RepoTransactionID: RepoTransactionID}
}
