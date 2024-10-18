package gen

import (
	"encoding/json"
	"fmt"
	"github.com/MarcoVitangeli/jsontypes/writer"
	"os/exec"
)

func Gen(input []byte, opts ...GenOption) error {
	conf := GenConfiguration{
		Depth: 10,
	}

	for _, opt := range opts {
		opt(&conf)
	}

	var c interface{}

	if err := json.Unmarshal(input, &c); err != nil {
		return fmt.Errorf("error parsing json object: %w", err)
	}

	switch v := c.(type) {
	case map[string]interface{}:
		return generateObjectType(v, &conf)
	case []interface{}:
		return fmt.Errorf("json arrays are not supported as root")
	default:
		return fmt.Errorf("invalid json object provided")
	}
}

func generateObjectType(v map[string]interface{}, conf *GenConfiguration) error {
	outPath := "type_gen.go"
	tw := writer.NewTypeWriter(outPath)
	tw.Init()

	err := tw.Write(v, conf.Depth)
	if err != nil {
		return err
	}

	err = tw.Close()
	if err != nil {
		return err
	}

	cmd := exec.Command("gofmt", "-w", outPath)

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
