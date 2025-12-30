package entity

import "time"

type Todo struct {
	ID          string
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}
