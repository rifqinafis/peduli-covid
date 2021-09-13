package cities

import (
	"context"
	"fmt"
	"peduli-covid/app/middleware"
	"peduli-covid/businesses/rsbedcovid"
	"peduli-covid/helpers/messages"
	"strconv"
	"strings"
	"time"
)

type CityUsecase struct {
	CityRepository       Repository
	rsbedcovidRepository rsbedcovid.Repository
	contextTimeout       time.Duration
	jwtAuth              *middleware.ConfigJWT
}

func NewCityUsecase(ur Repository, rsRepo rsbedcovid.Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &CityUsecase{
		CityRepository:       ur,
		rsbedcovidRepository: rsRepo,
		jwtAuth:              jwtauth,
		contextTimeout:       timeout,
	}
}

func (uc *CityUsecase) Store(ctx context.Context, CityDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedUser, err := uc.CityRepository.GetByID(ctx, CityDomain.ID)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
	}
	if existedUser != (Domain{}) {
		return messages.ErrDuplicateData
	}

	err = uc.CityRepository.Store(ctx, CityDomain)
	if err != nil {
		return err
	}

	return nil
}

func (uc *CityUsecase) StoreFromAPI(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	provinceData, err := uc.rsbedcovidRepository.GetProvince(ctx)
	fmt.Println(provinceData)
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
			existedUser, _ := uc.CityRepository.GetByID(ctx, id)
			if existedUser != (Domain{}) {
				continue
			} else {
				City := &Domain{
					ID:           id,
					ProvinceCode: dataProv.ID,
					Code:         data.Name,
					Name:         data.Name,
				}
				err = uc.CityRepository.Store(ctx, City)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
