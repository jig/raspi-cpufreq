package cpufreq

import "testing"

func TestBasic(t *testing.T) {
	freq, err := Get()
	if err != nil {
		t.Fatal(err)
	}
	if freq < 600_000 {
		t.Fatal("to low")
	}
	if freq > 10_000_000 {
		t.Fatal("to high")
	}
	t.Logf("freq is %d Hz", freq)
}
