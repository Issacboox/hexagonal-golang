// route/routes.go
package route

import (
	"bam/internal/app/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, userService handler.UserActions, prodService handler.ProductActions, excelService handler.ExcelActions, approveService handler.ApproveActions) {
	userHandler := handler.NewUserHandler(userService)
	prodHandler := handler.NewProductHandler(prodService)
	excelHandler := handler.NewExcelHandler(excelService)
	approveHandler := handler.NewApproveHandler(approveService)

	v1 := app.Group("/api/v1")
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
	v1.Delete("/ordination/:id", approveHandler.DeleteOrdination)
	// เอาทั้งหมด
	v1.Get("/ordination", approveHandler.FindOrdinations)
	// ค้นหาจากชื่อ - นามสกุล
	v1.Get("/ordination/name", approveHandler.FindOrdinationByName)
	// ค้นหาจาก สถานะ
	v1.Get("/ordination/status", approveHandler.FindOrdinationByStatus)
	// update สถานะ ถ้า cancel , reject ต้องใส่ comment ด้วย
	v1.Put(("/ordination/status/:id"), approveHandler.UpdateOrdinationStatus)
}
