package cities

import (
	"context"
	"peduli-covid/businesses/rsbedcovid"
	"peduli-covid/helpers/messages"
	"strconv"
	"time"
)

type cityUsecase struct {
	cityRepository       Repository
	rsbedcovidRepository rsbedcovid.Repository
	contextTimeout       time.Duration
}

func NewCityUsecase(ur Repository, rsRepo rsbedcovid.Repository, timeout time.Duration) Usecase {
	return &cityUsecase{
		cityRepository:       ur,
		rsbedcovidRepository: rsRepo,
		contextTimeout:       timeout,
	}
}

func (uc *cityUsecase) FindAll(ctx context.Context) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.cityRepository.FindAll(ctx)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *cityUsecase) FindByProvinceCode(ctx context.Context, provinceCode string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.cityRepository.FindByProvinceCode(ctx, provinceCode)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *cityUsecase) StoreFromAPI(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	provinceData, err := uc.rsbedcovidRepository.GetProvince(ctx)
	if err != nil {
		return messages.ErrNotFound
	}

	for _, dataProv := range provinceData {
		CityData, err := uc.rsbedcovidRepository.GetCity(ctx, dataProv.ID)
		if err != nil {
			continue
		}

		for _, data := range CityData {
			id, _ := strconv.Atoi(data.ID)
			existedUser, _ := uc.cityRepository.GetByID(ctx, id)
			if existedUser != (Domain{}) {
				continue
			} else {
				City := &Domain{
					ID:           id,
					ProvinceCode: dataProv.ID,
					Code:         data.Name,
					Name:         data.Name,
				}
				err = uc.cityRepository.Store(ctx, City)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
