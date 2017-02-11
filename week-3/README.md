# Week 3

- [Make a GET request to API](client/main.go#L19)
- Convert JSON string into Go `struct`:
  - `json.Unmarshal` needs a slice of bytes and a 'schema' in the form of a [blank interface](client/main.go#L44)
  - interface can be anything(?)
  - lowest form of object
  - Could be a value, could be a reference
- Sort slice [by alphabetical order of names](client/main.go#L53)

#### Bubble sorting

Given this slice/array: `[c, b, a]`:

```go
// is c 'bigger than' b?
=> [b, c, a]

// is c 'bigger than' a?
=> [b, a, c]

// is c 'bigger than' b?
=> [a, b, c]
```
