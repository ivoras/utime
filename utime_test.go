package utime

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

func Test00Now(t *testing.T) {
	unow := Now()
	now := time.Now()
	if unow != Time(now.Unix()) {
		t.Error("utime.Now() and time.Now() disagree")
	}
	fmt.Println("Sizeof time.Time:", unsafe.Sizeof(now), "Sizeof utime.Time:", unsafe.Sizeof(unow))
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

func BenchmarkUSince(b *testing.B) {
	t0 := Now()
	for i := 0; i < b.N; i++ {
		Since(t0)
	}
}

func BenchmarkSince(b *testing.B) {
	t0 := time.Now()
	for i := 0; i < b.N; i++ {
		time.Since(t0)
	}
}

func BenchmarkUSub(b *testing.B) {
	t0 := Now()
	t1 := Now()
	for i := 0; i < b.N; i++ {
		t1.Sub(t0)
	}
}

func BenchmarkSub(b *testing.B) {
	t0 := time.Now()
	t1 := time.Now()
	for i := 0; i < b.N; i++ {
		t1.Sub(t0)
	}
}
