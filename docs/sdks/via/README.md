# Via SDK

## Overview

Sample API: Optional multiline or single-line description in [CommonMark](http://commonmark.org/help/) or HTML.

### Available Operations

* [GetUsers](#getusers) - Returns a list of users.

## GetUsers

Optional extended description in CommonMark or HTML.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/via-go"
)

func main() {
    s := via.New()

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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetUsersResponse](../../models/operations/getusersresponse.md), error**

