package parser

import (
	"fmt"
	"ip-link-monitor/dto"
	"regexp"
	"strings"
)

type Ip6CommandStdoutParser struct {
	pattern string
}

func NewIp6CommandStdoutParser(
	deviceName string,
) *Ip6CommandStdoutParser {
	return &Ip6CommandStdoutParser{
		pattern: fmt.Sprintf(
			"^[0-9]:\\s%s\\s{4}inet6\\s(.*)\\/64\\sscope\\sglobal\\stentative\\sdynamic\\smngtmpaddr\\snoprefixroute.*",
			deviceName,
		),
	}
}

func (p *Ip6CommandStdoutParser) Parse(line string) *dto.ParsedIpAddress {
	hasMatch, _ := regexp.MatchString(p.pattern, line)
	if !hasMatch {
		return nil
	}

	exp := regexp.MustCompile(p.pattern)

	matches := exp.FindStringSubmatch(line)
	if len(matches) != 2 {
		return nil
	}

	// ignore local addresses
	if strings.Contains(matches[1], "fd00::") {
		return nil
	}
	if strings.Contains(matches[1], "fe80::") {
		return nil
	}

	return &dto.ParsedIpAddress{
		Value: matches[1],
	}
}
