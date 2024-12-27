package models

type NEOResponse struct {
	ElementCount int                       `json:"element_count"`
	NearObjects  map[string][]NEODayObject `json:"near_earth_objects"` // Date mapped to objects
}

type NEODayObject struct {
	ID                string  `json:"id"`
	Name              string  `json:"name"`
	Magnitude         float64 `json:"absolute_magnitude_h"`
	Hazardous         bool    `json:"is_potentially_hazardous_asteroid"`
	CloseApproachData []struct {
		CloseApproachDate string `json:"close_approach_date"`
		MissDistance      struct {
			Kilometers string `json:"kilometers"`
		} `json:"miss_distance"`
	} `json:"close_approach_data"`
}
