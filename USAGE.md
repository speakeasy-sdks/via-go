<!-- Start SDK Example Usage [usage] -->
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
	res, err := s.GetUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Strings != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->