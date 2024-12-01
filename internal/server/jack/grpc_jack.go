package jack

import (
	jack "server/api"
)

func NewServer() jack.JackServer {
	return &jackServer{}
}

type jackServer struct {
	jack.UnimplementedJackServer
}
