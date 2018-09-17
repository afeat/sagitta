package model

type Airport struct {
	Code string `json:"airport_code"`
	Name string `json:"name"`
	Iata string `json:"alternate_ident"`
	CountryCode string `json:"country_code"`
}

type AirportDoc struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Iata string `json:"iata"`
	CountryCode string `json:"countryCode"`
}

func (airport *Airport) ToAirportDoc() AirportDoc {
	return AirportDoc {
		airport.Code,
		airport.Name,
		airport.Iata,
		airport.CountryCode,
	}
}