package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestUserRoute(t *testing.T) {
  app := setup()

  // Define test cases
  tests := []struct {
    description  string
    requestBody  User
    expectStatus int
  }{
    {
      description:  "Valid input",
      requestBody:  User{"jane.doe@example.com", "Jane Doe", 30},
      expectStatus: fiber.StatusOK,
    },
    {
      description:  "Invalid email",
      requestBody:  User{"invalid-email", "Jane Doe", 30},
      expectStatus: fiber.StatusBadRequest,
    },
    {
      description:  "Invalid fullname",
      requestBody:  User{"jane.doe@example.com", "12345", 30},
      expectStatus: fiber.StatusBadRequest,
    },
    {
      description:  "Invalid age",
      requestBody:  User{"jane.doe@example.com", "Jane Doe", -5},
      expectStatus: fiber.StatusBadRequest,
    },
  }


	for _, tc := range tests {
    t.Run(tc.description, func(t *testing.T) {
      reqBody, _ := json.Marshal(tc.requestBody)
      req := httptest.NewRequest("POST", "/users", bytes.NewReader(reqBody))
      req.Header.Set("Content-Type", "application/json")
      resp, _ := app.Test(req)

      assert.Equal(t, tc.expectStatus, resp.StatusCode)
    })
  }
}