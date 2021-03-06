// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	users "peduli-covid/businesses/users"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// FindByHospitalID provides a mock function with given fields: ctx, hospitalID
func (_m *Repository) FindByHospitalID(ctx context.Context, hospitalID int) ([]users.Domain, error) {
	ret := _m.Called(ctx, hospitalID)

	var r0 []users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) []users.Domain); ok {
		r0 = rf(ctx, hospitalID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Domain)
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

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *Repository) GetByEmail(ctx context.Context, email string) (users.Domain, error) {
	ret := _m.Called(ctx, email)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) users.Domain); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Repository) GetByID(ctx context.Context, id int) (users.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) users.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, data
func (_m *Repository) Store(ctx context.Context, data *users.Domain) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *users.Domain) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
