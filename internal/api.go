package wrabbit

import (
	"bufio"
	"encoding/json"
	"net"
	"strings"

	"github.com/zp4rker/wrabbit/pkg/data"
)

func (wr *Wrabbit) Listen() {
	for {
		// accept request
		s, con, err := wr.acceptRequest()
		if err != nil {
			continue
		}

		// parse request
		req, err := parseRequest(s)
		if err != nil {
			resp := data.FailedResponse("Invalid request!")
			raw, _ := json.Marshal(resp)
			con.Write(raw)
			con.Close()
			continue
		}

		switch req.Action {
		case "output":
			// handle output
		case "input":
			// handle input
		case "kill":
			// handle kill
		case "status":
			// handle status
		default:
			resp := data.FailedResponse("Invalid action!")
			raw, _ := json.Marshal(resp)
			con.Write(raw)
		}

		con.Close()
	}
}

// accept incoming request
func (wr *Wrabbit) acceptRequest() (string, net.Conn, error) {
	// accept connection
	con, err := wr.Listener.Accept()
	if err != nil {
		return "", nil, err
	}

	// read request
	rd := bufio.NewReader(con)
	s, err := rd.ReadString('\n')
	if err != nil {
		return "", nil, err
	}

	return s, con, nil
}

// parse request from json to struct
func parseRequest(s string) (*data.Request, error) {
	var req data.Request
	err := json.Unmarshal([]byte(s), &req)
	if err != nil {
		return nil, err
	}

	req.Action = strings.ToLower(req.Action)
	req.Type = strings.ToLower(req.Type)

	return &req, nil
}
