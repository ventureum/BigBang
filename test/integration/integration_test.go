package integration

import (
	"BigBang/test/integration/TCR"
	"BigBang/test/integration/feed"
	"BigBang/test/integration/migrations/clear_tables"
	"testing"
)

func TestHandler(t *testing.T) {
	t.Run("clear tables for feed integration test", clear_tables_test.TestHandler)
	t.Run("feed integration test", feed_integration_test.TestHandler)
	t.Run("clear tables for tcr integration test", clear_tables_test.TestHandler)
	t.Run("tcr integration test", tcr_integration_test.TestHandler)
}
