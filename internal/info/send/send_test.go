package send

import (
	"os"
	"testing"
)

func TestJSON(t *testing.T) {
	filePath := "/opt/tenzir/bin/tenzir"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Fatalf("cannot start tenzir pipeline: %v", err)
	}
}
