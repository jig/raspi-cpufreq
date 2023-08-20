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
	log.Printf("CPU frequency is %.1f GHz", float64(freq)/1_000_000)
}
