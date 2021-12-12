package account_test

import (
	"testing"

	"github.com/madeindra/devoria-workshop-to-challenge/domain/account"
	"github.com/madeindra/devoria-workshop-to-challenge/domain/account/mocks"
	bcryptMocks "github.com/madeindra/devoria-workshop-to-challenge/internal/bcrypt/mocks"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/exception"
	jsonWebTokenMocks "github.com/madeindra/devoria-workshop-to-challenge/internal/jwt/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestUsecaseRegister_Success(t *testing.T) {
	jsonWebToken := new(jsonWebTokenMocks.JSONWebToken)
	bcrypt := new(bcryptMocks.Bcrypt)
	accountRepository := new(mocks.AccountRepository)

	bcrypt.On("HashPassword", mock.AnythingOfType("string")).Return("hashed password", nil)

	accountRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(account.Account{}, exception.ErrNotFound)
	accountRepository.On("Create", mock.Anything, mock.AnythingOfType("account.Account")).Return(int64(1), nil)

	accountUsecase := account.NewAccountUsecase(
		accountRepository,
		bcrypt,
		jsonWebToken,
	)

	ctx := context.TODO()
	params := account.AccountRegisterRequest{
		Email:     "user@example.com",
		Password:  "secret",
		FirstName: "test",
		LastName:  "test",
	}
	resp := accountUsecase.Register(ctx, params)

	assert.NoError(t, resp.Err())

	accountRepository.AssertExpectations(t)
	jsonWebToken.AssertExpectations(t)
	bcrypt.AssertExpectations(t)
}

func TestUsecaseRegister_CreateError(t *testing.T) {
	jsonWebToken := new(jsonWebTokenMocks.JSONWebToken)
	bcrypt := new(bcryptMocks.Bcrypt)
	accountRepository := new(mocks.AccountRepository)

	bcrypt.On("HashPassword", mock.AnythingOfType("string")).Return("hashed password", nil)

	accountRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(account.Account{}, exception.ErrNotFound)
	accountRepository.On("Create", mock.Anything, mock.AnythingOfType("account.Account")).Return(int64(0), exception.ErrInternalServer)

	accountUsecase := account.NewAccountUsecase(
		accountRepository,
		bcrypt,
		jsonWebToken,
	)

	ctx := context.TODO()
	params := account.AccountRegisterRequest{
		Email:     "user@example.com",
		Password:  "secret",
		FirstName: "test",
		LastName:  "test",
	}
	resp := accountUsecase.Register(ctx, params)

	assert.Error(t, resp.Err())

	accountRepository.AssertExpectations(t)
	jsonWebToken.AssertExpectations(t)
	bcrypt.AssertExpectations(t)
}

func TestUsecaseRegister_Conflict(t *testing.T) {
	jsonWebToken := new(jsonWebTokenMocks.JSONWebToken)
	bcrypt := new(bcryptMocks.Bcrypt)
	accountRepository := new(mocks.AccountRepository)

	accountRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(account.Account{}, nil)

	accountUsecase := account.NewAccountUsecase(
		accountRepository,
		bcrypt,
		jsonWebToken,
	)

	ctx := context.TODO()
	params := account.AccountRegisterRequest{
		Email:     "user@example.com",
		Password:  "secret",
		FirstName: "test",
		LastName:  "test",
	}
	resp := accountUsecase.Register(ctx, params)

	assert.Error(t, resp.Err())

	accountRepository.AssertExpectations(t)
	jsonWebToken.AssertExpectations(t)
	bcrypt.AssertExpectations(t)
}

func TestUsecaseRegister_InternalError(t *testing.T) {
	jsonWebToken := new(jsonWebTokenMocks.JSONWebToken)
	bcrypt := new(bcryptMocks.Bcrypt)
	accountRepository := new(mocks.AccountRepository)

	accountRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(account.Account{}, exception.ErrInternalServer)

	accountUsecase := account.NewAccountUsecase(
		accountRepository,
		bcrypt,
		jsonWebToken,
	)

	ctx := context.TODO()
	params := account.AccountRegisterRequest{
		Email:     "user@example.com",
		Password:  "secret",
		FirstName: "test",
		LastName:  "test",
	}
	resp := accountUsecase.Register(ctx, params)

	assert.Error(t, resp.Err())

	accountRepository.AssertExpectations(t)
	jsonWebToken.AssertExpectations(t)
	bcrypt.AssertExpectations(t)
}
