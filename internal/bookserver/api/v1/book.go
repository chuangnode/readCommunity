package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"readCommunity/internal/bookserver/model"
	"readCommunity/internal/bookserver/service"
	"readCommunity/internal/pkg/app"
	"readCommunity/internal/pkg/errcode"
	"strconv"
)
// @Summary  创建书籍
// @Produce  json
// @Param    book_name  body      string  true  "book_name"
// @Param    description  body      string  true  "description"
// @Param    identify  body      string  true  "identify"
// @Param    author  body      string  true  "author"
// @Success  200       {string}  json    "{"code": 200,"data":{},"msg":"ok"}"
//@Router    /api/v1/book [post]
func AddBook(c *gin.Context) {
	var bookParam model.BookParams
	if err := c.ShouldBindJSON(&bookParam); err != nil {
		fmt.Printf("addbook shouldBindJson failed, err: %s", err.Error())
		app.ResponseInfo(c, errcode.ErrBind, nil)
		return
	}
	err := service.ReleaseBook(bookParam)
	if err != nil {
		app.ResponseInfo(c, errcode.ErrUnknown, nil)
	} else {
		app.ResponseInfo(c, errcode.Success, nil)
	}
}
// @Summary  修改书籍信息
// @Produce  json
// @Param    ID  param   int  true  "id"
// @Param    book_name  body      string  true  "book_name"
// @Param    description  body      string  true  "description"
// @Param    identify  body      string  true  "identify"
// @Param    author  body      string  true  "author"
// @Success  200       {string}  json    "{"code": 200,"data":{},"msg":"ok"}"
//@Router    /api/v1/book/:id [put]
func EditBook(c *gin.Context)  {
	var bookParams model.BookParams
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("id param convert to int failed,err: %v\n", err)
		app.ResponseInfo(c, errcode.ErrDecodingFailed, nil)
	}
	if err = c.ShouldBindJSON(&bookParams); err != nil {
		fmt.Printf("editBook shoudbind failed,err:%v\n", err)
		app.ResponseInfo(c, errcode.ErrBind, nil)
	}
	err = service.EditBook(id, bookParams)
	if err != nil {
		fmt.Printf("edit book failed, err:%v\n", err)
		app.ResponseInfo(c, errcode.ErrUnknown, nil)
	}
	app.ResponseInfo(c, errcode.Success, nil)

}

// @Summary  删除书籍信息
// @Produce  json
// @Param    ID  param   int  true  "id"
// @Success  200       {string}  json    "{"code": 200,"data":{},"msg":"ok"}"
//@Router    /api/v1/book/:id [delete]
func DeleteBook(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		app.ResponseInfo(c, errcode.ErrDecodingFailed, nil)
		return
	}
	err = service.DeleteBook(id)
	if err != nil {
		fmt.Printf("delete book failed,err: %v\n", err)
		app.ResponseInfo(c, errcode.ErrUnknown, nil)
	} else {
		app.ResponseInfo(c, errcode.Success, nil)
	}
}
// BookList
// @Summary  删除书籍信息
// @Produce  json
// @Success  200       {string}  json    "{"code": 200,"data":{},"msg":"ok"}"
// @Router    /api/v1/book [get]
func BookList(c *gin.Context) {
	data, err := service.BookList()
	fmt.Printf("data list: %#v, type: %T\n", data)
	if err != nil {
		zap.L().Error("get book list failed")
		fmt.Println("book list err:", err)
		app.ResponseInfo(c, errcode.ErrBookGetFailed, nil)
	} else {
		zap.L().Debug("get book list success")
		app.ResponseInfo(c, errcode.Success, data)
	}
}

//BookInfo: 获取单本书籍信息
func BookInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error("book id convert to int error")
		app.ResponseInfo(c, errcode.ErrIdInvalid, nil)
		return
	}
	data, err := service.GetBookInfo(id)
	if err != nil {
		zap.L().Error("get book info failed")
		app.ResponseInfo(c, errcode.ErrBookGetFailed, nil)
	} else {
		app.ResponseInfo(c, errcode.Success, data)
	}
}