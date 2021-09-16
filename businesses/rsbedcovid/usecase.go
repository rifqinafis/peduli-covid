package rsbedcovid

import (
	"context"
	"time"
)

type rsbedcovidUsecase struct {
	rsbedcovidRepository Repository
	contextTimeout       time.Duration
}

func NewRSBedCovid(ur Repository, timeout time.Duration) Usecase {
	return &rsbedcovidUsecase{
		rsbedcovidRepository: ur,
		contextTimeout:       timeout,
	}
}

func (uc *rsbedcovidUsecase) GetProvince(ctx context.Context) ([]ProvinceDomain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.rsbedcovidRepository.GetProvince(ctx)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *rsbedcovidUsecase) GetCity(ctx context.Context, provinceID string) ([]CityDomain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.rsbedcovidRepository.GetCity(ctx, provinceID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *rsbedcovidUsecase) GetHospital(ctx context.Context, provinceID, cityID, types string) ([]HospitalDomain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.rsbedcovidRepository.GetHospital(ctx, provinceID, cityID, types)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *rsbedcovidUsecase) GetBedDetail(ctx context.Context, hospitalID, types string) ([]BedDetailDomain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.rsbedcovidRepository.GetBedDetail(ctx, hospitalID, types)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *rsbedcovidUsecase) GetHospitalLocation(ctx context.Context, hospitalID string) (HospitalLocationDomain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.rsbedcovidRepository.GetHospitalLocation(ctx, hospitalID)
	if err != nil {
		return res, err
	}

	return res, nil
}
