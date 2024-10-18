package writer

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"unicode"
    "time"
)

var (
	re = regexp.MustCompile("[^a-zA-Z0-9]+")
)

type TypeWriter struct {
	OutPath string
	bf      *bufio.Writer
	wr      io.WriteCloser
}

func NewTypeWriter(op string) *TypeWriter {
	return &TypeWriter{
		OutPath: op,
		bf:      nil,
	}
}

// Init creates the initial boilerplate for package
func (tw *TypeWriter) Init() error {
	f, err := os.OpenFile(tw.OutPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}

	tw.wr = f
	tw.bf = bufio.NewWriterSize(f, 1024)

	return nil
}

func (tw *TypeWriter) beginType() error {
	s := "type JsonResult "
	_, err := tw.bf.Write([]byte(s))
	return err
}

func (tw *TypeWriter) Write(rawJson map[string]interface{}, maxDepth uint) error {
	_, err := tw.bf.Write([]byte("package types\n\n"))
	if err != nil {
		return err
	}

	if err := tw.beginType(); err != nil {
		return err
	}

	sb, err := tw.consumeObject(rawJson, maxDepth, 0)

	if err != nil {
		return err
	}

	_, err = tw.bf.Write([]byte(sb.String()))

	return err
}

func (tw *TypeWriter) consumeObject(raw map[string]interface{}, maxDepth uint, depth uint) (strings.Builder, error) {
	var (
		sb strings.Builder
	)

	_, err := sb.Write([]byte("struct {\n"))

	if err != nil {
		return sb, err
	}

	var r string

	for k, v := range raw {
		t := getType(v)
		if depth == maxDepth {
			r = string(t)
		} else {
			switch t {
			case Array:
				arr := v.([]interface{})
				if len(arr) == 0 {
					r = string(Array)
				} else {
					one := arr[0]
					if c, ok := one.(map[string]interface{}); ok {
						sb, err := tw.consumeObject(c, maxDepth, depth+1)
						if err != nil {
							return sb, err
						}
						r = sb.String()
					} else {
						t := getType(one)
						r = fmt.Sprintf("[]%s", string(t))
					}
				}
			case Object:
				sb, err := tw.consumeObject(v.(map[string]interface{}), maxDepth, depth+1)
				if err != nil {
					return sb, err
				}
				r = sb.String()
			default:
				r = string(t)
			}
		}
		s := fmt.Sprintf("\t%s %s `json:\"%s\"`\n", toPublicIdentifierName(k), r, k)
		_, err := sb.Write([]byte(s))
		if err != nil {
			return sb, err
		}
	}

	end := "}"

	_, err = sb.Write([]byte(end))

	return sb, err
}

func (tw *TypeWriter) Close() error {
	if err := tw.bf.Flush(); err != nil {
		return err
	}

	return tw.wr.Close()
}

func toPublicIdentifierName(input string) string {
	cleanedInput := re.ReplaceAllString(input, " ")

	words := strings.Fields(cleanedInput)

	if len(words) == 0 {
		return ""
	}

	for i := range words {
		words[i] = capitalize(words[i])
	}

	result := strings.Join(words, "")

	if len(result) > 0 && !unicode.IsLetter(rune(result[0])) {
		result = "Var" + result
	}

	return result
}

func capitalize(word string) string {
	if word == "" {
		return ""
	}
	return string(unicode.ToUpper(rune(word[0]))) + strings.ToLower(word[1:])
}

func getType(v interface{}) WriterType {
	if _, ok := v.(bool); ok {
		return Bool
	}
	if _, ok := v.(float64); ok {
		return Float64
	}
	if _, ok := v.([]interface{}); ok {
		return Array
	}
	if _, ok := v.(map[string]interface{}); ok {
		return Object
	}

	if val, ok := v.(string); ok {
        if _, err := time.Parse(time.RFC3339, val); err == nil {
            return DateTime
        }
		return String
	}
	return Any
}
