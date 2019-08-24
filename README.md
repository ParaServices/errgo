# errgo

Is an opinionated and requirement specific implementation for error handling
with __Para's__ _golang_ applications.


This package wraps the [errorx](https://github.com/magicalbanana/errorx.git)
which allows [paralog](https://github.com/ParaServices/paralog.git) to log
the error as a `zapcode.Object` which represents the error in a human readable
format (JSON or string).

# features

- Wrap [errorx]
- Assign error ID
- Assign error code
- Assign message
- allow to add __details__ for more information
- verbosely log `pq.Error`
- verbosely log `googleapi.Error`
- verbosely log `amqp.Error`

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/ParaServices/errgo"
)

func main() {
    var i interface{}
    err := json.Unmarshal([]byte(`{`), &i)
    if err != nil {
      errx := errgo.New(err)
      errx.Code = "some code"
      errx.Message = "some message"
      fmt.Printf("ID: %s", errx.ID)
    }
}
```

[errorx]: https://github.com/magicalbanana/errorx.git
