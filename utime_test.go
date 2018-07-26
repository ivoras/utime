package utime

import (
	"testing"
	"time"
)

func Test00Now(t *testing.T) {
	unow := Now()
	now := time.Now()
	if unow != Time(now.Unix()) {
		t.Error("utime.Now() and time.Now() disagree")
	}
}

func BenchmarkUNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Now()
	}
}

func BenchmarkNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Now()
	}
}
