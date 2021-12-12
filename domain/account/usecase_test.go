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

func TestUsecaseRegister_HashError(t *testing.T) {
	jsonWebToken := new(jsonWebTokenMocks.JSONWebToken)
	bcrypt := new(bcryptMocks.Bcrypt)
	accountRepository := new(mocks.AccountRepository)

	bcrypt.On("HashPassword", mock.AnythingOfType("string")).Return("", exception.ErrInternalServer)

	accountRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(account.Account{}, exception.ErrNotFound)

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

func TestUsecaseLogin_NotRegistered(t *testing.T) {
	jsonWebToken := new(jsonWebTokenMocks.JSONWebToken)
	bcrypt := new(bcryptMocks.Bcrypt)
	accountRepository := new(mocks.AccountRepository)

	accountRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(account.Account{}, exception.ErrNotFound)

	accountUsecase := account.NewAccountUsecase(
		accountRepository,
		bcrypt,
		jsonWebToken,
	)

	ctx := context.TODO()
	params := account.AccountLoginRequest{
		Email: "user@example.com",
	}
	resp := accountUsecase.Login(ctx, params)

	assert.Error(t, resp.Err())

	accountRepository.AssertExpectations(t)
	jsonWebToken.AssertExpectations(t)
	bcrypt.AssertExpectations(t)
}

func TestUsecaseLogin_QueryError(t *testing.T) {
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
	params := account.AccountLoginRequest{
		Email: "user@example.com",
	}
	resp := accountUsecase.Login(ctx, params)

	assert.Error(t, resp.Err())

	accountRepository.AssertExpectations(t)
	jsonWebToken.AssertExpectations(t)
	bcrypt.AssertExpectations(t)
}

func TestUsecaseLogin_InvalidPassword(t *testing.T) {
	jsonWebToken := new(jsonWebTokenMocks.JSONWebToken)
	bcrypt := new(bcryptMocks.Bcrypt)
	accountRepository := new(mocks.AccountRepository)

	password := "hashed"

	mockAccount := account.Account{
		Password: &password,
	}

	bcrypt.On("ComparePasswordHash", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(false)

	accountRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(mockAccount, nil)

	accountUsecase := account.NewAccountUsecase(
		accountRepository,
		bcrypt,
		jsonWebToken,
	)

	ctx := context.TODO()
	params := account.AccountLoginRequest{
		Email: "user@example.com",
	}
	resp := accountUsecase.Login(ctx, params)

	assert.NoError(t, resp.Err())

	accountRepository.AssertExpectations(t)
	jsonWebToken.AssertExpectations(t)
	bcrypt.AssertExpectations(t)
}

func TestUsecaseLogin_SignFailed(t *testing.T) {
	jsonWebToken := new(jsonWebTokenMocks.JSONWebToken)
	bcrypt := new(bcryptMocks.Bcrypt)
	accountRepository := new(mocks.AccountRepository)

	password := "hashed"

	mockAccount := account.Account{
		Password: &password,
	}

	jsonWebToken.On("Sign", mock.Anything, mock.AnythingOfType("jwt.AccountStandardJWTClaims")).Return("", exception.ErrInternalServer)

	bcrypt.On("ComparePasswordHash", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(true)

	accountRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(mockAccount, nil)

	accountUsecase := account.NewAccountUsecase(
		accountRepository,
		bcrypt,
		jsonWebToken,
	)

	ctx := context.TODO()
	params := account.AccountLoginRequest{
		Email: "user@example.com",
	}
	resp := accountUsecase.Login(ctx, params)

	assert.Error(t, resp.Err())

	accountRepository.AssertExpectations(t)
	jsonWebToken.AssertExpectations(t)
	bcrypt.AssertExpectations(t)
}

func TestUsecaseLogin_Success(t *testing.T) {
	jsonWebToken := new(jsonWebTokenMocks.JSONWebToken)
	bcrypt := new(bcryptMocks.Bcrypt)
	accountRepository := new(mocks.AccountRepository)

	password := "hashed"

	mockAccount := account.Account{
		Password: &password,
	}

	jsonWebToken.On("Sign", mock.Anything, mock.AnythingOfType("jwt.AccountStandardJWTClaims")).Return("jwttoken", nil)

	bcrypt.On("ComparePasswordHash", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(true)

	accountRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(mockAccount, nil)

	accountUsecase := account.NewAccountUsecase(
		accountRepository,
		bcrypt,
		jsonWebToken,
	)

	ctx := context.TODO()
	params := account.AccountLoginRequest{
		Email: "user@example.com",
	}
	resp := accountUsecase.Login(ctx, params)

	assert.NoError(t, resp.Err())

	accountRepository.AssertExpectations(t)
	jsonWebToken.AssertExpectations(t)
	bcrypt.AssertExpectations(t)
}
