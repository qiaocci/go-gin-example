package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {

	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	fmt.Printf("==========pageNum=%v, pageSize=%v, maps=%#v\n", pageNum, pageSize, maps)
	fmt.Println(tags)
	return
}

func AddMyTag() {
	db.Create(&Tag{
		Name:       "c",
		CreatedBy:  "qiaocc",
		ModifiedBy: "qiaocc",
		State:      0,
	})
	fmt.Println("create success")
}

func GetTagTotal(maps interface{}) (count int64) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}
