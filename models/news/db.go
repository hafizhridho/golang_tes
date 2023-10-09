package news

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	ID        uint 			`gorm:"primarykey"`
    CreatedAt time.Time		`gorm:"autoCreateTime"`
    UpdatedAt time.Time		`gorm:"autoUpdateTime:milli"`
    DeletedAt gorm.DeletedAt 	`gorm:"index"`
	Title   string
	Author  string
	Content string
}

func (news News) MapFromRequestAdd(newsRequestAdd NewsRequestAdd) News{
	var newsDB News
	newsDB.Title = newsRequestAdd.Title
	newsDB.Author = newsRequestAdd.Author
	newsDB.Content = newsRequestAdd.Content
	return newsDB
}