package store

import (
	"testing"
	"time"
)

func TestAddAndList(t *testing.T) {
	s, _ := Open(t.TempDir() + "/habit.db")
	defer s.Close()
	id, err := s.Add("drink water")
	if err != nil { t.Fatal(err) }
	if id == 0 { t.Fatal("no id") }
	hs, _ := s.List()
	if len(hs) != 1 || hs[0].Name != "drink water" { t.Fatalf("list %v", hs) }
}

func TestCheckAndStreak(t *testing.T) {
	s, _ := Open(t.TempDir() + "/habit.db")
	defer s.Close()
	id, _ := s.Add("run")
	today := time.Now().Format("2006-01-02")
	if err := s.Check(id, today); err != nil { t.Fatal(err) }
	if streak, _ := s.Streak(id); streak != 1 { t.Fatalf("want 1 got %d", streak) }
	// duplicate check same day should not increase
	_ = s.Check(id, today)
	if streak, _ := s.Streak(id); streak != 1 { t.Fatalf("dup should stay 1 got %d", streak) }
	// yesterday + today = 2 streak
	yesterday := time.Now().AddDate(0,0,-1).Format("2006-01-02")
	s2, _ := Open(t.TempDir() + "/habit2.db")
	defer s2.Close()
	id2, _ := s2.Add("run2")
	s2.Check(id2, yesterday)
	s2.Check(id2, today)
	if streak, _ := s2.Streak(id2); streak != 2 { t.Fatalf("want 2 got %d", streak) }
}

func TestCheckInvalid(t *testing.T) {
	s, _ := Open(t.TempDir() + "/habit.db")
	defer s.Close()
	if err := s.Check(999, "2026-01-01"); err == nil { t.Fatal("want error for missing habit") }
}
