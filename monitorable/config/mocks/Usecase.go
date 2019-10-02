// Code generated by mockery v1.0.0. DO NOT EDIT.

// If you want to rebuild this file, make mock-monitorable

package mocks

import builder "github.com/monitoror/monitoror/pkg/monitoror/builder"

import mock "github.com/stretchr/testify/mock"
import models "github.com/monitoror/monitoror/monitorable/config/models"
import monitorormodels "github.com/monitoror/monitoror/models"
import validator "github.com/monitoror/monitoror/pkg/monitoror/validator"

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// GetConfig provides a mock function with given fields: params
func (_m *Usecase) GetConfig(params *models.ConfigParams) (*models.Config, error) {
	ret := _m.Called(params)

	var r0 *models.Config
	if rf, ok := ret.Get(0).(func(*models.ConfigParams) *models.Config); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Config)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.ConfigParams) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Hydrate provides a mock function with given fields: _a0
func (_m *Usecase) Hydrate(_a0 *models.Config) {
	_m.Called(_a0)
}

// RegisterDynamicTile provides a mock function with given fields: tileType, _a1, _a2
func (_m *Usecase) RegisterDynamicTile(tileType monitorormodels.TileType, _a1 validator.Validator, _a2 builder.DynamicTileBuilder) {
	_m.Called(tileType, _a1, _a2)
}

// RegisterDynamicTileWithConfigVariant provides a mock function with given fields: tileType, configVariant, _a2, _a3
func (_m *Usecase) RegisterDynamicTileWithConfigVariant(tileType monitorormodels.TileType, configVariant string, _a2 validator.Validator, _a3 builder.DynamicTileBuilder) {
	_m.Called(tileType, configVariant, _a2, _a3)
}

// RegisterTile provides a mock function with given fields: tileType, _a1, path
func (_m *Usecase) RegisterTile(tileType monitorormodels.TileType, _a1 validator.Validator, path string) {
	_m.Called(tileType, _a1, path)
}

// RegisterTileWithConfigVariant provides a mock function with given fields: tileType, variant, _a2, path
func (_m *Usecase) RegisterTileWithConfigVariant(tileType monitorormodels.TileType, variant string, _a2 validator.Validator, path string) {
	_m.Called(tileType, variant, _a2, path)
}

// Verify provides a mock function with given fields: _a0
func (_m *Usecase) Verify(_a0 *models.Config) {
	_m.Called(_a0)
}
