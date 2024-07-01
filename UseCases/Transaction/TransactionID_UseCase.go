package Transaction

import (
	"errors"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository/Address"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository/Product"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Repository/Transaction"
	"net/url"
	"strconv"
	"strings"
)

type TransactionIDUseCase struct {
	Repo        Transaction.ITransactionIDRepo
	RepoProduct Product.IProductRepo
	RepoAddress Address.IAddressRepo
}

func (o TransactionIDUseCase) GetAllTransactionIDs() ([]Entities.TransactionID, error) {
	return o.Repo.GetAllTransactionIDs()
}

func (o TransactionIDUseCase) GetOrderByTransactionID(transactionID string) (Entities.TransactionID, error) {
	return o.Repo.GetOrderByTransactionID(transactionID)
}

func (o TransactionIDUseCase) CreateTransactionID(transactionInfo Entities.TransactionID) (Entities.TransactionID, error) {
	var totalPrice float64
	var tempProductList []string

	productList := strings.Split(transactionInfo.TProductList, ", ")
	for _, product := range productList {
		parts := strings.Split(product, ":")

		if len(parts) == 2 {
			PID := strings.TrimSpace(parts[0])
			PQuantity, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil {
				return transactionInfo, errors.New("invalid Quantity")
			}

			if PQuantity <= 0 {
				return transactionInfo, errors.New("quantity Must Greater than 0")
			}

			for _, id := range tempProductList {
				if id == PID {
					return transactionInfo, errors.New("product ID Must Not Duplicated")
				}
			}

			tempProductList = append(tempProductList, PID)

			if temp, err := o.RepoProduct.GetProductByID(PID); err != nil {
				return transactionInfo, errors.New("product Id Not Found")
			} else {
				totalPrice += temp.PPrice * float64(PQuantity)
			}
		} else {
			return transactionInfo, errors.New("invalid Product Format, Should be Like This 'ProductID:Quantity'")
		}
	}

	tempAddress, err := url.QueryUnescape(transactionInfo.TDestination)
	if err != nil {
		return transactionInfo, errors.New("invalid Destination")
	}

	address, err := o.RepoAddress.GetAddressByCity(tempAddress)
	if err != nil {
		return transactionInfo, errors.New("address City Not Found")
	}

	totalPrice += address.APrice

	transactionInfo.TPrice = totalPrice
	transactionInfo.TID = transactionInfo.GenerateTransactionID(totalPrice)

	return o.Repo.CreateTransactionID(transactionInfo)
}

func (o TransactionIDUseCase) DeleteTransactionID(transactionID string) error {
	_, err := o.Repo.GetOrderByTransactionID(transactionID)
	if err != nil {
		return err
	}

	return o.Repo.DeleteTransactionID(transactionID)
}

func NewTransactionIDUseCase(repo Transaction.ITransactionIDRepo, repoProduct Product.IProductRepo, repoAddress Address.IAddressRepo) ITransactionIDUseCase {
	return TransactionIDUseCase{Repo: repo, RepoProduct: repoProduct, RepoAddress: repoAddress}
}
