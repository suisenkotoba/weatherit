// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	users "weatherit/usecases/users"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
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

// GetByID provides a mock function with given fields: ctx, userId
func (_m *Repository) GetByID(ctx context.Context, userId int) (users.Domain, error) {
	ret := _m.Called(ctx, userId)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) users.Domain); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, data
func (_m *Repository) Store(ctx context.Context, data *users.Domain) (int, error) {
	ret := _m.Called(ctx, data)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *users.Domain) int); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *users.Domain) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, data
func (_m *Repository) Update(ctx context.Context, data *users.Domain) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *users.Domain) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
