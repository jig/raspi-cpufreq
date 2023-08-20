# raspi-cpufreq

Minimal package to read ARM CPU frequency of the Raspberry Pi 4.


```go
import (
	"log"

	cpufreq "github.com/jig/raspi-cpufreq"
)

func main() {
	freq, err := cpufreq.Get()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CPU frequency is %.1f GHz (range from %.1f GHz to %.1f GHz)", float64(freq)/1_000_000, float64(minFreq)/1_000_000, float64(maxFreq)/1_000_000)
}

```

# Command line

```bash
cd cmd/freq
go build
go install

freq
```

Example output:

```bash
2023/08/20 18:14:59 CPU frequency is 1.5 GHz (range from 0.6 GHz to 1.5 GHz)
```
