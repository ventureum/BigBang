package get_project

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/cmd/lambda/TCR/tcr_table_creator/config"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    response lambda_tcr_table_creator_config.Response
    err    error
  }{
    {
      response: lambda_tcr_table_creator_config.Response {
        Ok: true,
      },
      err: nil,
    },
  }
  for _, test := range tests {
    result, err := lambda_tcr_table_creator_config.Handler()
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response.Ok, result.Ok)
  }
}
