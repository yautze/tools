package st

import "google.golang.org/grpc/codes"

// Errors -
type Errors interface {
	GetCode() int32
	GetGRPCCode() codes.Code
	GetMsg() string
	Err() error
}