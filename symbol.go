package spin

import "bytes"

const (
	green = iota
	blue
	yellow
	red
)
const (
	successSymbol = "✔"
	infoSymbol    = "ℹ"
	warningSymbol = "⚠"
	errorSymbol   = "✖"
)

func newSymbol(color int, text string) string {
	buffer := new(bytes.Buffer)
	switch color {
	case blue:
		buffer.WriteString("\u001b[1;34m")

	case green:
		buffer.WriteString("\u001b[1;32m")

	case yellow:
		buffer.WriteString("\u001b[1;33m")

	case red:
		buffer.WriteString("\u001b[1;31m")
	}
	buffer.WriteString(text)
	buffer.WriteString("\u001b[0m")
	return buffer.String()
}
