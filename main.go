package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/Pastalikek65/habit/store"
	"github.com/Pastalikek65/habit/tui"
)

var version = "0.1.0"

func defaultDB() string {
	if xdg := os.Getenv("XDG_CACHE_HOME"); xdg != "" {
		return filepath.Join(xdg, "habit", "habit.db")
	}
	home, _ := os.UserHomeDir()
	if home == "" { home = os.Getenv("HOME") }
	if home == "" { home = "/tmp" }
	return filepath.Join(home, ".cache", "habit", "habit.db")
}

func main() {
	root := &cobra.Command{Use: "habit", Short: "TUI habit tracker for Termux", Version: version}
	root.AddCommand(&cobra.Command{
		Use: "add [name]", Args: cobra.ExactArgs(1), Short: "add habit",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, _ := store.Open(defaultDB())
			defer s.Close()
			id, err := s.Add(args[0])
			if err != nil { return err }
			fmt.Println(id)
			return nil
		},
	})
	root.AddCommand(&cobra.Command{
		Use: "list", Short: "list habits",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, _ := store.Open(defaultDB())
			defer s.Close()
			hs,_:=s.List()
			for _,h:=range hs {
				streak,_:=s.Streak(h.ID)
				fmt.Printf("%d\t%s\t🔥 %d\n", h.ID, h.Name, streak)
			}
			return nil
		},
	})
	root.AddCommand(&cobra.Command{
		Use: "check [id]", Args: cobra.ExactArgs(1), Short: "check today",
		RunE: func(cmd *cobra.Command, args []string) error {
			var id int64
			fmt.Sscan(args[0], &id)
			s,_:=store.Open(defaultDB())
			defer s.Close()
			return s.Check(id, time.Now().Format("2006-01-02"))
		},
	})
	root.AddCommand(&cobra.Command{
		Use: "tui", Short: "open TUI",
		RunE: func(cmd *cobra.Command, args []string) error {
			s,_:=store.Open(defaultDB())
			defer s.Close()
			m:=tui.New(s)
			p:=tea.NewProgram(m, tea.WithAltScreen())
			_,err:=p.Run()
			return err
		},
	})
	if err:=root.Execute(); err!=nil { os.Exit(1) }
}
