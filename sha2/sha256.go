// Copyright © 2019 Weald Technology Trading
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

// Package sha2 provides hashing using the sha2 system.
package sha2

import (
	"crypto/sha256"
)

const _256hashlength = 32

// SHA256 is the 256-bit SHA2 hashing method.
type SHA256 struct{}

// New256 creates a new 256-bit SHA2 hashing method.
func New256() *SHA256 {
	return &SHA256{}
}

// HashLength returns the length of hashes generated by Hash() in bytes.
func (*SHA256) HashLength() int {
	return _256hashlength
}

// HashName returns the name of this hash.
func (*SHA256) HashName() string {
	return "sha256"
}

// Hash generates a SHA2 hash from input byte arrays.
func (*SHA256) Hash(data ...[]byte) []byte {
	var hash [_256hashlength]byte
	if len(data) == 1 {
		hash = sha256.Sum256(data[0])
	} else {
		concatDataLen := 0
		for _, d := range data {
			concatDataLen += len(d)
		}
		concatData := make([]byte, concatDataLen)
		curOffset := 0
		for _, d := range data {
			copy(concatData[curOffset:], d)
			curOffset += len(d)
		}
		hash = sha256.Sum256(concatData)
	}

	return hash[:]
}
