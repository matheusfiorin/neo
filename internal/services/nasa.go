package services

import (
	"encoding/json"
	"fmt"
	"go-rest-api/internal/models"
	"net/http"
	"os"
)

type NASAService struct {
	apiKey string
	client *http.Client
}

func NewNASAService() *NASAService {
	return &NASAService{
		apiKey: os.Getenv("NASA_API_KEY"),
		client: &http.Client{},
	}
}

func (s *NASAService) GetNEOs(startDate, endDate string) (*models.NEOResponse, error) {
	url := fmt.Sprintf(
		"https://api.nasa.gov/neo/rest/v1/feed?start_date=%s&end_date=%s&api_key=%s",
		startDate,
		endDate,
		s.apiKey,
	)

	resp, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result models.NEOResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
