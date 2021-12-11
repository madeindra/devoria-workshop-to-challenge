package article

import (
	"context"

	"github.com/madeindra/devoria-workshop-to-challenge/domain/account"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/response"
)

type ArticleUsecase interface {
	CreateArticle(ctx context.Context, params CreateArticleRequest) response.Response
	UpdateArticle(ctx context.Context, params UpdateArticleRequest) response.Response
	GetOneArticle(ctx context.Context, ID int64) response.Response
	GetAllArticles(ctx context.Context) response.Response
}

type articleUsecaseImpl struct {
	repository  ArticleRepository
	accountRepo account.AccountRepository
}

func NewArticleUsecase(
	repository ArticleRepository,
	accountRepo account.AccountRepository,
) ArticleUsecase {
	return &articleUsecaseImpl{
		repository:  repository,
		accountRepo: accountRepo,
	}
}

func (uc *articleUsecaseImpl) CreateArticle(ctx context.Context, params CreateArticleRequest) response.Response {
	return response.Success(response.StatusCreated, nil)
}

func (uc *articleUsecaseImpl) UpdateArticle(ctx context.Context, params UpdateArticleRequest) response.Response {
	return response.Success(response.StatusOK, nil)
}

func (uc *articleUsecaseImpl) GetOneArticle(ctx context.Context, ID int64) response.Response {
	return response.Success(response.StatusOK, nil)
}

func (uc *articleUsecaseImpl) GetAllArticles(ctx context.Context) response.Response {
	return response.Success(response.StatusOK, nil)
}
