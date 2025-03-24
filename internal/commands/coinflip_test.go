package commands

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestFlipCoin(t *testing.T) {
	numFlips := 10
	results := flipCoin(numFlips)

	if len(results) != numFlips {
		t.Errorf("Expected %d results, got %d", numFlips, len(results))
	}

	for _, result := range results {
		if result != "Heads" && result != "Tails" {
			t.Errorf("Unexpected result: %s", result)
		}
	}
}

func TestDisplayFlipResults(t *testing.T) {
	results := []string{"Heads", "Tails", "Heads"}
	expectedOutput := "3 coin flips:\n  - Flip 1: Heads\n  - Flip 2: Tails\n  - Flip 3: Heads\n"

	output := CaptureOutput(func() {
		displayFlipResults(results)
	})

	if output != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, output)
	}
}

func TestNewFlipCommand(t *testing.T) {
	cmd := NewFlipCommand()

	// Test default behavior (1 flip)
	output := CaptureOutput(func() {
		cmd.SetArgs([]string{})
		cmd.Execute()
	})
	if !strings.Contains(output, "1 coin flips:") {
		t.Errorf("Expected output to contain '1 coin flips:', got: %s", output)
	}

	// Test with valid number of flips
	output = CaptureOutput(func() {
		cmd.SetArgs([]string{"5"})
		cmd.Execute()
	})
	if !strings.Contains(output, "5 coin flips:") {
		t.Errorf("Expected output to contain '5 coin flips:', got: %s", output)
	}

	// Test with invalid input
	// exitCode := captureExitCode(func() {
	// 	cmd.SetArgs([]string{"invalid"})
	// 	cmd.Execute()
	// })
	// if exitCode != 1 {
	// 	t.Errorf("Expected exit code 1 for invalid input, got: %d", exitCode)
	// }
}

// Helper function to capture output
func CaptureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	done := make(chan struct{})
	var buf bytes.Buffer
	go func() {
		buf.ReadFrom(r)
		close(done)
	}()

	f()

	w.Close()
	<-done
	os.Stdout = old

	return buf.String()
}

// Helper function to capture exit code
// func captureExitCode(f func()) int {
// 	var code int
// 	oldExit := os.Exit
// 	defer func() { os.Exit = oldExit }()
//
// 	os.Exit = func(c int) {
// 		code = c
// 		panic("exit") // Use panic to simulate os.Exit
// 	}
//
// 	defer func() {
// 		if r := recover(); r != nil && r != "exit" {
// 			panic(r) // Re-throw unexpected panics
// 		}
// 	}()
//
// 	f()
// 	return code
// }
