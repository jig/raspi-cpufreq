package main

import (
	"log"

	cpufreq "github.com/jig/raspi-cpufreq"
)

func main() {
	freq, err := cpufreq.Get()
	if err != nil {
		log.Fatal(err)
	}
	maxFreq, err := cpufreq.Max()
	if err != nil {
		log.Fatal(err)
	}
	minFreq, err := cpufreq.Min()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CPU frequency is %.1f GHz (range from %.1f GHz to %.1f GHz)", float64(freq)/1_000_000, float64(minFreq)/1_000_000, float64(maxFreq)/1_000_000)
}
