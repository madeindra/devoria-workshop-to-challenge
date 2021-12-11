package article

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/madeindra/devoria-workshop-to-challenge/internal/constant"
	"github.com/madeindra/devoria-workshop-to-challenge/internal/exception"
)

type (
	ArticleRepository interface {
		FindByID(ctx context.Context, ID int64) (Article, error)
		FindAll(ctx context.Context) ([]Article, error)
		FindAllByAuthorId(ctx context.Context, authorId int64) ([]Article, error)
		Create(ctx context.Context, article Article) (int64, error)
		Update(ctx context.Context, ID int64, aritcle Article) error
	}

	articleRepositoryImpl struct {
		db        *sql.DB
		tableName string
	}
)

func NewAccountRepository(db *sql.DB, tableName string) ArticleRepository {
	return &articleRepositoryImpl{
		db:        db,
		tableName: tableName,
	}
}

func (repo *articleRepositoryImpl) FindByID(ctx context.Context, ID int64) (Article, error) {
	article := Article{}

	query := fmt.Sprintf(`SELECT id, title, subtitle, content, status, createdAt, lastModifiedAt, firstName, lastName, email, authorId FROM %s WHERE id = ? JOIN %s ON %s.authorId =%s.id`, repo.tableName, constant.TableAccount, repo.tableName, constant.TableAccount)
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return article, exception.ErrInternalServer
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, ID)

	var publishedAt sql.NullTime
	var lastModifiedAt sql.NullTime

	err = row.Scan(
		&article.ID,
		&article.Title,
		&article.Subtitle,
		&article.Content,
		&article.CreatedAt,
		&publishedAt,
		&lastModifiedAt,
		&article.Author.ID,
	)

	if err != nil {
		log.Println(err)
		return article, exception.ErrNotFound
	}

	if publishedAt.Valid {
		article.PublishedAt = &publishedAt.Time
	}

	if lastModifiedAt.Valid {
		article.LastModifiedAt = &lastModifiedAt.Time
	}

	return article, nil
}

func (repo *articleRepositoryImpl) FindAll(ctx context.Context) ([]Article, error) {
	return nil, nil
}

func (repo *articleRepositoryImpl) FindAllByAuthorId(ctx context.Context, authorId int64) ([]Article, error) {
	// query := fmt.Sprintf(`SELECT id, authorId, title, subtitle, content, status, createdAt, lastModifiedAt, firstName, lastName, email FROM %s WHERE id = ? JOIN %s ON %s.authorId =%s.id`, repo.tableName, constant.TableAccount, repo.tableName, constant.TableAccount)
	return nil, nil
}

func (repo *articleRepositoryImpl) Create(ctx context.Context, article Article) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, subtitle, content, status, createdAt, authorId) VALUES (?, ?, ?, ?, ?)", repo.tableName)
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		ctx,
		article.Title,
		article.Subtitle,
		article.Content,
		article.Status,
		article.CreatedAt,
		article.Author.ID,
	)

	if err != nil {
		return 0, err
	}

	ID, _ := result.LastInsertId()

	return ID, nil
}

func (repo *articleRepositoryImpl) Update(ctx context.Context, ID int64, aritcle Article) error {
	return nil
}
