package handler

import (
	m "bam/internal/app/model"

	"github.com/gofiber/fiber/v2"
)

type ProductActions interface {
	CreateProduct(prod *m.Product) error
	GetProductByID(id uint) (*m.Product, error)
	UpdateProduct(prod *m.Product) error
	DeleteProduct(id uint) error
	FindProductByName(name string) ([]*m.Product, error)
	GetProducts() ([]*m.Product, error)
}

type ProductHandler struct {
	prodService ProductActions
}

func NewProductHandler(service ProductActions) *ProductHandler {
	return &ProductHandler{prodService: service}
}

// func (h *ProductHandler) RegisterProductRoutes(app *fiber.App) {
// 	app.Post("/products", h.CreateProduct)
// 	app.Get("/products/:id", h.GetProductByID)
// 	app.Put("/products/:id", h.UpdateProduct)
// 	app.Delete("/products/:id", h.DeleteProduct)
// 	app.Get("/products", h.GetProducts)
// 	app.Get("/name", h.FindProductByName)

// }

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {

	prod := new(m.Product)
	if err := c.BodyParser(prod); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.prodService.CreateProduct(prod)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(prod)

}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	prod, err := h.prodService.GetProductByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(prod)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	prod := new(m.Product)
	if err := c.BodyParser(prod); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	prod.ID = uint(id)

	if err := h.prodService.UpdateProduct(prod); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(prod)

}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.prodService.DeleteProduct(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// return c.SendStatus(fiber.StatusNoContent)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"complete": "Deleted"})
}

// Handler that finds a product by name using query parameters.
func (h *ProductHandler) FindProductByName(c *fiber.Ctx) error {
	name := c.Query("name")
	if name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name is required"})
	}

	prod, err := h.prodService.FindProductByName(name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(prod)
}

func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	prods, err := h.prodService.GetProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"products": prods, "total": len(prods), "message": "success", "status": 200, "success": true})
}
