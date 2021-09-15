package bedtypes

import (
	"context"
	"peduli-covid/app/middleware"
	"peduli-covid/businesses/hospitals"
	"peduli-covid/businesses/rsbedcovid"
	"peduli-covid/helpers/messages"
	"strconv"
	"time"
)

type bedtypeUsecase struct {
	bedtypeRepository    Repository
	hospitalRepository   hospitals.Repository
	rsbedcovidRepository rsbedcovid.Repository
	contextTimeout       time.Duration
	jwtAuth              *middleware.ConfigJWT
}

func NewBedtypeUsecase(ur Repository, cityRepo hospitals.Repository, rsRepo rsbedcovid.Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &bedtypeUsecase{
		bedtypeRepository:    ur,
		hospitalRepository:   cityRepo,
		rsbedcovidRepository: rsRepo,
		jwtAuth:              jwtauth,
		contextTimeout:       timeout,
	}
}

func (uc *bedtypeUsecase) FindByHospitalID(ctx context.Context, hospitalID int) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.bedtypeRepository.FindByHospitalID(ctx, hospitalID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *bedtypeUsecase) StoreFromAPI(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	hospitalData, err := uc.hospitalRepository.FindAll(ctx)
	if err != nil {
		return messages.ErrNotFound
	}

	for _, dataHospital := range hospitalData {
		bedtypeData, err := uc.rsbedcovidRepository.GetBedDetail(ctx, strconv.Itoa(dataHospital.ID), "1")
		if err != nil {
			continue
		}

		for _, data := range bedtypeData {
			existedUser, _ := uc.bedtypeRepository.FindByTitleAndHospitalID(ctx, data.Title, dataHospital.ID)
			if existedUser != (Domain{}) {
				continue
			} else {
				bedtype := &Domain{
					HospitalID:   dataHospital.ID,
					Name:         data.Title,
					BedAvailable: data.BedAvailable,
					BedEmpty:     data.BedEmpty,
				}
				err = uc.bedtypeRepository.Store(ctx, bedtype)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
