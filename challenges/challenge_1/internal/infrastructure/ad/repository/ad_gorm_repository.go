package repository

import (
	"challenges/challenge_1/internal/domain/ad"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type adGorm struct {
	Id          string `gorm:"column:id"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	Price       int    `gorm:"column:price"`
	Timestamp   string `gorm:"column:time_stamp"`
}

func (adGorm) TableName() string {
	return "ads"
}

type AdGormRepository struct {
	db *gorm.DB
}

func NewAdGormRepository(gormDB *gorm.DB) ad.AdRepository {
	return &AdGormRepository{
		db: gormDB,
	}
}

func (repo *AdGormRepository) Save(ad ad.Ad) (ad.Ad, error) {
	_ad := mapAdToAdGorm(ad)
	err := repo.db.Create(_ad).Error
	if err != nil {
		return ad, err
	}
	return ad, nil
}

func (repo *AdGormRepository) FindById(id uuid.UUID) (*ad.Ad, error) {
	var _ad adGorm
	err := repo.db.First(&_ad, id).Error
	if err != nil {
		return nil, err
	}
	ad, err := mapAdGormToAds(_ad)
	if err != nil {
		return nil, err
	}
	return ad, nil
}

func (repo *AdGormRepository) FindAll() (*[]ad.Ad, error) {
	var _ads []adGorm
	err := repo.db.Find(&_ads).Error
	if err != nil {
		return nil, err
	}
	if len(_ads) == 0 {
		fmt.Println("Not found ads...")
		return &[]ad.Ad{}, nil
	}
	ads, err := mapAdsGormToAds(_ads)
	if err != nil {
		return nil, err
	}
	return &ads, nil
}

func mapAdGormToAds(foundAd adGorm) (*ad.Ad, error) {
	ad, err := ad.NewAd(foundAd.Id, foundAd.Title, foundAd.Description, foundAd.Price, foundAd.Timestamp)
	if err != nil {
		return nil, err
	}
	return &ad, nil
}

func mapAdsGormToAds(foundAds []adGorm) ([]ad.Ad, error) {
	var ads []ad.Ad
	for _, foundAd := range foundAds {
		ad, err := mapAdGormToAds(foundAd)
		if err != nil {
			return nil, err
		}
		ads = append(ads, *ad)
	}
	return ads, nil
}

func mapAdToAdGorm(ad ad.Ad) adGorm {
	return adGorm{
		Id:          ad.Id.String(),
		Title:       ad.Title,
		Description: ad.Description.Value,
		Price:       ad.Price,
		Timestamp:   ad.Timestamp.Format("2006-01-02"),
	}
}
