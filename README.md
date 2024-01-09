# sase-go
SDK for interacting with Secure Access Edge Service

[![GoDoc](https://godoc.org/github.com/PaloAltoNetworks/sase-go?status.svg)](https://godoc.org/github.com/PaloAltoNetworks/sase-go)

Package `sase` is a golang SDK for interacting with the Palo Alto Networks Secure Access Edge Service.

**NOTE**: Use the new [paloaltonetworks/scm-go](https://github.com/paloaltonetworks/scm-go) repo instead.

## Using sase-go

To start, create a client connection, then invoke `Setup()`, and retrieve the JWT with `RefreshJwt()`.  JWTs expire after some time, but the SDK will catch the failed auth and automatically refresh the JWT when a 401 is returned from the API:

```go
package main

import (
    "context"
    "log"

    "github.com/paloaltonetworks/sase-go"
)

func main() {
    var err error
    ctx := context.TODO()

    con := sase.Client{
        ClientId: "1234-56-78-9",
        ClientSecret: "a-b-c-d",
        Scope: "tsg_id:123456789",
        CheckEnvironment: true,
    }

    if err = con.Setup(); err != nil {
        log.Fatal(err)
    } else if err = con.RefreshJwt(ctx); err != nil {
        log.Fatal(err)
    }

    log.Printf("Authenticated ok")
}
```
