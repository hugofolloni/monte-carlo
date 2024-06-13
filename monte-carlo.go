package main;

import (
	"fmt"
	"time"
	"runtime"
	"math/rand"
)

func Sequential (points int) float64 {
	var inside int = 0

	for i := 0; i < points; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if x * x + y * y < 1 {
			inside++
		}
	}

	return 4 * float64(inside) / float64(points)
}

func MonteCarlo(points int) float64 {
	cores := runtime.NumCPU()

	sample := points / cores
	results := make(chan float64, cores)

	for core := 0; core < cores; core++ {
		go func() {
			var inside int = 0
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			for point := 0; point < sample; point++ {
				x, y := r.Float64(), r.Float64()

				if x * x + y * y <= 1 {
					inside++
				}
			}
			results <- 4 * float64(inside) / float64(sample) 
		}()
	}

	var total float64
	for i := 0; i < cores; i++ {
		total += <-results
	}

	return total / float64(cores)
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