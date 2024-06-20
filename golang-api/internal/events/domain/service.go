package domain

import (
	"errors"
	"fmt"
)

type spotService struct{}

var (
	ErrInvalidQuantity = errors.New("event capacity must be greater than zero")
)

func NewSpotService() *spotService {
	return &spotService{}
}

func (s *spotService) GenerateSpots(event *Event, quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	for i := range quantity {
		spotName := fmt.Sprintf("%c%d", 'A'+i/10, i+i%10)
		spot, err := NewSpot(event, spotName)

		if err != nil {
			return err
		}

		event.Spots = append(event.Spots, *spot)
	}

	return nil
}
