# restless (skeleton)

Single-binary, CLI-first API client with two faces:

- **TUI**: polished terminal UI (wizard + request + stream)
- **GUI**: will wrap the same TUI in a native window (next milestone)

## Run

```bash
go mod tidy
go run ./cmd/restless
```

## Modes

```bash
restless --mode tui
restless --mode gui
restless --quiet
```

Keys:
- `tab` / `shift+tab`: switch tabs
- `q`: quit
