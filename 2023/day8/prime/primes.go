package prime

import "fmt"

type Primes struct {
	Primes []int
}

// create a new Primes object with all primes up to max
func NewPrimes(max int) *Primes {
	primes := &Primes{Primes: []int{2, 3, 5, 7, 11, 13}}
	for i := 17; i <= max; i += 2 {

		hasFactor := false
		for _, p := range primes.Primes {
			if i%p == 0 {
				hasFactor = true
				break
			}
		}

		if !hasFactor {
			primes.Primes = append(primes.Primes, i)
		}
	}
	fmt.Printf("Found %d primes up to %d\n", len(primes.Primes), max)
	return primes
}

type Factor struct {
	Prime int
	Count int
}

// get the prime factors (sorted) of a number
func (primes *Primes) GetFactors(number int) []Factor {
	factors := []Factor{}
	for _, p := range primes.Primes {
		count := 0
		for number%p == 0 {
			count++
			number /= p
		}
		if count > 0 {
			factors = append(factors, Factor{Prime: p, Count: count})
		}
	}
	return factors
}

// merge prime factors of two numbers (GCD)
func MergeFactors(factors1 []Factor, factors2 []Factor) []Factor {
	factors := []Factor{}

	// factores are sorted by prime
	i1, i2 := 0, 0
	for i1 < len(factors1) && i2 < len(factors2) {
		f1, f2 := factors1[i1], factors2[i2]
		if f1.Prime == f2.Prime {
			factors = append(factors, Factor{Prime: f1.Prime, Count: max(f1.Count, f2.Count)})
			i1++
			i2++
		} else if f1.Prime < f2.Prime {
			factors = append(factors, f1)
			i1++
		} else {
			factors = append(factors, f2)
			i2++
		}
	}
	for ; i1 < len(factors1); i1++ {
		factors = append(factors, factors1[i1])
	}
	for ; i2 < len(factors2); i2++ {
		factors = append(factors, factors2[i2])
	}
	return factors
}

func MultiplyFactors(factors []Factor) int {
	product := 1
	for _, factor := range factors {
		for i := 0; i < factor.Count; i++ {
			product *= factor.Prime
		}
	}
	return product
}

func (primes *Primes) GetGreatestCommonDivisor(numbers []int) int {
	factors := []Factor{}
	for _, number := range numbers {
		factors = MergeFactors(factors, primes.GetFactors(number))
	}
	return MultiplyFactors(factors)
}
