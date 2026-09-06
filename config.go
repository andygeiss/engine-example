package main

import (
	"cmp"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

// errUsage means the command line was wrong. The message was already printed
// where the problem was found, so main only has to pick the exit code.
var errUsage = errors.New("usage error")

// Config is every knob this binary has. After parseConfig returns, nothing
// else reads os.Getenv — the struct is the whole contract.
type Config struct {
	Height  int
	Title   string
	Verbose bool // let raylib print its own log
	Version bool // print the version and exit, instead of opening a window
	Width   int
}

// parseConfig turns the command line into a Config. It returns flag.ErrHelp
// for -h, and errUsage for anything the caller has to fix.
func parseConfig(args []string, stderr io.Writer) (Config, error) {
	fs := flag.NewFlagSet("engine-example", flag.ContinueOnError)
	fs.SetOutput(stderr)

	var c Config
	// Width and height arrive as strings so a bad environment variable fails
	// the same way a bad flag does, with the same message.
	width := fs.String("width", cmp.Or(os.Getenv("ENGINE_EXAMPLE_WIDTH"), "800"), "window width in pixels (env ENGINE_EXAMPLE_WIDTH)")
	height := fs.String("height", cmp.Or(os.Getenv("ENGINE_EXAMPLE_HEIGHT"), "600"), "window height in pixels (env ENGINE_EXAMPLE_HEIGHT)")
	fs.StringVar(&c.Title, "title", cmp.Or(os.Getenv("ENGINE_EXAMPLE_TITLE"), "Example Engine"), "window title (env ENGINE_EXAMPLE_TITLE)")
	fs.BoolVar(&c.Verbose, "v", false, "print raylib's own log, which it writes to stdout")
	fs.BoolVar(&c.Version, "version", false, "print the version and exit")

	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return Config{}, err // -h: usage printed, exit 0
		}
		return Config{}, errUsage // fs already printed the message and the usage
	}

	var err error
	if c.Width, err = pixels(*width); err != nil {
		fmt.Fprintf(stderr, "engine-example: width %s\n", err)
		return Config{}, errUsage
	}
	if c.Height, err = pixels(*height); err != nil {
		fmt.Fprintf(stderr, "engine-example: height %s\n", err)
		return Config{}, errUsage
	}
	return c, nil
}

// pixels reads a window edge length. The error says what to type instead, and
// quotes what was given, so a stray space shows up.
func pixels(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil || n <= 0 {
		return 0, fmt.Errorf("%q: want a whole number of pixels above 0", s)
	}
	return n, nil
}
