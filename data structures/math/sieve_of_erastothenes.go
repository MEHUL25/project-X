package math

// SieveOfEratosthenes function to find all prime numbers up to 'n'
func SieveOfEratosthenes(n int) []int {
	// Step 1: Create a boolean array "isPrime" of size (n+1) and initialize all entries as true
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// Step 2: Apply the sieve algorithm
	for p := 2; p*p <= n; p++ {
		if isPrime[p] {
			// Mark all multiples of p as false
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}

	// Step 3: Collect all prime numbers
	primes := []int{}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
