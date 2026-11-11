# RayID for Go

Generate short, readable IDs for logs, events, requests, and traces.

```go
package main

import (
	"fmt"

	"gitlab.com/blackiq-packages/go/rayid"
)

func main() {
	ray, err := rayid.New(rayid.Digit)
	if err != nil {
		panic(err)
	}

	id, err := ray.Gen(16)
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
}
```

|        |                                                                                                            |
| ------ | ---------------------------------------------------------------------------------------------------------- |
| Module | [pkg.go.dev/gitlab.com/blackiq-packages/go/rayid](https://pkg.go.dev/gitlab.com/blackiq-packages/go/rayid) |
| Source | [https://gitlab.com/blackiq-packages/go/rayid](https://gitlab.com/blackiq-packages/go/rayid)               |

## Next

- [Install](learn/install.md)
- [Usage](learn/usage.md)
