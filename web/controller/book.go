package controller

import (
	"bookstore-go/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookService *service.BookService
}

func NewBookController() *BookController {
	return &BookController{
		BookService: service.NewBookService(),
	}
}

func (b *BookController) GetBooks(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	hotBooks, err := b.BookService.GetHotBooks(limit)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    -1,
			"message": "获取热门书籍失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "获取热门书籍成功",
		"data":    hotBooks,
	})
}

func (b *BookController) GetNewBooks(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	newBooks, err := b.BookService.GetNewBooks(limit)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    -1,
			"message": "获取最新书籍失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "获取最新书籍成功",
		"data":    newBooks,
	})
}
