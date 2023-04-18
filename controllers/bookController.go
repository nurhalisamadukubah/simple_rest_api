package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookId int    `json:"bookId"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var bookDatas = []Book{}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return

	}
	newBook.BookId = len(bookDatas) + 1
	bookDatas = append(bookDatas, newBook)

	ctx.JSON(http.StatusCreated, gin.H{"message": "Created"})

}

func UpdateBook(ctx *gin.Context) {
	bookId := ctx.Param("BookId")
	isFound := false
	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	bookIDInt, _ := strconv.Atoi(bookId)
	for i, book := range bookDatas {
		if bookIDInt == book.BookId {
			isFound = true
			bookDatas[i] = updatedBook
			bookDatas[i].BookId = bookIDInt
			break
		}
	}

	if isFound == false {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookId),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been successfully updated", bookId),
	})
}

func GetBook(ctx *gin.Context) {
	bookId := ctx.Param("BookId")
	condition := false
	var BookData Book

	bookIDInt, _ := strconv.Atoi(bookId)
	for i, book := range bookDatas {
		if bookIDInt == book.BookId {
			condition = true
			BookData = bookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookId),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Book": BookData,
	})

}

func DeleteBook(ctx *gin.Context) {
	bookId := ctx.Param("BookId")
	condition := false
	var bookIndex int

	bookIDInt, _ := strconv.Atoi(bookId)
	for i, book := range bookDatas {
		if bookIDInt == book.BookId {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookId),
		})
		return
	}

	copy(bookDatas[bookIndex:], bookDatas[bookIndex+1:])
	bookDatas[len(bookDatas)-1] = Book{}
	bookDatas = bookDatas[:len(bookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been successfully deletd", bookId),
	})
}
