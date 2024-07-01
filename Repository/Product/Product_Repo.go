package Product

import (
	"context"
	"errors"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Model"
	"gorm.io/gorm"
	"log"
)

type ProductRepo struct {
	Db *gorm.DB
}

func (p *ProductRepo) GetProductByID(ctx context.Context, productID string) (Entities.Product, error) {
	//_, span := Ptracer.Start(ctx, "GetProductByID")
	//defer span.End()

	var product Entities.Product

	err := p.Db.Where("p_id = ?", productID).First(&product).Error
	if err != nil {
		//span.RecordError(err)
		return product, errors.New("product not found")
	}

	return product, nil
}

func (p *ProductRepo) UpdateProduct(ctx context.Context, product Entities.Product, productId string) (Entities.Product, error) {
	//_, span := Ptracer.Start(ctx, "UpdateProduct")
	//defer span.End()

	err := p.Db.Where("p_id = ?", productId).Save(&product).Error
	if err != nil {
		//span.RecordError(err)
		return product, errors.New("failed to Update Product")
	}

	return product, nil
}

func (p *ProductRepo) DeleteProduct(ctx context.Context, productID string) error {
	//_, span := Ptracer.Start(ctx, "DeleteProduct")
	//defer span.End()

	err := p.Db.Where("p_id = ?", productID).Delete(&Entities.Product{}).Error
	if err != nil {
		//span.RecordError(err)
		return errors.New("failed to Delete Product")
	}

	return nil
}

func (p *ProductRepo) CreateProduct(ctx context.Context, product Entities.Product) (Entities.Product, error) {
	//_, span := Ptracer.Start(ctx, "CreateProduct")
	//defer span.End()

	err := p.Db.Create(&product).Error
	if err != nil {
		//span.RecordError(err)
		return product, errors.New("failed to Create Product")
	}

	return product, nil
}

func (p *ProductRepo) GetAllProducts(ctx context.Context) ([]Entities.Product, error) {
	//_, span := Ptracer.Start(ctx, "GetAllProducts")
	//defer span.End()

	var products []Entities.Product

	err := p.Db.Order("CAST(p_id AS INTEGER)").Find(&products).Error
	if err != nil {
		//span.RecordError(err)
		return products, err
	}

	return products, nil
}

func NewProductRepo(db *gorm.DB) IProductRepo {
	err := db.AutoMigrate(&Model.Product{})
	if err != nil {
		log.Fatalf("Failed to auto migrate Product: %v", err)
	}
	return &ProductRepo{Db: db}
}
