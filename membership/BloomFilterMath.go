package membership

import "math"

//Calculate numberOfHashFunctions(K)
func numberOfHashFunctions(falsePositiveRate float64) int {
	return int(math.Ceil(math.Log2(1.0 / falsePositiveRate)))
}

//Calculate bitVectorSize(M)
func bitVectorSize(capacity int, falsePositiveRate float64) int {
	//ln22 = ln2^2
	ln22 := math.Pow(math.Ln2, 2)
	return int(float64(capacity) * math.Abs(math.Log(falsePositiveRate)) / ln22)
}
