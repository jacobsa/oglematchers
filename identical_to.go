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

package oglematchers

import (
	"errors"
	"fmt"
	"reflect"
)

// Is the type comparable according to the definition here?
//
//     http://weekly.golang.org/doc/go_spec.html#Comparison_operators
//
func isComparable(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Array:
		return isComparable(t.Elem())

	case reflect.Slice, reflect.Map, reflect.Func:
		return false
	}

	return true
}

// Should the supplied type be allowed as an argument to IdenticalTo?
func isLegalForIdenticalTo(t reflect.Type) (bool, error) {
	// Reject containers with non-comparable elements.
	if t.Kind() == reflect.Array && !isComparable(t.Elem()) {
		return false, errors.New("TODO")
	}

	return true, nil
}

// IdenticalTo(x) returns a matcher that matches values v such that all of the
// following hold:
//
//  *  v and x have identical types.
//
//  *  If v and x are of a reference type (slice, map, function, channel), then
//     they are either both nil or are references to the same object.
//
//  *  If v and x are not of a reference type, then v == x.
//
// This function will panic if x is of a value type that is not comparable. For
// example, x cannot be an array of functions.
func IdenticalTo(x interface{}) Matcher {
	t := reflect.TypeOf(x)

	// Reject illegal arguments.
	if ok, err := isLegalForIdenticalTo(t); !ok {
		panic("IdenticalTo: " + err.Error())
	}

	return &identicalToMatcher{x}
}

type identicalToMatcher struct {
	x interface{}
}

func (m *identicalToMatcher) Description() string {
	t := reflect.TypeOf(m.x)
	return fmt.Sprintf("identical to <%v> %v", t, m.x)
}

func (m *identicalToMatcher) Matches(c interface{}) error {
	// Make sure the candidate's type is correct.
	t := reflect.TypeOf(m.x)
	if ct := reflect.TypeOf(c); t != ct {
		return NewFatalError(fmt.Sprintf("which is of type %v", ct))
	}

	return errors.New("TODO")
}
