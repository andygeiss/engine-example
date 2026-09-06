package main

import (
	"errors"
	"strings"
	"testing"
)

// run is the whole command line. Opening a window needs a screen, so the
// happy path is not here — internal/game and internal/systems cover what it
// wires. These are the paths that return before any window exists.
func TestRun(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantErr    error
		wantStdout string
		wantStderr string
	}{
		{
			name:       "-version prints the version to stdout",
			args:       []string{"-version"},
			wantStdout: version(),
		},
		{
			name:       "-h prints the usage to stderr and succeeds",
			args:       []string{"-h"},
			wantStderr: "-width",
		},
		{
			name:       "an unknown flag is a usage error",
			args:       []string{"-bogus"},
			wantErr:    errUsage,
			wantStderr: "-bogus",
		},
		{
			name:       "a bad width is a usage error",
			args:       []string{"-width", "0"},
			wantErr:    errUsage,
			wantStderr: "want a whole number of pixels above 0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var stdout, stderr strings.Builder
			err := run(t.Context(), tt.args, &stdout, &stderr)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error = %v, want %v", err, tt.wantErr)
			}
			if got := strings.TrimSpace(stdout.String()); got != strings.TrimSpace(tt.wantStdout) {
				t.Errorf("stdout = %q, want %q", got, tt.wantStdout)
			}
			if !strings.Contains(stderr.String(), tt.wantStderr) {
				t.Errorf("stderr %q does not mention %q", stderr.String(), tt.wantStderr)
			}
		})
	}
}
