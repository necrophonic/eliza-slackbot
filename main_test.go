package main

import (
	"testing"
	"os"
)

func TestNoToken (t *testing.T) {
	_, _ , err := startRTM()

	if err.Error() != "Must set ELIZA_BOT_TOKEN environment variable" {
		t.Errorf("Expecting 'Must set ELIZA_BOT_TOKEN environment variable' as error")
	}
}

func TestConnectionWithSlack (t *testing.T) {
	os.Setenv("ELIZA_BOT_TOKEN", "TEST_KEY")
	_, _ , err := startRTM()

	if err == nil {
		t.Errorf("Expecting 'Slack error, invalid token")
	}
}

