package service

import (
	"errors"
	"readCommunity/internal/bookserver/model"
)



// 新增书籍
func ReleaseBook(bookParam model.BookParams) error {
	book := model.Book{}
	return book.AddBook(bookParam)
}

//编辑书籍
func EditBook(id int, params model.BookParams) error {
	var book model.Book
	isExist, err := book.ExistedBookById(id)
	if err != nil {
		return errors.New("查询书籍发生错误")
	}
	if !isExist {
		return errors.New("该书籍不存在")
	}
	err = book.UpdateBook(id, params)
	return err
}

// 删除书籍
func DeleteBook(id int) error {
	var book model.Book
	return book.DeleteBook(id, true)
}

// 获取书籍列表
func BookList() (data []model.Book, err error) {
	var book model.Book
	data, err = book.GetBookList()
	return
}

func GetBookInfo(id int) (data model.Book, err error) {
	var book model.Book
	data, err = book.GetBookInfo(id)
	return
}