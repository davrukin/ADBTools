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

import "os/exec"

// https://developer.android.com/distribute/best-practices/develop/restrictions-non-sdk-interfaces#general-questions

// set

func set10() *exec.Cmd {
	return exec.Command(adb, shell, settings, put, global, hiddenApiPolicy, one)
}

func set9PreP() *exec.Cmd {
	return exec.Command(adb, shell, settings, put, global, hiddenApiPolicyPreP, one)
}

func set9P() *exec.Cmd {
	return exec.Command(adb, shell, settings, put, global, hiddenApiPolicyP, one)
}

// reset

func reset10() *exec.Cmd {
	return exec.Command(adb, shell, settings, delete, global, hiddenApiPolicy)
}

func reset9PreP() *exec.Cmd {
	return exec.Command(adb, shell, settings, delete, global, hiddenApiPolicyPreP)
}

func reset9P() *exec.Cmd {
	return exec.Command(adb, shell, settings, delete, global, hiddenApiPolicyP)
}
