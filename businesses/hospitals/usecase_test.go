package hospitals_test

import (
	"context"
	"errors"
	"os"
	"peduli-covid/businesses/cities"
	city_mock "peduli-covid/businesses/cities/mocks"
	"peduli-covid/businesses/hospitals"
	hospital_mock "peduli-covid/businesses/hospitals/mocks"
	"peduli-covid/businesses/rsbedcovid"
	rsbedcovid_mock "peduli-covid/businesses/rsbedcovid/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	hospitalRepository   hospital_mock.Repository
	cityRepository       city_mock.Repository
	rsbedcovidRepository rsbedcovid_mock.Repository
	hospitalUsecase      hospitals.Usecase
)

func setup() {
	hospitalUsecase = hospitals.NewHospitalUsecase(&hospitalRepository, &cityRepository, &rsbedcovidRepository, 2)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindAll(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		hospitalDomain := hospitals.Domain{
			ID:      1101015,
			CityID:  1101,
			Name:    "RS Umum Daerah Simeulue",
			Address: "Jl. Teuku Raja Mahmud Desa Amiria Bahagia Kecamatan Simeulue Timur",
			Phone:   "082365706161",
		}
		hospitalRepository.On("FindAll", mock.Anything).Return([]hospitals.Domain{hospitalDomain}, nil).Once()

		result, err := hospitalUsecase.FindAll(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, []hospitals.Domain{hospitalDomain}, result)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		hospitalRepository.On("FindAll", mock.Anything).Return([]hospitals.Domain{}, errors.New("error")).Once()

		result, err := hospitalUsecase.FindAll(context.Background())

		assert.Equal(t, []hospitals.Domain{}, result)
		assert.Equal(t, errors.New("error"), err)
	})
}

func TestFindByCityID(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		hospitalDomain := hospitals.Domain{
			ID:      1101015,
			CityID:  1101,
			Name:    "RS Umum Daerah Simeulue",
			Address: "Jl. Teuku Raja Mahmud Desa Amiria Bahagia Kecamatan Simeulue Timur",
			Phone:   "082365706161",
		}
		hospitalRepository.On("FindByCityID", mock.Anything, mock.AnythingOfType("int")).Return([]hospitals.Domain{hospitalDomain}, nil).Once()

		result, err := hospitalUsecase.FindByCityID(context.Background(), 1101)

		assert.Nil(t, err)
		assert.Equal(t, []hospitals.Domain{hospitalDomain}, result)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		hospitalRepository.On("FindByCityID", mock.Anything, mock.AnythingOfType("int")).Return([]hospitals.Domain{}, errors.New("error")).Once()

		result, err := hospitalUsecase.FindByCityID(context.Background(), 1101)

		assert.Equal(t, []hospitals.Domain{}, result)
		assert.Equal(t, errors.New("error"), err)
	})
}

func TestStoreFromAPI(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		cityDomain := cities.Domain{
			ID:           1101,
			ProvinceCode: "11prop",
			Code:         "Simeulue",
			Name:         "Simeulue",
		}
		cityRepository.On("FindAll", mock.Anything).Return([]cities.Domain{cityDomain}, nil).Once()
		rsbedHospital := rsbedcovid.HospitalDomain{
			ID:              "1101015",
			Name:            "RS Umum Daerah Simeulue",
			Address:         "Jl. Teuku Raja Mahmud Desa Amiria Bahagia Kecamatan Simeulue Timur",
			Phone:           "082365706161",
			Queue:           0,
			BedAvailability: 0,
			Info:            "Updated",
		}
		rsbedcovidRepository.On("GetHospital", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]rsbedcovid.HospitalDomain{rsbedHospital}, nil).Once()
		hospitalDomain := hospitals.Domain{
			ID:      1101015,
			CityID:  1101,
			Name:    "RS Umum Daerah Simeulue",
			Address: "Jl. Teuku Raja Mahmud Desa Amiria Bahagia Kecamatan Simeulue Timur",
			Phone:   "082365706161",
		}
		hospitalRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(hospitalDomain, nil).Once()
		hospitalRepository.On("Store", mock.Anything, hospitalDomain).Return(nil).Once()

		err := hospitalUsecase.StoreFromAPI(context.Background())

		assert.Equal(t, nil, err)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		cityRepository.On("FindAll", mock.Anything).Return([]cities.Domain{}, errors.New("error")).Once()
		rsbedcovidRepository.On("GetHospital", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return().Once()
		hospitalRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(hospitals.Domain{}, errors.New("error")).Once()
		hospitalRepository.On("Store", mock.Anything, hospitals.Domain{}).Return(errors.New("error")).Once()

		err := hospitalUsecase.StoreFromAPI(context.Background())

		assert.Equal(t, errors.New("data not found"), err)
	})
}
