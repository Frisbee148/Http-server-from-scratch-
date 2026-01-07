package request

import (
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

var ERROr_MALFORMED_REQUEST_LINE= fmt.Errorf("malformed requestline")
var INCOMPLETE_START_LINE=fmt.Errorf("")
var SEPARATOR="\r\n"

func parseRequestLine(b []byte) (*RequestLine, string, error) {
	idx := strings.Index(b, SEPARATOR)
	if idx == -1{
		return nil,b,nil

	}
	startLine:=b[:idx]
	restOfMsg:=b[idx+len(SEPARATOR):]
	parts :=strings.Split(startLine," ")
	if len(parts)!=3{
			return nil,restOfMsg,ERROr_MALFORMED_REQUEST_LINE
		}
	return &RequestLine{
		Method: parts[0],
		RequestTarget:parts[1],
	},restOfMsg,nil

func RequestFromReader(reader io.Reader) (*Request, error) {

}
