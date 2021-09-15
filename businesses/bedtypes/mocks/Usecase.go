// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	bedtypes "peduli-covid/businesses/bedtypes"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// FindByHospitalID provides a mock function with given fields: ctx, hospitalID
func (_m *Usecase) FindByHospitalID(ctx context.Context, hospitalID int) ([]bedtypes.Domain, error) {
	ret := _m.Called(ctx, hospitalID)

	var r0 []bedtypes.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) []bedtypes.Domain); ok {
		r0 = rf(ctx, hospitalID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bedtypes.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, hospitalID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreFromAPI provides a mock function with given fields: ctx
func (_m *Usecase) StoreFromAPI(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
