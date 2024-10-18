build:
	go build cmd/cli/jsontypes.go

clean:
	rm type_gen.go
	rm jsontypes
