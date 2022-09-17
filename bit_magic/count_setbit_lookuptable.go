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

package bit_magic

func CountSetBitLookUpTable(a int) int {
	bitsettable := make([]int, 256)
	for i := 0; i < 256; i++ {
		bitsettable[i] = (i & 1) + bitsettable[i/2]
	}

	return bitsettable[a&0xFF] +
		bitsettable[(a>>8)&0xFF] +
		bitsettable[(a>>16)&0xFF] +
		bitsettable[(a>>24&0xFF)]
}
