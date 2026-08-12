package gobustertftp

import (
	"bytes"
	"fmt"

	"github.com/fatih/color"
)

var green = color.New(color.FgGreen).FprintfFunc()

// Result represents a single result
type Result struct {
	Filename     string
	Size         int64
	ErrorMessage string
}

// ResultToJSON returns a JSON-encodable representation of the result.
func (r Result) ResultToJSON() (interface{}, error) {
	return map[string]interface{}{
		"filename":      r.Filename,
		"size":          r.Size,
		"error_message": r.ErrorMessage,
	}, nil
}

// ResultToString converts the Result to its textual representation
func (r Result) ResultToString() (string, error) {
	buf := &bytes.Buffer{}

	green(buf, "%s", r.Filename)
	if r.Size > 0 {
		if _, err := fmt.Fprintf(buf, " [%d]", r.Size); err != nil {
			return "", err
		}
	}

	if _, err := fmt.Fprintf(buf, "\n"); err != nil {
		return "", err
	}

	s := buf.String()
	return s, nil
}
