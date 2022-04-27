package monitor

import (
	"bufio"
	"ip-link-monitor/dto"
	parser2 "ip-link-monitor/service/parser"
	"log"
	"os/exec"
	"strings"
)

type StdoutMonitor struct {
	parser        parser2.StdoutParser
	command       string
	outputChannel chan *dto.ParsedIpAddress
}

func NewStdoutMonitor(parser parser2.StdoutParser, command string) *StdoutMonitor {
	return &StdoutMonitor{
		parser:        parser,
		command:       command,
		outputChannel: make(chan *dto.ParsedIpAddress),
	}
}

func (m *StdoutMonitor) GetOutputChannel() chan *dto.ParsedIpAddress {
	return m.outputChannel
}

func (m *StdoutMonitor) Run() error {
	defer close(m.outputChannel)

	args := strings.Fields(m.command)

	cliCommand := exec.Command(args[0], args[1:]...)
	stdout, _ := cliCommand.StdoutPipe()

	if err := cliCommand.Start(); err != nil {
		return err
	}

	buf := bufio.NewReader(stdout)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return err
		}

		strLine := string(line)
		log.Printf("[DEBUG] got stdout line: %s\n", strLine)

		parsedIp := m.parser.Parse(strLine)
		if nil == parsedIp {
			continue
		}

		log.Printf("[DEBUG] found ip address: %s\n", parsedIp.Value)
		m.outputChannel <- parsedIp
		log.Printf("[DEBUG] sent ip address to channel: %s\n", parsedIp.Value)
	}

	return nil
}
