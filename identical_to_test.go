// Copyright 2012 Aaron Jacobs. All Rights Reserved.
// Author: aaronjjacobs@gmail.com (Aaron Jacobs)
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

package oglematchers_test

import (
	. "github.com/jacobsa/oglematchers"
	. "github.com/jacobsa/ogletest"
)

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type IdenticalToTest struct {
}

func init() { RegisterTestSuite(&IdenticalToTest{}) }

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *IdenticalToTest) TypesNotIdentical() {
	var m Matcher
	var res bool
	var err error

	type intAlias int

	// Type alias expected value
	m = IdenticalTo(intAlias(17))
	res, err = m.Matches(int(17))

	ExpectFalse(res)
	ExpectThat(err, Error(Equals("which is of type int")))

	// Type alias candidate
	m = IdenticalTo(int(17))
	res, err = m.Matches(intAlias(17))

	ExpectFalse(res)
	ExpectThat(err, Error(Equals("which is of type intAlias")))

	// int and uint
	m = IdenticalTo(int(17))
	res, err = m.Matches(uint(17))

	ExpectFalse(res)
	ExpectThat(err, Error(Equals("which is of type uint")))
}

func (t *IdenticalToTest) NilExpectedValue() {
}

func (t *IdenticalToTest) NilCandidate() {
}

func (t *IdenticalToTest) Slices() {
}

func (t *IdenticalToTest) Maps() {
}

func (t *IdenticalToTest) Functions() {
}

func (t *IdenticalToTest) Channels() {
}

func (t *IdenticalToTest) Bools() {
}

func (t *IdenticalToTest) Ints() {
}

func (t *IdenticalToTest) Int8s() {
}

func (t *IdenticalToTest) Int16s() {
}

func (t *IdenticalToTest) Int32s() {
}

func (t *IdenticalToTest) Int64s() {
}

func (t *IdenticalToTest) Uints() {
}

func (t *IdenticalToTest) Uint8s() {
}

func (t *IdenticalToTest) Uint16s() {
}

func (t *IdenticalToTest) Uint32s() {
}

func (t *IdenticalToTest) Uint64s() {
}

func (t *IdenticalToTest) Float32s() {
}

func (t *IdenticalToTest) Float64s() {
}

func (t *IdenticalToTest) Complex64s() {
}

func (t *IdenticalToTest) Complex128s() {
}

func (t *IdenticalToTest) ComparableArrays() {
}

func (t *IdenticalToTest) NonComparableArrays() {
}

func (t *IdenticalToTest) ComparableInterfaces() {
}

func (t *IdenticalToTest) NonComparableInterfaces() {
}

func (t *IdenticalToTest) Pointers() {
}

func (t *IdenticalToTest) Strings() {
}

func (t *IdenticalToTest) ComparableStructs() {
}

func (t *IdenticalToTest) NonComparableStructs() {
}

func (t *IdenticalToTest) UnsafePointers() {
}

func (t *IdenticalToTest) IntAlias() {
}
