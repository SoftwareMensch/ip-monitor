package notifier

import (
	"errors"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"ip-link-monitor/dto"
	"ip-link-monitor/service/monitor"
	"testing"
)

func TestNewCliCmdNotifier(t *testing.T) {
	executor := &MockedCommandExecutor{}
	mon := &monitor.MockedMonitor{}
	currIp := faker.IPv6()
	targetCmd := faker.Word()

	n := NewCliCmdNotifier(executor, mon, currIp, targetCmd)
	assert.NotNil(t, n)
	assert.Equal(t, executor, n.cmdExecutor)
	assert.Equal(t, mon, n.monitor)
	assert.Equal(t, currIp, n.currentIp)
	assert.Equal(t, targetCmd, n.targetCommand)
}

func TestCliCmdNotifier_Observe(t *testing.T) {
	// monitor run error test
	{
		monErr := errors.New(faker.Sentence())
		mockedMonitor := new(monitor.MockedMonitor)
		mockedMonitor.On("Run").Return(monErr)
		mockedMonitor.On("GetOutputChannel").Return(make(chan *dto.ParsedIpAddress))

		n := &CliCmdNotifier{
			monitor: mockedMonitor,
		}

		err := n.Observe()
		assert.Equal(t, monErr, err)
	}

	// no ip change test
	{
		ipAddress := faker.IPv6()

		monitorOutputCh := make(chan *dto.ParsedIpAddress)
		mockedMonitor := new(monitor.MockedMonitor)
		mockedMonitor.On("Run").Return(nil)
		mockedMonitor.On("GetOutputChannel").Return(monitorOutputCh)

		n := &CliCmdNotifier{
			monitor:   mockedMonitor,
			currentIp: ipAddress,
		}

		go func() {
			err := n.Observe()
			assert.Nil(t, err)
		}()

		monitorOutputCh <- &dto.ParsedIpAddress{
			Value: ipAddress,
		}
		close(monitorOutputCh)
	}

	// executor error test
	{
		currentAddress := faker.IPv6()

		monitorOutputCh := make(chan *dto.ParsedIpAddress)
		mockedMonitor := new(monitor.MockedMonitor)
		mockedMonitor.On("Run").Return(nil)
		mockedMonitor.On("GetOutputChannel").Return(monitorOutputCh)

		mockedExecutor := new(MockedCommandExecutor)
		mockedExecutor.On("Execute", mock.IsType("")).Once().Return(errors.New(faker.Sentence()))

		n := &CliCmdNotifier{
			cmdExecutor:   mockedExecutor,
			monitor:       mockedMonitor,
			currentIp:     currentAddress,
			targetCommand: faker.Word(),
		}

		go func() {
			err := n.Observe()
			assert.Nil(t, err)
		}()

		monitorOutputCh <- &dto.ParsedIpAddress{
			Value: faker.IPv6(),
		}
		close(monitorOutputCh)
	}

	// success test
	{
		currentAddress := faker.IPv6()

		monitorOutputCh := make(chan *dto.ParsedIpAddress)
		mockedMonitor := new(monitor.MockedMonitor)
		mockedMonitor.On("Run").Return(nil)
		mockedMonitor.On("GetOutputChannel").Return(monitorOutputCh)

		mockedExecutor := new(MockedCommandExecutor)
		mockedExecutor.On("Execute", mock.IsType("")).Once().Return(nil)

		n := &CliCmdNotifier{
			cmdExecutor: mockedExecutor,
			monitor:     mockedMonitor,
			currentIp:   currentAddress,
		}

		go func() {
			err := n.Observe()
			assert.Nil(t, err)
		}()

		monitorOutputCh <- &dto.ParsedIpAddress{
			Value: faker.IPv6(),
		}
		close(monitorOutputCh)
	}
}
