package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// "wansanjou/adapters"
	"wansanjou/adapters"
	"wansanjou/core"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)


func main() {
	app := fiber.New()

	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
  "password=%s dbname=%s sslmode=disable",
  host, port, user, password, dbname)
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	})
  if err != nil {
    panic("failed to connect to database")
  }

	
	db.AutoMigrate(&core.Order{})

	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := core.NewOrderService(orderRepo)
	orderHanler := adapters.NewHttpOrderHandler(orderService)

	app.Post("/order" , orderHanler.CreateOrder)


	app.Listen(":8080")
}