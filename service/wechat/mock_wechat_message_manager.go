// Code generated by mockery v2.43.2. DO NOT EDIT.

package wechat

import (
	message "github.com/silenceper/wechat/v2/officialaccount/message"
	mock "github.com/stretchr/testify/mock"
)

// mockwechatMessageManager is an autogenerated mock type for the wechatMessageManager type
type mockwechatMessageManager struct {
	mock.Mock
}

type mockwechatMessageManager_Expecter struct {
	mock *mock.Mock
}

func (_m *mockwechatMessageManager) EXPECT() *mockwechatMessageManager_Expecter {
	return &mockwechatMessageManager_Expecter{mock: &_m.Mock}
}

// Send provides a mock function with given fields: msg
func (_m *mockwechatMessageManager) Send(msg *message.CustomerMessage) error {
	ret := _m.Called(msg)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*message.CustomerMessage) error); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockwechatMessageManager_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type mockwechatMessageManager_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - msg *message.CustomerMessage
func (_e *mockwechatMessageManager_Expecter) Send(msg interface{}) *mockwechatMessageManager_Send_Call {
	return &mockwechatMessageManager_Send_Call{Call: _e.mock.On("Send", msg)}
}

func (_c *mockwechatMessageManager_Send_Call) Run(run func(msg *message.CustomerMessage)) *mockwechatMessageManager_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*message.CustomerMessage))
	})
	return _c
}

func (_c *mockwechatMessageManager_Send_Call) Return(_a0 error) *mockwechatMessageManager_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockwechatMessageManager_Send_Call) RunAndReturn(run func(*message.CustomerMessage) error) *mockwechatMessageManager_Send_Call {
	_c.Call.Return(run)
	return _c
}

// newMockwechatMessageManager creates a new instance of mockwechatMessageManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockwechatMessageManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockwechatMessageManager {
	mock := &mockwechatMessageManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
