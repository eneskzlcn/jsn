package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	indent = "    "
	prefix = ""
)

func GetFormattedJSON(data []byte) (string, error) {
	var out bytes.Buffer
	if err := json.Indent(&out, data, "", "    "); err != nil {
		return "", err
	}
	resp := out.String()
	return fmt.Sprintf("\n%s\n", resp), nil
}
