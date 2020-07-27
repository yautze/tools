package srv

import "net/http"

// Service -
type Service interface {
	// The service id
	ID() string
	// The service name
	Name() string
	// Initialize
	Init(...Option)
	// Version
	Version() string
	// Run the service
	Run() error
	// Options returns the current options
	Options() Options
	// The server
	Server() http.Handler
	// The server type
	ServerType() string
}

// Option -
type Option func(*Options)

// NewService - creates and returns a new Service
func NewService(opts ...Option) Service {
	return newService(opts...)
}