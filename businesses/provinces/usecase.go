package provinces

import (
	"context"
	"peduli-covid/app/middleware"
	"peduli-covid/businesses/rsbedcovid"
	"peduli-covid/helpers/messages"
	"time"
)

type provinceUsecase struct {
	provinceRepository   Repository
	rsbedcovidRepository rsbedcovid.Repository
	contextTimeout       time.Duration
	jwtAuth              *middleware.ConfigJWT
}

func NewProvinceUsecase(ur Repository, rsRepo rsbedcovid.Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &provinceUsecase{
		provinceRepository:   ur,
		rsbedcovidRepository: rsRepo,
		jwtAuth:              jwtauth,
		contextTimeout:       timeout,
	}
}

func (uc *provinceUsecase) FindAll(ctx context.Context) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.provinceRepository.FindAll(ctx)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *provinceUsecase) StoreFromAPI(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	provinceData, err := uc.rsbedcovidRepository.GetProvince(ctx)
	if err != nil {
		return messages.ErrNotFound
	}

	for _, data := range provinceData {
		existedUser, _ := uc.provinceRepository.GetByCode(ctx, data.ID)
		if existedUser != (Domain{}) {
			continue
		} else {
			province := &Domain{
				Code: data.ID,
				Name: data.Name,
			}
			err = uc.provinceRepository.Store(ctx, province)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
