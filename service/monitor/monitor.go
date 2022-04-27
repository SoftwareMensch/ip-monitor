package monitor

import (
	"github.com/stretchr/testify/mock"
	"ip-link-monitor/dto"
)

type Monitor interface {
	Run() error
	GetOutputChannel() chan *dto.ParsedIpAddress
}

type MockedMonitor struct {
	mock.Mock
}

func (m *MockedMonitor) Run() error {
	args := m.Called()

	return args.Error(0)
}

func (m *MockedMonitor) GetOutputChannel() chan *dto.ParsedIpAddress {
	args := m.Called()
	if nil == args.Get(0) {
		return nil
	}

	return args.Get(0).(chan *dto.ParsedIpAddress)
}
