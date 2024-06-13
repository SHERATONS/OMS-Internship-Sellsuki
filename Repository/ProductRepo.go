package Repository

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"gorm.io/gorm"
)

type ProductRepo struct {
	Db *gorm.DB
}

func (p *ProductRepo) GetProductByID(productId string) (Entities.Product, error) {
	var product Entities.Product
	err := p.Db.Where("p_id = ?", productId).First(&product).Error
	return product, err
}

func (p *ProductRepo) UpdateProduct(product Entities.Product, productId string) (Entities.Product, error) {
	err := p.Db.Where("p_id = ?", productId).Save(&product).Error
	return product, err
}

func (p *ProductRepo) DeleteProduct(productId string) error {
	err := p.Db.Where("p_id = ?", productId).Delete(&Entities.Product{}).Error
	return err
}

func (p *ProductRepo) CreateProduct(product Entities.Product) (Entities.Product, error) {
	err := p.Db.Create(&product).Error
	return product, err
}

func (p *ProductRepo) GetAllProducts() ([]Entities.Product, error) {
	var products []Entities.Product
	err := p.Db.Order("p_id").Find(&products).Error
	return products, err
}

func NewProductRepo(db *gorm.DB) IProductRepo {
	return &ProductRepo{Db: db}
}
