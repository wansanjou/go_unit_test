package main

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// User struct with validation tags.
type User struct {
  Email    string `json:"email" validate:"required,email"`
  Fullname string `json:"fullname" validate:"required,fullname"`
  Age      int    `json:"age" validate:"required,numeric,min=1"`
}

// setup function initializes the Fiber app.
func setup() *fiber.App {
  app := fiber.New()

  // Register the custom validation function for 'fullname'
  validate.RegisterValidation("fullname", validateFullname)

  app.Post("/users", func(c *fiber.Ctx) error {
    user := new(User)

    if err := c.BodyParser(user); err != nil {
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    if err := validate.Struct(user); err != nil {
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusOK).JSON(user)
  })

  return app
}

// validateFullname checks if the value contains only alphabets and spaces.
func validateFullname(fl validator.FieldLevel) bool {
  return regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(fl.Field().String())
}

func main() {
  app := setup()
  app.Listen(":8000")
}