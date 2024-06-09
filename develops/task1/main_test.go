package main

import (
	"testing"
	"time"
)

func Test_exactTime(t *testing.T) {
	ntpTime := exactTime()
	correctTime := time.Now()

	if int(correctTime.Sub(ntpTime)) > correctTime.Second() {
		t.Errorf("time Time - %q, NTP Time - %q", correctTime, ntpTime)
	}
}
