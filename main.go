package main

import (
	"flag"
	"log"
)

func main() {
	var start uint
	var end uint
	flag.UintVar(&start, "start", 0, "usage")
	flag.UintVar(&end, "end", 10, "usage")
	flag.Parse()
	primes := make([]uint, end)
	isPrimes := make([]bool, end)
	for prime := range primes {
		primes[prime] = uint(prime)
		isPrimes[prime] = true
	}
	for prime := range primes {
		if prime < 2 {
			continue
		}
		if isPrimes[primes[prime]] {
			if primes[prime] >= start {
				log.Println(prime)
			}
			go func() {
				for i := prime * 2; i < int(end); i += prime {
					isPrimes[i] = false
				}
			}()
		}
	}
}
