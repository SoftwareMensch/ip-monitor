package notifier

import (
	"fmt"
	"ip-link-monitor/dto"
	"ip-link-monitor/service/monitor"
	"log"
	"strings"
)

type CliCmdNotifier struct {
	cmdExecutor   CommandExecutor
	monitor       monitor.Monitor
	currentIp     string
	targetCommand string
}

func NewCliCmdNotifier(
	cmdExecutor CommandExecutor,
	monitor monitor.Monitor,
	currentId,
	targetCommand string,
) *CliCmdNotifier {
	return &CliCmdNotifier{
		cmdExecutor:   cmdExecutor,
		monitor:       monitor,
		currentIp:     currentId,
		targetCommand: targetCommand,
	}
}

func (n *CliCmdNotifier) Observe() error {
	outputCh := n.monitor.GetOutputChannel()
	log.Printf("[DEBUG] notifer got output channel from monitor")

	go func(ch chan *dto.ParsedIpAddress) {
		for ipAddress := range ch {
			if 0 != strings.Compare(n.currentIp, ipAddress.Value) {
				log.Printf("[DEBUG] passing found ip address (%s) to \"%s\"", ipAddress.Value, n.targetCommand)
				if err := n.cmdExecutor.Execute(fmt.Sprintf("%s %s", n.targetCommand, ipAddress.Value)); err != nil {
					fmt.Printf("execute \"%s\" error: %s\n", n.targetCommand, err)
				}
				log.Printf("[DEBUG] passed successful")
			}
		}
	}(outputCh)

	if err := n.monitor.Run(); err != nil {
		return err
	}

	return nil
}
