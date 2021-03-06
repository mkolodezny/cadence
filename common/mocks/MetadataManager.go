// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	persistence "github.com/uber/cadence/common/persistence"
)

// MetadataManager is an autogenerated mock type for the MetadataManager type
type MetadataManager struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *MetadataManager) Close() {
	_m.Called()
}

// CreateDomain provides a mock function with given fields: ctx, request
func (_m *MetadataManager) CreateDomain(ctx context.Context, request *persistence.CreateDomainRequest) (*persistence.CreateDomainResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.CreateDomainResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.CreateDomainRequest) *persistence.CreateDomainResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.CreateDomainResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.CreateDomainRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDomain provides a mock function with given fields: ctx, request
func (_m *MetadataManager) DeleteDomain(ctx context.Context, request *persistence.DeleteDomainRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.DeleteDomainRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDomainByName provides a mock function with given fields: ctx, request
func (_m *MetadataManager) DeleteDomainByName(ctx context.Context, request *persistence.DeleteDomainByNameRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.DeleteDomainByNameRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDomain provides a mock function with given fields: ctx, request
func (_m *MetadataManager) GetDomain(ctx context.Context, request *persistence.GetDomainRequest) (*persistence.GetDomainResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.GetDomainResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.GetDomainRequest) *persistence.GetDomainResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetDomainResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.GetDomainRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMetadata provides a mock function with given fields: ctx
func (_m *MetadataManager) GetMetadata(ctx context.Context) (*persistence.GetMetadataResponse, error) {
	ret := _m.Called(ctx)

	var r0 *persistence.GetMetadataResponse
	if rf, ok := ret.Get(0).(func(context.Context) *persistence.GetMetadataResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetMetadataResponse)
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

// GetName provides a mock function with given fields:
func (_m *MetadataManager) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ListDomains provides a mock function with given fields: ctx, request
func (_m *MetadataManager) ListDomains(ctx context.Context, request *persistence.ListDomainsRequest) (*persistence.ListDomainsResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.ListDomainsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.ListDomainsRequest) *persistence.ListDomainsResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.ListDomainsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.ListDomainsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDomain provides a mock function with given fields: ctx, request
func (_m *MetadataManager) UpdateDomain(ctx context.Context, request *persistence.UpdateDomainRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.UpdateDomainRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
