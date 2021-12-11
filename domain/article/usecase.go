package article

import (
	"context"
	"fmt"
	"time"

	"github.com/madeindra/devoria-workshop-to-challenge/domain/account"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/exception"
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
	email := ctx.Value(account.EmailContex).(string)

	author, err := uc.accountRepo.FindByEmail(ctx, email)

	if err == exception.ErrNotFound {
		return response.Error(response.StatusNotFound, exception.ErrNotFound)
	}

	article := Article{
		Title:     params.Title,
		Subtitle:  params.Subtitle,
		Content:   params.Content,
		CreatedAt: time.Now(),
		Author:    author,
	}

	if params.IsPublished {
		article.Status = ArticleStatusPublished
		article.PublishedAt = &article.CreatedAt
	} else {
		article.Status = ArticleStatusDraft
	}

	ID, err := uc.repository.Create(ctx, article)
	if err != nil {
		fmt.Println(err)
		return response.Error(response.StatusInternalServerError, err)
	}

	article.ID = ID
	article.Author.Password = nil

	return response.Success(response.StatusCreated, article)
}

func (uc *articleUsecaseImpl) UpdateArticle(ctx context.Context, params UpdateArticleRequest) response.Response {
	return response.Success(response.StatusOK, nil)
}

func (uc *articleUsecaseImpl) GetOneArticle(ctx context.Context, ID int64) response.Response {
	article := Article{}

	article, err := uc.repository.FindByID(ctx, ID)

	if err == exception.ErrNotFound {
		return response.Error(response.StatusNotFound, exception.ErrNotFound)
	}

	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	return response.Success(response.StatusOK, article)
}

func (uc *articleUsecaseImpl) GetAllArticles(ctx context.Context) response.Response {
	articles, err := uc.repository.FindAll(ctx)
	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	var mapped []Article

	for _, item := range articles {
		article := Article{}
		account := account.Account{}

		article.ID = item.ID
		article.Title = item.Title
		article.Content = item.Content
		article.Status = item.Status
		article.CreatedAt = item.CreatedAt
		article.LastModifiedAt = item.LastModifiedAt

		account.ID = item.Author.ID
		account.Email = item.Author.Email
		account.FirstName = item.Author.FirstName
		account.LastName = item.Author.LastName
		account.CreatedAt = item.Author.CreatedAt
		account.LastModifiedAt = item.Author.LastModifiedAt

		article.Author = account

		mapped = append(mapped, article)
	}

	return response.Success(response.StatusOK, mapped)
}
