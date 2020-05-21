package libgen_cli

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Skip()
	// Create command
	cmd := searchCmd
	b := bytes.NewBufferString("")
	// Set command output to our bytes
	cmd.SetOut(b)
	// Add arguments
	cmd.SetArgs([]string{"test", "-r", "1"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("error executing command: %v", err)
	}
	// Read bytes outputted by command
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatalf("error reading command output: %v", err)
	}
	// Confirm values are expected
	if strings.Contains(string(out), "The Turing Test and the Frame Problem: AI's Mistaken Understanding of "+
		"Intelligence") {

		t.Fatalf("expected \"%s\" got \"%s\"", "The Turing Test and the Frame Problem: AI's "+
			"Mistaken Understanding of Intelligence", string(out))
	}
}