package log

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

func Console() io.Writer {
	return zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    true,
		TimeFormat: time.RFC3339,
		FormatLevel: func(i interface{}) string {
			return fmt.Sprintf("[%s]", i)
		},
		FormatMessage: func(i interface{}) string {
			if i != nil {
				return fmt.Sprintf("| %s ", i)
			}

			return "|"
		},
		FormatCaller: func(i interface{}) string {
			paths := strings.Split(i.(string), "/")
			size := len(paths)
			caller := fmt.Sprintf("%s/%s/%s", paths[size-3], paths[size-2], paths[size-1])
			return fmt.Sprintf("%-32s", caller)
		},
	}
}
