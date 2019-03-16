package main

import "time"

type Benchmark struct {
	Format string
	ReqMethod string
	Iterations int
	AvgExecTime float32
	StartTime time.Time
	EndTime time.Time
	TotalExecTime time.Duration
}
