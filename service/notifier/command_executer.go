package notifier

import (
	"github.com/stretchr/testify/mock"
	"os/exec"
	"strings"
)

type CommandExecutor interface {
	Execute(cmd string) error
}

type MockedCommandExecutor struct {
	mock.Mock
}

func (m *MockedCommandExecutor) Execute(cmd string) error {
	args := m.Called(cmd)

	if nil == args.Get(0) {
		return nil
	}

	return args.Get(0).(error)
}

type ShellCommandExecutor struct {
}

func NewShellCommandExecutor() *ShellCommandExecutor {
	return &ShellCommandExecutor{}
}

func (s *ShellCommandExecutor) Execute(cmd string) error {
	args := strings.Fields(cmd)

	// wie cannot simply mock this for unit testing
	cliCommand := exec.Command(args[0], args[1:]...)

	defer func() { _ = cliCommand.Wait() }()
	if err := cliCommand.Start(); err != nil {
		return err
	}

	return nil
}
