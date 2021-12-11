package account

import (
	"time"

	"github.com/madeindra/devoria-workshop-to-challenge/internal/bcrypt"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/exception"
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
}

func NewAccountUsecase(repository AccountRepository, bcrypt bcrypt.Bcrypt) AccountUsecase {
	return &accountUsecaseImpl{
		repository: repository,
		bcrypt:     bcrypt,
	}
}

// Registration usecase
func (uc *accountUsecaseImpl) Register(ctx context.Context, params AccountRegisterRequest) response.Response {
	_, err := uc.repository.FindByEmail(ctx, params.Email)
	if err == nil {
		return response.Error(response.StatusConflicted, exception.ErrConflicted)
	}

	if err != exception.ErrNotFound {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

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
	return response.Success(response.StatusOK, ctx)
}

// Get Account usecase
func (uc *accountUsecaseImpl) GetAccount(ctx context.Context, ID int64) response.Response {
	account := Account{}

	account, err := uc.repository.FindByID(ctx, ID)
	if err == exception.ErrNotFound {
		return response.Error(response.StatusNotFound, exception.ErrNotFound)
	}

	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	account.Password = nil

	return response.Success(response.StatusOK, account)
}
