package proto_test

import (
	proto "github.com/percona/cloud-protocol"
	. "launchpad.net/gocheck"
	"testing"
)

// Hook gocheck into the "go test" runner.
// http://labix.org/gocheck
func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s *TestSuite) TestReply(t *C) {
	cmd := &proto.Cmd{
		CmdId:     "0x02",        // normally set by User
		RelayId:   "0x01",        // normally set by API
		AgentUuid: "abc-123-def", // fake uuid
		Cmd:       "StartService",
		Data:      []byte("..."),
	}
	got := cmd.Reply(nil, nil)
	expect := &proto.Reply{
		CmdId:   "0x02",
		RelayId: "0x01",
		Cmd:     "StartService",
		Error:   "",
		Data:    nil,
	}
	t.Check(got, DeepEquals, expect)
}
