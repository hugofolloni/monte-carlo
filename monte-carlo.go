package main;

import (
	"fmt"
	"math/rand"
	"time"
	"runtime"
)

func Sequential (samples int) float64 {
	var inside int = 0

	for i := 0; i < samples; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if (x*x + y*y) < 1 {
			inside++
		}
	}

	ratio := float64(inside) / float64(samples)

	return ratio * 4
}

func MonteCarlo(samples int) float64 {
	cpus := runtime.NumCPU()

	threadSamples := samples / cpus
	results := make(chan float64, cpus)

	for j := 0; j < cpus; j++ {
		go func() {
			var inside int
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			for i := 0; i < threadSamples; i++ {
				x, y := r.Float64(), r.Float64()

				if x*x+y*y <= 1 {
					inside++
				}
			}
			results <- float64(inside) / float64(threadSamples) * 4
		}()
	}

	var total float64
	for i := 0; i < cpus; i++ {
		total += <-results
	}

	return total / float64(cpus)
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())
}

func main(){
	var SAMPLES int = 10000

	start_concurrency := time.Now()
	monte_carlo_concurrency := MonteCarlo(SAMPLES)
	end_concurrency := time.Now()
	fmt.Println("Valor de pi no algoritmo concorrente:", monte_carlo_concurrency)
	fmt.Println("Tempo concorrente:",  end_concurrency.Sub(start_concurrency))

	fmt.Println()

	start_sequential := time.Now()
	monte_carlo_sequential := Sequential(SAMPLES)
	end_sequential := time.Now()
	fmt.Println("Valor de pi no algoritmo sequencial:", monte_carlo_sequential)
	fmt.Println("Tempo sequencial:", end_sequential.Sub(start_sequential))

}