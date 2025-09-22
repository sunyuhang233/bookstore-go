package model

import "time"

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Price       int       `json:"price"`
	Discount    int       `json:"discount"`
	Type        string    `json:"type"`
	Stock       int       `json:"stock"`
	Description string    `json:"description"`
	CoverURL    string    `json:"cover_url"`
	ISBN        string    `json:"isbn"`
	Publisher   string    `json:"publisher"`
	Pages       int       `json:"pages"`
	Language    string    `json:"language"`
	Format      string    `json:"format"`
	CategoryID  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
