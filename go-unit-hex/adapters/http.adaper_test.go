package adapters

import (
  "bytes"
  "errors"
  "wansanjou/core"
  "net/http/httptest"
  "testing"

  "github.com/gofiber/fiber/v2"
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/mock"
)

// MockOrderService is a mock implementation of core.OrderService
type MockOrderService struct {
  mock.Mock
}

func (m *MockOrderService) CreateOrder(order core.Order) error {
  args := m.Called(order)
  return args.Error(0)
}

// TestCreateOrderHandler tests the CreateOrder handler of HttpOrderHandler
func TestCreateOrderHandler(t *testing.T) {
  mockService := new(MockOrderService)
  handler := NewHttpOrderHandler(mockService)

  app := fiber.New()
  app.Post("/orders", handler.CreateOrder)

  // Test case: Successful order creation
  t.Run("successful order creation", func(t *testing.T) {
    mockService.On("CreateOrder", mock.AnythingOfType("core.Order")).Return(nil)

    req := httptest.NewRequest("POST", "/orders", bytes.NewBufferString(`{"total": 100}`))
    req.Header.Set("Content-Type", "application/json")
    resp, err := app.Test(req)

    assert.NoError(t, err)
    assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
    mockService.AssertExpectations(t)
  })

  t.Run("fail order creation (total less than 0)", func(t *testing.T) {
    mockService.ExpectedCalls = nil
    mockService.On("CreateOrder", mock.AnythingOfType("core.Order")).Return(errors.New("total must be positive"))

    req := httptest.NewRequest("POST", "/orders", bytes.NewBufferString(`{"total": -200}`))
    req.Header.Set("Content-Type", "application/json")
    resp, err := app.Test(req)

    assert.NoError(t, err)
    assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
    mockService.AssertExpectations(t)
  })

  // Test case: Invalid request body
  t.Run("invalid request body", func(t *testing.T) {
    req := httptest.NewRequest("POST", "/orders", bytes.NewBufferString(`{"total": "invalid"}`))
    req.Header.Set("Content-Type", "application/json")
    resp, err := app.Test(req)

    assert.NoError(t, err)
    assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
  })

  // Test case: Order service returns an error
  t.Run("order service error", func(t *testing.T) {
    mockService.ExpectedCalls = nil
    mockService.On("CreateOrder", mock.AnythingOfType("core.Order")).Return(errors.New("service error"))

    req := httptest.NewRequest("POST", "/orders", bytes.NewBufferString(`{"total": 100}`))
    req.Header.Set("Content-Type", "application/json")
    resp, err := app.Test(req)

    assert.NoError(t, err)
    assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
    mockService.AssertExpectations(t)
  })
}