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
)

// Given a list of arguments M, ElementsAre returns a matcher that matches
// arrays and slices A where all of the following hold:
//
//  *  A is the same length as M.
//
//  *  For each i < len(A) where M[i] is a matcher, A[i] matches M[i].
//
//  *  For each i < len(A) where M[i] is not a matcher, A[i] matches
//     Equals(M[i]).
//
func ElementsAre(M ...interface{}) Matcher {
	return nil
}
