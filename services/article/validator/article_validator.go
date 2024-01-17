package validator

import (
	"github.com/loak155/microservices/services/article/domain"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IArticleValidator interface {
	ArticleValidate(article domain.Article) error
}

type articleValidator struct{}

func NewArticleValidator() IArticleValidator {
	return &articleValidator{}
}

func (av *articleValidator) ArticleValidate(article domain.Article) error {
	return validation.ValidateStruct(&article,
		validation.Field(
			&article.Title,
			validation.Required.Error("title is required"),
		),
		validation.Field(
			&article.Url,
			validation.Required.Error("url is required"),
			is.URL.Error("is not valid url format"),
		),
	)
}
