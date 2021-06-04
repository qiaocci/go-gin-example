package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model

	// belong to关系
	TagID int `json:"tag_id" gorm:"index"` // 外键
	Tag   Tag `json:"tag"`                 // 通过Tag关联查询

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id=?", id).First(&article)
	if article.ID > 0 {
		return true
	} else {
		return false
	}
}

func GetArticles(PageOffset int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(PageOffset).Limit(pageSize).Find(&articles)
	return
}

func GetArticleTotal(maps interface{}) (count int64) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticle(id int) (article Article) {
	db.Preload("Tag").Where("id=?", id).First(&article)
	return
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["CreatedBy"].(string),
		State:     data["State"].(int),
		TagID:     data["TagID"].(int),
	})
	return true
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id=?", id).Updates(data)
	return true
}

func DeleteArticle(id int) bool {
	db.Delete(&Article{}, id)
	return true
}
