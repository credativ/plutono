package vali

import "time"

type ValiQuery struct {
	Expr         string
	Step         time.Duration
	LegendFormat string
	Start        time.Time
	End          time.Time
	RefId        string
}
