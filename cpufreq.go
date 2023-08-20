package cpufreq

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const cpuFreqPath = "/sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq"

// Get returns the current frequency of the Raspberry Pi CPU
func Get() (int, error) {
	raw, err := os.ReadFile(cpuFreqPath)
	if err != nil {
		return 0, fmt.Errorf("error reading %q: %v", cpuFreqPath, err)
	}

	cpuFreq, err := strconv.Atoi(strings.TrimSpace(string(raw)))
	if err != nil {
		return 0, errors.New("invalid format")
	}
	return cpuFreq, nil
}
