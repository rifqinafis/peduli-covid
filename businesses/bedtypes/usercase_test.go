package bedtypes_test

import (
	"context"
	"errors"
	"os"
	"peduli-covid/businesses/bedtypes"
	bedtype_mock "peduli-covid/businesses/bedtypes/mocks"
	"peduli-covid/businesses/hospitals"
	hospital_mock "peduli-covid/businesses/hospitals/mocks"
	"peduli-covid/businesses/rsbedcovid"
	rsbedcovid_mock "peduli-covid/businesses/rsbedcovid/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	bedtypeRepository    bedtype_mock.Repository
	hospitalRepository   hospital_mock.Repository
	rsbedcovidRepository rsbedcovid_mock.Repository
	bedtypeUsecase       bedtypes.Usecase
)

func setup() {
	bedtypeUsecase = bedtypes.NewBedtypeUsecase(&bedtypeRepository, &hospitalRepository, &rsbedcovidRepository, 2)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindByHospitalID(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		bedtypeDomain := bedtypes.Domain{
			ID:           1,
			HospitalID:   1101015,
			Name:         "IGD Khusus Covid",
			BedAvailable: 0,
			BedEmpty:     0,
		}
		bedtypeRepository.On("FindByHospitalID", mock.Anything, mock.AnythingOfType("int")).Return([]bedtypes.Domain{bedtypeDomain}, nil).Once()

		result, err := bedtypeUsecase.FindByHospitalID(context.Background(), 1101015)

		assert.Nil(t, err)
		assert.Equal(t, []bedtypes.Domain{bedtypeDomain}, result)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		bedtypeRepository.On("FindByHospitalID", mock.Anything, mock.AnythingOfType("int")).Return([]bedtypes.Domain{}, errors.New("error")).Once()

		result, err := bedtypeUsecase.FindByHospitalID(context.Background(), 1)

		assert.Equal(t, []bedtypes.Domain{}, result)
		assert.Equal(t, errors.New("error"), err)
	})
}

func TestStoreFromAPI(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		hospitalDomain := hospitals.Domain{
			ID:      1101015,
			CityID:  1101,
			Name:    "RS Umum Daerah Simeulue",
			Address: "Jl. Teuku Raja Mahmud Desa Amiria Bahagia Kecamatan Simeulue Timur",
			Phone:   "082365706161",
		}
		hospitalRepository.On("FindAll", mock.Anything).Return([]hospitals.Domain{hospitalDomain}, nil).Once()
		rsbedDomain := rsbedcovid.BedDetailDomain{
			Title:        "IGD Khusus Covid",
			BedAvailable: 0,
			BedEmpty:     0,
			Queue:        0,
		}
		rsbedcovidRepository.On("GetBedDetail", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]rsbedcovid.BedDetailDomain{rsbedDomain}, nil).Once()
		bedtypeDomain := bedtypes.Domain{
			ID:           1,
			HospitalID:   1101015,
			Name:         "IGD Khusus Covid",
			BedAvailable: 0,
			BedEmpty:     0,
		}
		bedtypeRepository.On("FindByTitleAndHospitalID", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(bedtypeDomain, nil).Once()
		bedtypeRepository.On("Store", mock.Anything, bedtypeDomain).Return(nil).Once()

		err := bedtypeUsecase.StoreFromAPI(context.Background())

		assert.Equal(t, nil, err)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		hospitalRepository.On("FindAll", mock.Anything).Return([]hospitals.Domain{}, errors.New("error")).Once()
		rsbedcovidRepository.On("GetBedDetail", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]rsbedcovid.BedDetailDomain{}, errors.New("error")).Once()
		bedtypeRepository.On("FindByTitleAndHospitalID", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(rsbedcovid.BedDetailDomain{}, errors.New("error")).Once()
		bedtypeRepository.On("Store", mock.Anything, rsbedcovid.BedDetailDomain{}).Return(errors.New("error")).Once()

		err := bedtypeUsecase.StoreFromAPI(context.Background())

		assert.Equal(t, errors.New("data not found"), err)
	})
}
