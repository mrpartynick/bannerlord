package pgprovider

import (
	"bannerlord/internal/api/models"
	"context"
)

func (p *pgProvider) CheckUser(login string) (bool, error) {
	var result *bool
	_, err := p.db.QueryContext(
		context.Background(),
		result,
		CheckUser,
		login,
	)
	if err != nil {
		return false, err
	}
	return *result, nil
}

func (p *pgProvider) CheckAdmin(login string) (bool, error) {
	var result *bool
	_, err := p.db.QueryContext(
		context.Background(),
		result,
		CheckUser,
		login,
	)
	if err != nil {
		return false, err
	}
	return *result, nil
}

func (p *pgProvider) CreateUser(login string, password string) error {
	_, err := p.db.QueryContext(
		context.Background(),
		CheckUser,
		login,
		password,
	)
	if err != nil {
		return err
	}
	return nil
}

func (p *pgProvider) AuthUser(login string, password string) (bool, error) {
	var result *bool
	_, err := p.db.QueryContext(
		context.Background(),
		result,
		AuthUser,
		login,
		password,
	)
	if err != nil {
		return false, err
	}
	return *result, nil
}

func (p *pgProvider) AuthAdmin(login string, password string) (bool, error) {
	var result *bool
	_, err := p.db.QueryContext(
		context.Background(),
		result,
		AuthAdmin,
		login,
		password,
	)
	if err != nil {
		return false, err
	}
	return *result, nil
}

func (p *pgProvider) GetBanners(feature int, tag int, isActive bool) ([]models.Banner, error) {
	query := GetBannersForTagAndFeature
	var result []models.Banner
	if isActive {
		query += "and is_active='true'"
	}
	_, err := p.db.Query(&result, query, feature, tag)
	return result, err
}

func (p *pgProvider) GetAdminBanners(feature int, tag int) ([]models.Banner, error) {
	query := GetBannersForTagAndFeature
	var result []models.Banner
	_, err := p.db.Query(&result, query, feature, tag)
	return result, err
}

func (p *pgProvider) GetByFeature(feature int) ([]models.Banner, error) {
	query := GetByFeature
	var result []models.Banner
	_, err := p.db.Query(&result, query, feature)
	return result, err
}

func (p *pgProvider) GetByTag(tag int) ([]models.Banner, error) {
	query := GetByTag
	var result []models.Banner
	_, err := p.db.Query(&result, query, tag)
	return result, err
}

func (p *pgProvider) GetAll() ([]models.Banner, error) {
	query := GetAll
	var result []models.Banner
	_, err := p.db.Query(&result, query)
	return result, err
}

func (p *pgProvider) UpdateBanner(banner *models.BannerPatch) error {
	t, err := p.db.Begin()
	defer t.Close()

	if err != nil {
		return err
	}

	_, err = t.Exec(UpdateBanner,
		banner.FeatureID != nil,
		*banner.FeatureID,
		banner.Content != nil,
		*banner.Content,
		banner.IsActive != nil,
		*banner.IsActive)
	if err != nil {
		return err
	}
	if banner.TagIDs != nil {
		for tag := range *banner.TagIDs {
			_, err = t.Exec(InsertTag, banner.ID, tag)
			if err != nil {
				return err
			}
		}
	}

	if err = t.Commit(); err != nil {
		return err
	}
	return nil
}

func (p *pgProvider) DeleteBanner(id int) error {
	_, err := p.db.Exec(DeleteBanner, id)
	return err
}
