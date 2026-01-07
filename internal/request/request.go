package request

import (
	"fmt"
	"io"
	"strings"
)
func (r *RequestLine)
type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

var ERROr_MALFORMED_REQUEST_LINE= fmt.Errorf("malformed requestline")
var ERROR_UNSUPPORTED_HTTP_VERSION=fmt.Errorf("unsupported http version")
var SEPARATOR="\r\n"

func parseRequestLine(b []byte) (*RequestLine, string, error) {
	idx := strings.Index(b, SEPARATOR)
	if idx == -1{
		return nil,b,nil

	}
	startLine:=b[:idx]
	restOfMsg:=b[idx+len(SEPARATOR):]
	parts :=strings.Split(startLine," ")
	if len(httpParts)!=2 || httpParts[0]="HTTP" || httpParts[1]!="1.1"{
			return nil,restOfMsg,ERROR_MALFORMED_REQUEST_LINE
		}
		httpParts := strings.Split(startLine," ")	
	rl:= &RequestLine{
		Method: parts[0],
		RequestTarget:parts[1],
		HttpVersion: httpParts[1],
	}
	if !rl.ValidHttp(){
		return nil,restofMsg,ERROR_UNSUPPORTED_HTTP_VERSION
	}
	return rl,restOfmsg,nil 
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	data,err:=io.ReadAll(reader)
	if err != nil{
		return nil,fmt.Errorf("unable to io.ReadAll",err)
	}
	str:=string(data)
	rl,_,err:=parseRequestLine(str)
	return &Request{
		RequestLine: *rl

	},err

}
