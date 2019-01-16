package integration

import (
	"testing"
	"BigBang/test/integration/feed"
	"BigBang/test/integration/TCR"
)

func TestHandler(t *testing.T) {
	t.Run("feed integration test", feed_integration_test.TestHandler)
	t.Run("tcr integration test", tcr_integration_test.TestHandler)
}
