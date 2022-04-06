# sase-go
SDK for interacting with Secure Access Edge Service

[![GoDoc](https://godoc.org/github.com/PaloAltoNetworks/sase-go?status.svg)](https://godoc.org/github.com/PaloAltoNetworks/sase-go)

Package `sase` is a golang SDK for interacting with the Palo Alto Networks Secure Access Edge Service.

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

## Support

Community Supported

The software and templates in the repo are released under an as-is, best effort, support policy. This software should be seen as community supported and Palo Alto Networks will contribute our expertise as and when possible. We do not provide technical support or help in using or troubleshooting the components of the project through our normal support options such as Palo Alto Networks support teams, or ASC (Authorized Support Centers) partners and backline support options. The underlying product used (the SASE API) by the scripts or templates are still supported, but the support is only for the product functionality and not for help in deploying or using the template or script itself. Unless explicitly tagged, all projects or work posted in our GitHub repository (at https://github.com/PaloAltoNetworks) or sites other than our official Downloads page on https://support.paloaltonetworks.com are provided under the best effort policy.
