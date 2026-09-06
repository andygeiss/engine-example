# SPEC

**Job:** Show how to build a small game in Go with
[ecs](https://github.com/andygeiss/ecs), as a program someone can install and run in
one command.

**Why:** The ecs README explains the idea; this repository is where a reader sees it
work. An example that only runs from its own directory, or that ignores the
[baseline](https://github.com/andygeiss/baseline), teaches the wrong habits along with
the right library.

**Guardrails:**

- The game stays small. It is read more often than it is played, so a feature that
  makes the code harder to follow is not worth it.
- Two dependencies, ecs and raylib-go, and no more. The raylib one is justified in
  [README.md](README.md).
- Named decisions and any waived baseline rule live in [README.md](README.md).
- No release binaries. The channel is `go install`.

**Done means:**

- `checklists/cli-tool.md` in the [baseline](https://github.com/andygeiss/baseline) is
  walked, and every box is checked or waived on the record in the README.
- `make ci` is green on the commit being pushed.
- `go install github.com/andygeiss/engine-example@latest` gives a binary that opens
  the window from any directory.
