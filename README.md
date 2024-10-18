# jsontypes

CLI tool to convert automatically a JSON file to a Go struct with json tags. It admits the following data types:
- float64
- string
- nested objects
- arrays
- time.Time

In the case of arrays, it will construct the type using the first element in the collection. Maybe later a smarter algorithm could come.

## Running the example
```sh
make
./jsontypes sample.json
```

That will generate a `types_gen.go` with the generated golang struct. In order to see usage just call `jsontypes` with no argumnets.
