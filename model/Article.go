package model

import (
	"userservice/database"
	"userservice/schema"
)

type Article schema.Article
type UpdateArticle schema.UpdateArticle

func (entry *Article) Save() (*Article, error) {
	err := database.Database.Create(&entry).Error
	return entry, err
}

func (update_data *Article) ChangeData(id string, ua UpdateArticle) (Article, error) {
	err := database.Database.Model(Article{}).Where("id = ?", id).Updates(ua).Error
	if err != nil {
		return *update_data, err
	}
	res, _ := ArticleFindById(id)
	return res, nil
}

func ArticleFindById(id string) (Article, error) {
	var ar Article
	err := database.Database.Where("id = ?", id).First(&ar).Error
	return ar, err
}
