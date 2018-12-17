package tracker

import (
	"testing"
	"time"
)

func TestRawTime(t *testing.T) {
	rawTimeFirst := RawTime().Nanoseconds()
	goTimeFirst := time.Now()
	t.Logf("rawTimeFirst = %v", rawTimeFirst)
	t.Logf("goTimeFirst = %v", goTimeFirst)
	time.Sleep(1000 * time.Millisecond)
	t.Logf("rawTime diff = %v", float64(RawTime().Nanoseconds() - rawTimeFirst)/1000000000.0)
	t.Logf("goTime diff = %v", float64(time.Now().Sub(goTimeFirst).Nanoseconds())/1000000000.0)
	t.Logf("rawTime diff = %v", float64(RawTime().Nanoseconds() - rawTimeFirst)/1000000000.0)
	t.Logf("goTime diff = %v", float64(time.Now().Sub(goTimeFirst).Nanoseconds())/1000000000.0)

}

func BenchmarkRawTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RawTime().Nanoseconds()
	}
}

func BenchmarkGoTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Now()
	}
}