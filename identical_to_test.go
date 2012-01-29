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
	var err error

	type intAlias int

	// Type alias expected value
	m = IdenticalTo(intAlias(17))
	err = m.Matches(int(17))
	ExpectThat(err, Error(Equals("which is of type int")))

	// Type alias candidate
	m = IdenticalTo(int(17))
	err = m.Matches(intAlias(17))
	ExpectThat(err, Error(Equals("which is of type intAlias")))

	// int and uint
	m = IdenticalTo(int(17))
	err = m.Matches(uint(17))
	ExpectThat(err, Error(Equals("which is of type uint")))
}

func (t *IdenticalToTest) InvalidTypeExpectedValue() {
	f := func() { IdenticalTo(nil) }
	ExpectThat(f, Panics(AllOf(HasSubstr("IdenticalTo"), HasSubstr("invalid"))))
}

func (t *IdenticalToTest) InvalidTypeCandidate() {
	var m Matcher
	var err error

	// Nil chan expected value
	m = IdenticalTo((chan int)(nil))
	err = m.Matches(nil)
	ExpectThat(err, Error(Equals("which is of type <nil>")))

	// Non-nil chan expected value
	m = IdenticalTo(make(chan int))
	err = m.Matches(nil)
	ExpectThat(err, Error(Equals("which is of type <nil>")))
}

func (t *IdenticalToTest) Slices() {
	var m Matcher
	var err error

	// Nil expected value
	m = IdenticalTo(([]int)(nil))

	err = m.Matches(([]int)(nil))
	ExpectEq(nil, err)

	err = m.Matches([]int{})
	ExpectThat(err, Equals("which is not an identical reference"))

	// Non-nil expected value
	o1 := []int{}
	o2 := []int{}
	m = IdenticalTo(o1)

	err = m.Matches(o1)
	ExpectEq(nil, err)

	err = m.Matches(o2)
	ExpectThat(err, Equals("which is not an identical reference"))
}

func (t *IdenticalToTest) Maps() {
	var m Matcher
	var err error

	// Nil expected value
	m = IdenticalTo((map[int]int)(nil))

	err = m.Matches((map[int]int)(nil))
	ExpectEq(nil, err)

	err = m.Matches(map[int]int{})
	ExpectThat(err, Equals("which is not an identical reference"))

	// Non-nil expected value
	o1 := map[int]int{}
	o2 := map[int]int{}
	m = IdenticalTo(o1)

	err = m.Matches(o1)
	ExpectEq(nil, err)

	err = m.Matches(o2)
	ExpectThat(err, Equals("which is not an identical reference"))
}

func (t *IdenticalToTest) Functions() {
	var m Matcher
	var err error

	// Nil expected value
	m = IdenticalTo((func())(nil))

	err = m.Matches((func())(nil))
	ExpectEq(nil, err)

	err = m.Matches(func(){})
	ExpectThat(err, Equals("which is not an identical reference"))

	// Non-nil expected value
	o1 := func() {}
	o2 := func() {}
	m = IdenticalTo(o1)

	err = m.Matches(o1)
	ExpectEq(nil, err)

	err = m.Matches(o2)
	ExpectThat(err, Equals("which is not an identical reference"))
}

func (t *IdenticalToTest) Channels() {
	var m Matcher
	var err error

	// Nil expected value
	m = IdenticalTo((chan int)(nil))

	err = m.Matches((chan int)(nil))
	ExpectEq(nil, err)

	err = m.Matches(make(chan int))
	ExpectThat(err, Equals("which is not an identical reference"))

	// Non-nil expected value
	o1 := make(chan int)
	o2 := make(chan int)
	m = IdenticalTo(o1)

	err = m.Matches(o1)
	ExpectEq(nil, err)

	err = m.Matches(o2)
	ExpectThat(err, Equals("which is not an identical reference"))
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
