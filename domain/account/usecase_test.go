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
	jsonWebToken.On("Sign", mock.Anything, mock.AnythingOfType("jwt.AccountStandardJWTClaims")).Return("mock token", nil)

	bcrypt := new(bcryptMocks.Bcrypt)
	bcrypt.On("HashPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("hashed password")

	accountRepository := new(mocks.AccountRepository)
	accountRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(account.Account{}, exception.ErrNotFound)
	accountRepository.On("Save", mock.Anything, mock.AnythingOfType("account.Account")).Return(int64(1), nil)

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
