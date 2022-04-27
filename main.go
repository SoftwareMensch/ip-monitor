package main

import (
	"fmt"
	monitor2 "ip-link-monitor/service/monitor"
	notifier2 "ip-link-monitor/service/notifier"
	parser2 "ip-link-monitor/service/parser"
	"os"
)

func usage() string {
	return `
USAGE:

    cmd [NET DEVICE] [CURRENT IP6 AT DEVICE] [TARGET CMD]
`
}

func main() {
	args := os.Args
	if len(args) != 4 {
		fmt.Println(usage())
		os.Exit(1)
	}

	deviceName := args[1]
	currentIp := args[2]
	targetCmd := args[3]

	parser := parser2.NewIp6CommandStdoutParser(deviceName)
	monitor := monitor2.NewStdoutMonitor(parser, fmt.Sprintf("ip -6 monitor dev %s", deviceName))
	shellExecutor := notifier2.NewShellCommandExecutor()
	notifier := notifier2.NewCliCmdNotifier(shellExecutor, monitor, currentIp, targetCmd)

	if err := notifier.Observe(); err != nil {
		panic(err)
	}
}
