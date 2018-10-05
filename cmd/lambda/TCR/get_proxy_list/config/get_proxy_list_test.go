package config

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/pkg/utils"
)

var EmptyStringLIst []string

func TestHandler(t *testing.T) {
  tests := []struct{
    request Request
    response Response
    err    error
  }{
    {
      request: Request {
        Limit: 0,
      },
      response: Response {
        Proxies: &EmptyStringLIst,
        NextCursor: utils.Base64EncodeInt64(6),
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        Limit: 2,
      },
      response: Response {
        Proxies: &[]string{ "0xProxy006", "0xProxy005",},
        NextCursor: utils.Base64EncodeInt64(4),
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        Limit: 2,
        Cursor: utils.Base64EncodeInt64(4),
      },
      response: Response {
        Proxies: &[]string{ "0xProxy004", "0xProxy003",},
        NextCursor: utils.Base64EncodeInt64(2),
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        Limit: 5,
        Cursor: utils.Base64EncodeInt64(4),
      },
      response: Response {
        Proxies: &[]string{ "0xProxy004", "0xProxy003", "0xProxy002", "0xProxy001"},
        NextCursor: "",
        Ok: true,
      },
      err: nil,
    },
  }

  for _, test := range tests {
    result, err := Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}

