package model

import "time"

type Book struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Author      string    `json:"author"`
	Price       int       `json:"price"`       // 价格（元）
	Discount    int       `json:"discount"`    // 折扣（百分比，100表示无折扣）
	Type        string    `json:"type"`        // 图书类型
	Stock       int       `json:"stock"`       // 库存数量
	Status      int       `json:"status"`      // 图书状态：0-下架，1-上架
	Description string    `json:"description"` // 图书描述
	CoverURL    string    `json:"cover_url"`
	ISBN        string    `json:"isbn"`         // ISBN号
	Publisher   string    `json:"publisher"`    // 出版社
	PublishDate string    `json:"publish_date"` // 出版日期
	Pages       int       `json:"pages"`        // 页数
	Language    string    `json:"language"`     // 语言
	Format      string    `json:"format"`       // 装帧格式
	CategoryID  uint      `json:"category_id"`  // 分类ID
	Sale        int       `json:"sale"`         // 销售量
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (b *Book) TableName() string {
	return "books"
}

// BookCreateRequest 创建图书请求
type BookCreateRequest struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Price       int    `json:"price" binding:"required,min=0"`
	Discount    int    `json:"discount" binding:"required,min=0,max=100"`
	Type        string `json:"type" binding:"required,min=1"`
	Stock       int    `json:"stock" binding:"required,min=0"`
	Status      int    `json:"status" binding:"min=0,max=1"`
	CoverURL    string `json:"cover_url"`
	Description string `json:"description"`
	ISBN        string `json:"isbn"`
	Publisher   string `json:"publisher"`
	PublishDate string `json:"publish_date"`
	Pages       int    `json:"pages"`
	Language    string `json:"language"`
	Format      string `json:"format"`
	CategoryID  uint   `json:"category_id"`
	Sale        int    `json:"sale" binding:"min=0"` // 销售量
}

// BookUpdateRequest 更新图书请求
type BookUpdateRequest struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Price       int    `json:"price" binding:"min=0"`
	Discount    int    `json:"discount" binding:"min=0,max=100"`
	Type        string `json:"type"`
	Stock       int    `json:"stock" binding:"min=0"`
	CoverURL    string `json:"cover_url"`
	Description string `json:"description"`
	ISBN        string `json:"isbn"`
	Publisher   string `json:"publisher"`
	PublishDate string `json:"publish_date"`
	Pages       int    `json:"pages"`
	Language    string `json:"language"`
	Format      string `json:"format"`
	Status      int    `json:"status" binding:"min=0,max=1"`
	CategoryID  uint   `json:"category_id"`
	Sale        int    `json:"sale" binding:"min=0"` // 销售量
}

// BookListRequest 图书列表请求
type BookListRequest struct {
	Page       int    `form:"page" binding:"min=1"`
	PageSize   int    `form:"page_size" binding:"min=1,max=100"`
	Title      string `form:"title"`
	Author     string `form:"author"`
	Type       string `form:"type"`
	Status     *int   `form:"status"` // nil表示不过滤状态
	CategoryID uint   `form:"category_id"`
}

// BookListResponse 图书列表响应
type BookListResponse struct {
	Books       []Book `json:"books"`
	Total       int64  `json:"total"`
	TotalPage   int    `json:"total_page"`
	CurrentPage int    `json:"current_page"`
}
