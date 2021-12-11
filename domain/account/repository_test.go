package account

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/constant"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

var account = &Account{
	ID:             1,
	FirstName:      "User",
	Email:          "user@example.com",
	LastName:       "08123456789",
	Password:       nil,
	CreatedAt:      time.Date(2021, 12, 12, 0, 0, 0, 0, &time.Location{}),
	LastModifiedAt: nil,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestFindByID(t *testing.T) {
	db, mock := NewMock()
	repo := NewAccountRepository(db, constant.TableAccount)

	defer db.Close()

	query := fmt.Sprintf(`SELECT id, email, password, firstName, lastName, createdAt, lastModifiedAt FROM %s WHERE id = ?`, constant.TableAccount)
	rows := sqlmock.NewRows([]string{"id", "email", "password", "firstName", "lastName", "createdAt", "lastModifiedAt"}).AddRow(account.ID, account.Email, account.Password, account.FirstName, account.LastName, account.CreatedAt, account.LastModifiedAt)

	ctx := context.Background()

	mock.ExpectPrepare(query).ExpectQuery().WithArgs(account.ID).WillReturnRows(rows)

	user, err := repo.FindByID(ctx, account.ID)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}
