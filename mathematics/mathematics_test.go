// Copyright (C) 2022   Anurag Dhadse. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mathematics

import (
	"testing"
)

func TestLenOfString(t *testing.T) {
	var val int
	if val = LenOfString(1234); val != 4 {
		t.Error("Method returned value other than 4 for input 1234:", val)
	}
	if val = LenOfString(-1234); val != 4 {
		t.Error("Method returned value other than 4 for input -1234:", val)
	}
	if val = LenOfString(0); val != 1 {
		t.Error("Method returned value other than 1 for input 0:", val)
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
		t.Error("Method returned false for input \"aibohphobia\"")
	}
	if val = Palindrome("string"); val != false {
		t.Error("Method returned true for input \"string\"")
	}
}

func BenchmarkPalindrome(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Palindrome("aibohphobia")
	}
}

func TestGCD(t *testing.T) {
	if val := GCD(-123, -32); val != 1 {
		t.Error("Method returned value other than 1 for input (-123, -32):", val)
	}
	if val := GCD(12, 18); val != 6 {
		t.Error("Method returned value other than 6 for input (12, 18):", val)
	}
	if val := GCD(-4, 14); val != 2 {
		t.Error("Method returned value other than 2 for input (-4, 14):", val)
	}
}

func BenchmarkGCD(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		GCD(3, 9)
	}
}

func TestLCM(t *testing.T) {
	if val := LCM(2, 3); val != 6 {
		t.Error("Method returned value other than 6 for input (2, 3):", val)
	}
	if val := LCM(-2, 3); val != 6 {
		t.Error("Method returned value other than 6 for input (-2, 3):", val)
	}
	if val := LCM(2, -3); val != 6 {
		t.Error("Method returned value other than 6 for input (2, -3):", val)
	}
	if val := LCM(-2, -3); val != 6 {
		t.Error("Method returned value other than 6 for input (-2, -3):", val)
	}
}

func BenchmarkLCM(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		LCM(3, 9)
	}
}

func TestFactorial(t *testing.T) {
	if val := Factorial(1); val != 1 {
		t.Error("Method returned value other than 1 for input 1:", val)
	}
	if val := Factorial(0); val != 1 {
		t.Error("Method returned value other than 1 for input 0:", val)
	}
	if val := Factorial(5); val != 120 {
		t.Error("Method returned value other than 120 for input 5:", val)
	}
	if val := Factorial(2); val != 2 {
		t.Error("Method returned value other than 2 for input 2:", val)
	}
	if val := Factorial(3); val != 6 {
		t.Error("Method returned value other than 6 for input 3:", val)
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Factorial(5)
	}
}

func TestIsPrime(t *testing.T) {
	if val := IsPrime(13); val != true {
		t.Error("Method returned false for input 13")
	}
	if val := IsPrime(-1); val != false {
		t.Error("Method returned true for input -1")
	}
	if val := IsPrime(0); val != false {
		t.Error("Method returned true for input 0")
	}
	if val := IsPrime(1); val != false {
		t.Error("Method returned true for input 0")
	}
	if val := IsPrime(3); val != true {
		t.Error("Method returned false for input 0")
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		IsPrime(13)
	}
}

func TestPrimeFactors(t *testing.T) {
	if val := PrimeFactors(315); !testEq(val, []int{3, 5, 7}) {
		t.Error("Function returned values other [3 5 7] for input 315:", val)
	}
	if val := PrimeFactors(2); !testEq(val, []int{}) {
		t.Error("Function returned values other [] for input 2:", val)
	}
	if val := PrimeFactors(4); !testEq(val, []int{2}) {
		t.Error("Function returned values other [2] for input 4:", val)
	}
	if val := PrimeFactors(6); !testEq(val, []int{2, 3}) {
		t.Error("Function returned values other [2, 3] for input 6:", val)
	}
	if val := PrimeFactors(15); !testEq(val, []int{3, 5}) {
		t.Error("Function returned values other [3, 5] for input 15:", val)
	}
	if val := PrimeFactors(-10); !testEq(val, []int{}) {
		t.Error("Function returned values other [] for input -10:", val)
	}
}

func BenchmarkPrimeFactors(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		PrimeFactors(315)
	}
}
