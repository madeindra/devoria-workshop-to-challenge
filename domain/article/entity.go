package article

import (
	"time"

	"github.com/madeindra/devoria-workshop-to-challenge/domain/account"
)

// Enum for article status
type ArticleStatus string

const (
	ArticleStatusDraft     ArticleStatus = "DRAFT"
	ArticleStatusPublished ArticleStatus = "PUBLISHED"
	ArticleStatusArchived  ArticleStatus = "ARCHIVED"
)

// properties of account
// json attributes will set the field name on json form
type Article struct {
	ID             int64           `json:"id"`
	Title          string          `json:"title"`
	Subtitle       string          `json:"subtitle"`
	Content        string          `json:"content"`
	Status         ArticleStatus   `json:"status"`
	CreatedAt      time.Time       `json:"createdAt"`
	PublishedAt    *time.Time      `json:"publishedAt"`
	LastModifiedAt *time.Time      `json:"lastModifiedAt"`
	Author         account.Account `json:"author"`
}
