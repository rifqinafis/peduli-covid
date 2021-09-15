package hospitals

import (
	"context"
	"peduli-covid/app/middleware"
	"peduli-covid/businesses/cities"
	"peduli-covid/businesses/rsbedcovid"
	"peduli-covid/helpers/messages"
	"strconv"
	"time"
)

type hospitalUsecase struct {
	hospitalRepository   Repository
	cityRepository       cities.Repository
	rsbedcovidRepository rsbedcovid.Repository
	contextTimeout       time.Duration
	jwtAuth              *middleware.ConfigJWT
}

func NewHospitalUsecase(ur Repository, cityRepo cities.Repository, rsRepo rsbedcovid.Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &hospitalUsecase{
		hospitalRepository:   ur,
		cityRepository:       cityRepo,
		rsbedcovidRepository: rsRepo,
		jwtAuth:              jwtauth,
		contextTimeout:       timeout,
	}
}

func (uc *hospitalUsecase) FindAll(ctx context.Context) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.hospitalRepository.FindAll(ctx)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *hospitalUsecase) FindByCityID(ctx context.Context, cityID int) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.hospitalRepository.FindByCityID(ctx, cityID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *hospitalUsecase) StoreFromAPI(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	cityData, err := uc.cityRepository.FindAll(ctx)
	if err != nil {
		return messages.ErrNotFound
	}

	for _, dataCity := range cityData {
		hospitalData, err := uc.rsbedcovidRepository.GetHospital(ctx, dataCity.ProvinceCode, strconv.Itoa(dataCity.ID), "1")
		if err != nil {
			continue
		}

		for _, data := range hospitalData {
			id, _ := strconv.Atoi(data.ID)
			existedUser, _ := uc.hospitalRepository.GetByID(ctx, id)
			if existedUser != (Domain{}) {
				continue
			} else {
				hospital := &Domain{
					ID:      id,
					CityID:  dataCity.ID,
					Name:    data.Name,
					Address: data.Address,
					Phone:   data.Phone,
				}
				err = uc.hospitalRepository.Store(ctx, hospital)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
