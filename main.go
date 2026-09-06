// Command engine-example is a small game that shows how to build an engine
// with the ecs package: entities carry components, systems do the work.
//
// It opens a window. WASD moves the logo, ESC closes it.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/andygeiss/ecs"
	"github.com/andygeiss/engine-example/internal/game"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	err := run(ctx, os.Args[1:], os.Stdout, os.Stderr)
	switch {
	case err == nil:
	case errors.Is(err, errUsage):
		os.Exit(2) // message already printed where the error was detected
	default:
		fmt.Fprintf(os.Stderr, "engine-example: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string, stdout, stderr io.Writer) error {
	cfg, err := parseConfig(args, stderr)
	switch {
	case errors.Is(err, flag.ErrHelp):
		return nil // -h: usage already printed
	case err != nil:
		return err
	}
	if cfg.Version {
		fmt.Fprintln(stdout, version())
		return nil
	}

	g := game.New(cfg.Width, cfg.Height, cfg.Title, cfg.Verbose)
	engine := ecs.NewDefaultEngine(g.Entities, g.Systems)
	engine.Setup()
	defer engine.Teardown()
	engine.Run(ctx)

	// Run stops for three reasons and only the first is success: the window
	// was closed, a system failed, or Ctrl-C cancelled the context.
	if err := g.Err(); err != nil {
		return err
	}
	return ctx.Err()
}

// version reports the tag this binary was built from. The toolchain stamps it,
// so there is no ldflags ceremony.
func version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}
	if v := info.Main.Version; v != "" && v != "(devel)" {
		return v // go install @version, or VCS-derived (Go 1.24+)
	}
	return "unknown" // no VCS metadata to fall back on
}
