# gimme

gimme allows you to reduce a lot of repetitive tasks in Go to simple one liners, at some cost of stability (errors will panic by default)

## Usage

To use, simply import the library and call one of the functions:

```go
content := gimme.URL("http://httpbin.org/get")
```

By default if you have not specificed http:// in the url, gimme will just add http://, so if you need https it needs to be specificed in the url

## Todo

- [ ] add configuration and options without getting in the way of simplicity
- [ ] options for what to do on failure/errors