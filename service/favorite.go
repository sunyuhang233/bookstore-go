package service

import (
	"bookstore-go/model"
	"bookstore-go/repository"
)

type FavoriteService struct {
	FavoriteDAO *repository.FavoriteDAO
}

func NewRepositoryService() *FavoriteService {
	return &FavoriteService{
		FavoriteDAO: repository.NewFavoriteDAO(),
	}
}

func (s *FavoriteService) AddFavorite(userId, bookId int) error {
	return s.FavoriteDAO.AddFavorite(userId, bookId)
}

func (s *FavoriteService) RemoveFavorite(userId, bookId int) error {
	return s.FavoriteDAO.RemoveFavorite(userId, bookId)
}

func (s *FavoriteService) GetFavorites(userId, page, pageSize int) ([]*model.Favorite, int, error) {
	favorites, total, err := s.FavoriteDAO.GetFavorites(userId, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return favorites, int(total), nil
}
