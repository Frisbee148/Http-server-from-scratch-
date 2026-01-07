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
	Method        string
	RequestTarget string
	HttpVersion   string
}

var ERROR_MALFORMED_REQUEST_LINE = fmt.Errorf("malformed request line")
var ERROR_UNSUPPORTED_HTTP_VERSION = fmt.Errorf("unsupported http version")

const SEPARATOR = "\r\n"

func (r *RequestLine) ValidHttp() bool {
	return r.HttpVersion == "HTTP/1.1"
}

func parseRequestLine(b string) (*RequestLine, string, error) {
	idx := strings.Index(b, SEPARATOR)
	if idx == -1 {
		return nil, b, nil
	}

	startLine := b[:idx]
	restOfMsg := b[idx+len(SEPARATOR):]

	parts := strings.Split(startLine, " ")
	if len(parts) != 3 {
		return nil, restOfMsg, ERROR_MALFORMED_REQUEST_LINE
	}

	rl := &RequestLine{
		Method:        parts[0],
		RequestTarget: parts[1],
		HttpVersion:   parts[2],
	}

	if !rl.ValidHttp() {
		return nil, restOfMsg, ERROR_UNSUPPORTED_HTTP_VERSION
	}

	return rl, restOfMsg, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("unable to io.ReadAll: %w", err)
	}

	rl, _, err := parseRequestLine(string(data))
	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: *rl,
	}, nil
}
