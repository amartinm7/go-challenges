package ad

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Ad struct {
	Id          uuid.UUID
	Title       string
	Description Description
	Price       int
	Timestamp   time.Time
}

var InvalidStringToUUID = errors.New("Invalid String to UUID")
var InvalidDescription = errors.New("Description can be more than 50 chars")

var InvalidTimeStamp = errors.New("Timestamp malformated")

type Description struct {
	Value string `validate:"len=50" message:"Description can be more than 50 chars"`
}

func NewDescription(value string) (Description, error) {
	description := Description{Value: value}
	err := description.validate()
	if err != nil {
		return Description{}, err
	}
	return description, nil
}

func (description *Description) validate() error {
	if len(description.Value) >= 50 {
		return InvalidDescription
	}
	return nil
}

func NewAd(id, title, description string, price int, timestamp string) (Ad, error) {
	uuidVO, err := uuid.Parse(id)
	if err != nil {
		return Ad{}, InvalidStringToUUID
	}
	timestampVO, err := time.Parse("2006-01-02", timestamp)
	if err != nil {
		return Ad{}, InvalidTimeStamp
	}
	descVO, err := NewDescription(description)
	if err != nil {
		return Ad{}, err
	}
	return Ad{
		Id:          uuidVO,
		Title:       title,
		Description: descVO,
		Price:       price,
		Timestamp:   timestampVO,
	}, nil
}
