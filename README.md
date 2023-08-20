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
	log.Printf("CPU frequency is %.1f GHz", float64(freq)/1_000_000)
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
2023/08/20 18:01:05 CPU frequency is 0.6 GHz
```
