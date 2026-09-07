<p align="center">
<img src="https://github.com/andygeiss/ecs/blob/master/logo.png?raw=true" />
</p>

# Example Engine

[![License](https://img.shields.io/github/license/andygeiss/engine-example)](https://github.com/andygeiss/engine-example/blob/main/LICENSE)

A small game that shows how to build an engine with
[ecs](https://github.com/andygeiss/ecs). It is for anyone who has read that library's
README and wants to see the pieces fit together: entities carry components, systems do
the work, and the engine calls the systems over and over.

## Install and play

```bash
go install github.com/andygeiss/engine-example@latest
engine-example
```

A window opens. **WASD** moves the logo, **ESC** closes it. The sprites are built into
the binary, so it runs from any directory.

Working on the code instead?

```bash
git clone https://github.com/andygeiss/engine-example.git
cd engine-example
make run     # start it
make         # run every gate before you commit
```

## Options

Every flag has an environment variable, and the flag wins:

```
  -height string
    	window height in pixels (env ENGINE_EXAMPLE_HEIGHT) (default "600")
  -title string
    	window title (env ENGINE_EXAMPLE_TITLE) (default "Example Engine")
  -v	print raylib's own log, which it writes to stdout
  -version
    	print the version and exit
  -width string
    	window width in pixels (env ENGINE_EXAMPLE_WIDTH) (default "800")
```

## How it is put together

```
config.go              ← the Config struct and its parser
go.mod
internal/
├── components/        ← the data: position, size, state, texture, velocity
├── game/              ← the world: which entities exist, which systems run
│   ├── assets.go      ← //go:embed of the sprites beside it
│   ├── game.go
│   └── resources/     ← the two sprites, built into the binary
└── systems/           ← the behaviour: input, movement, collision, state, rendering
main.go                ← wiring only: parse, build the engine, run it
Makefile               ← the command runner, copied from the baseline
README.md
SPEC.md                ← this project's brief: job, why, guardrails, done means
```

Each component owns one bit of a mask, so a system finds its work with a single
`FilterByMask` call. `internal/game` is the only place that says which entities the game
starts with, and it carries the sprites they are drawn with; `main.go` hands the two
managers to `ecs.NewDefaultEngine` and runs it until the window closes or Ctrl-C
arrives.

Three systems never touch the graphics card — movement, collision, and state — which is
why they are the ones with tests.

## Why raylib

[stack/go.md](https://github.com/andygeiss/baseline/blob/main/stack/go.md) says every
dependency outside its approved list needs a written justification. This one is
[gen2brain/raylib-go](https://github.com/gen2brain/raylib-go), and drawing a window is
the reason: the standard library has no graphics, and a game engine example with nothing
on screen would not be an example of anything. Since v0.60.0 it builds without cgo — it
loads an embedded raylib through purego — so `CGO_ENABLED=0 go build` still produces one
static binary, exactly as the baseline wants.

## Baseline deviations

This project follows the
[engineering baseline](https://github.com/andygeiss/baseline), specifically
`project-types/cli-tool.md` and `checklists/cli-tool.md`. Where it does not, it says so
here.

**Waived**

- **No interactivity — no prompts, no TUI**
  ([project-types/cli-tool.md](https://github.com/andygeiss/baseline/blob/main/project-types/cli-tool.md)
  *Architecture defaults*) — waived 2026-09-06 by Andy. A game reads the keyboard while
  it runs; that is the thing being demonstrated. Contained: the window is the only
  interactive surface, the terminal gets no prompt, and every setting still arrives as a
  flag or an environment variable.
- **stdout carries data only**
  ([checklists/cli-tool.md](https://github.com/andygeiss/baseline/blob/main/checklists/cli-tool.md)
  *The command-line contract*) — waived 2026-09-06 by Andy. raylib writes its own log to
  stdout and v0.60.1 exports no callback to redirect it. Contained: the log level is set
  to warnings before the window opens, so a normal run writes nothing to stdout at all;
  `-v` turns the log on for whoever asks for it.
- **Assets embedded from a root-level `assets.go`**
  ([patterns/go-project-layout.md](https://github.com/andygeiss/baseline/blob/main/patterns/go-project-layout.md)
  rule 5) — waived 2026-09-06 by Andy. That rule puts the embed at the module root
  because `//go:embed` cannot reach `../web` from `cmd/server`; this module's `main`
  package already sits at the root, so nothing forces the file up there, and the sprites
  read better beside the only package that draws them. Contained:
  `internal/game/assets.go` embeds `internal/game/resources`, the variable stays
  unexported because nothing outside that package needs it, and `main.go` no longer
  touches a file system at all.
- **`run()` is table-tested: happy path**
  ([checklists/cli-tool.md](https://github.com/andygeiss/baseline/blob/main/checklists/cli-tool.md)
  *Writing a test*) — waived 2026-09-06 by Andy. The single command opens a window and
  blocks until someone closes it, and there is no headless raylib to run that against.
  Contained: `run`'s other paths are table-tested, and `internal/game` builds the whole
  world and system list in a test without a window, so everything `run` wires is covered
  up to `engine.Setup`.

**Conformance, by a different route**

- The one binary this module ships lives in the `main` package at the module root, so
  there is no `cmd/` directory to look for.
- The tool keeps no state between runs, so "partial work is safe" is met by there being
  no partial work: it reads two embedded files and writes nothing.

**Unexercised**

- No secrets, no HTTP, no database, no LLM. Those checklist sections never fire.

## Known gaps

Small bugs, kept out of the migration on purpose so the diff stayed one thing:

- Walking left or up leaves the window for good. `CollisionSystem` wraps at the right and
  bottom edge only.
- Diagonal movement is impossible, and holding two keys gives whichever the system checks
  last: D beats S beats A beats W.
