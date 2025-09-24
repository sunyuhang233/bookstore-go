package service

import (
	"bookstore-go/model"
	"bookstore-go/repository"
)

type BookService struct {
	BookDAO *repository.BookDAO
}

func NewBookService() *BookService {
	return &BookService{
		BookDAO: repository.NewBookDAO(),
	}
}

func (s *BookService) GetHotBooks(limit int) ([]model.Book, error) {
	books, err := s.BookDAO.GetHotBooks(limit)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *BookService) GetNewBooks(limit int) ([]model.Book, error) {
	books, err := s.BookDAO.GetNewBooks(limit)
	if err != nil {
		return nil, err
	}
	return books, nil
}
