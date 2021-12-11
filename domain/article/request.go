package article

// Request body for account registration
type CreateArticleRequest struct {
	Title    string `json:"title" validate:"required"`
	Subtitle string `json:"subtitle" validate:"required"`
	Content  string `json:"content" validate:"required"`
}

// Request body for account login
type UpdateArticleRequest struct {
	ID       int64         `json:"id"`
	Title    string        `json:"title" validate:"required"`
	Subtitle string        `json:"subtitle" validate:"required"`
	Content  string        `json:"content" validate:"required"`
	Status   ArticleStatus `json:"status"`
}
