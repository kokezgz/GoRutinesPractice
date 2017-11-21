package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type PrimeResult struct {
	number int64
	prime  bool
}

func primes(min int64, max int64) []int64 {
	var primeNums []int64
	var res PrimeResult

	channel := make(chan PrimeResult)
	defer close(channel)

	go fireCalculation(min, max, channel)

	for i := min; i <= max; i++ {
		res = <-channel

		if res.prime {
			fmt.Printf("%d prime. \n", res.number)
			primeNums = append(primeNums, res.number)
		}
	}

	return primeNums
}

func fireCalculation(min int64, max int64, channel chan PrimeResult) {
	var i int64
	for i = min; i <= max; i++ {
		go isPrimeAsyc(i, channel)
	}
}

func isPrimeAsyc(number int64, channel chan PrimeResult) {
	time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

	result := new(PrimeResult)
	result.number = number
	result.prime = isPrime(number)
	channel <- *result
}

func isPrime(candidate int64) bool {
	var i, limit int64

	if candidate == 2 {
		return true
	}

	if math.Mod(float64(candidate), 2) == 0 {
		return false
	}

	limit = int64(math.Ceil(math.Sqrt(float64(candidate))))
	for i = 3; i <= limit; i += 2 { //Only odd numbers
		if math.Mod(float64(candidate), float64(i)) == 0 {
			return false
		}
	}
	return true
}

func main() {
	primes(1, 1000)
}
