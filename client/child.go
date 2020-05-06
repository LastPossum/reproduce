package client

import "github.com/LastPossum/reproduce/server"

func PublicFunc() string {
	return server.PublicFunc()
}
