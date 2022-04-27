package monitor

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"ip-link-monitor/dto"
	parser2 "ip-link-monitor/service/parser"
	path2 "path"
	"runtime"
	"testing"
)

func TestNewStdoutMonitor(t *testing.T) {
	parser := &parser2.MockedStdoutParser{}
	command := faker.Word()

	monitor := NewStdoutMonitor(parser, command)
	assert.NotNil(t, monitor)
	assert.Equal(t, parser, monitor.parser)
	assert.Equal(t, command, monitor.command)
}

func TestStdoutMonitor_Run(t *testing.T) {
	// command start error
	{
		monitor := NewStdoutMonitor(&parser2.MockedStdoutParser{}, "/command/does/not/exists")
		err := monitor.Run()
		assert.NotNil(t, err)
		assert.Equal(t, "fork/exec /command/does/not/exists: no such file or directory", err.Error())
	}

	// nil parsing test
	{
		_, thisFileName, _, _ := runtime.Caller(0)
		dummyCommand := fmt.Sprintf("%s/../../test/bin/dummy-stdout-monitor-command.sh -6 monitor dev eth0", path2.Dir(thisFileName))

		parser := new(parser2.MockedStdoutParser)

		parser.On("Parse").Return(nil)

		monitor := NewStdoutMonitor(parser, dummyCommand)
		ipWasParsed := false

		go func(ch chan *dto.ParsedIpAddress, wasParsed *bool) {
			for range ch {
				*wasParsed = true
			}
		}(monitor.GetOutputChannel(), &ipWasParsed)

		assert.Nil(t, monitor.Run())
		assert.False(t, ipWasParsed)
	}

	// success test
	{
		_, thisFileName, _, _ := runtime.Caller(0)
		dummyCommand := fmt.Sprintf("%s/../../test/bin/dummy-stdout-monitor-command.sh -6 monitor dev eth0", path2.Dir(thisFileName))

		parser := new(parser2.MockedStdoutParser)

		fakedIp := faker.IPv6()
		parser.On("Parse").Return(&dto.ParsedIpAddress{
			Value: fakedIp,
		})

		monitor := NewStdoutMonitor(parser, dummyCommand)

		go func(ch chan *dto.ParsedIpAddress) {
			for parsedIp := range ch {
				assert.Equal(t, fakedIp, parsedIp.Value)
			}
		}(monitor.GetOutputChannel())

		assert.Nil(t, monitor.Run())
	}
}
