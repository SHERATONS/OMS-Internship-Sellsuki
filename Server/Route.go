package Server

import (
	"github.com/SHERATONS/OMS-Sellsuki-Internship/Handlers"
	"github.com/SHERATONS/OMS-Sellsuki-Internship/UseCases"
)

func (s *FiberServer) SetupRoute(u UseCases.IProductCase) {
	ProductHandler := Handlers.NewProductHandler(u)
	s.app.Get("/products/", ProductHandler.GetAllProducts)
	s.app.Get("/product/:id", ProductHandler.GetProductById)
	s.app.Post("/createProduct/", ProductHandler.CreateProduct)
	s.app.Put("/updateProduct/:id", ProductHandler.UpdateProductById)
	s.app.Delete("/deleteProduct/:id", ProductHandler.DeleteProductById)
}
