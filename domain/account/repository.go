package account

import (
	"errors"

	"github.com/madeindra/devoria-workshop-to-challenge/internal/exception"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type (
	AccountRepository interface {
		FindByID(ctx context.Context, ID int64) (Account, error)
		FindByEmail(ctx context.Context, email string) (Account, error)
		Create(ctx context.Context, account Account) (int64, error)
		Update(ctx context.Context, ID int64, account Account) (Account, error)
	}

	accountRepositoryImpl struct {
		db        *gorm.DB
		tableName string
	}
)

func NewAccountRepository(db *gorm.DB, tableName string) AccountRepository {
	return &accountRepositoryImpl{
		db:        db,
		tableName: tableName,
	}
}

// Select Account by ID (PK)
func (repo *accountRepositoryImpl) FindByID(ctx context.Context, ID int64) (Account, error) {
	account := Account{}
	result := repo.db.First(&account, ID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return account, exception.ErrNotFound
	}

	if result.Error != nil {
		return account, exception.ErrInternalServer
	}

	return account, nil
}

// Select Account by Email
func (repo *accountRepositoryImpl) FindByEmail(ctx context.Context, email string) (Account, error) {
	account := Account{}
	result := repo.db.Where(&Account{Email: email}).First(&account)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return account, exception.ErrNotFound
	}

	if result.Error != nil {
		return account, exception.ErrInternalServer
	}

	return account, nil
}

// Insert into Account
func (repo *accountRepositoryImpl) Create(ctx context.Context, account Account) (int64, error) {
	newAccount := Account{
		Email:     account.Email,
		Password:  account.Password,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		CreatedAt: account.CreatedAt,
	}

	result := repo.db.Create(&newAccount)

	if result.Error != nil {
		return 0, exception.ErrInternalServer
	}

	return newAccount.ID, nil
}

// Update Account
func (repo *accountRepositoryImpl) Update(ctx context.Context, ID int64, account Account) (Account, error) {
	updatedAccount := Account{}

	result := repo.db.Model(&account).UpdateColumns(account).Find(&updatedAccount)

	if result.Error != nil {
		return updatedAccount, exception.ErrInternalServer
	}

	return updatedAccount, nil
}
