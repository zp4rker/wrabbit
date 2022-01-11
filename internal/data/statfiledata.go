package data

import "time"

type StatfileData struct {
	Id      string
	Running bool
	Args    []string
	ProcId  int

	StartDate time.Time
	EndDate   time.Time
}
