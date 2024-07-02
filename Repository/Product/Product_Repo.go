package Product

import (
	"context"
	"errors"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Entities/Product"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Model"
	"gorm.io/gorm"
	"log"
)

type ProductRepo struct {
	Db *gorm.DB
}

func (p *ProductRepo) GetProductByID(ctx context.Context, productID string) (Product.Product, error) {
	ctx, span := tracer.Start(ctx, "GetProductByID_Repo")
	defer span.End()

	var product Product.Product

	err := p.Db.Where("p_id = ?", productID).First(&product).Error
	if err != nil {
		//span.RecordError(err)
		return product, errors.New("product not found")
	}

	return product, nil
}

func (p *ProductRepo) UpdateProduct(ctx context.Context, product Product.Product, productId string) (Product.Product, error) {
	ctx, span := tracer.Start(ctx, "UpdateProduct_Repo")
	defer span.End()

	err := p.Db.Where("p_id = ?", productId).Save(&product).Error
	if err != nil {
		//span.RecordError(err)
		return product, errors.New("failed to Update Product")
	}

	return product, nil
}

func (p *ProductRepo) DeleteProduct(ctx context.Context, productID string) error {
	ctx, span := tracer.Start(ctx, "DeleteProduct_Repo")
	defer span.End()

	err := p.Db.Where("p_id = ?", productID).Delete(&Product.Product{}).Error
	if err != nil {
		//span.RecordError(err)
		return errors.New("failed to Delete Product")
	}

	return nil
}

func (p *ProductRepo) CreateProduct(ctx context.Context, product Product.Product) (Product.Product, error) {
	ctx, span := tracer.Start(ctx, "CreateProduct_Repo")
	defer span.End()

	err := p.Db.Create(&product).Error
	if err != nil {
		//span.RecordError(err)
		return product, errors.New("failed to Create Product")
	}

	return product, nil
}

func (p *ProductRepo) GetAllProducts(ctx context.Context) ([]Product.Product, error) {
	ctx, span := tracer.Start(ctx, "GetAllProducts_Repo")
	defer span.End()

	var products []Product.Product

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
