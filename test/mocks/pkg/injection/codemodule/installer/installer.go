// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Installer is an autogenerated mock type for the Installer type
type Installer struct {
	mock.Mock
}

type Installer_Expecter struct {
	mock *mock.Mock
}

func (_m *Installer) EXPECT() *Installer_Expecter {
	return &Installer_Expecter{mock: &_m.Mock}
}

// InstallAgent provides a mock function with given fields: targetDir
func (_m *Installer) InstallAgent(targetDir string) (bool, error) {
	ret := _m.Called(targetDir)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(targetDir)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(targetDir)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(targetDir)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Installer_InstallAgent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InstallAgent'
type Installer_InstallAgent_Call struct {
	*mock.Call
}

// InstallAgent is a helper method to define mock.On call
//   - targetDir string
func (_e *Installer_Expecter) InstallAgent(targetDir interface{}) *Installer_InstallAgent_Call {
	return &Installer_InstallAgent_Call{Call: _e.mock.On("InstallAgent", targetDir)}
}

func (_c *Installer_InstallAgent_Call) Run(run func(targetDir string)) *Installer_InstallAgent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Installer_InstallAgent_Call) Return(_a0 bool, _a1 error) *Installer_InstallAgent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Installer_InstallAgent_Call) RunAndReturn(run func(string) (bool, error)) *Installer_InstallAgent_Call {
	_c.Call.Return(run)
	return _c
}

// NewInstaller creates a new instance of Installer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInstaller(t interface {
	mock.TestingT
	Cleanup(func())
}) *Installer {
	mock := &Installer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}