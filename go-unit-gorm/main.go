package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)




type User struct {
  gorm.Model
  Fullname string
  Email    string `gorm:"unique"`
  Age      int
}

// InitializeDB initializes the database and automigrates the User model.
func InitializeDB() *gorm.DB {
	const (
		host     = "localhost"  // or the Docker service name if running in another container
		port     = 5432         // default PostgreSQL port
		user     = "myuser"     // as defined in docker-compose.yml
		password = "mypassword" // as defined in docker-compose.yml
		dbname   = "mydatabase" // as defined in docker-compose.yml
	)

	// New logger for detailed SQL logging
  newLogger := logger.New(
    log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
    logger.Config{
      SlowThreshold: time.Second, // Slow SQL threshold
      LogLevel:      logger.Info, // Log level
      Colorful:      true,        // Enable color
    },
  )
	
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
  "password=%s dbname=%s sslmode=disable",
  host, port, user, password, dbname)
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger, // add Logger
	})
  if err != nil {
    panic("failed to connect to database")
  }
  db.AutoMigrate(&User{})
  return db
}

// AddUser adds a new user to the database.
func AddUser(db *gorm.DB, fullname, email string, age int) error {
  user := User{Fullname: fullname, Email: email, Age: age}

  // Check if email already exists
  var count int64
  db.Model(&User{}).Where("email = ?", email).Count(&count)
  if count > 0 {
    return errors.New("email already exists")
  }

  // Save the new user
  result := db.Create(&user)
  return result.Error
}


func main() {
  db := InitializeDB()
  // Your application code
  err := AddUser(db, "jj Doe", "jj.doe@example.com", 30)
	if err != nil {
		fmt.Println("err DB" , err)
	}
}