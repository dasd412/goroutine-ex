package contextcancel

import "testing"

func TestContextCancel(t *testing.T) {
	ContextCancel()
}

func TestParentChildCancel(t *testing.T) {
	results := ParentChildCancel()

	if len(results) != 2 {
		t.Errorf("expected 2 results, got %d", len(results))
	}
}
