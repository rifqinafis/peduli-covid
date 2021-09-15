// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	rsbedcovid "peduli-covid/businesses/rsbedcovid"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// GetBedDetail provides a mock function with given fields: ctx, hospitalID, types
func (_m *Usecase) GetBedDetail(ctx context.Context, hospitalID string, types string) ([]rsbedcovid.BedDetailDomain, error) {
	ret := _m.Called(ctx, hospitalID, types)

	var r0 []rsbedcovid.BedDetailDomain
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []rsbedcovid.BedDetailDomain); ok {
		r0 = rf(ctx, hospitalID, types)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]rsbedcovid.BedDetailDomain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, hospitalID, types)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCity provides a mock function with given fields: ctx, provinceID
func (_m *Usecase) GetCity(ctx context.Context, provinceID string) ([]rsbedcovid.CityDomain, error) {
	ret := _m.Called(ctx, provinceID)

	var r0 []rsbedcovid.CityDomain
	if rf, ok := ret.Get(0).(func(context.Context, string) []rsbedcovid.CityDomain); ok {
		r0 = rf(ctx, provinceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]rsbedcovid.CityDomain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, provinceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHospital provides a mock function with given fields: ctx, provinceID, cityID, types
func (_m *Usecase) GetHospital(ctx context.Context, provinceID string, cityID string, types string) ([]rsbedcovid.HospitalDomain, error) {
	ret := _m.Called(ctx, provinceID, cityID, types)

	var r0 []rsbedcovid.HospitalDomain
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) []rsbedcovid.HospitalDomain); ok {
		r0 = rf(ctx, provinceID, cityID, types)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]rsbedcovid.HospitalDomain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, provinceID, cityID, types)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHospitalLocation provides a mock function with given fields: ctx, hospitalID
func (_m *Usecase) GetHospitalLocation(ctx context.Context, hospitalID string) (rsbedcovid.HospitalLocationDomain, error) {
	ret := _m.Called(ctx, hospitalID)

	var r0 rsbedcovid.HospitalLocationDomain
	if rf, ok := ret.Get(0).(func(context.Context, string) rsbedcovid.HospitalLocationDomain); ok {
		r0 = rf(ctx, hospitalID)
	} else {
		r0 = ret.Get(0).(rsbedcovid.HospitalLocationDomain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, hospitalID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvince provides a mock function with given fields: ctx
func (_m *Usecase) GetProvince(ctx context.Context) ([]rsbedcovid.ProvinceDomain, error) {
	ret := _m.Called(ctx)

	var r0 []rsbedcovid.ProvinceDomain
	if rf, ok := ret.Get(0).(func(context.Context) []rsbedcovid.ProvinceDomain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]rsbedcovid.ProvinceDomain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
