package usecase

import (
	"github.com/loak155/microservices/services/article/domain"
	"github.com/loak155/microservices/services/article/repository"
	"github.com/loak155/microservices/services/article/validator"
)

type IArticleUsecase interface {
	CreateArticle(article domain.Article) (domain.Article, error)
	GetArticle(id int) (domain.Article, error)
	ListArticles() ([]domain.Article, error)
	UpdateArticle(article domain.Article) (bool, error)
	DeleteArticle(id int) (bool, error)
}

type articleUsecase struct {
	ar repository.IArticleRepository
	av validator.IArticleValidator
}

func NewArticleUsecase(ar repository.IArticleRepository, av validator.IArticleValidator) IArticleUsecase {
	return &articleUsecase{ar, av}
}

func (uu *articleUsecase) CreateArticle(article domain.Article) (domain.Article, error) {
	if err := uu.av.ArticleValidate(article); err != nil {
		return domain.Article{}, err
	}
	article.BookmarkCount = 0
	if err := uu.ar.CreateArticle(&article); err != nil {
		return domain.Article{}, err
	}
	return article, nil
}

func (uu *articleUsecase) GetArticle(id int) (domain.Article, error) {
	storedArticle := domain.Article{}
	if err := uu.ar.GetArticle(&storedArticle, id); err != nil {
		return domain.Article{}, err
	}
	return storedArticle, nil
}

func (uu *articleUsecase) ListArticles() ([]domain.Article, error) {
	storedArticles := []domain.Article{}
	if err := uu.ar.ListArticles(&storedArticles); err != nil {
		return []domain.Article{}, err
	}
	return storedArticles, nil
}

func (uu *articleUsecase) UpdateArticle(article domain.Article) (bool, error) {
	if err := uu.av.ArticleValidate(article); err != nil {
		return false, err
	}
	updatedArticle := domain.Article{
		ID:            article.ID,
		Title:         article.Title,
		Url:           article.Url,
		BookmarkCount: article.BookmarkCount,
	}
	if err := uu.ar.UpdateArticle(&updatedArticle); err != nil {
		return false, err
	}
	return true, nil
}

func (uu *articleUsecase) DeleteArticle(id int) (bool, error) {
	if err := uu.ar.DeleteArticle(id); err != nil {
		return false, err
	}
	return true, nil
}
