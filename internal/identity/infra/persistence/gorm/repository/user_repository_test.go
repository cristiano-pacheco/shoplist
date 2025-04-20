package repository

// import (
// 	"context"
// 	"database/sql/driver"
// 	"fmt"
// 	"regexp"
// 	"testing"
// 	"time"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/cristiano-pacheco/shoplist/internal/shared/model"
// 	"github.com/cristiano-pacheco/shoplist/internal/shared/test/dbtest"
// 	"github.com/stretchr/testify/assert"
// )

// func TestUserRepository_Create_Success(t *testing.T) {
// 	// Arrange
// 	sqldb, db, mock := dbtest.NewDBMock(t)
// 	defer dbtest.CloseWithErrorCheck(sqldb)

// 	sut := NewUserRepository(db)

// 	query := `INSERT INTO "users"`
// 	query += ` ("name","email","password_hash","is_activated")`
// 	query += ` VALUES ($1,$2,$3,$4) RETURNING "id","created_at","updated_at"`

// 	createdAt := time.Now()

// 	mock.ExpectBegin()
// 	mock.ExpectQuery(regexp.QuoteMeta(query)).
// 		WithArgs("John Doe", "john@example.com", "hashedpassword", true).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(1, createdAt, createdAt))
// 	mock.ExpectCommit()

// 	ctx := context.Background()
// 	user := model.UserModel{
// 		Name:         "John Doe",
// 		Email:        "john@example.com",
// 		PasswordHash: "hashedpassword",
// 		IsActivated:  true,
// 	}

// 	// Act
// 	createdUser, err := sut.Create(ctx, user)

// 	// Assert
// 	assert.NoError(t, err)
// 	assert.NotNil(t, createdUser)
// 	assert.Equal(t, "John Doe", createdUser.Name)
// 	assert.Equal(t, "john@example.com", createdUser.Email)
// 	assert.True(t, createdUser.IsActivated)
// 	assert.Equal(t, createdAt, createdUser.CreatedAt)
// 	assert.Equal(t, createdAt, createdUser.UpdatedAt)

// 	err = mock.ExpectationsWereMet()
// 	assert.NoError(t, err)
// }

// func TestUserRepository_Create_Error(t *testing.T) {
// 	// Arrange
// 	sqldb, db, mock := dbtest.NewDBMock(t)
// 	defer dbtest.CloseWithErrorCheck(sqldb)

// 	sut := NewUserRepository(db)

// 	query := `INSERT INTO "users"`
// 	query += ` ("name","email","password_hash","is_activated")`
// 	query += ` VALUES ($1,$2,$3,$4) RETURNING "id","created_at","updated_at"`

// 	mock.ExpectBegin()
// 	mock.ExpectQuery(regexp.QuoteMeta(query)).
// 		WithArgs("John Doe", "john@example.com", "hashedpassword", true).
// 		WillReturnError(fmt.Errorf("error"))
// 	mock.ExpectRollback()

// 	ctx := context.Background()
// 	user := model.UserModel{
// 		Name:         "John Doe",
// 		Email:        "john@example.com",
// 		PasswordHash: "hashedpassword",
// 		IsActivated:  true,
// 	}

// 	// Act
// 	createdUser, err := sut.Create(ctx, user)

// 	// Assert
// 	assert.Error(t, err)
// 	assert.Nil(t, createdUser)
// 	err = mock.ExpectationsWereMet()
// 	assert.NoError(t, err)
// }

// func TestUserRepository_Update_Success(t *testing.T) {
// 	// Arrange
// 	sqldb, db, mock := dbtest.NewDBMock(t)
// 	defer dbtest.CloseWithErrorCheck(sqldb)

// 	sut := NewUserRepository(db)

// 	query := `UPDATE "users"`
// 	query += ` SET "name"=$1,"email"=$2,"password_hash"=$3,"is_activated"=$4`
// 	query += ` WHERE "id" = $5`

// 	mock.ExpectBegin()
// 	mock.ExpectExec(regexp.QuoteMeta(query)).
// 		WithArgs("John Doe", "john@example.com", "hashedpassword", true, 1).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()

// 	ctx := context.Background()
// 	user := model.UserModel{
// 		Base: model.Base{
// 			ID: 1,
// 		},
// 		Name:         "John Doe",
// 		Email:        "john@example.com",
// 		PasswordHash: "hashedpassword",
// 		IsActivated:  true,
// 	}

// 	// Act
// 	err := sut.Update(ctx, user)

// 	// Assert
// 	assert.NoError(t, err)
// 	err = mock.ExpectationsWereMet()
// 	assert.NoError(t, err)
// }

// func TestUserRepository_Update_Error(t *testing.T) {
// 	// Arrange
// 	sqldb, db, mock := dbtest.NewDBMock(t)
// 	defer dbtest.CloseWithErrorCheck(sqldb)

// 	sut := NewUserRepository(db)

// 	query := `UPDATE "users"`
// 	query += ` SET "name"=$1,"email"=$2,"password_hash"=$3,"is_activated"=$4`
// 	query += ` WHERE "id" = $5`

// 	mock.ExpectBegin()
// 	mock.ExpectExec(regexp.QuoteMeta(query)).
// 		WithArgs("John Doe", "john@example.com", "hashedpassword", true, 1).
// 		WillReturnError(fmt.Errorf("error"))
// 	mock.ExpectRollback()

// 	ctx := context.Background()
// 	user := model.UserModel{
// 		Base: model.Base{
// 			ID: 1,
// 		},
// 		Name:         "John Doe",
// 		Email:        "john@example.com",
// 		PasswordHash: "hashedpassword",
// 		IsActivated:  true,
// 	}

// 	// Act
// 	err := sut.Update(ctx, user)

// 	// Assert
// 	assert.Error(t, err)
// 	err = mock.ExpectationsWereMet()
// 	assert.NoError(t, err)
// }

// func TestUserRepository_FindOneByID_Success(t *testing.T) {
// 	// Arrange
// 	sqldb, db, mock := dbtest.NewDBMock(t)
// 	defer dbtest.CloseWithErrorCheck(sqldb)

// 	sut := NewUserRepository(db)

// 	var (
// 		id           uint64 = 1
// 		createdAt           = time.Now()
// 		updatedAt           = time.Now()
// 		name                = "John Doe"
// 		email               = "test@gmail.com"
// 		passwordHash        = "hashedpassword"
// 		isActivated         = true
// 		limit               = 1
// 	)

// 	fields := []string{
// 		"id",
// 		"created_at",
// 		"updated_at",
// 		"name",
// 		"email",
// 		"password_hash",
// 		"is_activated",
// 	}

// 	values := []driver.Value{}
// 	values = append(values, id, createdAt, updatedAt, name, email, passwordHash, isActivated)

// 	query := `SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT $2`
// 	mock.ExpectQuery(regexp.QuoteMeta(query)).
// 		WithArgs(id, limit).
// 		WillReturnRows(sqlmock.NewRows(fields).AddRow(values...))

// 	// Act
// 	user, err := sut.FindOneByID(context.Background(), id)

// 	// Assert
// 	assert.NoError(t, err)
// 	assert.NotNil(t, user)
// 	assert.Equal(t, id, user.ID)
// 	assert.Equal(t, createdAt, user.CreatedAt)
// 	assert.Equal(t, updatedAt, user.UpdatedAt)
// 	assert.Equal(t, name, user.Name)
// 	assert.Equal(t, email, user.Email)
// 	assert.Equal(t, passwordHash, user.PasswordHash)
// 	assert.Equal(t, isActivated, user.IsActivated)
// 	assert.NoError(t, mock.ExpectationsWereMet())
// }

// func TestUserRepository_FindOneByID_Error(t *testing.T) {
// 	// Arrange
// 	sqldb, db, mock := dbtest.NewDBMock(t)
// 	defer dbtest.CloseWithErrorCheck(sqldb)

// 	sut := NewUserRepository(db)

// 	var (
// 		id    uint64 = 1
// 		limit        = 1
// 	)

// 	query := `SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT $2`
// 	mock.ExpectQuery(regexp.QuoteMeta(query)).
// 		WithArgs(id, limit).
// 		WillReturnError(fmt.Errorf("error"))

// 	// Act
// 	user, err := sut.FindOneByID(context.Background(), id)

// 	// Assert
// 	assert.Error(t, err)
// 	assert.Nil(t, user)
// 	assert.NoError(t, mock.ExpectationsWereMet())
// }
