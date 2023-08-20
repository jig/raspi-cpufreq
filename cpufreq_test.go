package cpufreq

import "testing"

func TestGet(t *testing.T) {
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

func TestMax(t *testing.T) {
	freq, err := Max()
	if err != nil {
		t.Fatal(err)
	}
	if freq < 600_000 {
		t.Fatal("to low")
	}
	if freq > 10_000_000 {
		t.Fatal("to high")
	}
	t.Logf("max freq is %d Hz", freq)
}

func TestMin(t *testing.T) {
	freq, err := Min()
	if err != nil {
		t.Fatal(err)
	}
	if freq < 600_000 {
		t.Fatal("to low")
	}
	if freq > 10_000_000 {
		t.Fatal("to high")
	}
	t.Logf("min freq is %d Hz", freq)
}
