 # Gofish - Redfish and Swordfish client library

 [![Go Report Card](https://goreportcard.com/badge/github.com/stmcginnis/gofish?branch=master)](https://goreportcard.com/report/github.com/stmcginnis/gofish)
[![Releases](https://img.shields.io/github/release/stmcginnis/gofish/all.svg?style=flat-square)](https://github.com/stmcginnis/gofish/releases)
[![LICENSE](https://img.shields.io/github/license/stmcginnis/gofish.svg?style=flat-square)](https://github.com/stmcginnis/gofish/blob/master/LICENSE)

## Introduction

Gofish is a Golang library for interacting with [DMTF
Redfish](https://www.dmtf.org/standards/redfish) and [SNIA
Swordfish](https://www.snia.org/forums/smi/swordfish) enabled devices.

## Usage ##

Basic usage would be:

```go

package main

import (
    "fmt"

    "github.com/stmcginnis/gofish/school"
)

func main() {
    c := gofish.APIClient("localhost", 5000, false)
    service, _ := school.ServiceRoot(c)

    chassis, _ := service.Chassis()
    for _, chass := range chassis {
        fmt.Printf("Chassis: %#v\n\n", chass)
    }
}
```
