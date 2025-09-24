package controller

import (
	"bookstore-go/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavoriteController struct {
	FavoriteService *service.FavoriteService
}

func NewFavoriteController() *FavoriteController {
	return &FavoriteController{
		FavoriteService: service.NewRepositoryService(),
	}
}

func (f *FavoriteController) AddFavorite(c *gin.Context) {
	userId, exists := c.Get("admin_user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    -1,
			"message": "未登录",
		})
		return
	}
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "参数错误",
			"error":   err.Error(),
		})
		return
	}

	err = f.FavoriteService.FavoriteDAO.AddFavorite(userId.(int), bookId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "添加收藏失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "添加收藏成功",
	})
}

func (f *FavoriteController) RemoveFavorite(c *gin.Context) {
	userId, exists := c.Get("admin_user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    -1,
			"message": "未登录",
		})
		return
	}
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "参数错误",
			"error":   err.Error(),
		})
		return
	}

	err = f.FavoriteService.FavoriteDAO.RemoveFavorite(userId.(int), bookId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "移除收藏失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "移除收藏成功",
	})
}

func (f *FavoriteController) GetFavorites(c *gin.Context) {
	userId, exists := c.Get("admin_user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    -1,
			"message": "未登录",
		})
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	favorites, total, err := f.FavoriteService.GetFavorites(userId.(int), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "获取收藏列表失败",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     0,
		"message":  "获取收藏列表成功",
		"data":     favorites,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}
