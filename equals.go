// Copyright 2011 Aaron Jacobs. All Rights Reserved.
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

package oglematchers

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

// Equals(x) returns a matcher that matches values v such that any of the
// following hold:
//
//  1. v and x are of the same type, and v == x.
//
//  2. v's type is an alias for x's type or vice versa, and when casted to the
//     basic type v and x are equal.
//
//  3. v and x are both of numeric types, and their values are equivalent.
//
// If you want a stricter matcher that works just like the built-in == operator
// (including rejecting even slightly different types), see StrictEquals
// instead.
func Equals(x interface{}) Matcher {
	v := reflect.ValueOf(x)

	// The == operator is not defined for array or struct types.
	if v.Kind() == reflect.Array || v.Kind() == reflect.Struct {
		panic(fmt.Sprintf("oglematchers.Equals: unsupported kind %v", v.Kind()))
	}

	// The == operator is not defined for non-nil slices.
	if v.Kind() == reflect.Slice && v.Pointer() != uintptr(0) {
		panic(fmt.Sprintf("oglematchers.Equals: non-nil slice"))
	}

	return &equalsMatcher{v}
}

type equalsMatcher struct {
	expectedValue reflect.Value
}

////////////////////////////////////////////////////////////
// Numeric types
////////////////////////////////////////////////////////////

func isSignedInteger(v reflect.Value) bool {
	k := v.Kind()
	return k >= reflect.Int && k <= reflect.Int64
}

func isUnsignedInteger(v reflect.Value) bool {
	k := v.Kind()
	return k >= reflect.Uint && k <= reflect.Uint64
}

func isInteger(v reflect.Value) bool {
	return isSignedInteger(v) || isUnsignedInteger(v)
}

func isFloat(v reflect.Value) bool {
	k := v.Kind()
	return k == reflect.Float32 || k == reflect.Float64
}

func isComplex(v reflect.Value) bool {
	k := v.Kind()
	return k == reflect.Complex64 || k == reflect.Complex128
}

func checkAgainstInt64(e int64, c reflect.Value) (res bool, err error) {
	res = false
	err = errors.New("")

	switch {
	case isSignedInteger(c):
		if c.Int() == e {
			res = true
			err = nil
		}

	case isUnsignedInteger(c):
		u := c.Uint()
		if u <= math.MaxInt64 && int64(u) == e {
			res = true
			err = nil
		}

	// Turn around the various floating point types so that the checkAgainst*
	// functions for them can deal with precision issues.
	case isFloat(c), isComplex(c):
		return Equals(c.Interface()).Matches(e)

	default:
		res = false
		err = NewFatalError("which is not numeric")
	}

	return
}

func checkAgainstUint64(e uint64, c reflect.Value) (res bool, err error) {
	res = false
	err = errors.New("")

	switch {
	case isSignedInteger(c):
		i := c.Int()
		if i >= 0 && uint64(i) == e {
			res = true
			err = nil
		}

	case isUnsignedInteger(c):
		if c.Uint() == e {
			res = true
			err = nil
		}

	// Turn around the various floating point types so that the checkAgainst*
	// functions for them can deal with precision issues.
	case isFloat(c), isComplex(c):
		return Equals(c.Interface()).Matches(e)

	default:
		res = false
		err = NewFatalError("which is not numeric")
	}

	return
}

func checkAgainstFloat32(e float32, c reflect.Value) (res bool, err error) {
	res = false
	err = errors.New("")

	switch {
	case isSignedInteger(c):
		if float32(c.Int()) == e {
			res = true
			err = nil
		}

	case isUnsignedInteger(c):
		if float32(c.Uint()) == e {
			res = true
			err = nil
		}

	case isFloat(c):
		// Compare using float32 to avoid a false sense of precision; otherwise
		// e.g. Equals(float32(0.1)) won't match float32(0.1).
		if float32(c.Float()) == e {
			res = true
			err = nil
		}

	case isComplex(c):
		comp := c.Complex()
		rl := real(comp)
		im := imag(comp)

		// Compare using float32 to avoid a false sense of precision; otherwise
		// e.g. Equals(float32(0.1)) won't match (0.1 + 0i).
		if im == 0 && float32(rl) == e {
			res = true
			err = nil
		}

	default:
		res = false
		err = NewFatalError("which is not numeric")
	}

	return
}

func checkAgainstFloat64(e float64, c reflect.Value) (res bool, err error) {
	res = false
	err = errors.New("")

	ck := c.Kind()

	switch {
	case isSignedInteger(c):
		if float64(c.Int()) == e {
			res = true
			err = nil
		}

	case isUnsignedInteger(c):
		if float64(c.Uint()) == e {
			res = true
			err = nil
		}

	// If the actual value is lower precision, turn the comparison around so we
	// apply the low-precision rules. Otherwise, e.g. Equals(0.1) may not match
	// float32(0.1).
	case ck == reflect.Float32 || ck == reflect.Complex64:
		return Equals(c.Interface()).Matches(e)

		// Otherwise, compare with double precision.
	case isFloat(c):
		if c.Float() == e {
			res = true
			err = nil
		}

	case isComplex(c):
		comp := c.Complex()
		rl := real(comp)
		im := imag(comp)

		if im == 0 && rl == e {
			res = true
			err = nil
		}

	default:
		res = false
		err = NewFatalError("which is not numeric")
	}

	return
}

func checkAgainstComplex64(e complex64, c reflect.Value) (res bool, err error) {
	res = false
	err = errors.New("")
	realPart := real(e)
	imaginaryPart := imag(e)

	switch {
	case isInteger(c) || isFloat(c):
		// If we have no imaginary part, then we should just compare against the
		// real part. Otherwise, we can't be equal.
		if imaginaryPart != 0 {
			return
		}

		return checkAgainstFloat32(realPart, c)

	case isComplex(c):
		// Compare using complex64 to avoid a false sense of precision; otherwise
		// e.g. Equals(0.1 + 0i) won't match float32(0.1).
		if complex64(c.Complex()) == e {
			res = true
			err = nil
		}

	default:
		res = false
		err = NewFatalError("which is not numeric")
	}

	return
}

func checkAgainstComplex128(e complex128, c reflect.Value) (res bool, err error) {
	res = false
	err = errors.New("")
	realPart := real(e)
	imaginaryPart := imag(e)

	switch {
	case isInteger(c) || isFloat(c):
		// If we have no imaginary part, then we should just compare against the
		// real part. Otherwise, we can't be equal.
		if imaginaryPart != 0 {
			return
		}

		return checkAgainstFloat64(realPart, c)

	case isComplex(c):
		if c.Complex() == e {
			res = true
			err = nil
		}

	default:
		res = false
		err = NewFatalError("which is not numeric")
	}

	return
}

////////////////////////////////////////////////////////////
// Other types
////////////////////////////////////////////////////////////

func checkAgainstBool(e bool, c reflect.Value) (res bool, err error) {
	if c.Kind() != reflect.Bool {
		res = false
		err = NewFatalError("which is not a bool")
		return
	}

	res = false
	err = errors.New("")
	if c.Bool() == e {
		res = true
		err = nil
	}
	return
}

func checkAgainstUintptr(e uintptr, c reflect.Value) (res bool, err error) {
	if c.Kind() != reflect.Uintptr {
		res = false
		err = NewFatalError("which is not a uintptr")
		return
	}

	res = false
	err = errors.New("")
	if uintptr(c.Uint()) == e {
		res = true
		err = nil
	}
	return
}

func checkAgainstChan(e reflect.Value, c reflect.Value) (res bool, err error) {
	// Create a description of e's type, e.g. "chan int".
	typeStr := fmt.Sprintf("%s %s", e.Type().ChanDir(), e.Type().Elem())

	// Make sure c is a chan of the correct type.
	if c.Kind() != reflect.Chan ||
		c.Type().ChanDir() != e.Type().ChanDir() ||
		c.Type().Elem() != e.Type().Elem() {
		res = false
		err = NewFatalError(fmt.Sprintf("which is not a %s", typeStr))
		return
	}

	res = false
	err = errors.New("")
	if c.Pointer() == e.Pointer() {
		res = true
		err = nil
	}
	return
}

func checkAgainstFunc(e reflect.Value, c reflect.Value) (res bool, err error) {
	// Make sure c is a function.
	if c.Kind() != reflect.Func {
		res = false
		err = NewFatalError("which is not a function")
		return
	}

	res = false
	err = errors.New("")
	if c.Pointer() == e.Pointer() {
		res = true
		err = nil
	}
	return
}

func checkAgainstMap(e reflect.Value, c reflect.Value) (res bool, err error) {
	// Make sure c is a map.
	if c.Kind() != reflect.Map {
		res = false
		err = NewFatalError("which is not a map")
		return
	}

	res = false
	err = errors.New("")
	if c.Pointer() == e.Pointer() {
		res = true
		err = nil
	}
	return
}

func checkAgainstPtr(e reflect.Value, c reflect.Value) (res bool, err error) {
	// Create a description of e's type, e.g. "*int".
	typeStr := fmt.Sprintf("*%v", e.Type().Elem())

	// Make sure c is a pointer of the correct type.
	if c.Kind() != reflect.Ptr ||
		c.Type().Elem() != e.Type().Elem() {
		res = false
		err = NewFatalError(fmt.Sprintf("which is not a %s", typeStr))
		return
	}

	res = false
	err = errors.New("")
	if c.Pointer() == e.Pointer() {
		res = true
		err = nil
	}
	return
}

func checkAgainstSlice(e reflect.Value, c reflect.Value) (res bool, err error) {
	// Create a description of e's type, e.g. "[]int".
	typeStr := fmt.Sprintf("[]%v", e.Type().Elem())

	// Make sure c is a slice of the correct type.
	if c.Kind() != reflect.Slice ||
		c.Type().Elem() != e.Type().Elem() {
		res = false
		err = NewFatalError(fmt.Sprintf("which is not a %s", typeStr))
		return
	}

	res = false
	err = errors.New("")
	if c.Pointer() == e.Pointer() {
		res = true
		err = nil
	}
	return
}

func checkAgainstString(e reflect.Value, c reflect.Value) (res bool, err error) {
	// Make sure c is a string.
	if c.Kind() != reflect.String {
		res = false
		err = NewFatalError("which is not a string")
		return
	}

	res = false
	err = errors.New("")
	if c.String() == e.String() {
		res = true
		err = nil
	}
	return
}

func checkAgainstUnsafePointer(e reflect.Value, c reflect.Value) (res bool, err error) {
	// Make sure c is a pointer.
	if c.Kind() != reflect.UnsafePointer {
		res = false
		err = NewFatalError("which is not a unsafe.Pointer")
		return
	}

	res = false
	err = errors.New("")
	if c.Pointer() == e.Pointer() {
		res = true
		err = nil
	}
	return
}

func checkForNil(c reflect.Value) (res bool, err error) {
	res = false
	err = errors.New("")

	// Make sure it is legal to call IsNil.
	switch c.Kind() {
	case reflect.Invalid:
	case reflect.Chan:
	case reflect.Func:
	case reflect.Interface:
	case reflect.Map:
	case reflect.Ptr:
	case reflect.Slice:

	default:
		res = false
		err = NewFatalError("which cannot be compared to nil")
		return
	}

	// Ask whether the value is nil. Handle a nil literal (kind Invalid)
	// specially, since it's not legal to call IsNil there.
	if c.Kind() == reflect.Invalid || c.IsNil() {
		res = true
		err = nil
	}
	return
}

////////////////////////////////////////////////////////////
// Public implementation
////////////////////////////////////////////////////////////

func (m *equalsMatcher) Matches(candidate interface{}) (bool, error) {
	e := m.expectedValue
	c := reflect.ValueOf(candidate)
	ek := e.Kind()

	switch {
	case ek == reflect.Bool:
		return checkAgainstBool(e.Bool(), c)

	case isSignedInteger(e):
		return checkAgainstInt64(e.Int(), c)

	case isUnsignedInteger(e):
		return checkAgainstUint64(e.Uint(), c)

	case ek == reflect.Uintptr:
		return checkAgainstUintptr(uintptr(e.Uint()), c)

	case ek == reflect.Float32:
		return checkAgainstFloat32(float32(e.Float()), c)

	case ek == reflect.Float64:
		return checkAgainstFloat64(e.Float(), c)

	case ek == reflect.Complex64:
		return checkAgainstComplex64(complex64(e.Complex()), c)

	case ek == reflect.Complex128:
		return checkAgainstComplex128(complex128(e.Complex()), c)

	case ek == reflect.Chan:
		return checkAgainstChan(e, c)

	case ek == reflect.Func:
		return checkAgainstFunc(e, c)

	case ek == reflect.Map:
		return checkAgainstMap(e, c)

	case ek == reflect.Ptr:
		return checkAgainstPtr(e, c)

	case ek == reflect.Slice:
		return checkAgainstSlice(e, c)

	case ek == reflect.String:
		return checkAgainstString(e, c)

	case ek == reflect.UnsafePointer:
		return checkAgainstUnsafePointer(e, c)

	case ek == reflect.Invalid:
		return checkForNil(c)
	}

	panic(fmt.Sprintf("equalsMatcher.Matches: unexpected kind: %v", ek))
}

func (m *equalsMatcher) Description() string {
	// Special case: handle nil.
	if !m.expectedValue.IsValid() {
		return "is nil"
	}

	return fmt.Sprintf("%v", m.expectedValue.Interface())
}
