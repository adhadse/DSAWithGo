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
// limitations under the License.

package recursion

// TailRecursive is a special form of recursion where
// the last operation of a function is a recursive call.
func TailRecursive(n, a int) int {
	if n == 0 {
		return a
	}
	return TailRecursive(n-1, n*a)
}

// NonTailRecursive is not tail recursive
// because the value returned by NonTailRecursive(n-1)
// is used in NonTailRecursive(n) and call to NonTailRecursive(n-1)
// is not the last thing done by NonTailRecursive(n)
func NonTailRecursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * NonTailRecursive(n-1)
}
