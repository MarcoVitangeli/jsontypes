package writer

type WriterType string

const (
	Float64 WriterType = "float64"
	String  WriterType = "string"
	Bool    WriterType = "bool"
	Any     WriterType = "interface{}"
	Object  WriterType = "map[string]interface{}"
	Array   WriterType = "[]interface{}"
)
