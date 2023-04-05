package log

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/rs/zerolog"
)

func ColorConsole() io.Writer {
	return zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
		FormatLevel: func(i interface{}) string {
			level := strings.ToUpper(i.(string)[:1])
			switch level {
			case "T": // trace
				return fmt.Sprintf("[%s]", level)
			case "D": // debug
				return color.GreenString("[%s]", level)
			case "I": // info
				return color.BlueString("[%s]", level)
			case "W": // warning
				return color.YellowString("[%s]", level)
			case "E": // error
				return color.RedString("[%s]", level)
			case "C": // critical
				return color.MagentaString("[%s]", level)

			default:
				return fmt.Sprintf("[%s]", level)
			}
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
