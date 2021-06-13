package article_service

import "github.com/qiaocco/go-gin-example/models"

type Article struct {
	ID int
}

func (a Article) ExistArticleByID() (bool, error) {
	return models.ExistArticleByID(a.ID), nil
}

func (a Article) Get() (models.Article, error) {
	return models.GetArticle(a.ID), nil
}
