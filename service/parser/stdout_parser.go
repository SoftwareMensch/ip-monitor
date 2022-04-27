package parser

import (
	"github.com/stretchr/testify/mock"
	"ip-link-monitor/dto"
)

type StdoutParser interface {
	Parse(line string) *dto.ParsedIpAddress
}

type MockedStdoutParser struct {
	mock.Mock
}

func (m *MockedStdoutParser) Parse(string) *dto.ParsedIpAddress {
	args := m.Called()
	if nil == args.Get(0) {
		return nil
	}

	return args.Get(0).(*dto.ParsedIpAddress)
}
