package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetArticles(PageOffset int, pageSize int, maps interface{}) (articles []Article) {
	db.Where(maps).Offset(PageOffset).Limit(pageSize).Find(&articles)
	return
}

func GetArticleTotal(maps interface{}) (count int64) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticle(id int) (article Article) {
	db.Where("id=?", id).First(&article)
	return
}

func AddArticle(title, desc, content, createdBy string, tagID int) bool {
	db.Create(&Article{
		Title:     title,
		Desc:      desc,
		Content:   content,
		CreatedBy: createdBy,
		TagID:     tagID,
	})
	return true
}

func EditArticle(id int, data interface{}) bool {
	fmt.Printf("id=%v, data=%v\n", id, data)
	db.Model(&Article{}).Where("id=?", id).Updates(data)
	return true
}

func DeleteArticle(id int) bool {
	db.Delete(&Article{}, id)
	return true
}
