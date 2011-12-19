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

package ogletest

import (
	"fmt"
	"math"
	"reflect"
)

// Equals returns a matcher that matches any value v such that v == x, with the
// exception that if x is a numeric type, Equals(x) will match equivalent
// numeric values of any type.
func Equals(x interface{}) Matcher {
	return &equalsMatcher{x}
}

type equalsMatcher struct {
	expected interface{}
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

func checkAgainstInt64(e int64, c reflect.Value) (res MatchResult, err string) {
	res = MATCH_FALSE

	switch {
	case isSignedInteger(c):
		if c.Int() == e {
			res = MATCH_TRUE
		}

	case isUnsignedInteger(c):
		u := c.Uint()
		if u <= math.MaxInt64 && int64(u) == e {
			res = MATCH_TRUE
		}

	// Turn around the various floating point types so that the checkAgainst*
	// functions for them can deal with precision issues.
	case isFloat(c), isComplex(c):
		return Equals(c.Interface()).Matches(e)

	default:
		res = MATCH_UNDEFINED
		err = "which is not numeric"
	}

	return
}

func checkAgainstFloat32(e float32, c reflect.Value) (res MatchResult, err string) {
	res = MATCH_FALSE

	switch {
	case isSignedInteger(c):
		if float32(c.Int()) == e {
			res = MATCH_TRUE
		}

	case isUnsignedInteger(c):
		if float32(c.Uint()) == e {
			res = MATCH_TRUE
		}

	case isFloat(c):
		// Compare using float32 to avoid a false sense of precision; otherwise
		// e.g. Equals(float32(0.1)) won't match float32(0.1).
		if float32(c.Float()) == e {
			res = MATCH_TRUE
		}

	case isComplex(c):
		comp := c.Complex()
		rl := real(comp)
		im := imag(comp)

		// Compare using float32 to avoid a false sense of precision; otherwise
		// e.g. Equals(float32(0.1)) won't match (0.1 + 0i).
		if im == 0 && float32(rl) == e {
			res = MATCH_TRUE
		}

	default:
		res = MATCH_UNDEFINED
		err = "which is not numeric"
	}

	return
}

func checkAgainstFloat64(e float64, c reflect.Value) (res MatchResult, err string) {
	res = MATCH_FALSE

	switch {
	case isSignedInteger(c):
		if float64(c.Int()) == e {
			res = MATCH_TRUE
		}

	case isUnsignedInteger(c):
		if float64(c.Uint()) == e {
			res = MATCH_TRUE
		}

	case isFloat(c):
		if c.Float() == e {
			res = MATCH_TRUE
		}

	case isComplex(c):
		comp := c.Complex()
		rl := real(comp)
		im := imag(comp)

		if im == 0 && rl == e {
			res = MATCH_TRUE
		}

	default:
		res = MATCH_UNDEFINED
		err = "which is not numeric"
	}

	return
}

func checkAgainstComplex64(e complex64, c reflect.Value) (res MatchResult, err string) {
	res = MATCH_FALSE
	realPart := real(e)
	imaginaryPart := imag(e)

	switch {
	case isInteger(c) || isFloat(c):
		// If we have no imaginary part, then we should just compare against the
		// real part. Otherwise, we can't be equal.
		if imaginaryPart != 0 {
			res = MATCH_FALSE
			return
		}

		return checkAgainstFloat32(realPart, c)

	case isComplex(c):
		// Compare using complex64 to avoid a false sense of precision; otherwise
		// e.g. Equals(0.1 + 0i) won't match float32(0.1).
		if complex64(c.Complex()) == e {
			res = MATCH_TRUE
		}

	default:
		res = MATCH_UNDEFINED
		err = "which is not numeric"
	}

	return
}

func checkAgainstComplex128(e complex128, c reflect.Value) (res MatchResult, err string) {
	res = MATCH_FALSE
	realPart := real(e)
	imaginaryPart := imag(e)

	switch {
	case isInteger(c) || isFloat(c):
		// If we have no imaginary part, then we should just compare against the
		// real part. Otherwise, we can't be equal.
		if imaginaryPart != 0 {
			res = MATCH_FALSE
			return
		}

		return checkAgainstFloat64(realPart, c)

	case isComplex(c):
		if c.Complex() == e {
			res = MATCH_TRUE
		}

	default:
		res = MATCH_UNDEFINED
		err = "which is not numeric"
	}

	return
}

////////////////////////////////////////////////////////////
// Public implementation
////////////////////////////////////////////////////////////

func (m *equalsMatcher) Matches(candidate interface{}) (MatchResult, string) {
	e := reflect.ValueOf(m.expected)
	c := reflect.ValueOf(candidate)
	ek := e.Kind()

	switch {
	case isSignedInteger(e):
		return checkAgainstInt64(e.Int(), c)

	case ek == reflect.Float32:
		return checkAgainstFloat32(float32(e.Float()), c)

	case ek == reflect.Float64:
		return checkAgainstFloat64(e.Float(), c)

	case ek == reflect.Complex64:
		return checkAgainstComplex64(complex64(e.Complex()), c)

	case ek == reflect.Complex128:
		return checkAgainstComplex128(complex128(e.Complex()), c)
	}

	return MATCH_UNDEFINED, "TODO"
}

func (m *equalsMatcher) Description() string {
	return fmt.Sprintf("%v", m.expected)
}
