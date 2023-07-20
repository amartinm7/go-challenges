package repository

import (
	"challenges/challenge_1/internal/domain/ad"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/huandu/go-sqlbuilder"
)

const (
	ID_COLUMN          = "id"
	TITLE_COLUMN       = "title"
	DESCRIPTION_COLUMN = "description"
	PRICE_COLUMN       = "price"
	TIME_STAMP_COLUMN  = "time_stamp"
	ADS_TABLE_         = "ADS"
)

type adJpa struct {
	Id          string `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Price       int    `db:"price"`
	Timestamp   string `db:"time_stamp"`
}

type adRepository struct {
	dbConnection *sql.DB
}

func NewAdRepository(dbConnection *sql.DB) ad.AdRepository {
	return &adRepository{
		dbConnection: dbConnection,
	}
}

func (adRepository *adRepository) Save(ad ad.Ad) (ad.Ad, error) {
	builder := sqlbuilder.PostgreSQL.NewInsertBuilder()
	builder.InsertInto(ADS_TABLE_)
	builder.Cols(ID_COLUMN, TITLE_COLUMN, DESCRIPTION_COLUMN, PRICE_COLUMN, TIME_STAMP_COLUMN)
	builder.Values(ad.Id.String(), ad.Title, ad.Description.Value, ad.Price, ad.Timestamp.Format("2006-01-02"))
	sql, args := builder.Build()
	_, err := adRepository.dbConnection.ExecContext(context.Background(), sql, args...)
	if err != nil {
		return ad, fmt.Errorf("error trying to persist an AD on database: %v", err)
	}

	return ad, nil
}

func (adRepository *adRepository) FindById(id uuid.UUID) (*ad.Ad, error) {
	var adJpaStruct = sqlbuilder.NewStruct(adJpa{})

	builder := sqlbuilder.PostgreSQL.NewSelectBuilder()
	builder.Select(ID_COLUMN, TITLE_COLUMN, DESCRIPTION_COLUMN, PRICE_COLUMN, TIME_STAMP_COLUMN)
	builder.From(ADS_TABLE_)
	builder.Where(
		builder.Equal(ID_COLUMN, id.String()),
	)

	sql, args := builder.Build()
	row := adRepository.dbConnection.QueryRowContext(context.Background(), sql, args...)
	var foundAd adJpa
	err := row.Scan(adJpaStruct.Addr(&foundAd)...)
	if err != nil {
		return nil, fmt.Errorf("element not found %s:  %v", id.String(), err)
	}
	ad, err := mapAdJpaToAd(foundAd)
	if err != nil {
		return nil, fmt.Errorf("invalid format %s:  %v", id.String(), err)
	}
	return &ad, nil
}

func mapAdJpaToAd(foundAd adJpa) (ad.Ad, error) {
	return ad.NewAd(foundAd.Id, foundAd.Title, foundAd.Description, foundAd.Price, foundAd.Timestamp)
}

func (adRepository *adRepository) FindAll() (*[]ad.Ad, error) {
	var adJpaStruct = sqlbuilder.NewStruct(adJpa{})

	builder := sqlbuilder.PostgreSQL.NewSelectBuilder()
	builder.Select(ID_COLUMN, TITLE_COLUMN, DESCRIPTION_COLUMN, PRICE_COLUMN, TIME_STAMP_COLUMN)
	builder.From(ADS_TABLE_)

	sql, args := builder.Build()
	rows, err := adRepository.dbConnection.QueryContext(context.Background(), sql, args...)
	if err != nil {
		return nil, fmt.Errorf("element not found:  %v", err)
	}
	defer rows.Close()

	var ads []ad.Ad
	for rows.Next() {
		var foundAd = adJpa{}
		err := rows.Scan(adJpaStruct.Addr(&foundAd)...)
		if err != nil {
			return nil, fmt.Errorf("element not found:  %v", err)
		}
		ad, err := mapAdJpaToAd(foundAd)
		if err != nil {
			return nil, fmt.Errorf("invalid format:  %v", err)
		}
		ads = append(ads, ad)
	}
	return &ads, nil
}
