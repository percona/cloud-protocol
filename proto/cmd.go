package proto

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
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

// Sent by user to agent via API (API relays Cmd to agent and Reply to user)
type Cmd struct {
	User      string
	Ts        time.Time
	AgentUuid string
	Service   string
	Cmd       string
	Data      []byte `json:",omitempty"` // struct for Cmd, if any
	// --
	CmdId   string `json:",omitempty"` // set by User
	RelayId string `json:",omitempty"` // set by API
}

// Sent by agent to user in response to every command
type Reply struct {
	Cmd   string // original Cmd.Cmd
	Error string // success if empty
	Data  []byte `json:",omitempty"`
	// --
	CmdId   string `json:",omitempty"` // set by User
	RelayId string // set by API
}

// Data for StartService and StopService command replies
type ServiceData struct {
	Name   string
	Config []byte `json:",omitempty"` // cloud-tools/<service>/config.go
	// --
	ConfigId uint // set by API
}

// Data for Status command reply
type StatusData map[string]string

/*
	Agent           string
	AgentCmdHandler string
	AgentCmdQueue   []string
	Qan             string
	QanLogParser    string
	Mm              string
	MmMonitors      map[string]string
}
*/

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

func (cmd *Cmd) Reply(err error, data interface{}) *Reply {
	// todo: encoding/json or websocket.JSON doesn't seem to handle error type
	reply := &Reply{
		Cmd:     cmd.Cmd,
		CmdId:   cmd.CmdId,
		RelayId: cmd.RelayId,
	}
	if err != nil {
		reply.Error = err.Error()
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

// Used by pct.Logger and pct.Status to stringify Cmd related to log entries and status updates
func (cmd *Cmd) String() string {
	cmdWithoutData := *cmd
	cmdWithoutData.Data = []byte{}
	bytes, err := json.Marshal(cmdWithoutData)
	if err != nil {
		return ""
	}
	return string(bytes)
}
