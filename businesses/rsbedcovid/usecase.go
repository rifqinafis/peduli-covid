package rsbedcovid

import (
	"context"
	"peduli-covid/app/middleware"
	"time"
)

type rsbedcovidUseCase struct {
	rsbedcovidRepository Repository
	contextTimeout       time.Duration
	jwtAuth              *middleware.ConfigJWT
}

func NewRSBedCovid(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &rsbedcovidUseCase{
		rsbedcovidRepository: ur,
		jwtAuth:              jwtauth,
		contextTimeout:       timeout,
	}
}

func (uc *rsbedcovidUseCase) GetProvince(ctx context.Context) ([]ProvinceDomain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.rsbedcovidRepository.GetProvince(ctx)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *rsbedcovidUseCase) GetCity(ctx context.Context, provinceID string) ([]CityDomain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.rsbedcovidRepository.GetCity(ctx, provinceID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *rsbedcovidUseCase) GetHospital(ctx context.Context, provinceID, cityID, types string) ([]HospitalDomain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.rsbedcovidRepository.GetHospital(ctx, provinceID, cityID, types)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *rsbedcovidUseCase) GetBedDetail(ctx context.Context, hospitalID, types string) ([]BedDetailDomain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.rsbedcovidRepository.GetBedDetail(ctx, hospitalID, types)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *rsbedcovidUseCase) GetHospitalLocation(ctx context.Context, hospitalID string) (HospitalLocationDomain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.rsbedcovidRepository.GetHospitalLocation(ctx, hospitalID)
	if err != nil {
		return res, err
	}

	return res, nil
}
