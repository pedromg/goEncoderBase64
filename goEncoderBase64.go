package goEncoderBase64

import (
	"encoding/base64"
)

var delimiter = []byte("\r\n")
const maxLen = 76

// The Base64MimeEncoder function returns a Base64 encoded []byte for 
// MIME (RFC 2045) with max line length maxLen (+CR/LF)
func Base64MimeEncoder(b string) (m string) {

	m = base64.StdEncoding.EncodeToString([]byte(b))
	the_len := len(m)

	if (the_len <= maxLen) {
		return m
	}

	new_m := []byte(m)

	// set the slice capacity to the slice len + each newline delimiters
	m1 := make([]byte, 0, the_len+(len(delimiter)*int(the_len/maxLen)))
	ii := 0
	for i := 0; i < int(the_len/maxLen); i++ {
		m1 = append(m1, new_m[i*maxLen:(i+1)*maxLen]...)
		m1 = append(m1, delimiter...)
		ii++
	}
	m1 = append(m1, new_m[ii*maxLen:the_len]...)
	m = string(m1)
	return m
}

