package proto

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

/**
 * Service commands
 */

var Commands map[string][]string = map[string][]string{
	"": []string{ // standard (all services)
		"SetConfig",
		"GetConfig",
		"Status",
	},
	"agent": []string{
		"StartService",
		"StopService",
		"Restart",
		"Stop",
		"Abort",
		"Update",
	},
	"data": []string{},
	"log":  []string{},
	"qan":  []string{},
	"mm": []string{
		"StartService", // monitor
		"StopService",  // monitor
	},
}

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

// Sent by agent to user in response to every command
type Reply struct {
	Cmd    string            // original Cmd.Cmd
	Error  string            // success if empty
	Data   []byte            `json:",omitempty"`
	Status map[string]string `json:",omitempty"`
	// --
	RelayId string // set by API
}

// Data for StartService and StopService command replies
type ServiceData struct {
	Name   string
	Config []byte `json:",omitempty"` // cloud-tools/<service>/config.go
	// --
	ConfigId uint // set by API
}

/**
 * Functions
 */

func (cmd *Cmd) Validate() error {
	cmds, ok := Commands[cmd.Service]
	if !ok {
		return errors.New(fmt.Sprintf("Invalid service: %s", cmd.Service))
	}

	validCmd := false
	for _, val := range cmds {
		if cmd.Cmd == val {
			validCmd = true
			break
		}
	}
	if !validCmd {
		return errors.New(fmt.Sprintf("Invalid command for %s: %s", cmd.Service, cmd.Cmd))
	}

	return nil // is valid
}

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
		if cmd.Cmd == "Status" {
			reply.Status = data.(map[string]string)
		} else {
			codedData, jsonErr := json.Marshal(data)
			if jsonErr != nil {
				// This shouldn't happen.
				log.Fatal(jsonErr)
			}
			reply.Data = codedData
		}
	}
	return reply
}

// [2013-01-01T01:01:03 user@example.com 00000000-0000-0000-0000-000000000000 mm StartService]
func (cmd *Cmd) String() string {
	cmdx := *cmd
	cmdx.Data = nil
	return fmt.Sprintf("[%s %s %s %s %s]",
		cmdx.Ts, cmdx.User, cmdx.AgentUuid, cmdx.Service, cmdx.Cmd)
}
