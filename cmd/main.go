package main

import (
	r "bam/internal/app/repository"
	s "bam/internal/app/service"

	"bam/internal/infrastructure/database"
	p "bam/route"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	DBUser string `json:"db_user"`
	DBPass string `json:"db_pass"`
	DBHost string `json:"db_host"`
	DBPort string `json:"db_port"`
	DBName string `json:"db_name"`
}

func main() {
	app := fiber.New()

	configFile, err := os.Open("C:/Users/Sirin/OneDrive/เอกสาร/go/v2/internal/config/config.json")

	if err != nil {
		panic(err)
	}
	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			panic(err)
		}
	}(configFile)

	var config Config
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)

	db, err := database.ConnectDB(dsn)
	if err != nil {
		panic(err)
	}

	userRepo := r.NewUserRepository(db)
	userService := s.NewUserService(userRepo)

	prodRepo := r.NewProductRepository(db)
	prodService := s.NewProductService(prodRepo)

	excelRepo := r.NewExcelRepository(db)
	excelService := s.NewExcelService(excelRepo)

	approveRepo := r.NewApproveRepository(db)
	approveService := s.NewApproveService(approveRepo)

	p.RegisterRoutes(app, userService, prodService, excelService, approveService)

	err = app.Listen(":8080")
	if err != nil {
		return
	}
}
