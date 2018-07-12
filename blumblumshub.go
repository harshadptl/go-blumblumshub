package blumblumshub


var p uint64 = 5651
var q uint64 = 5623
var M uint64 = p * q


/* ChangePrimes is used to set custom values for primes p & q 
used in the algorithm. It throws an error if the primes are not
congruent to 3mod4 as part of the specification to keep this 
algorithm strong and unbreakable.
*/
func ChangePrimes(p1, q1 uint64) {
    if (p1 % 4 != 3) || (p2 % 4 != 3){
		errors.New("Prime not valid")
		return
	}

	p = p1
	q = q1
	
	return
}

/*
gcd returns the GCD of two numbers by the Euclidean algorithm.
*/
func gcd(a, b uint64) uint64 {
	for a != b {
		if(a > b) {
			a = a - b;
		} else {
			b = b - a;
		}
	}
	return a;
}
