# Contributing to habit

```bash
git clone https://github.com/Pastalikek65/habit.git
cd habit
go vet ./...
go test ./... -count=1 -timeout 30s -cover
CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o habit .
./habit add "drink water" && ./habit check 1 && ./habit list
```

PRs: fork → feat/name → test → open PR.
