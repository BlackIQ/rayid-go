# RayID (Go)

Generate short IDs for logs, events, and traces.

```bash
go get gitlab.com/blackiq-packages/go/rayid@latest
```

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

**Types:** `digit` · `lower` · `upper` · `all`

## Docs

Full documentation: [rayid-go.docs.amirhossein.info](https://rayid-go.docs.amirhossein.info)

## Links

|        |                                                                                                            |
| ------ | ---------------------------------------------------------------------------------------------------------- |
| Module | [pkg.go.dev/gitlab.com/blackiq-packages/go/rayid](https://pkg.go.dev/gitlab.com/blackiq-packages/go/rayid) |
| Source | [https://gitlab.com/blackiq-packages/go/rayid](https://gitlab.com/blackiq-packages/go/rayid)               |

## License

MIT
