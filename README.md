# RayID (Go)

Generate short IDs for logs, events, and traces.

```bash
go get github.com/BlackIQ/rayid-go@latest
```

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

**Types:** `digit` · `lower` · `upper` · `all`

## Docs

Full documentation: [rayid-go.amirhossein.info](https://rayid-go.amirhossein.info)

## Links

| Title  | URL                                            |
| ------ | ---------------------------------------------- |
| Module | https://pkg.go.dev/github.com/BlackIQ/rayid-go |
| Source | https://github.com/BlackIQ/rayid-go            |
| Docs   | https://rayid-go.amirhossein.info              |

## License

MIT
