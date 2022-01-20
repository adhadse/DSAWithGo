package Mathematics

import (
	"testing"
)

func TestLenOfString(t *testing.T) {
	var val int
	if val = LenOfString(1234567890); val != 10 {
		t.Error("Method returned value other than 10:", val)
	}
}

func BenchmarkLenOfString(b *testing.B) {
	testVal := 1234567890
	for i := 0; i <= b.N; i++ {
		LenOfString(testVal)
	}
}

func TestPalindrome(t *testing.T) {
	var val bool
	if val = Palindrome("aibohphobia"); val != true {
		t.Error("Method returned false")
	}
}

func BenchmarkPalindrome(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Palindrome("aibohphobia")
	}
}

func TestGCD(t *testing.T) {
	if val := GCD(3, 9); val != 3 {
		t.Error("Method returned value other than 3:", val)
	}
}

func BenchmarkGCD(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		GCD(3, 9)
	}
}

func TestFactorial(t *testing.T) {
	if val := Factorial(5); val != 120 {
		t.Error("Method returned value other than 120:", val)
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Factorial(5)
	}
}

func TestIsPrime(t *testing.T) {
	if val := IsPrime(13); val != true {
		t.Error("Method returned value other than true")
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		IsPrime(13)
	}
}

func TestPrimeFactors(t *testing.T) {
	if val := PrimeFactors(315); !testEq(val, []int{3, 5, 7}) {
		t.Error("Function returned values other [3 5 7]:", val)
	}
}

func BenchmarkPrimeFactors(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		PrimeFactors(315)
	}
}
