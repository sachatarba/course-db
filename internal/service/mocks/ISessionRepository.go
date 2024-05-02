// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/sachatarba/course-db/internal/entity"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ISessionRepository is an autogenerated mock type for the ISessionRepository type
type ISessionRepository struct {
	mock.Mock
}

// CreateNewSession provides a mock function with given fields: ctx, session
func (_m *ISessionRepository) CreateNewSession(ctx context.Context, session entity.Session) error {
	ret := _m.Called(ctx, session)

	if len(ret) == 0 {
		panic("no return value specified for CreateNewSession")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Session) error); ok {
		r0 = rf(ctx, session)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSession provides a mock function with given fields: ctx, clientID
func (_m *ISessionRepository) DeleteSession(ctx context.Context, clientID uuid.UUID) error {
	ret := _m.Called(ctx, clientID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSession")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, clientID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSessionBySessionID provides a mock function with given fields: ctx, sessionID
func (_m *ISessionRepository) DeleteSessionBySessionID(ctx context.Context, sessionID uuid.UUID) error {
	ret := _m.Called(ctx, sessionID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSessionBySessionID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, sessionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSessionBySessionID provides a mock function with given fields: ctx, sessionID
func (_m *ISessionRepository) GetSessionBySessionID(ctx context.Context, sessionID uuid.UUID) (entity.Session, error) {
	ret := _m.Called(ctx, sessionID)

	if len(ret) == 0 {
		panic("no return value specified for GetSessionBySessionID")
	}

	var r0 entity.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (entity.Session, error)); ok {
		return rf(ctx, sessionID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) entity.Session); ok {
		r0 = rf(ctx, sessionID)
	} else {
		r0 = ret.Get(0).(entity.Session)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, sessionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSessionsByClientID provides a mock function with given fields: ctx, clientID
func (_m *ISessionRepository) GetSessionsByClientID(ctx context.Context, clientID uuid.UUID) ([]entity.Session, error) {
	ret := _m.Called(ctx, clientID)

	if len(ret) == 0 {
		panic("no return value specified for GetSessionsByClientID")
	}

	var r0 []entity.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]entity.Session, error)); ok {
		return rf(ctx, clientID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []entity.Session); ok {
		r0 = rf(ctx, clientID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, clientID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewISessionRepository creates a new instance of ISessionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewISessionRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ISessionRepository {
	mock := &ISessionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}