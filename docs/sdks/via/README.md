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
	viago "github.com/speakeasy-sdks/via-go"
	"context"
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetUsersResponse](../../pkg/models/operations/getusersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
