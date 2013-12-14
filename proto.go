package proto

import (
	"log"
	"encoding/json"
)

/**
 * Service and command names
 */

var Services []string = []string{
	"qan",	// Query Analytics
	"mm",	// Metrics Monitor
}

var Commands []string = []string{
	// Cmd.Cmd          Cmd.Data     Reply.Data
	// ===============  ===========  ===========
	"SetConifg",    //  
	"StartService", //  ServiceData
	"StopService",  //  ServiceData
	"Status",       //               StatusReply
	"Update",       // 
	"Exit",         //
}

/**
 * JSON message structures
 */

// Sent by user to API
type Cmd struct {
	RelayId string	`json:",omitempty"`		// set by API
	Agent string							// UUID
	Cmd string								// one of Commands
	Data []byte		`json:",omitempty"`		// struct for Cmd, if any
}

// Sent by agent to user (via API)
type Reply struct {
	RelayId string							// for API, user can ignore
	Cmd string								// original Cmd.Cmd 
	Error string							// success if empty
	Data []byte		`json:",omitempty"`
}

// Data for StartService and StopService command replies
type ServiceData struct {
	Name	string						// one of Services
	Config	[]byte `json:",omitempty"`	// cloud-tools/<service>/config.go
}

// Data for Status command reply
type StatusData struct {
	Agent string				// agent internals
	CmdQueue []string			// Cmd.Cmd agent has queued
	Service map[string]string	// keyed on Services
}

/**
 * Functions
 */

// Create Reply from Cmd
func (cmd *Cmd) Reply(err string, data interface {}) *Reply {
	reply := &Reply{
		RelayId: cmd.RelayId,
		Cmd: cmd.Cmd,
		Error: err,
	}
	if data != nil {
		codedData, jsonErr := json.Marshal(data)
		if jsonErr != nil {
			log.Panic(jsonErr) // @todo
		}
		reply.Data = codedData
	}
	return reply
}
