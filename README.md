# DateTime

Convenience functions for working with [JavaScript date time strings](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date#date_time_string_format) in Go.

```bash
$ go get github.com/simonward87/datetime
```

## Example

```go
package main

import (
    "fmt"
    "os"

    "github.com/simonward87/datetime"
)

func main() {
    if err := run("2025-01-26T12:58:14.521Z"); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func run(dateTime string) error {
    t, err := datetime.Parse(dateTime)
    if err != nil {
        return err
    }
    // do something with t

    dt := datetime.String(t)
    // do something with dt

    fmt.Println(dt)
}
```
