
# Dup

 [![Build Status](https://travis-ci.org/segmentio/go-dup.svg?branch=master)](https://travis-ci.org/segmentio/go-dup)

 `Dup(fd, name)` returns a RW pipe for the given `fd`. For example
 this allows you to proxy stdio, logging to your preferred logger
 and then re-write to the original stdio stream.

## Example

```go
r, w, err := Dup(1, "stdout")
check(err)

go func() {
  fmt.Printf("testing\n")
}()

buf := bufio.NewReader(r)
line, err := buf.ReadString('\n')
check(err)

fmt.Fprintf(w, line)
```

# License

 MIT