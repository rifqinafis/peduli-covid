package rsbedcovid

import (
	"context"
	"encoding/json"
	"net/http"
	"peduli-covid/businesses/rsbedcovid"
)

type RSBedCovid struct {
	httpClient http.Client
}

func NewRSBedCovid() rsbedcovid.Repository {
	return &RSBedCovid{
		httpClient: http.Client{},
	}
}

func (rs *RSBedCovid) GetProvince(ctx context.Context) (res []rsbedcovid.ProvinceDomain, err error) {
	req, _ := http.NewRequest("GET", "https://rs-bed-covid-api.vercel.app/api/get-provinces", nil)
	resp, err := rs.httpClient.Do(req)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	data := ProvinceRespon{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return res, err
	}

	for _, value := range data.Province {
		temp := rsbedcovid.ProvinceDomain{
			ID:   value.ID,
			Name: value.Name,
		}
		res = append(res, temp)
	}

	return res, nil
}

func (rs *RSBedCovid) GetCity(ctx context.Context, provinceID string) (res []rsbedcovid.CityDomain, err error) {
	req, _ := http.NewRequest("GET", "https://rs-bed-covid-api.vercel.app/api/get-cities", nil)
	q := req.URL.Query()
	q.Add("provinceid", provinceID)
	req.URL.RawQuery = q.Encode()

	resp, err := rs.httpClient.Do(req)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	data := CityRespon{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return res, err
	}
	for _, value := range data.City {
		temp := rsbedcovid.CityDomain{
			ID:   value.ID,
			Name: value.Name,
		}
		res = append(res, temp)
	}

	return res, nil
}

func (rs *RSBedCovid) GetHospital(ctx context.Context, provinceID, cityID, types string) (res []rsbedcovid.HospitalDomain, err error) {
	req, _ := http.NewRequest("GET", "https://rs-bed-covid-api.vercel.app/api/get-hospitals", nil)
	q := req.URL.Query()
	q.Add("provinceid", provinceID)
	q.Add("cityid", cityID)
	q.Add("type", types)
	req.URL.RawQuery = q.Encode()

	resp, err := rs.httpClient.Do(req)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	data := HospitalRespon{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return res, err
	}
	for _, value := range data.Hospital {
		temp := rsbedcovid.HospitalDomain{
			ID:              value.ID,
			Name:            value.Name,
			Address:         value.Address,
			Phone:           value.Phone,
			Queue:           value.Queue,
			BedAvailability: value.BedAvailability,
			Info:            value.Info,
		}
		res = append(res, temp)
	}

	return res, nil
}

func (rs *RSBedCovid) GetBedDetail(ctx context.Context, hospitalID, types string) (res []rsbedcovid.BedDetailDomain, err error) {
	req, _ := http.NewRequest("GET", "https://rs-bed-covid-api.vercel.app/api/get-bed-detail", nil)
	q := req.URL.Query()
	q.Add("hospitalid", hospitalID)
	q.Add("type", types)
	req.URL.RawQuery = q.Encode()

	resp, err := rs.httpClient.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	data := BedDetailRespon{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return res, err
	}

	for _, value := range data.Data.BedDetail {
		temp := rsbedcovid.BedDetailDomain{
			Title:        value.Stats.Title,
			BedAvailable: value.Stats.BedAvailable,
			BedEmpty:     value.Stats.BedEmpty,
			Queue:        value.Stats.Queue,
		}
		res = append(res, temp)
	}

	return res, nil
}

func (rs *RSBedCovid) GetHospitalLocation(ctx context.Context, hospitalID string) (res rsbedcovid.HospitalLocationDomain, err error) {
	req, _ := http.NewRequest("GET", "https://rs-bed-covid-api.vercel.app/api/get-hospital-map", nil)
	q := req.URL.Query()
	q.Add("hospitalid", hospitalID)
	req.URL.RawQuery = q.Encode()

	resp, err := rs.httpClient.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	data := HospitalLocationRespon{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return res, err
	}

	res = rsbedcovid.HospitalLocationDomain{
		ID:      data.Data.ID,
		Name:    data.Data.Name,
		Address: data.Data.Address,
		Lat:     data.Data.Lat,
		Long:    data.Data.Long,
		Gmaps:   data.Data.Gmaps,
	}

	return res, nil
}
