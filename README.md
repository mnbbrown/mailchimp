# Mailchimp 

A golang SDK for Mailchimp API v3.

[![Build Status](https://ci.matthewbrown.io/api/badge/github.com/mnbbrown/mailchimp/status.svg?branch=master)](https://ci.matthewbrown.io/github.com/mnbbrown/mailchimp)
[![GoDoc](https://godoc.org/github.com/mnbbrown/mailchimp?status.svg)](https://godoc.org/github.com/mnbbrown/mailchimp)
[![Coverage Status](https://coveralls.io/repos/mnbbrown/mailchimp/badge.svg?branch=master&service=github)](https://coveralls.io/github/mnbbrown/mailchimp?branch=master)

## Usage

```go
package main

import (
    "github.com/mnbbrown/mailchimp"
)


func main() {

    client := mailchimp.NewClient("apixyc-us11", nil)
    _, err := client.Subscribe("me@matthewbrown.io", "listidxyz")
    if err != nil {
        panic(err)
    }

}
```