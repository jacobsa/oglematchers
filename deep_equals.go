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

// DeepEquals returns a matcher that matches based on 'deep equality', as
// defined by the reflect package. This matcher requires that values have
// identical types to x.
func DeepEquals(x interface{}) Matcher {
	return &deepEqualsMatcher{x}
}

type deepEqualsMatcher struct {
	x interface{}
}

func (m *deepEqualsMatcher) Description() string {
	return fmt.Sprintf("deep equals: %v", m.x)
}

func (m *deepEqualsMatcher) Matches(c interface{}) error {
	// Make sure the types match.
	ct := reflect.TypeOf(c)
	xt := reflect.TypeOf(m.x)

	if ct != xt {
		return NewFatalError(fmt.Sprintf("which is of type %v", ct))
	}

	// Defer to the reflect package.
	if reflect.DeepEqual(m.x, c) {
		return nil
	}

	return errors.New("")
}
