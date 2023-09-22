<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	viago "github.com/speakeasy-sdks/via-go"
)

func main() {
    s := viago.New()

    ctx := context.Background()
    res, err := s.GetUsers(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetUsers200ApplicationJSONStrings != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->