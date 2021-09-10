package rsbedcovid

import "context"

type ProvinceDomain struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CityDomain struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type HospitalDomain struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Address         string `json:"address"`
	Phone           string `json:"phone"`
	Queue           int    `json:"queue"`
	BedAvailability int    `json:"bed_availability"`
	Info            string `json:"info"`
}

type BedDetailDomain struct {
	Title        string `json:"title"`
	BedAvailable int    `json:"bed_available"`
	BedEmpty     int    `json:"bed_empty"`
	Queue        int    `json:"queue"`
}

type HospitalLocationDomain struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Lat     string `json:"lat"`
	Long    string `json:"long"`
	Gmaps   string `json:"gmaps"`
}

type Usecase interface {
	GetProvince(ctx context.Context) ([]ProvinceDomain, error)
	GetCity(ctx context.Context, provinceID string) ([]CityDomain, error)
	GetHospital(ctx context.Context, provinceID, cityID, types string) ([]HospitalDomain, error)
	GetBedDetail(ctx context.Context, hospitalID, types string) ([]BedDetailDomain, error)
	GetHospitalLocation(ctx context.Context, hospitalID string) (HospitalLocationDomain, error)
}

type Repository interface {
	GetProvince(ctx context.Context) ([]ProvinceDomain, error)
	GetCity(ctx context.Context, provinceID string) ([]CityDomain, error)
	GetHospital(ctx context.Context, provinceID, cityID, types string) ([]HospitalDomain, error)
	GetBedDetail(ctx context.Context, hospitalID, types string) ([]BedDetailDomain, error)
	GetHospitalLocation(ctx context.Context, hospitalID string) (HospitalLocationDomain, error)
}
