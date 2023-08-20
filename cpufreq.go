package cpufreq

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	cpuFreqPath    = "/sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq"
	cpuMaxFreqPath = "/sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_max_freq"
	cpuMinFreqPath = "/sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_min_freq"
)

// Get returns the current frequency of the Raspberry Pi CPU
func Get() (int, error) {
	return get(cpuFreqPath)
}

// Max returns the maximum frequency of the Raspberry Pi CPU
func Max() (int, error) {
	return get(cpuMaxFreqPath)
}

// Min returns the minimum frequency of the Raspberry Pi CPU
func Min() (int, error) {
	return get(cpuMinFreqPath)
}

func get(path string) (int, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("error reading %q: %v", cpuFreqPath, err)
	}

	cpuFreq, err := strconv.Atoi(strings.TrimSpace(string(raw)))
	if err != nil {
		return 0, errors.New("invalid format")
	}
	return cpuFreq, nil
}
