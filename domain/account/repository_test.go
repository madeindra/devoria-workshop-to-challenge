package account

import (
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/constant"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/mock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

var accountPassword = "secret"
var currentTime = time.Date(2021, 12, 12, 0, 0, 0, 0, &time.Location{})

var account = &Account{
	ID:             1,
	FirstName:      "User",
	Email:          "user@example.com",
	LastName:       "08123456789",
	Password:       &accountPassword,
	CreatedAt:      currentTime,
	LastModifiedAt: &currentTime,
}

func TestFindByID(t *testing.T) {
	db, mock := mock.NewMock()
	repo := NewAccountRepository(db, constant.TableAccount)

	defer db.Close()

	query := fmt.Sprintf(`SELECT id, email, password, firstName, lastName, createdAt, lastModifiedAt FROM %s WHERE id = ?`, constant.TableAccount)
	rows := sqlmock.NewRows([]string{"id", "email", "password", "firstName", "lastName", "createdAt", "lastModifiedAt"}).AddRow(account.ID, account.Email, account.Password, account.FirstName, account.LastName, account.CreatedAt, account.LastModifiedAt)

	ctx := context.TODO()

	mock.ExpectPrepare(query).ExpectQuery().WithArgs(account.ID).WillReturnRows(rows)

	user, err := repo.FindByID(ctx, account.ID)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestFindByIDError(t *testing.T) {
	db, mock := mock.NewMock()
	repo := NewAccountRepository(db, constant.TableAccount)
	defer db.Close()

	query := fmt.Sprintf(`SELECT id, email, password, firstName, lastName, createdAt, lastModifiedAt FROM %s WHERE id = ?`, constant.TableAccount)
	rows := sqlmock.NewRows([]string{"id", "email", "password", "firstName", "lastName", "createdAt", "lastModifiedAt"})

	ctx := context.TODO()

	mock.ExpectPrepare(query).ExpectQuery().WithArgs(account.ID).WillReturnRows(rows)

	user, err := repo.FindByID(ctx, account.ID)
	assert.Empty(t, user)
	assert.Error(t, err)
}

func TestFindByEmail(t *testing.T) {
	db, mock := mock.NewMock()
	repo := NewAccountRepository(db, constant.TableAccount)

	defer db.Close()

	query := fmt.Sprintf(`SELECT id, email, password, firstName, lastName, createdAt, lastModifiedAt FROM %s WHERE email = ?`, constant.TableAccount)
	rows := sqlmock.NewRows([]string{"id", "email", "password", "firstName", "lastName", "createdAt", "lastModifiedAt"}).AddRow(account.ID, account.Email, account.Password, account.FirstName, account.LastName, account.CreatedAt, account.LastModifiedAt)

	ctx := context.TODO()

	mock.ExpectPrepare(query).ExpectQuery().WithArgs(account.Email).WillReturnRows(rows)

	user, err := repo.FindByEmail(ctx, account.Email)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestFindByEmailError(t *testing.T) {
	db, mock := mock.NewMock()
	repo := NewAccountRepository(db, constant.TableAccount)
	defer db.Close()

	query := fmt.Sprintf(`SELECT id, email, password, firstName, lastName, createdAt, lastModifiedAt FROM %s WHERE id = ?`, constant.TableAccount)
	rows := sqlmock.NewRows([]string{"id", "email", "password", "firstName", "lastName", "createdAt", "lastModifiedAt"})

	ctx := context.TODO()

	mock.ExpectPrepare(query).ExpectQuery().WithArgs(account.Email).WillReturnRows(rows)

	account, err := repo.FindByEmail(ctx, account.Email)
	assert.Empty(t, account)
	assert.Error(t, err)
}

func TestCreate(t *testing.T) {
	db, mock := mock.NewMock()
	repo := NewAccountRepository(db, constant.TableAccount)

	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO %s`, constant.TableAccount)

	ctx := context.TODO()

	mock.ExpectPrepare(query).ExpectExec().WithArgs(account.Email, account.Password, account.FirstName, account.LastName, account.CreatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

	ID, err := repo.Create(ctx, *account)
	assert.Equal(t, int64(1), ID)
	assert.NoError(t, err)
}

func TestCreateError(t *testing.T) {
	db, mock := mock.NewMock()
	repo := NewAccountRepository(db, constant.TableAccount)

	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO %s`, constant.TableAccount)

	ctx := context.TODO()

	mock.ExpectPrepare(query).ExpectExec().WithArgs(account.Email, account.Password, account.FirstName, account.LastName, account.CreatedAt).WillReturnResult(sqlmock.NewResult(0, 0))

	ID, err := repo.Create(ctx, *account)
	assert.Equal(t, int64(0), ID)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	db, mock := mock.NewMock()
	repo := NewAccountRepository(db, constant.TableAccount)

	defer db.Close()

	query := fmt.Sprintf(`UPDATE %s SET`, constant.TableAccount)

	ctx := context.TODO()

	mock.ExpectPrepare(query).ExpectExec().WithArgs(account.Password, account.FirstName, account.LastName, account.LastModifiedAt).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Update(ctx, account.ID, *account)

	assert.NoError(t, err)
}

func TestUpdateError(t *testing.T) {
	db, mock := mock.NewMock()
	repo := NewAccountRepository(db, constant.TableAccount)

	defer db.Close()

	query := fmt.Sprintf(`UPDATE %s SET`, constant.TableAccount)

	ctx := context.TODO()

	mock.ExpectPrepare(query).ExpectExec().WithArgs(account.Password, account.FirstName, account.LastName, account.LastModifiedAt).WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.Update(ctx, account.ID, *account)

	assert.Error(t, err)
}
