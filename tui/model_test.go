package tui

import (
	"testing"

	"github.com/Pastalikek65/habit/store"
)

func TestView(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/habit.db")
	defer s.Close()
	s.Add("drink water")
	m := New(s)
	v := m.View()
	if v == "" { t.Fatal("empty") }
}
