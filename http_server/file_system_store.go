package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

// file_system_store.go
func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, io.SeekStart)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var Wins int

	for _, Player := range f.GetLeague() {
		if Player.Name == name {
			Wins = Player.Wins
			break
		}
	}
	return Wins
}
