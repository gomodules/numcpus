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

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	sysfsCPUBasePath = "/sys/devices/system/cpu"
)

func readCPURange(file string) (int, error) {
	buf, err := ioutil.ReadFile(filepath.Join(sysfsCPUBasePath, file))
	if err != nil {
		return 0, err
	}
	return parseCPURange(strings.Trim(string(buf), "\n "))
}

func parseCPURange(cpus string) (int, error) {
	n := int(0)
	for _, cpuRange := range strings.Split(cpus, ",") {
		if len(cpuRange) == 0 {
			continue
		}
		rangeOp := strings.SplitN(cpuRange, "-", 2)
		first, err := strconv.ParseUint(rangeOp[0], 10, 32)
		if err != nil {
			return 0, err
		}
		if len(rangeOp) == 1 {
			n++
			continue
		}
		last, err := strconv.ParseUint(rangeOp[1], 10, 32)
		if err != nil {
			return 0, err
		}
		n += int(last - first + 1)
	}
	return n, nil
}

func getKernelMax() (int, error) {
	return readCPURange("kernel_max")
}

func getOffline() (int, error) {
	return readCPURange("offline")
}

func getOnline() (int, error) {
	return readCPURange("online")
}

func getPossible() (int, error) {
	return readCPURange("possible")
}

func getPresent() (int, error) {
	return readCPURange("present")
}