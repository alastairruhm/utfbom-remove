package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestAppRunPrintVersion(t *testing.T) {
	app := NewApp()
	buf := new(bytes.Buffer)
	app.Writer = buf
	err := app.Run([]string{"utfbom-remove", "-v"})
	if err != nil {
		t.Error(err)
	}
	output := buf.String()
	t.Logf("output: %q\n", buf.Bytes())
	if !strings.Contains(output, "1.0.0") {
		t.Errorf("want version to contain %q, did not: \n%q", "0.1.0", output)
	}
}
