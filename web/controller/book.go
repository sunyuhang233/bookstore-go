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

func (b *BookController) GetHotBooks(c *gin.Context) {
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

func (b *BookController) GetBooks(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	books, total, err := b.BookService.GetBooks(page, pageSize)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    -1,
			"message": "获取书籍列表失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":     0,
		"message":  "获取书籍列表成功",
		"data":     books,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

func (b *BookController) SearchBooks(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(400, gin.H{
			"code":    -1,
			"message": "参数错误",
		})
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	books, total, err := b.BookService.SearchBooks(page, pageSize, query)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    -1,
			"message": "获取书籍列表失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":     0,
		"message":  "获取书籍列表成功",
		"data":     books,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

func (b *BookController) GetBookDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := b.BookService.GetBookDetail(uint(id))
	if err != nil {
		c.JSON(500, gin.H{
			"code":    -1,
			"message": "获取书籍详情失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "获取书籍详情成功",
		"data":    book,
	})
}
