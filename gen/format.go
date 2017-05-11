package gen

import (
	"bytes"
	"strconv"
)

// returns string with byte representation of string
func stringAsBytes(s string) string {
	buf := &bytes.Buffer{}
	buf.Write([]byte("[]byte{"))
	totalLen := len(s)

	for i := 0; i < totalLen; i++ {
		if i != 0 {
			buf.WriteByte(',')
		}

		buf.Write([]byte("byte("))
		buf.WriteString(strconv.Itoa(int(s[i])))
		buf.WriteByte(')')
	}

	buf.WriteByte('}')

	return buf.String()
}
