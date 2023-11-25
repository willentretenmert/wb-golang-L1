package main

import (
	"fmt"
	"strings"
)

// Интерфейс нового музыкального плеера
type MusicPlayer interface {
	PlaySong(song string)
}

// Старый плеер
type OldMusicPlayer struct{}

// Старый метод воспроизведения
func (p *OldMusicPlayer) PlayFile(file string) {
	fmt.Println("Воспроизведение (старый формат):", file)
}

// Адаптер адаптирует ( =) ) OldMusicPlayer к новому интерфейсу MusicPlayer, реализую через структуру
type MusicPlayerAdapter struct {
	player *OldMusicPlayer
}

func (adapter *MusicPlayerAdapter) PlaySong(song string) {
	// Конвертируем название песни в формат, совместимый со старым плеером
	file := strings.ReplaceAll(song, " ", "_") + ".mp3"
	adapter.player.PlayFile(file)
}
