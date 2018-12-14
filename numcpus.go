// Copyright 2018 Tobias Klauser
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

package numcpus

// GetKernelMax returns the maximum number of CPUs allowed by the kernel
// configuration.
func GetKernelMax() (int, error) {
	return getKernelMax()
}

// GetOffline returns the number of offline CPUs, i.e. CPUs that are not online
// because they have been hotplugged off or exceed the limit of CPUs allowed by
// the kernel configuration (see GetKernelMax).
func GetOffline() (int, error) {
	return getOffline()
}

// GetOnline returns the number of CPUs that are online and being scheduled.
func GetOnline() (int, error) {
	return getOnline()
}

// GetPossible returns the number of possible CPUs, i.e. CPUs that
// have been allocated resources and can be brought online if they are present.
// The number is retrieved by parsing /sys/device/system/cpu/possible.
func GetPossible() (int, error) {
	return getPossible()
}

// GetPresent returns the number of CPUs present in the system.
func GetPresent() (int, error) {
	return getPresent()
}
