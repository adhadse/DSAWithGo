// Copyright 2022 The DSAWithGo Authors. All Rights Reserved.
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
// limitations under the License.package backtracking

package bit_magic

import (
	"testing"
)

func TestAnd(t *testing.T) {
	if val := And(5, 3); val != 1 {
		t.Errorf("Function returned value other than 1 (1) for AND of 5 and 3: %b (%d)", val, val)
	}
	if val := And(25, 30); val != 24 {
		t.Errorf("Function returned value other than 11000 (24) for AND of 25 and 30: %b (%d)", val, val)
	}
	if val := And(25, 30); val == 16 {
		t.Errorf("Function returned value other than 11000 (24) for AND of 25 and 30: %b (%d)", val, val)
	}
	if val := And(-25, -30); val != -30 {
		t.Errorf("Function returned value other than -11110 (-30) for AND of -25 and -30: %b (%d)", val, val)
	}
}

func TestOr(t *testing.T) {
	if val := Or(60, 13); val != 61 {
		t.Errorf("Function returned value other than 111101 (61) for OR of 60 and 13: %b (%d)", val, val)
	}
	if val := Or(60, 13); val == 64 {
		t.Errorf("Function returned value other than 111101 (61) for OR of 60 and 13: %b (%d)", val, val)
	}
	if val := Or(-10, -12); val != -10 {
		t.Errorf("Function returned value other than -1010 (-10) for OR of -10 and -12: %b (%d)", val, val)
	}
}

func TestXOr(t *testing.T) {
	if val := XOr(-25, -30); val != 5 {
		t.Errorf("Function returned value other than 101 (5) for XOR of -25 and -30: %b (%d)", val, val)
	}
	if val := XOr(10, 12); val != 6 {
		t.Errorf("Function returned value other than 110 (6) for XOR of 10 and 12: %b (%d)", val, val)
	}
	if val := XOr(25, 30); val != 7 {
		t.Errorf("Function returned value other than 111 (7) for XOR of 25 and 30: %b (%d)", val, val)
	}
}

func TestComplement(t *testing.T) {
	if val := Complement(60); val != -61 {
		t.Errorf("Function returned value other than -111101 (-61) for Complement of 60: %b (%d)", val, val)
	}
	if val := Complement(61); val == -60 {
		t.Errorf("Function returned value other than -111100 (-60) for Complement of -61: %b (%d)", val, val)
	}
}

func TestLeftShift(t *testing.T) {
	if val := LeftShift(60, 2); val != 240 {
		t.Errorf("Function returned value other than 11110000 (240) for Leftshift of 60 by 2 bits: %b (%d)", val, val)
	}
}

func TestRightShift(t *testing.T) {
	if val := RightShift(60, 2); val != 15 {
		t.Errorf("Function returned value other than 1111 (15) for Leftshift of 60 by 2 bits: %b (%d)", val, val)
	}
}

func TestCountSetBit(t *testing.T) {
	if val := CountSetBit(10); val != 2 {
		t.Errorf("Function returned value other than 2 for CountSetBit of 10: %d", val)
	}
	if val := CountSetBit(0); val != 0 {
		t.Errorf("Function returned value other than 0 for CountSetBit of 0: %d", val)
	}
	if val := CountSetBit(1); val != 1 {
		t.Errorf("Function returned value other than 1 for CountSetBit of 1: %d", val)
	}
	if val := CountSetBit(60); val != 4 {
		t.Errorf("Function returned value other than 4 for CountSetBit of 60: %d", val)
	}
}

func TestCountSetbitBrianKerningham(t *testing.T) {
	if val := CountSetbitBrianKerningham(10); val != 2 {
		t.Errorf("Function returned value other than 2 for CountSetBit of 10: %d", val)
	}
	if val := CountSetbitBrianKerningham(0); val != 0 {
		t.Errorf("Function returned value other than 0 for CountSetBit of 0: %d", val)
	}
	if val := CountSetbitBrianKerningham(1); val != 1 {
		t.Errorf("Function returned value other than 1 for CountSetBit of 1: %d", val)
	}
	if val := CountSetbitBrianKerningham(60); val != 4 {
		t.Errorf("Function returned value other than 4 for CountSetBit of 60: %d", val)
	}
}

func TestCountSetBitLookUpTable(t *testing.T) {
	if val := CountSetBitLookUpTable(10); val != 2 {
		t.Errorf("Function returned value other than 2 for CountSetBit of 10: %d", val)
	}
	if val := CountSetBitLookUpTable(0); val != 0 {
		t.Errorf("Function returned value other than 0 for CountSetBit of 0: %d", val)
	}
	if val := CountSetBitLookUpTable(1); val != 1 {
		t.Errorf("Function returned value other than 1 for CountSetBit of 1: %d", val)
	}
	if val := CountSetBitLookUpTable(60); val != 4 {
		t.Errorf("Function returned value other than 4 for CountSetBit of 60: %d", val)
	}
}
