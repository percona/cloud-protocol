package proto_test

import (
	"testing"
	. "launchpad.net/gocheck"
	. "github.com/percona/cloud-protocol"
)

// Hook gocheck into the "go test" runner.
// http://labix.org/gocheck
func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}
var _ = Suite(&TestSuite{})

func (s *TestSuite) TestReply(t *C) {
	cmd := &Cmd{
		RelayId: "0x01",		// normally set by API
		Agent: "abc-123-def",	// fake uuid
		Cmd: "StartService",
		Data: []byte("..."),
	}
	got := cmd.Reply("", nil)
	expect := &Reply{
		RelayId: "0x01",
		Cmd: "StartService",
		Error: "",
		Data: nil,
	}
	t.Check(got, DeepEquals, expect)
}
