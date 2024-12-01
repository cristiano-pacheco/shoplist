package dbtest

import (
	"database/sql"
	"io"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cristiano-pacheco/shoplist/internal/shared/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDBMock(t *testing.T) (*sql.DB, *database.ShoplistDB, sqlmock.Sqlmock) {
	sqldb, mock, err := sqlmock.New()
	assert.NoError(t, err)

	gormdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	assert.NoError(t, err)

	db := database.NewFromGorm(gormdb)

	return sqldb, db, mock
}

func CloseWithErrorCheck(closer io.Closer) {
	if err := closer.Close(); err != nil {
		log.Printf("failed to close resource: %v", err)
	}
}
