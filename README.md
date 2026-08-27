# habit

[![Go 1.25](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Termux](https://img.shields.io/badge/Termux-friendly-brightgreen?logo=android)](https://termux.dev)

Single-binary habit tracker for Termux — SQLite WAL, `bubbletea` TUI, 6.8M.

I built `habit` because phone habit apps need internet and ads. This works offline, shows streaks, and lives in your terminal.

```bash
go install github.com/Pastalikek65/habit@latest
habit add "drink water"
habit check 1
habit list              # 1  drink water  🔥 1
habit tui               # j/k nav, c check, q quit
```

Single file `~/.cache/habit/habit.db` WAL `0700`, survives reboot, no daemon.

![demo](https://raw.githubusercontent.com/Pastalikek65/habit/main/demo.gif)

## Why

Termux kills GUI habit apps. `habit` keeps streaks in SQLite and shows `🔥` in TUI. No `cron` needed — check when you open terminal.

## Install (Termux)

```bash
git clone https://github.com/Pastalikek65/habit.git
cd habit
CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o habit .
./habit add "run"
./habit tui
```

## How it works

`store` → `habits(id,name,created)` + `checks(habit_id,date)` WAL. `Streak` counts consecutive days from today via `checks` set. `tui` `bubbletea` shows `habit` + `🔥`.

## Dev

```bash
go vet ./...
go test ./... -count=1 -timeout 30s -cover
CGO_ENABLED=0 go build -o habit .
```

## License

MIT
