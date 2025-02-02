# go-yespower

Go bindings of [yespower](https://www.openwall.com/yespower) hashing algorithm written in C

## Testing

```bash
$ go test -v

# ns/op -> s/op -> h/s
$ go test -bench=Yes -benchtime 4s
```

## LICENSE

BSD 2-Clause, as per written from yespower files