package account

import (
	"time"

	"github.com/madeindra/devoria-workshop-to-challenge/internal/constant"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/response"
	"golang.org/x/net/context"
)

type AccountUsecase interface {
	Register(ctx context.Context, params AccountRegisterRequest) response.Response
	Login(ctx context.Context, params AccountLoginRequest) response.Response
	GetAccount(ctx context.Context) response.Response
}

type accountUsecaseImpl struct {
	repository AccountRepository
}

func NewAccountUsecase(repository AccountRepository) AccountUsecase {
	return &accountUsecaseImpl{
		repository: repository,
	}
}

// Registration usecase
func (uc *accountUsecaseImpl) Register(ctx context.Context, params AccountRegisterRequest) response.Response {
	account := Account{
		Email:     params.Email,
		Password:  &params.Password,
		FirstName: params.FirstName,
		LastName:  params.LastName,
		CreatedAt: time.Now(),
	}

	ID, err := uc.repository.Create(ctx, account)
	if err != nil {
		return response.Error(err)
	}
	account.ID = ID

	return response.Success(constant.MessageRegisterSuccess, account)
}

// Login usecase
func (uc *accountUsecaseImpl) Login(ctx context.Context, params AccountLoginRequest) response.Response {
	return response.Success(constant.MessageLoginSuccess, ctx)
}

// Get Account usecase
func (uc *accountUsecaseImpl) GetAccount(ctx context.Context) response.Response {
	return response.Success(constant.MessageGeneralSuccess, ctx)
}
