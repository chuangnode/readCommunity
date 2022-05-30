package model

import (
	"fmt"
	"gorm.io/gorm"
	"readCommunity/global"
)

type Book struct {
	gorm.Model
	BookName      string `gorm:"column:book_name" json:"book_name"`
	Identify      string `gorm:"column:identify" json:"identify"`
	OrderIndex    int    `gorm:"column:order_index" json:"order_index"`
	Description   string `gorm:"column:description" json:"description"`
	Cover         string `gorm:"column:cover" json:"cover"`
	Editor        string `gorm:"column:editor" json:"editor"`
	Status        uint8  `gorm:"column:status" json:"status"`
	PrivateOwn    int    `gorm:"column:private_own" json:"private_own"`
	UserId        int    `gorm:"column:user_id" json:"user_id"`
	DocCount      int    `gorm:"column:doc_count" json:"doc___count"`
	CommentCount  int    `gorm:"column:comment_count" json:"comment_count"`
	ReadCount     int    `gorm:"column:read_count" json:"read_count"`
	Star          int    `gorm:"column:star" json:"star"`
	Score         int    `gorm:"column:score" json:"score"`
	ScoreCount    int    `gorm:"column:score_count" json:"score_count"`
	CommentUcount int    `gorm:"column:comment_ucount" json:"comment_ucount"`
	Author        string `gorm:"author" json:"author"`
	AuthorUrl     string `gorm:"author_url" json:"author_url"`
}

type BookParams struct {
	BookName    string `validate:"required" json:"book_name"`
	Identify    string `validate:"required" json:"identify"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

func (b Book) TableName() string {
	return "rd_books"
}

func (b Book) AddBook(params BookParams) (err error) {
	err = global.DB.Create(&Book{
		BookName:    params.BookName,
		Description: params.Description,
		Identify:    params.Identify,
		Author:      params.Author,
	}).Error
	return err
}

func (b Book) ExistedBookById(id int) (bool, error) {
	err := global.DB.First(&b, "id = ?", id).Error
	if err != nil {
		fmt.Printf("existed book by id failed, err;%v\n", err)
		return false, err
	}
	if b.ID > 0 {
		return true, nil
	}
	return false, nil
}

func (b Book) UpdateBook(id int, params BookParams) (err error) {
	db := global.DB
	err = db.Model(&b).Where("id=?", id).Updates(Book{BookName: params.BookName, Identify: params.Identify,
		Description: params.Description, Author: params.Author}).Error
	return
}

func (b Book) DeleteBook(id int, data interface{}) error {
	db := global.DB
	if data.(bool) {
		return db.Where("id=?", id).Unscoped().Delete(&b).Error
	} else {
		return db.Where("id=?", id).Delete(&b).Error
	}

}

func (b Book) GetBookList() (books []Book, err error) {
	db := global.DB
	result := db.Find(&books)
	return books, result.Error
}

func (b Book) GetBookInfo(id int) (book Book, err error) {
	db := global.DB
	return book, db.First(&book, id).Error
}
