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
		Update(ctx context.Context, ID int64, article Article) error
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
	articles := []Article{}

	query := fmt.Sprintf(`SELECT %s.id, title, subtitle, content, status, %s.createdAt, publishedAt, %s.lastModifiedAt, authorId, email, firstName, lastName, accounts.createdAt, %s.lastModifiedAt  FROM %s JOIN %s ON %s.authorId = %s.id`, repo.tableName, repo.tableName, repo.tableName, constant.TableAccount, repo.tableName, constant.TableAccount, repo.tableName, constant.TableAccount)

	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return articles, exception.ErrInternalServer
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Println(err)
		return articles, exception.ErrInternalServer
	}
	defer rows.Close()

	for rows.Next() {
		article := Article{}

		var publishedAt sql.NullTime
		var lastModifiedAt sql.NullTime

		err = rows.Scan(
			&article.ID,
			&article.Title,
			&article.Subtitle,
			&article.Content,
			&article.Status,
			&article.CreatedAt,
			&publishedAt,
			&lastModifiedAt,
			&article.Author.ID,
			&article.Author.Email,
			&article.Author.FirstName,
			&article.Author.LastName,
			&article.Author.CreatedAt,
			&article.Author.LastModifiedAt,
		)

		if err != nil {
			log.Println(err)
			return articles, exception.ErrNotFound
		}

		if publishedAt.Valid {
			article.PublishedAt = &publishedAt.Time
		}

		if lastModifiedAt.Valid {
			article.LastModifiedAt = &lastModifiedAt.Time
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func (repo *articleRepositoryImpl) FindAllByAuthorId(ctx context.Context, authorId int64) ([]Article, error) {
	articles := []Article{}

	query := fmt.Sprintf(`SELECT id, authorId, title, subtitle, content, status, createdAt, lastModifiedAt, firstName, lastName, email FROM %s WHERE authorId = ?`, repo.tableName)

	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return articles, exception.ErrInternalServer
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Println(err)
		return articles, exception.ErrInternalServer
	}
	defer rows.Close()

	for rows.Next() {
		article := Article{}

		var publishedAt sql.NullTime
		var lastModifiedAt sql.NullTime

		err = rows.Scan(
			&article.ID,
			&article.Title,
			&article.Subtitle,
			&article.Content,
			&article.Status,
			&article.CreatedAt,
			&publishedAt,
			&lastModifiedAt,
			&article.Author.ID,
		)

		if err != nil {
			log.Println(err)
			return articles, exception.ErrNotFound
		}

		if publishedAt.Valid {
			article.PublishedAt = &publishedAt.Time
		}

		if lastModifiedAt.Valid {
			article.LastModifiedAt = &lastModifiedAt.Time
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func (repo *articleRepositoryImpl) Create(ctx context.Context, article Article) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, subtitle, content, status, createdAt, publishedAt, authorId) VALUES (?, ?, ?, ?, ?, ?, ?)", repo.tableName)

	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		article.Title,
		article.Subtitle,
		article.Content,
		article.Status,
		article.CreatedAt,
		article.PublishedAt,
		article.Author.ID,
	)

	if err != nil {
		return 0, err
	}

	ID, _ := result.LastInsertId()

	return ID, nil
}

func (repo *articleRepositoryImpl) Update(ctx context.Context, ID int64, article Article) error {
	query := fmt.Sprintf("UPDATE %s SET title = ?, subtitle = ?, content = ?, status=?, lastModifiedAt = ? WHERE id = ?", repo.tableName)

	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return exception.ErrInternalServer
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		article.Title,
		article.Subtitle,
		article.Content,
		article.Status,
		*article.LastModifiedAt,
		ID,
	)

	if err != nil {
		log.Println(err)
		return exception.ErrInternalServer
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected < 1 {
		return exception.ErrNotFound
	}

	return nil
}
