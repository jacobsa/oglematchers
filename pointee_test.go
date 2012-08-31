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
	"testing"
)

////////////////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////////////////

type PointeeTest struct {}
func init()                     { RegisterTestSuite(&NotTest{}) }

func TestPointee(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////////////////

func (t *NotTest) Description() {
	ExpectEq("TODO", "")
}

func (t *NotTest) CandidateIsNotAPointer() {
	ExpectEq("TODO", "")
}

func (t *NotTest) CandidateIsANilPointer() {
	ExpectEq("TODO", "")
}

func (t *NotTest) CallsWrapped() {
	ExpectEq("TODO", "")
}

func (t *NotTest) WrappedReturnsTrue() {
	ExpectEq("TODO", "")
}

func (t *NotTest) WrappedReturnsNonFatalError() {
	ExpectEq("TODO", "")
}

func (t *NotTest) WrappedReturnsFatalError() {
	ExpectEq("TODO", "")
}
