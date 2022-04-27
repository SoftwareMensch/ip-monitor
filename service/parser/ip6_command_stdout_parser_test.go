package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewIp6CommandStdoutParser(t *testing.T) {
	p := NewIp6CommandStdoutParser("foobar")
	assert.NotNil(t, p)
	assert.Equal(t, "^[0-9]:\\sfoobar\\s{4}inet6\\s(.*)\\/64\\sscope\\sglobal\\stentative\\sdynamic\\smngtmpaddr\\snoprefixroute.*", p.pattern)
}

func TestIp6CommandStdoutParser_Parse(t *testing.T) {
	// failed test
	{
		lines := []string{
			"Deleted aaaa:aaaa:aaaa:aaa::/64 proto ra metric 1024 expires 7191sec pref medium \n",
			"Deleted 2: eth0    inet6 aaaa:aaaa:aaaa:aaa:bbbb:bbbb:bbbb:bbbb/64 scope global deprecated dynamic mngtmpaddr noprefixroute \n",
			"valid_lft 7191sec preferred_lft 0sec \n",
			"2: eth0    inet6 fd00::aaaa:aaaa:aaaa:aaaa/64 scope global deprecated dynamic mngtmpaddr noprefixroute \n",
			"2: eth0    inet6 fd00::aaaa:aaaa:aaaa:aaaa/64 scope global tentative dynamic mngtmpaddr noprefixroute \n",
			"what ever \n",
		}

		parser := NewIp6CommandStdoutParser("eth0")

		for _, l := range lines {
			parsedIp := parser.Parse(l)
			assert.Nil(t, parsedIp)
		}
	}

	// success test
	{
		line := "2: eth0    inet6 aaaa:aaaa:aaaa:aaaa:bbbb:bbbb:bbbb:bbbb/64 scope global tentative dynamic mngtmpaddr noprefixroute \n"

		parser := NewIp6CommandStdoutParser("eth0")
		parsedIp := parser.Parse(line)
		assert.NotNil(t, parsedIp)
		assert.Equal(t, "aaaa:aaaa:aaaa:aaaa:bbbb:bbbb:bbbb:bbbb", parsedIp.Value)
	}
}
