package entity

import "time"

type Resposta struct {
	TempoRequest time.Duration
	StatusCode int
}