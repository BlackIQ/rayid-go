# RayID for Go

Generate short, readable IDs for logs, events, requests, and traces.

```go
package main

import (
	"fmt"

	"github.com/BlackIQ/rayid-go"
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

|        |                                                |
| ------ | ---------------------------------------------- |
| Module | https://pkg.go.dev/github.com/BlackIQ/rayid-go |
| Source | https://github.com/BlackIQ/rayid-go            |

## Next

- [Install](learn/install.md)
- [Usage](learn/usage.md)
