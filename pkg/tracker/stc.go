package tracker

import (
	"github.com/spacemonkeygo/monotime"
	"time"
)

// TODO Record the current monotonic raw clock immediately (if no time is provided)
// TODO Lock to mux clock
// TODO try different PLL models
// TODO validate STC is smooth, no big jitter
// TODO validate freq diff is stable
// TODO support pcap, instead of get current time, use the time from pcap

func RawTime() time.Duration {
	return monotime.Monotonic()
}