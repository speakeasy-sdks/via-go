<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	viago "github.com/speakeasy-sdks/via-go"
	"log"
)

func main() {
	s := viago.New()

	ctx := context.Background()
	res, err := s.Via.GetUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.GetUsers200ApplicationJSONStrings != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->