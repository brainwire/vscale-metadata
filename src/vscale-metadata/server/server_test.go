package server

import (
	"testing"
	_ "time"
)

func TestConnectionClient(t *testing.T) {
	config := &Config{
		Port:          32000,
		Host:          "0.0.0.0",
		ApiKey:				 "",
	}

	srv := NewServer(config)
	go srv.Serve()
}
