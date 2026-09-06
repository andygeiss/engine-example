package main

import (
	"errors"
	"flag"
	"io"
	"strings"
	"testing"
)

func TestParseConfig(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		env     map[string]string
		want    Config
		wantErr error
	}{
		{
			name: "no arguments gives a working 800x600 window",
			want: Config{Width: 800, Height: 600, Title: "Example Engine"},
		},
		{
			name: "flags win",
			args: []string{"-width", "1024", "-height", "768", "-title", "Test"},
			want: Config{Width: 1024, Height: 768, Title: "Test"},
		},
		{
			name: "the environment is the default when there is no flag",
			env:  map[string]string{"ENGINE_EXAMPLE_WIDTH": "640", "ENGINE_EXAMPLE_TITLE": "From the environment"},
			want: Config{Width: 640, Height: 600, Title: "From the environment"},
		},
		{
			name: "a flag beats the environment",
			args: []string{"-width", "1280"},
			env:  map[string]string{"ENGINE_EXAMPLE_WIDTH": "640"},
			want: Config{Width: 1280, Height: 600, Title: "Example Engine"},
		},
		{
			name: "-version is a flag like any other",
			args: []string{"-version"},
			want: Config{Width: 800, Height: 600, Title: "Example Engine", Version: true},
		},
		{
			name: "-v asks raylib for its log",
			args: []string{"-v"},
			want: Config{Width: 800, Height: 600, Title: "Example Engine", Verbose: true},
		},
		{
			name:    "an unknown flag is a usage error",
			args:    []string{"-bogus"},
			wantErr: errUsage,
		},
		{
			name:    "a width that is not a number is a usage error",
			args:    []string{"-width", "wide"},
			wantErr: errUsage,
		},
		{
			name:    "a width of zero is a usage error",
			args:    []string{"-width", "0"},
			wantErr: errUsage,
		},
		{
			name:    "a negative height is a usage error",
			args:    []string{"-height", "-10"},
			wantErr: errUsage,
		},
		{
			name:    "a bad environment variable fails like a bad flag",
			env:     map[string]string{"ENGINE_EXAMPLE_HEIGHT": "tall"},
			wantErr: errUsage,
		},
		{
			name:    "-h asks for help",
			args:    []string{"-h"},
			wantErr: flag.ErrHelp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.env {
				t.Setenv(k, v)
			}
			got, err := parseConfig(tt.args, io.Discard)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error = %v, want %v", err, tt.wantErr)
			}
			if tt.wantErr == nil && got != tt.want {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

// A bad value names the setting and says what to type instead, so the reader
// does not have to guess. The message goes to stderr, never to stdout.
func TestParseConfigExplainsABadValue(t *testing.T) {
	var stderr strings.Builder
	if _, err := parseConfig([]string{"-width", "nope"}, &stderr); !errors.Is(err, errUsage) {
		t.Fatalf("error = %v, want errUsage", err)
	}
	got := stderr.String()
	for _, want := range []string{"width", `"nope"`, "want a whole number of pixels above 0"} {
		if !strings.Contains(got, want) {
			t.Errorf("stderr %q does not mention %q", got, want)
		}
	}
}
