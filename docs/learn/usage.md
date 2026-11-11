# Usage

## Create a generator

```go
import "gitlab.com/blackiq-packages/go/rayid"

ray, err := rayid.New(rayid.Lower)
if err != nil {
    panic(err)
}
```

Or use the panic version:

```go
ray := rayid.MustNew(rayid.Lower)
```

### Types

| Type    | Characters                |
| ------- | ------------------------- |
| `digit` | `0-9`                     |
| `lower` | `a-z`                     |
| `upper` | `A-Z`                     |
| `all`   | all of the above combined |

```go
digits := rayid.MustNew(rayid.Digit)
letters := rayid.MustNew(rayid.Lower)
mixed   := rayid.MustNew(rayid.All)
```

## Generate an ID

```go
id, err := ray.Gen(12)

if err != nil {
    panic(err)
}
```

Or:

```go
id := ray.MustGen(12)
```

For the `digit` type, the first character is never `0`.

```go
package main

import (
	"fmt"

	"gitlab.com/blackiq-packages/go/rayid"
)

func main() {
	fmt.Println(rayid.MustNew(rayid.Digit).MustGen(10))
	fmt.Println(rayid.MustNew(rayid.Lower).MustGen(10))
	fmt.Println(rayid.MustNew(rayid.Upper).MustGen(10))
	fmt.Println(rayid.MustNew(rayid.All).MustGen(16))
}
```

## API

### `New(t Type) (*RayID, error)`

| Argument | Description                         |
| -------- | ----------------------------------- |
| `t`      | `Digit`, `Lower`, `Upper`, or `All` |

Returns an error if the type is invalid.

### `MustNew(t Type) *RayID`

Same as `New`, but panics on invalid type.

### `Gen(length int) (string, error)`

Returns a string of the given length from the selected alphabet.
Returns an error if `length <= 0`.

### `MustGen(length int) string`

Same as `Gen`, but panics on error.

### `Characters() string`

Returns the character set used by this generator.
