package core

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock implementation of OrderRepository
type mockOrderRepo struct {
  saveFunc func(order Order) error
}

func (m *mockOrderRepo) Save(order Order) error {
  return m.saveFunc(order)
}

func TestCreateOrder(t *testing.T) {
  // Success case
  t.Run("success", func(t *testing.T) {
    repo := &mockOrderRepo{
      saveFunc: func(order Order) error {
        // Simulate successful save
        return nil
      },
    }
    service := NewOrderService(repo)

    err := service.CreateOrder(Order{Total: 100})
    assert.NoError(t, err)
  })

  // Failure case: Total less than 0
	t.Run("total less than 0", func(t *testing.T) {
    // Arrange
    repo := &mockOrderRepo{} // ไม่ต้องใส่ฟังก์ชัน saveFunc ตรงนี้
    service := NewOrderService(repo)

    // Act
    err := service.CreateOrder(Order{Total: -10})

    // Assert
    if err == nil {
        t.Error("Expected an error but got nil")
    } else {
        assert.Equal(t, "Total must be positive", err.Error())
    }
})


	t.Run("repository error", func(t *testing.T) {
    repo := &mockOrderRepo{
      saveFunc: func(order Order) error {
        return errors.New("database error")
      },
    }
    service := NewOrderService(repo)

    err := service.CreateOrder(Order{Total: 10})
    assert.Error(t, err)
    assert.Equal(t, "database error", err.Error())
  })
}