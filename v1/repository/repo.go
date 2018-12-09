package repository

import (
	articleEntity "github.com/alamin-mahamud/go-bootstrap/v1/entity/article"
)

// IArticleRepo abstracts the DB Layer with Implementation.
type IArticleRepo interface {
	List() ([]articleEntity.Article, error)
}
