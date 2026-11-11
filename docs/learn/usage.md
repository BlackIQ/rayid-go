# Usage

## Create a generator

```javascript
import RayID from "rayid";

const ray = new RayID("lower");
```

### Types

| Type               | Characters                |
| ------------------ | ------------------------- |
| `digit`            | `0-9`                     |
| `lower`            | `a-z`                     |
| `upper`            | `A-Z`                     |
| `symbol`           | common symbols            |
| _(omit / unknown)_ | all of the above combined |

```javascript
const digits = new RayID("digit");
const symbols = new RayID("symbol");
const mixed = new RayID(); // all character sets
```

## Generate an ID

```javascript
ray.gen(12);
```

If the result would start with `0`, generation is retried so the first character is non-zero.

```javascript
import RayID from "rayid";

console.log(new RayID("digit").gen(10));
console.log(new RayID("lower").gen(10));
console.log(new RayID("symbol").gen(10));
console.log(new RayID().gen(16));
```

## API

### `new RayID(type?)`

| Argument | Description                                          |
| -------- | ---------------------------------------------------- |
| `type`   | `digit`, `lower`, `upper`, `symbol`, or omit for all |

### `gen(length) → string`

Returns a string of the given length from the selected alphabet.
