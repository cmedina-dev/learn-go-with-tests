package dependency_injection

import "bytes"

func Greet(buffer *bytes.Buffer, name string) {
	buffer.WriteString("Hello, " + name)
}
