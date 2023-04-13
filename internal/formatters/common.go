package formatters

import (
	"io"

	"github.com/spf13/viper"
)

type Formattable interface {
	WriteAsTable(io.Writer) error
	WriteAsJSON(io.Writer) error
}

// PrintTo ensures that all printable responses implement all possible Formattable methods
func PrintTo(w io.Writer, value Formattable) error {

	switch viper.GetString("output") { // just dispatch the correct method depending on cfg
	case "json":
		return value.WriteAsJSON(w)
	default:
		return value.WriteAsTable(w)
	}

}