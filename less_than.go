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
	"reflect"
)

// LessThan returns a matcher that matches integer, floating point, or strings
// values v such that v < x. Comparison is not defined between numeric and
// string types, but is defined between all integer and floating point types.
//
// x must itself be an integer, floating point, or string type; otherwise,
// LessThan will panic.
func LessThan(x interface{}) Matcher {
	v := reflect.ValueOf(x)
	kind := v.Kind()

	switch {
	case isInteger(v):
	case isFloat(v):
	case kind == reflect.String:

	default:
		panic(fmt.Sprintf("LessThan: unexpected kind %v", kind))
	}

	return &lessThanMatcher{v}
}

type lessThanMatcher struct {
	limit reflect.Value
}

func (m *lessThanMatcher) Description() string {
	return fmt.Sprintf("less than %v", m.limit)
}

func compareIntegers(v1, v2 reflect.Value) (MatchResult, string) {
	return MATCH_UNDEFINED, "TODO"
}

func (m *lessThanMatcher) Matches(c interface{}) (res MatchResult, err string) {
	v1 := reflect.ValueOf(c)
	v2 := m.limit

	res = MATCH_FALSE

	switch {
	case isInteger(v1) && isInteger(v2):
		return compareIntegers(v1, v2)

	case (isInteger(v1) || isFloat(v1)) && v2.Kind() == reflect.Float32:
		if float32(getFloat(v1)) < float32(v2.Float()) {
			res = MATCH_TRUE
		}
		return

	case (isInteger(v1) || isFloat(v1)) && v2.Kind() == reflect.Float64:
		if getFloat(v1) < v2.Float() {
			res = MATCH_TRUE
		}
		return

	case v1.Kind() == reflect.Float32 && isNumeric(v2):
		if float32(getFloat(v1)) < float32(getFloat(v2)) {
			res = MATCH_TRUE
		}
		return

	case v1.Kind() == reflect.Float64 && isNumeric(v2):
		if getFloat(v1) < getFloat(v2) {
			res = MATCH_TRUE
		}
		return

	case v1.Kind() == reflect.String && v2.Kind() == reflect.String:
		if v1.String() < v2.String() {
			res = MATCH_TRUE
		}
		return
	}

	res = MATCH_UNDEFINED
	err = "which is not comparable"
}
