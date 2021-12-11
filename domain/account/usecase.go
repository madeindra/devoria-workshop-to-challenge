package account

import (
	"fmt"
	"time"

	"github.com/madeindra/devoria-workshop-to-challenge/internal/bcrypt"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/exception"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/jwt"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/response"
	"golang.org/x/net/context"
)

type AccountUsecase interface {
	Register(ctx context.Context, params AccountRegisterRequest) response.Response
	Login(ctx context.Context, params AccountLoginRequest) response.Response
	GetAccount(ctx context.Context, ID int64) response.Response
}

type accountUsecaseImpl struct {
	repository AccountRepository
	bcrypt     bcrypt.Bcrypt
	jwt        jwt.JSONWebToken
}

func NewAccountUsecase(repository AccountRepository, bcrypt bcrypt.Bcrypt, jwt jwt.JSONWebToken) AccountUsecase {
	return &accountUsecaseImpl{
		repository: repository,
		bcrypt:     bcrypt,
		jwt:        jwt,
	}
}

// Registration usecase
func (uc *accountUsecaseImpl) Register(ctx context.Context, params AccountRegisterRequest) response.Response {
	// query account
	_, err := uc.repository.FindByEmail(ctx, params.Email)

	// account found, no error
	if err == nil {
		return response.Error(response.StatusConflicted, exception.ErrConflicted)
	}

	// error querying db
	if err != exception.ErrNotFound {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	// hash password
	hashedPassword, err := uc.bcrypt.HashPassword(params.Password)
	if err != nil {
		return response.Error(response.StatusInternalServerError, err)
	}

	account := Account{
		Email:     params.Email,
		Password:  &hashedPassword,
		FirstName: params.FirstName,
		LastName:  params.LastName,
		CreatedAt: time.Now(),
	}

	// save to db
	ID, err := uc.repository.Create(ctx, account)
	if err != nil {
		return response.Error(response.StatusInternalServerError, err)
	}

	// assign new ID & omit password
	account.ID = ID
	account.Password = nil

	return response.Success(response.StatusCreated, account)
}

// Login usecase
func (uc *accountUsecaseImpl) Login(ctx context.Context, params AccountLoginRequest) response.Response {
	// query account
	account, err := uc.repository.FindByEmail(ctx, params.Email)

	// account not registered
	if err == exception.ErrNotFound {
		return response.Error(response.StatusNotFound, exception.ErrNotFound)
	}

	// error querying
	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	// match password
	isPasswordValid := uc.bcrypt.ComparePasswordHash(params.Password, *account.Password)
	if !isPasswordValid {
		return response.Error(response.StatusUnauthorized, err)
	}

	// omit password
	account.Password = nil

	// generate token
	claims := jwt.AccountStandardJWTClaims{}
	claims.Email = account.Email
	claims.Subject = fmt.Sprintf("%d", account.ID)
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Hour * 24 * 1).Unix()

	token, err := uc.jwt.Sign(ctx, claims)
	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	res := AccountAuthenticationResponse{
		Token:   token,
		Profile: account,
	}

	return response.Success(response.StatusOK, res)
}

// Get Account usecase
func (uc *accountUsecaseImpl) GetAccount(ctx context.Context, ID int64) response.Response {
	// query account
	account, err := uc.repository.FindByID(ctx, ID)

	// account not registered
	if err == exception.ErrNotFound {
		return response.Error(response.StatusNotFound, exception.ErrNotFound)
	}

	// error querying
	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	// omit password
	account.Password = nil

	return response.Success(response.StatusOK, account)
}
