package main

func integers() chan uint64 {
	out := make(chan uint64)

	go func() {
		var i uint64 = 2
		for {
			out <- i
			i++
		}
	}()

	return out
}

func filter(prime uint64, in chan uint64) chan uint64 {
	out := make(chan uint64)

	go func() {
		for i := range in {
			if (i % prime) != 0 {
				out <- i
			}
		}
	}()

	return out
}

func Primes() chan uint64 {
	out := make(chan uint64)
	in := integers()
	go func() {
		for {
			i := <-in
			out <- i
			in = filter(i, in)
		}
	}()

	return out
}

var primes []uint64 = nil
var pgen = Primes()

func fillPrimes(n uint64) {
	for {
		if len(primes) > 0 && primes[len(primes)-1] > n {
			return
		}
		p := <-pgen
		primes = append(primes, p)
	}
}

func search(needle uint64, haystack []uint64) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}

	return false
}

func IsPrime(n uint64) bool {
	fillPrimes(n)
	return search(n, primes)
}
