// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"

import mock "github.com/stretchr/testify/mock"
import model "github.com/ddoffy/shopping/model"

// ERepository is an autogenerated mock type for the ERepository type
type ERepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, username
func (_m *ERepository) Delete(ctx context.Context, username string) (bool, error) {
	ret := _m.Called(ctx, username)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchLoginwithUsername provides a mock function with given fields: ctx, query
func (_m *ERepository) FetchLoginwithUsername(ctx context.Context, query string) ([]*model.Login, error) {
	ret := _m.Called(ctx, query)

	var r0 []*model.Login
	if rf, ok := ret.Get(0).(func(context.Context, string) []*model.Login); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Login)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, a
func (_m *ERepository) Store(ctx context.Context, a *model.Login) (int64, error) {
	ret := _m.Called(ctx, a)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *model.Login) int64); ok {
		r0 = rf(ctx, a)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.Login) error); ok {
		r1 = rf(ctx, a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
