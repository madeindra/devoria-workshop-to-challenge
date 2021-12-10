package account

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/madeindra/devoria-workshop-to-challenge/internal/exception"
	"golang.org/x/net/context"
)

type (
	AccountRepository interface {
		FindByID(ctx context.Context, ID int64) (Account, error)
		FindByEmail(ctx context.Context, email string) (Account, error)
		Create(ctx context.Context, account Account) (int64, error)
		Update(ctx context.Context, ID int64, account Account) error
	}

	accountRepositoryImpl struct {
		db        *sql.DB
		tableName string
	}
)

func NewAccountRepository(db *sql.DB, tableName string) AccountRepository {
	return &accountRepositoryImpl{
		db:        db,
		tableName: tableName,
	}
}

// Select Account by ID (PK)
func (repo *accountRepositoryImpl) FindByID(ctx context.Context, ID int64) (Account, error) {
	account := Account{}

	query := fmt.Sprintf(`SELECT id, email, password, firstName, lastName, createdAt, lastModifiedAt FROM %s WHERE id = ?`, repo.tableName)
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return account, exception.ErrInternalServer
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, ID)

	var password sql.NullString
	var lastModifiedAt sql.NullTime

	err = row.Scan(
		&account.ID,
		&account.Email,
		&password,
		&account.FirstName,
		&account.LastName,
		&account.CreatedAt,
		&lastModifiedAt,
	)

	if err != nil {
		log.Println(err)
		return account, exception.ErrNotFound
	}

	if password.Valid {
		account.Password = &password.String
	}

	if lastModifiedAt.Valid {
		account.LastModifiedAt = &lastModifiedAt.Time
	}

	return account, nil
}

// Select Account by Email
func (repo *accountRepositoryImpl) FindByEmail(ctx context.Context, email string) (Account, error) {
	account := Account{}

	query := fmt.Sprintf(`SELECT id, email, password, firstName, lastName, createdAt, lastModifiedAt FROM %s WHERE email = ?`, repo.tableName)
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return account, exception.ErrInternalServer
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, email)

	var password sql.NullString
	var lastModifiedAt sql.NullTime

	err = row.Scan(
		&account.ID,
		&account.Email,
		&password,
		&account.FirstName,
		&account.LastName,
		&account.CreatedAt,
		&lastModifiedAt,
	)

	if err != nil {
		log.Println(err)
		return account, exception.ErrNotFound
	}

	if password.Valid {
		account.Password = &password.String
	}

	if lastModifiedAt.Valid {
		account.LastModifiedAt = &lastModifiedAt.Time
	}

	return account, nil
}

// Insert into Account
func (repo *accountRepositoryImpl) Create(ctx context.Context, account Account) (int64, error) {
	command := fmt.Sprintf("INSERT INTO %s (email, password, firstName, lastName, createdAt) VALUES (?, ?, ?, ?, ?)", repo.tableName)
	stmt, err := repo.db.PrepareContext(ctx, command)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		account.Email,
		*account.Password,
		account.FirstName,
		account.LastName,
		account.CreatedAt,
	)

	if err != nil {
		return 0, err
	}

	ID, _ := result.LastInsertId()

	return ID, nil
}

// Update Account
func (repo *accountRepositoryImpl) Update(ctx context.Context, ID int64, account Account) error {
	command := fmt.Sprintf(`UPDATE %s SET password = ?, firstName = ?, lastName = ?, lastModifiedAt = ? WHERE id = ?`, repo.tableName)
	stmt, err := repo.db.PrepareContext(ctx, command)
	if err != nil {
		log.Println(err)
		return exception.ErrInternalServer
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		*account.Password,
		account.FirstName,
		account.LastName,
		account.LastModifiedAt,
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
