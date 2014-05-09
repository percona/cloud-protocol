package proto

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

/**
 * JSON message structures
 */

// Sent by API to agent
type Cmd struct {
	Ts        time.Time
	User      string
	AgentUuid string
	Service   string
	Cmd       string
	Data      []byte `json:",omitempty"`
	// --
	RelayId string `json:",omitempty"` // set by API
}

// Sent by agent in response to every command
type Reply struct {
	Cmd   string // original Cmd.Cmd
	Error string // success if empty
	Data  []byte `json:",omitempty"`
	// --
	RelayId string // set by API
}

// Data for StartService and StopService command replies
type ServiceData struct {
	Name   string
	Config []byte `json:",omitempty"` // cloud-tools/<service>/config.go
}

// Reply from agent to Version command.  The two can differ is agent self-update
// but hasn't be restarted yet.
type Version struct {
	Installed string
	Running   string
	Revision  string
}

/**
 * Functions
 */

func (cmd *Cmd) Reply(data interface{}, errs ...error) *Reply {
	// todo: encoding/json or websocket.JSON doesn't seem to handle error type

	reply := &Reply{
		Cmd:     cmd.Cmd,
		RelayId: cmd.RelayId,
	}
	if len(errs) > 0 {
		errmsgs := make([]string, len(errs))
		for i, err := range errs {
			if err == nil {
				continue
			}
			errmsgs[i] = err.Error()
		}
		reply.Error = strings.Join(errmsgs, "\n")
	}
	if data != nil {
		codedData, jsonErr := json.Marshal(data)
		if jsonErr != nil {
			// This shouldn't happen.
			log.Fatal(jsonErr)
		}
		reply.Data = codedData
	}
	return reply
}

func (cmd *Cmd) String() string {
	cmdx := *cmd
	cmdx.Data = nil
	return fmt.Sprintf("Cmd[Service:%s Cmd:%s Ts:'%s' User:%s AgentUuid:%s RelayId:%s]",
		cmdx.Service, cmdx.Cmd,
		cmdx.Ts, cmdx.User, cmd.AgentUuid,
		cmdx.RelayId)
}

func (reply *Reply) String() string {
	replyx := *reply
	replyx.Data = nil
	return fmt.Sprintf("Reply[Cmd:%s Error:'%s' RelayId:%s]", replyx.Cmd, replyx.Error, replyx.RelayId)
}
