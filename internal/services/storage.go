package services

import "bannerlord/internal/api/models"

type userStorage interface {
	CheckUser(login string) (bool, error)
	CheckAdmin(login string) (bool, error)
	CreateUser(login string, password string) error
	AuthUser(login string, password string) (bool, error)
	AuthAdmin(login string, password string) (bool, error)
}

type bannerStorage interface {
	GetBanners(feature int, tag int, isActive bool) ([]models.Banner, error)
	GetAdminBanners(feature int, tag int) ([]models.Banner, error)
	GetByFeature(feature int) ([]models.Banner, error)
	GetByTag(tag int) ([]models.Banner, error)
	GetAll() ([]models.Banner, error)
	UpdateBanner(banner *models.BannerPatch) error
	DeleteBanner(id int) error
}

type Storage interface {
	userStorage
	bannerStorage
	Connect() error
}
