# Mailchimp 

A golang SDK for Mailchimp API v3.

[![Build Status](https://ci.matthewbrown.io/api/badge/github.com/mnbbrown/mailchimp/status.svg?branch=master)](https://ci.matthewbrown.io/github.com/mnbbrown/mailchimp)
[![GoDoc](https://godoc.org/github.com/mnbbrown/mailchimp?status.svg)](https://godoc.org/github.com/mnbbrown/mailchimp)
[![Coverage Status](https://coveralls.io/repos/mnbbrown/mailchimp/badge.svg?branch=master&service=github)](https://coveralls.io/github/mnbbrown/mailchimp?branch=master)

<div style="margin: 25px;">
<a href="https://rapidapi.com/package/MailChimp/functions?utm_source=MailchimpGitHub&utm_medium=button&utm_content=Vendor_GitHub" style="
    all: initial;
    background-color: #498FE1;
    border-width: 0;
    border-radius: 5px;
    padding: 10px 20px;
    color: white;
    font-family: 'Helvetica';
    font-size: 12pt;
    background-image: url(https://scdn.rapidapi.com/logo-small.png);
    background-size: 25px;
    background-repeat: no-repeat;
    background-position-y: center;
    background-position-x: 10px;
    padding-left: 44px;
    cursor: pointer;">
  Run now on <b>RapidAPI</b>
</a>
</div>

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
