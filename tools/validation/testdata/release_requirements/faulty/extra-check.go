/*
Copyright 2024 Flant JSC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package test

import (
	"github.com/deckhouse/deckhouse/go_lib/dependency/requirements"
)

var (
	whatDidYouSay   = "what?"
	yetAnotherCheck = "testVer"
)

func test4() {
	requirements.RegisterCheck(whatDidYouSay, func(requirementValue string, getter requirements.ValueGetter) (bool, error) {
		return false, nil
	})
	requirements.RegisterCheck(yetAnotherCheck, func(requirementValue string, getter requirements.ValueGetter) (bool, error) {
		return false, nil
	})
}