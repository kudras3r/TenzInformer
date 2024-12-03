package grab

import (
	"os"
	"testing"
)

func TestPCInfo(t *testing.T) {
	validYAML := `
os:
  family: linux
  name: ubuntu
  kernel: 5.4.0-26-generic
  codename: focal
  type: desktop
  platform: x86_64
  version: "20.04"
name: MyPC
mac: "00:1A:2B:3C:4D:5E"
`
	tmpFile, err := os.CreateTemp("", "test_config_*.yaml")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write([]byte(validYAML)); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	info, err := PCInfo(tmpFile.Name())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedOSName := "ubuntu"
	if info.OS.OSName != expectedOSName {
		t.Errorf("expected OS name %s, got %s", expectedOSName, info.OS.OSName)
	}

	_, err = PCInfo("nonexistent.yaml")
	if err == nil {
		t.Fatal("expected error for nonexistent file")
	}
}
