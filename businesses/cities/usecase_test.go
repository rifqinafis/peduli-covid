package cities_test

import (
	"context"
	"errors"
	"os"
	"peduli-covid/businesses/cities"
	city_mock "peduli-covid/businesses/cities/mocks"
	"peduli-covid/businesses/rsbedcovid"
	rsbedcovid_mock "peduli-covid/businesses/rsbedcovid/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	cityRepository       city_mock.Repository
	rsbedcovidRepository rsbedcovid_mock.Repository
	cityUsecase          cities.Usecase
)

func setup() {
	cityUsecase = cities.NewCityUsecase(&cityRepository, &rsbedcovidRepository, 2)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindAll(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		cityDomain := cities.Domain{
			ID:           1101,
			ProvinceCode: "11prop",
			Code:         "Simeulue",
			Name:         "Simeulue",
		}
		cityRepository.On("FindAll", mock.Anything).Return([]cities.Domain{cityDomain}, nil).Once()

		result, err := cityUsecase.FindAll(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, []cities.Domain{cityDomain}, result)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		cityRepository.On("FindAll", mock.Anything).Return([]cities.Domain{}, errors.New("error")).Once()

		result, err := cityUsecase.FindAll(context.Background())

		assert.Equal(t, []cities.Domain{}, result)
		assert.Equal(t, errors.New("error"), err)
	})
}

func TestFindByProvinceCode(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		cityDomain := cities.Domain{
			ID:           1101,
			ProvinceCode: "11prop",
			Code:         "Simeulue",
			Name:         "Simeulue",
		}
		cityRepository.On("FindByProvinceCode", mock.Anything, mock.AnythingOfType("string")).Return([]cities.Domain{cityDomain}, nil).Once()

		result, err := cityUsecase.FindByProvinceCode(context.Background(), "11prop")

		assert.Nil(t, err)
		assert.Equal(t, []cities.Domain{cityDomain}, result)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		cityRepository.On("FindByProvinceCode", mock.Anything, mock.AnythingOfType("string")).Return([]cities.Domain{}, errors.New("error")).Once()

		result, err := cityUsecase.FindByProvinceCode(context.Background(), "11prop")

		assert.Equal(t, []cities.Domain{}, result)
		assert.Equal(t, errors.New("error"), err)
	})
}

func TestStoreFromAPI(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		rsbedProvince := rsbedcovid.ProvinceDomain{
			ID:   "11prop",
			Name: "Aceh",
		}
		rsbedcovidRepository.On("GetProvince", mock.Anything).Return([]rsbedcovid.ProvinceDomain{rsbedProvince}, nil).Once()
		rsbedCity := rsbedcovid.CityDomain{
			ID:   "1101",
			Name: "Simeulue",
		}
		rsbedcovidRepository.On("GetCity", mock.Anything, mock.AnythingOfType("string")).Return([]rsbedcovid.CityDomain{rsbedCity}, nil).Once()
		cityDomain := cities.Domain{
			ID:           1101,
			ProvinceCode: "11prop",
			Code:         "Simeulue",
			Name:         "Simeulue",
		}
		cityRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(cityDomain, nil).Once()
		cityRepository.On("Store", mock.Anything, cityDomain).Return(nil).Once()

		err := cityUsecase.StoreFromAPI(context.Background())

		assert.Equal(t, nil, err)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		rsbedcovidRepository.On("GetProvince", mock.Anything).Return([]rsbedcovid.ProvinceDomain{}, errors.New("error")).Once()
		rsbedcovidRepository.On("GetCity", mock.Anything, mock.AnythingOfType("string")).Return([]rsbedcovid.CityDomain{}, errors.New("error")).Once()
		cityRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(cities.Domain{}, errors.New("error")).Once()
		cityRepository.On("Store", mock.Anything, cities.Domain{}).Return(errors.New("error")).Once()

		err := cityUsecase.StoreFromAPI(context.Background())

		assert.Equal(t, errors.New("data not found"), err)
	})
}
