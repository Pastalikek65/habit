package store

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"
)

type Habit struct {
	ID      int64
	Name    string
	Created string
}

type Store struct {
	db *sql.DB
}

func Open(path string) (*Store, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite", path)
	if err != nil { return nil, err }
	if _, err := db.Exec("PRAGMA journal_mode=WAL"); err != nil { db.Close(); return nil, err }
	if _, err := db.Exec("PRAGMA busy_timeout=5000"); err != nil { db.Close(); return nil, err }
	_ = os.Chmod(path, 0600)
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS habits(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT UNIQUE, created TEXT)`); err != nil { return nil, err }
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS checks(habit_id INTEGER, date TEXT, PRIMARY KEY(habit_id, date))`); err != nil { return nil, err }
	return &Store{db: db}, nil
}

func (s *Store) Close() error { return s.db.Close() }

func (s *Store) Add(name string) (int64, error) {
	if name == "" { return 0, fmt.Errorf("empty name") }
	now := time.Now().Format(time.RFC3339)
	res, err := s.db.Exec("INSERT INTO habits(name, created) VALUES(?,?)", name, now)
	if err != nil { return 0, err }
	return res.LastInsertId()
}

func (s *Store) List() ([]Habit, error) {
	rows, err := s.db.Query("SELECT id, name, created FROM habits ORDER BY id")
	if err != nil { return nil, err }
	defer rows.Close()
	var out []Habit
	for rows.Next() {
		var h Habit
		if err := rows.Scan(&h.ID, &h.Name, &h.Created); err != nil { return nil, err }
		out = append(out, h)
	}
	return out, rows.Err()
}

func (s *Store) Check(id int64, date string) error {
	// verify habit exists
	var c int
	if err := s.db.QueryRow("SELECT COUNT(*) FROM habits WHERE id=?", id).Scan(&c); err != nil { return err }
	if c == 0 { return fmt.Errorf("habit %d not found", id) }
	if _, err := time.Parse("2006-01-02", date); err != nil { return fmt.Errorf("invalid date %s", date) }
	_, err := s.db.Exec("INSERT OR IGNORE INTO checks(habit_id, date) VALUES(?,?)", id, date)
	return err
}

func (s *Store) Streak(id int64) (int, error) {
	rows, err := s.db.Query("SELECT date FROM checks WHERE habit_id=? ORDER BY date DESC", id)
	if err != nil { return 0, err }
	defer rows.Close()
	var dates []string
	for rows.Next() {
		var d string
		rows.Scan(&d)
		dates = append(dates, d)
	}
	if len(dates)==0 { return 0, nil }
	// count consecutive from today backwards
	today := time.Now().Format("2006-01-02")
	// build set for quick lookup
	set := map[string]bool{}
	for _, d := range dates { set[d] = true }
	streak := 0
	cur, _ := time.Parse("2006-01-02", today)
	for {
		ds := cur.Format("2006-01-02")
		if !set[ds] { break }
		streak++
		cur = cur.AddDate(0,0,-1)
	}
	return streak, nil
}
