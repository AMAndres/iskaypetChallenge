package domain

import "time"

type Pet struct {
	ID       int64
	Name     string
	Species  int8
	Gender   int8
	Birthday time.Time
}
