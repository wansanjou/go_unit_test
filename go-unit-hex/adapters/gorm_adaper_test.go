package adapters

import (
	"errors"
	"testing"
	"wansanjou/core"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGormOrderRepository_Save(t *testing.T) {
  // Mock database
  db, mock, err := sqlmock.New()
  if err != nil {
    t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
  }
  defer db.Close()

  gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
  if err != nil {
    t.Fatalf("Failed to open gorm database: %v", err)
  }

	repo := NewGormOrderRepository(gormDB)

	// Success case
  t.Run("success", func(t *testing.T) {
    // Setup expectations
    mock.ExpectBegin()
    mock.ExpectQuery(`INSERT INTO "orders"`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
    mock.ExpectCommit()

    err := repo.Save(core.Order{Total: 100})
    assert.NoError(t, err)

    // Ensure all expectations were met
    assert.NoError(t, mock.ExpectationsWereMet())
  })

	// Failure case
  t.Run("failure", func(t *testing.T) {
    // Setup expectations
    mock.ExpectBegin()
    mock.ExpectQuery(`INSERT INTO "orders"`).WillReturnError(errors.New("database error"))
    mock.ExpectRollback()

    err := repo.Save(core.Order{Total: 100})
    assert.Error(t, err)

    // Ensure all expectations were met
    assert.NoError(t, mock.ExpectationsWereMet())
  })
}