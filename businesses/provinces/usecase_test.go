package provinces_test

import (
	"context"
	"errors"
	"os"
	"peduli-covid/businesses/provinces"
	province_mock "peduli-covid/businesses/provinces/mocks"
	"peduli-covid/businesses/rsbedcovid"
	rsbedcovid_mock "peduli-covid/businesses/rsbedcovid/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	provinceRepository   province_mock.Repository
	rsbedcovidRepository rsbedcovid_mock.Repository
	provinceUsecase      provinces.Usecase
)

func setup() {
	provinceUsecase = provinces.NewProvinceUsecase(&provinceRepository, &rsbedcovidRepository, 2)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindAll(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		provinceDomain := provinces.Domain{
			ID:   1,
			Code: "11prop",
			Name: "Aceh",
		}
		provinceRepository.On("FindAll", mock.Anything).Return([]provinces.Domain{provinceDomain}, nil).Once()

		result, err := provinceUsecase.FindAll(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, []provinces.Domain{provinceDomain}, result)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		provinceRepository.On("FindAll", mock.Anything).Return([]provinces.Domain{}, errors.New("error")).Once()

		result, err := provinceUsecase.FindAll(context.Background())

		assert.Equal(t, []provinces.Domain{}, result)
		assert.Equal(t, errors.New("error"), err)
	})
}

func TestStoreFromAPI(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		rsbedcovidProvince := rsbedcovid.ProvinceDomain{
			ID:   "11prop",
			Name: "Aceh",
		}
		rsbedcovidRepository.On("GetProvince", mock.Anything).Return([]rsbedcovid.ProvinceDomain{rsbedcovidProvince}, nil).Once()
		provinceDomain := provinces.Domain{
			ID:   1,
			Code: "11prop",
			Name: "Aceh",
		}
		provinceRepository.On("GetByCode", mock.Anything, mock.AnythingOfType("string")).Return(provinceDomain, nil).Once()
		provinceRepository.On("Store", mock.Anything, mock.Anything).Return(nil).Once()

		err := provinceUsecase.StoreFromAPI(context.Background())

		assert.Equal(t, nil, err)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		rsbedcovidRepository.On("GetProvince", mock.Anything).Return([]rsbedcovid.ProvinceDomain{}, errors.New("error")).Once()
		provinceRepository.On("GetByCode", mock.Anything, mock.AnythingOfType("string")).Return(provinces.Domain{}, errors.New("error")).Once()
		provinceRepository.On("Store", mock.Anything, mock.Anything).Return(errors.New("error")).Once()

		err := provinceUsecase.StoreFromAPI(context.Background())

		assert.Equal(t, errors.New("data not found"), err)
	})
}
