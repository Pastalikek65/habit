# r/termux — Showcase Post Draft for habit

**Status: DRAFT — copy-paste to Reddit. Use Showcase flair, Tue-Thu 14-18 UTC.**

## Title

```
[Showcase] Built habit on my phone in Termux — TUI habit tracker with streaks (6.8M, SQLite WAL, offline)
```

## Body

I built `habit` on my phone in Termux because habit apps need internet and ads. You get a terminal tracker you can check offline, with `🔥` streaks.

![demo](https://raw.githubusercontent.com/Pastalikek65/habit/main/demo.gif)

**Why Termux-related:**
- Built in Termux proot Debian, `CGO_ENABLED=0` pure Go `modernc.org/sqlite` WAL `0700`
- 6.8M binary, <30MB RAM, `go build` 3s on phone
- Single file `~/.cache/habit/habit.db`, survives reboot, no daemon
- `habits` + `checks` tables, `Streak` counts consecutive days from today

**Install (Termux):**

```bash
git clone https://github.com/Pastalikek65/habit.git
cd habit
CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o habit .
./habit add "drink water"
./habit check 1
./habit tui   # j/k nav, c check today, q quit
```

**Usage:**

```bash
habit add "run"
habit list              # 1  drink water  🔥 1
habit check 1
habit tui
```

Tech: Go 1.25, `bubbletea`, `lipgloss`, `cobra`. Repo: https://github.com/Pastalikek65/habit — MIT. Tested on Termux aarch64.

---

Checklist: flair `Showcase`, demo gif loads, no "please star", keep <300 words.
