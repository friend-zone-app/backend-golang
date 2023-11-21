package types

import "parties-app/backend/graph/customTypes"

type BingRes struct {
	AuthenticationResultCode string `json:"authenticationResultCode"`
	BrandLogoURI             string `json:"brandLogoUri"`
	Copyright                string `json:"copyright"`
	ResourceSets             []struct {
		EstimatedTotal int `json:"estimatedTotal"`
		Resources      []struct {
			Type  string    `json:"__type"`
			Bbox  []float64 `json:"bbox"`
			Name  string    `json:"name"`
			Point struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			} `json:"point"`
			Address struct {
				AddressLine       string                   `json:"addressLine"`
				AdminDistrict     string                   `json:"adminDistrict"`
				AdminDistrict2    string                   `json:"adminDistrict2"`
				CountryRegion     string                   `json:"countryRegion"`
				Interception      customTypes.Interception `json:"interception"`
				FormattedAddress  string                   `json:"formattedAddress"`
				Locality          string                   `json:"locality"`
				PostalCode        string                   `json:"postalCode"`
				CountryRegionIso2 string                   `json:"countryRegionIso2"`
			} `json:"address"`
			Confidence    string `json:"confidence"`
			EntityType    string `json:"entityType"`
			GeocodePoints []struct {
				Type              string    `json:"type"`
				Coordinates       []float64 `json:"coordinates"`
				CalculationMethod string    `json:"calculationMethod"`
				UsageTypes        []string  `json:"usageTypes"`
			} `json:"geocodePoints"`
			MatchCodes []string `json:"matchCodes"`
		} `json:"resources"`
	} `json:"resourceSets"`
	StatusCode        int    `json:"statusCode"`
	StatusDescription string `json:"statusDescription"`
	TraceID           string `json:"traceId"`
}
