package main;

import (
	"fmt"
	"time"
	"runtime"
	"math/rand"
	"sync"
	"math"
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

func MonteCarlo(points int, threads int) float64 {
	cores := threads

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


func MonteCarloBolsa(points int, threads int) float64 {
	var total int = 0
	cores := threads

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


func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num * output)) / output
}

func main(){
	samples := [3]int {100, 100000, 10000000}
	threads:=  [4]int {1, 2, 4, 8}
	for sample := 0; sample < 3; sample++{
		for thread := 0; thread < 4; thread++ {	
			SAMPLES := samples[sample]
			THREADS := threads[thread]
			var c1_time int64 = 0
			var c2_time int64 = 0
			var sequencial_time int64 = 0
			for i := 0; i < 10; i++ {
				start_concurrency := time.Now()
				MonteCarlo(SAMPLES, THREADS)
				end_concurrency := time.Now()
				start_pack := time.Now()
				MonteCarloBolsa(SAMPLES, THREADS)
				end_pack := time.Now()
				start_sequential := time.Now()
				Sequential(SAMPLES)
				end_sequential := time.Now()
				c1_time += (end_concurrency.Sub(start_concurrency)).Microseconds()
				c2_time +=  (end_pack.Sub(start_pack)).Microseconds()
				sequencial_time += (end_sequential.Sub(start_sequential)).Microseconds()
			}

			c1_time /= 10
			c2_time /= 10
			sequencial_time /= 10

			c1 := float64(c1_time) / 1000
			c2 := float64(c2_time) / 1000
			sq := float64(sequencial_time) / 1000
			var acc_1 float64 = float64(sequencial_time) /  float64(c1_time)
			var acc_2 float64 = float64(sequencial_time) /  float64(c2_time)

			fmt.Println(THREADS, "&", c1, "&", toFixed(acc_1, 3), "&",  c2, "&", toFixed(acc_2, 3), "&", sq, "\\\\ \\hline")

		}
		
		fmt.Println()
	}
}