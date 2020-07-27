package st

import (
	"google.golang.org/grpc/codes"
	gs "google.golang.org/grpc/status"
)

// err -
type err struct {
	gst *gs.Status
	st  *body
}

// body -
type body struct {
	Code int32  `json:"code" yaml:"code" `
	Msg  string `json:"msg" yaml:"msg"`
}

// parseMsg -
func (e *err) parseMsg() error {
	e.st = &body{}

	err := json.Unmarshal([]byte(e.gst.Message()), e.st)

	return err
}

func (e *err) Err() error {
	return e.gst.Err()
}

// get err.st.Code
func (e *err) GetCode() int32 {
	return e.st.Code
}

// get err.st.Msg
func (e *err) GetMsg() string {
	return e.st.Msg
}

// get err.gst.Code
func (e *err) GetGRPCCode() codes.Code {
	return e.gst.Code()
}
