# numcpus

[![Build Status][1]][2]
[![Go Report Card][3]][4]
[![GoDoc][5]][6]

Get the number of CPUs (online, offline, present, possible or kernel maximum)
on a Linux, FreeBSD or NetBSD system.

On Linux, the information is retrieved by reading the CPU topology files in
`/sys/devices/system/cpu`. On FreeBSD and netBSD, not all functions are
supported.

## Usage

```Go
package main

import (
	"fmt"

	"github.com/tklauser/numcpus"
)

func main() {
	online, err := numcpus.GetOnline()
	if err != nil {
		fmt.Printf("online CPUs: %v\n", online)
	}
}
```

## References

* [Linux kernel sysfs documenation for CPU attributes](https://www.kernel.org/doc/Documentation/ABI/testing/sysfs-devices-system-cpu)
* [Linux kernel CPU topology documentation](https://www.kernel.org/doc/Documentation/cputopology.txt)

[1]: https://travis-ci.com/tklauser/numcpus.svg?branch=master
[2]: https://travis-ci.com/tklauser/numcpus
[3]: https://goreportcard.com/badge/github.com/tklauser/numcpus
[4]: https://goreportcard.com/report/github.com/tklauser/numcpus
[5]: https://godoc.org/github.com/tklauser/numcpus?status.svg
[6]: https://godoc.org/github.com/tklauser/numcpus
