// route/routes.go
package route

import (
	h "bam/internal/app/handler"
	m "bam/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type AuthActions = h.AuthActions

func RegisterRoutes(app *fiber.App, userService h.UserActions, prodService h.ProductActions, excelService h.ExcelActions, approveService h.ApproveActions) {
	userHandler := h.NewUserHandler(userService)
	prodHandler := h.NewProductHandler(prodService)
	excelHandler := h.NewExcelHandler(excelService)
	approveHandler := h.NewApproveHandler(approveService)
	// authHandler := h.NewAuthHandler(userService, jwtService)

	v1 := app.Group("/api/v1")

	// Apply AuthRequired middleware to routes that require authentication
	// v1.Use(m.AuthRequired())

	v1.Post("/users", userHandler.CreateUser)
	v1.Get("/users/:id", userHandler.GetUser)
	v1.Put("/users/:id", userHandler.UpdateUser)
	v1.Delete("/users/:id", userHandler.DeleteUser)
	v1.Get("/users", userHandler.GetUsers)

	v1.Post("/products", prodHandler.CreateProduct)
	v1.Get("/products/:id", prodHandler.GetProductByID)
	v1.Put("/products/:id", prodHandler.UpdateProduct)
	v1.Delete("/products/:id", prodHandler.DeleteProduct)
	v1.Get("/products", prodHandler.GetProducts)
	v1.Get("/name", prodHandler.FindProductByName)
	// import from excel?
	v1.Post("/products/upload", prodHandler.InsertProductsFromExcel)
	// read from excel show as json
	v1.Post("/products/read", excelHandler.ReadExcel)
	// write to excel file
	v1.Get("/write", excelHandler.ExportDataToExcel)

	// ลงทะเบียนลาบวช
	v1.Post("/ordination", approveHandler.RegisterOrdination)
	// ค้นหาใช้ id
	v1.Get("/ordination/:id", approveHandler.FindOrdinationByID)
	// แก้ไข
	v1.Put("/ordination/:id", approveHandler.UpdateOrdination)
	// ลบ
	v1.Delete("/ordination/:id", m.AuthRequired(), approveHandler.DeleteOrdination)
	// เอาทั้งหมด
	v1.Get("/ordination", approveHandler.FindOrdinations)
	// ค้นหาจากชื่อ - นามสกุล
	v1.Get("/ordname", approveHandler.FindOrdinationByName)
	// ค้นหาจาก สถานะ
	v1.Get("/ordstatus", approveHandler.FindOrdinationByStatus)
	// update สถานะ ถ้า cancel , reject ต้องใส่ comment ด้วย
	v1.Put(("/ordstatus/:id"), approveHandler.UpdateOrdinationStatus)

	// Add login and logout routes
	// v1.Post("/login", authHandler.Login)
	// v1.Post("/logout", authHandler.Logout)
}
