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
	"fmt"
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
	ExpectEq("identical to <[]int> nil", m.Description())

	err = m.Matches(([]int)(nil))
	ExpectEq(nil, err)

	err = m.Matches([]int{})
	ExpectThat(err, Equals("which is not an identical reference"))

	// Non-nil expected value
	o1 := []int{}
	o2 := []int{}
	m = IdenticalTo(o1)
	ExpectEq(fmt.Sprintf("identical to <[]int> %v", o1), m.Description())

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
	ExpectEq("identical to <map[int]int> nil", m.Description())

	err = m.Matches((map[int]int)(nil))
	ExpectEq(nil, err)

	err = m.Matches(map[int]int{})
	ExpectThat(err, Equals("which is not an identical reference"))

	// Non-nil expected value
	o1 := map[int]int{}
	o2 := map[int]int{}
	m = IdenticalTo(o1)
	ExpectEq(fmt.Sprintf("identical to <map[int]int> %v", o1), m.Description())

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
	ExpectEq("identical to <func()> nil", m.Description())

	err = m.Matches((func())(nil))
	ExpectEq(nil, err)

	err = m.Matches(func(){})
	ExpectThat(err, Equals("which is not an identical reference"))

	// Non-nil expected value
	o1 := func() {}
	o2 := func() {}
	m = IdenticalTo(o1)
	ExpectEq(fmt.Sprintf("identical to <func()> %v", o1), m.Description())

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
	ExpectEq("identical to <chan int> nil", m.Description())

	err = m.Matches((chan int)(nil))
	ExpectEq(nil, err)

	err = m.Matches(make(chan int))
	ExpectThat(err, Equals("which is not an identical reference"))

	// Non-nil expected value
	o1 := make(chan int)
	o2 := make(chan int)
	m = IdenticalTo(o1)
	ExpectEq(fmt.Sprintf("identical to <chan int> %v", o1), m.Description())

	err = m.Matches(o1)
	ExpectEq(nil, err)

	err = m.Matches(o2)
	ExpectThat(err, Equals("which is not an identical reference"))
}

func (t *IdenticalToTest) Bools() {
	var m Matcher
	var err error

	// false
	m = IdenticalTo(false)
	ExpectEq("identical to <bool> false", m.Description())

	err = m.Matches(false)
	ExpectEq(nil, err)

	err = m.Matches(true)
	ExpectThat(err, Error(Equals("")))

	// true
	m = IdenticalTo(true)
	ExpectEq("identical to <bool> true", m.Description())

	err = m.Matches(false)
	ExpectThat(err, Error(Equals("")))

	err = m.Matches(true)
	ExpectEq(nil, err)
}

func (t *IdenticalToTest) Ints() {
	var m Matcher
	var err error

	m = IdenticalTo(int(17))
	ExpectEq("identical to <int> 17", m.Description())

	// Identical value
	err = m.Matches(int(17))
	ExpectEq(nil, err)

	// Type alias
	type myType int
	err = m.Matches(myType(17))
	ExpectThat(err, Error(Equals("which is of type myType")))

	// Completely wrong type
	err = m.Matches(int32(17))
	ExpectThat(err, Error(Equals("which is of type int32")))
}

func (t *IdenticalToTest) Int8s() {
	var m Matcher
	var err error

	m = IdenticalTo(int8(17))
	ExpectEq("identical to <int8> 17", m.Description())

	// Identical value
	err = m.Matches(int8(17))
	ExpectEq(nil, err)

	// Type alias
	type myType int8
	err = m.Matches(myType(17))
	ExpectThat(err, Error(Equals("which is of type myType")))

	// Completely wrong type
	err = m.Matches(int32(17))
	ExpectThat(err, Error(Equals("which is of type int32")))
}

func (t *IdenticalToTest) Int16s() {
	var m Matcher
	var err error

	m = IdenticalTo(int16(17))
	ExpectEq("identical to <int16> 17", m.Description())

	// Identical value
	err = m.Matches(int16(17))
	ExpectEq(nil, err)

	// Type alias
	type myType int16
	err = m.Matches(myType(17))
	ExpectThat(err, Error(Equals("which is of type myType")))

	// Completely wrong type
	err = m.Matches(int32(17))
	ExpectThat(err, Error(Equals("which is of type int32")))
}

func (t *IdenticalToTest) Int32s() {
	var m Matcher
	var err error

	m = IdenticalTo(int32(17))
	ExpectEq("identical to <int32> 17", m.Description())

	// Identical value
	err = m.Matches(int32(17))
	ExpectEq(nil, err)

	// Type alias
	type myType int32
	err = m.Matches(myType(17))
	ExpectThat(err, Error(Equals("which is of type myType")))

	// Completely wrong type
	err = m.Matches(int16(17))
	ExpectThat(err, Error(Equals("which is of type int32")))
}

func (t *IdenticalToTest) Int64s() {
	var m Matcher
	var err error

	m = IdenticalTo(int64(17))
	ExpectEq("identical to <int64> 17", m.Description())

	// Identical value
	err = m.Matches(int64(17))
	ExpectEq(nil, err)

	// Type alias
	type myType int64
	err = m.Matches(myType(17))
	ExpectThat(err, Error(Equals("which is of type myType")))

	// Completely wrong type
	err = m.Matches(int32(17))
	ExpectThat(err, Error(Equals("which is of type int32")))
}

func (t *IdenticalToTest) Uints() {
	var m Matcher
	var err error

	m = IdenticalTo(uint(17))
	ExpectEq("identical to <uint> 17", m.Description())

	// Identical value
	err = m.Matches(uint(17))
	ExpectEq(nil, err)

	// Type alias
	type myType uint
	err = m.Matches(myType(17))
	ExpectThat(err, Error(Equals("which is of type myType")))

	// Completely wrong type
	err = m.Matches(int32(17))
	ExpectThat(err, Error(Equals("which is of type int32")))
}

func (t *IdenticalToTest) Uint8s() {
	var m Matcher
	var err error

	m = IdenticalTo(uint8(17))
	ExpectEq("identical to <uint8> 17", m.Description())

	// Identical value
	err = m.Matches(uint8(17))
	ExpectEq(nil, err)

	// Type alias
	type myType uint8
	err = m.Matches(myType(17))
	ExpectThat(err, Error(Equals("which is of type myType")))

	// Completely wrong type
	err = m.Matches(int32(17))
	ExpectThat(err, Error(Equals("which is of type int32")))
}

func (t *IdenticalToTest) Uint16s() {
	var m Matcher
	var err error

	m = IdenticalTo(uint16(17))
	ExpectEq("identical to <uint16> 17", m.Description())

	// Identical value
	err = m.Matches(uint16(17))
	ExpectEq(nil, err)

	// Type alias
	type myType uint16
	err = m.Matches(myType(17))
	ExpectThat(err, Error(Equals("which is of type myType")))

	// Completely wrong type
	err = m.Matches(int32(17))
	ExpectThat(err, Error(Equals("which is of type int32")))
}

func (t *IdenticalToTest) Uint32s() {
	var m Matcher
	var err error

	m = IdenticalTo(uint32(17))
	ExpectEq("identical to <uint32> 17", m.Description())

	// Identical value
	err = m.Matches(uint32(17))
	ExpectEq(nil, err)

	// Type alias
	type myType uint32
	err = m.Matches(myType(17))
	ExpectThat(err, Error(Equals("which is of type myType")))

	// Completely wrong type
	err = m.Matches(int32(17))
	ExpectThat(err, Error(Equals("which is of type int32")))
}

func (t *IdenticalToTest) Uint64s() {
	var m Matcher
	var err error

	m = IdenticalTo(uint64(17))
	ExpectEq("identical to <uint64> 17", m.Description())

	// Identical value
	err = m.Matches(uint64(17))
	ExpectEq(nil, err)

	// Type alias
	type myType uint64
	err = m.Matches(myType(17))
	ExpectThat(err, Error(Equals("which is of type myType")))

	// Completely wrong type
	err = m.Matches(int32(17))
	ExpectThat(err, Error(Equals("which is of type int32")))
}

func (t *IdenticalToTest) Float32s() {
	var m Matcher
	var err error

	m = IdenticalTo(float32(17))
	ExpectEq("identical to <float32> 17", m.Description())

	// Identical value
	err = m.Matches(float32(17))
	ExpectEq(nil, err)

	// Type alias
	type myType float32
	err = m.Matches(myType(17))
	ExpectThat(err, Error(Equals("which is of type myType")))

	// Completely wrong type
	err = m.Matches(float32(17))
	ExpectThat(err, Error(Equals("which is of type float32")))
}

func (t *IdenticalToTest) Float64s() {
	var m Matcher
	var err error

	m = IdenticalTo(float64(17))
	ExpectEq("identical to <float64> 17", m.Description())

	// Identical value
	err = m.Matches(float64(17))
	ExpectEq(nil, err)

	// Type alias
	type myType float64
	err = m.Matches(myType(17))
	ExpectThat(err, Error(Equals("which is of type myType")))

	// Completely wrong type
	err = m.Matches(float64(17))
	ExpectThat(err, Error(Equals("which is of type float64")))
}

func (t *IdenticalToTest) Complex64s() {
	var m Matcher
	var err error

	m = IdenticalTo(complex64(17))
	ExpectEq("identical to <complex64> 17", m.Description())

	// Identical value
	err = m.Matches(complex64(17))
	ExpectEq(nil, err)

	// Type alias
	type myType complex64
	err = m.Matches(myType(17))
	ExpectThat(err, Error(Equals("which is of type myType")))

	// Completely wrong type
	err = m.Matches(complex64(17))
	ExpectThat(err, Error(Equals("which is of type complex64")))
}

func (t *IdenticalToTest) Complex128s() {
	var m Matcher
	var err error

	m = IdenticalTo(complex128(17))
	ExpectEq("identical to <complex128> 17", m.Description())

	// Identical value
	err = m.Matches(complex128(17))
	ExpectEq(nil, err)

	// Type alias
	type myType complex128
	err = m.Matches(myType(17))
	ExpectThat(err, Error(Equals("which is of type myType")))

	// Completely wrong type
	err = m.Matches(complex128(17))
	ExpectThat(err, Error(Equals("which is of type complex128")))
}

func (t *IdenticalToTest) EmptyComparableArrays() {
	var m Matcher
	var err error

	m = IdenticalTo([0]int{})
	ExpectEq("identical to <[0]int> []", m.Description())

	// Identical value
	err = m.Matches([0]int{})
	ExpectEq(nil, err)

	// Length too long
	err = m.Matches([1]int{17})
	ExpectThat(err, Error(Equals("")))

	// Element type alias
	type myType int
	err = m.Matches([0]myType{})
	ExpectThat(err, Error(Equals("which is of type [0]myType")))

	// Completely wrong element type
	err = m.Matches([0]int32{})
	ExpectThat(err, Error(Equals("which is of type [0]int32")))
}

func (t *IdenticalToTest) NonEmptyComparableArrays() {
	var m Matcher
	var err error

	m = IdenticalTo([2]int{17, 19})
	ExpectEq("identical to <[2]int> [17 19]", m.Description())

	// Identical value
	err = m.Matches([2]int{17, 19})
	ExpectEq(nil, err)

	// Length too short
	err = m.Matches([1]int{17})
	ExpectThat(err, Error(Equals("")))

	// Length too long
	err = m.Matches([3]int{17, 19, 23})
	ExpectThat(err, Error(Equals("")))

	// First element different
	err = m.Matches([2]int{13, 19})
	ExpectThat(err, Error(Equals("")))

	// Second element different
	err = m.Matches([2]int{17, 23})
	ExpectThat(err, Error(Equals("")))

	// Element type alias
	type myType int
	err = m.Matches([2]myType{17, 19})
	ExpectThat(err, Error(Equals("which is of type [2]myType")))

	// Completely wrong element type
	err = m.Matches([2]int32{17, 19})
	ExpectThat(err, Error(Equals("which is of type [2]int32")))
}

func (t *IdenticalToTest) NonEmptyArraysOfComparableArrays() {
	var m Matcher
	var err error

	x := [2][2]int{
		[2]int{17, 19},
		[2]int{23, 29},
	}
	m = IdenticalTo(x)
	ExpectEq("identical to <[2][2]int> [[17 19] [23 29]]", m.Description())

	// Identical value
	err = m.Matches([2][2]int{[2]int{17, 19}, [2]int{23, 29}})
	ExpectEq(nil, err)

	// Outer length too short
	err = m.Matches([1][2]int{[2]int{17, 19}})
	ExpectThat(err, Error(Equals("")))

	// Inner length too short
	err = m.Matches([2][1]int{[1]int{17}, [1]int{23}})
	ExpectThat(err, Error(Equals("")))

	// First element different
	err = m.Matches([2][2]int{[2]int{13, 19}, [2]int{23, 29}})
	ExpectThat(err, Error(Equals("")))

	// Element type alias
	type myType int
	err = m.Matches([2][2]myType{[2]myType{17, 19}, [2]myType{23, 29}})
	ExpectThat(err, Error(Equals("which is of type [2][2]myType")))
}

func (t *IdenticalToTest) NonComparableArrays() {
	x := [0]func(){}
	f := func() { IdenticalTo(x) }
	ExpectThat(f, Panics(HasSubstr("is not comparable")))
}

func (t *IdenticalToTest) ArraysOfNonComparableArrays() {
	x := [0][0]func(){}
	f := func() { IdenticalTo(x) }
	ExpectThat(f, Panics(HasSubstr("is not comparable")))
}

func (t *IdenticalToTest) Strings() {
	var m Matcher
	var err error

	m = IdenticalTo("taco")
	ExpectEq("identical to <string> taco", m.Description())

	// Identical value
	err = m.Matches("ta" + "co")
	ExpectEq(nil, err)

	// Type alias
	type myType string
	err = m.Matches(myType("taco"))
	ExpectThat(err, Error(Equals("which is of type myType")))

	// Completely wrong type
	err = m.Matches(int32(17))
	ExpectThat(err, Error(Equals("which is of type int32")))
}

func (t *IdenticalToTest) ComparableStructs() {
	var m Matcher
	var err error

	type subStruct struct {
		i int
	}

	type myStruct struct {
		u uint
		s subStruct
	}

	x := myStruct{17, subStruct{19}}
	m = IdenticalTo(x)
	ExpectEq("identical to <myStruct> {17 {19}}", m.Description())

	// Identical value
	err = m.Matches(myStruct{17, subStruct{19}})
	ExpectEq(nil, err)

	// Wrong outer field
	err = m.Matches(myStruct{13, subStruct{19}})
	ExpectThat(err, Error(Equals("")))

	// Wrong inner field
	err = m.Matches(myStruct{17, subStruct{23}})
	ExpectThat(err, Error(Equals("")))

	// Type alias
	type myType myStruct
	err = m.Matches(myType{17, subStruct{19}})
	ExpectThat(err, Error(Equals("which is of type myType")))

	// Completely wrong type
	err = m.Matches(int32(17))
	ExpectThat(err, Error(Equals("which is of type int32")))
}

func (t *IdenticalToTest) NonComparableStructs() {
	type subStruct struct {
		s []int
	}

	type myStruct struct {
		u uint
		s subStruct
	}

	x := myStruct{17, subStruct{[]int{19}}}
	f := func() { IdenticalTo(x) }
	ExpectThat(f, Panics(AllOf(HasSubstr("IdenticalTo"), HasSubstr("comparable"))))
}

func (t *IdenticalToTest) UnsafePointers() {
}

func (t *IdenticalToTest) IntAlias() {
}
