package main;

import (
	"fmt"
	"time"
	"runtime"
	"math/rand"
	"sync"
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

	var total_inside float64
	for i := 0; i < cores; i++ {
		total_inside += <-results
	}

	return total_inside / float64(cores)
}

var total int = 0

func MonteCarloBolsa(points int) float64 {
	cores := runtime.NumCPU()

	// var wg sync.WaitGroup
	var mutex sync.Mutex	

	results := make(chan float64, cores)

	for core := 0; core < cores; core++ {
		go func() {
			var inside int = 0
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			for {
				mutex.Lock()
				if total == points {
					mutex.Unlock()
					break
				}
				total++
				mutex.Unlock()
				x, y := r.Float64(), r.Float64()

				if x * x + y * y <= 1 {
					inside++
				}
			}
			results <- float64(inside)
		}()
	}

	var total_inside float64
	for i := 0; i < cores; i++ {
		total_inside += <-results
	}

	return 4 * float64(total_inside) / float64(total)
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())
}

func main(){
	var SAMPLES int = 1000000

	start_concurrency := time.Now()
	monte_carlo_concurrency := MonteCarlo(SAMPLES)
	end_concurrency := time.Now()
	fmt.Println("Valor de pi no algoritmo concorrente:", monte_carlo_concurrency)
	fmt.Println("Tempo concorrente:",  end_concurrency.Sub(start_concurrency))

	fmt.Println()

	start_pack := time.Now()
	monte_carlo_pack := MonteCarloBolsa(SAMPLES)
	end_pack := time.Now()
	fmt.Println("Valor de pi no algoritmo concorrente utilizando bolsa de tarefas:", monte_carlo_pack)
	fmt.Println("Tempo concorrente com bolsa de tarefas:",  end_pack.Sub(start_pack))

	fmt.Println()

	start_sequential := time.Now()
	monte_carlo_sequential := Sequential(SAMPLES)
	end_sequential := time.Now()
	fmt.Println("Valor de pi no algoritmo sequencial:", monte_carlo_sequential)
	fmt.Println("Tempo sequencial:", end_sequential.Sub(start_sequential))

	}