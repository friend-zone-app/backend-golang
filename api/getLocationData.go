package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"parties-app/backend/config"
	"parties-app/backend/types"
)

type Map struct {
	HttpClient *http.Client
}

var (
	MapURL     string = "https://maps.googleapis.com/maps/api/staticmap?center=%v,%v8&zoom=%v&size=600x600&maptype=%v&key=" + config.GOOGLE_MAP_API
	AddressURL string = "http://dev.virtualearth.net/REST/v1/Locations/%q?maxResults=1&key=" + config.BING_MAP_API
	PointURL   string = "http://dev.virtualearth.net/REST/v1/Locations/%q,%q?maxResults=1&key=" + config.BING_MAP_API
)

func NewMap() Map {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return Map{
		HttpClient: &client,
	}
}

func (handler Map) Get(lat float64, long float64, zoom int64, maptype string) (*bytes.Reader, error) {
	url := fmt.Sprintf(MapURL, lat, long, zoom, maptype)
	resp, err := handler.HttpClient.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return bytes.NewReader(body), nil
}

func (handler Map) GetAddress(query string) (*types.BingRes, error) {
	url := fmt.Sprintf(AddressURL, query)
	resp, err := handler.HttpClient.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	res := types.BingRes{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (handler Map) GetPoint(lat string, long string) (*types.BingRes, error) {
	url := fmt.Sprintf(PointURL, lat, long)

	fmt.Println(url)

	resp, err := handler.HttpClient.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	res := types.BingRes{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (handler Map) GetMapImage(query string, zoom int64, maptype string) (*bytes.Reader, error) {

	var point = []float64{0.0, 0.0}

	res, err := handler.GetAddress(query)
	if err != nil {
		return nil, err
	}
	if len(res.ResourceSets) > 0 {
		if len(res.ResourceSets[0].Resources) > 0 {
			point = res.ResourceSets[0].Resources[0].Point.Coordinates
		}
	}

	if point[0] == 0.0 || point[1] == 0.0 {
		return nil, errors.New("the point is invalid")
	}
	return handler.Get(point[0], point[1], zoom, maptype)
}
