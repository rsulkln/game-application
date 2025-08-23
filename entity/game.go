package entity

import "time"

type Game struct {
	ID          uint
	Category    uint
	QuestionIDs []uint
	Players     []uint
	WinnerID    uint
	StartTime   time.Time
}

type Player struct {
	ID     uint
	UserID uint
	GameID uint
	Score  uint
	Answer []PlayerAnswer
}

type PlayerAnswer struct {
	ID         uint
	PlayerID   uint
	QuestionID uint
}
