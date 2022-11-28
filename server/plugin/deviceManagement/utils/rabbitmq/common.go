package rabbitmq

import (
	"errors"
	"time"
)

const (
	// When reconnecting to the server after connection failure
	ReconnectDelay = 5 * time.Second

	// When setting up the channel after a channel exception
	ReInitDelay = 2 * time.Second

	// When resending messages the server didn't confirm
	ResendDelay = 5 * time.Second
)

var (
	//errNotConnected  = errors.New("not connected to a server")
	ErrAlreadyClosed = errors.New("already closed: not connected to the server")
	ErrShutdown      = errors.New("session is shutting down")
)
