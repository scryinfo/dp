// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	"github.com/scryinfo/dot/dot"
	"github.com/stretchr/testify/mock"
)

// Line is an autogenerated mock type for the Line type
type Line struct {
	mock.Mock
}

// AddNewerByLiveId provides a mock function with given fields: liveid, newDot
func (m *Line) AddNewerByLiveId(liveid dot.LiveId, newDot func([]byte) (dot.Dot, error)) error {
	return nil
}

// AddNewerByTypeId provides a mock function with given fields: typeid, newDot
func (m *Line) AddNewerByTypeId(typeid dot.TypeId, newDot func([]byte) (dot.Dot, error)) error {
	args := m.Called(typeid, newDot)
	return args.Error(0)
}

// Config provides a mock function with given fields:
func (m *Line) Config() *dot.Config {
	args := m.Called()
	return args.Get(0).(*dot.Config)
}

// GetDotConfig provides a mock function with given fields: liveid
func (m *Line) GetDotConfig(liveid dot.LiveId) *dot.LiveConfig {
	args := m.Called(liveid)
	return args.Get(0).(*dot.LiveConfig)
}

// GetLineBuilder provides a mock function with given fields:
func (m *Line) GetLineBuilder() *dot.Builder {
	args := m.Called()
	return args.Get(0).(*dot.Builder)
}

// Id provides a mock function with given fields:
func (m *Line) Id() string {
	args := m.Called()
	return args.String(0)
}

// InfoAllTypeAdnLives provides a mock function with given fields:
func (m *Line) InfoAllTypeAdnLives() {
	_ = m.Called()
}

// PreAdd provides a mock function with given fields: typeLives
func (m *Line) PreAdd(typeLives ...*dot.TypeLives) error {
	args := m.Called(typeLives)
	return args.Error(0)
}

// RemoveNewerByLiveId provides a mock function with given fields: liveid
func (m *Line) RemoveNewerByLiveId(liveid dot.LiveId) {
	_ = m.Called(liveid)
}

// RemoveNewerByTypeId provides a mock function with given fields: typeid
func (m *Line) RemoveNewerByTypeId(typeid dot.TypeId) {
	_ = m.Called(typeid)
}

// SConfig provides a mock function with given fields:
func (m *Line) SConfig() dot.SConfig {
	args := m.Called()
	return args.Get(0).(dot.SConfig)
}

// SLogger provides a mock function with given fields:
func (m *Line) SLogger() dot.SLogger {
	args := m.Called()
	return args.Get(0).(dot.SLogger)
}

// ToDotEventer provides a mock function with given fields:
func (m *Line) ToDotEventer() dot.Eventer {
	args := m.Called()
	return args.Get(0).(dot.Eventer)
}

// ToInjecter provides a mock function with given fields:
func (m *Line) ToInjecter() dot.Injecter {
	args := m.Called()
	return args.Get(0).(dot.Injecter)
}

// ToLifer provides a mock function with given fields:
func (m *Line) ToLifer() dot.Lifer {
	args := m.Called()
	return args.Get(0).(dot.Lifer)
}
