/*
Copyright 2020 Daniel Avrukin
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

package main

func allowHiddenApiAccess() {

	println("About to ALLOW access to hidden APIs.")

	_, err1 := set10().Output()
	_, err2 := set9PreP().Output()
	_, err3 := set9P().Output()

	if err1 != nil {
		println("Step 1 error:")
		println(err1.Error())
	} else {
		println("Step 1 complete.")
	}

	if err2 != nil {
		println("Step 2 error:")
		println(err2.Error())
	} else {
		println("Step 2 complete.")
	}

	if err3 != nil {
		println("Step 3 error:")
		println(err3.Error())
	} else {
		println("Step 3 complete.")
	}
}

func denyHiddenApiAccess() {

	println("About to DENY access to hidden APIs.")

	_, err1 := reset10().Output()
	_, err2 := reset9PreP().Output()
	_, err3 := reset9P().Output()

	if err1 != nil {
		println("Step 1 error:")
		println(err1.Error())
	} else {
		println("Step 1 complete.")
	}

	if err2 != nil {
		println("Step 2 error:")
		println(err2.Error())
	} else {
		println("Step 2 complete.")
	}

	if err3 != nil {
		println("Step 3 error:")
		println(err3.Error())
	} else {
		println("Step 3 complete.")
	}
}
