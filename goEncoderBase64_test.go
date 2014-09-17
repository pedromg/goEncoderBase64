package goEncoderBase64

import (
	"testing"
	"encoding/base64"
	"github.com/pedromg/goEncoderBase64"
)

const msg_short = "This is Your Amiga Speaking."
const msg_long  = "I think computer viruses should count as life. I think it says something about human nature that the only form of life we have created so far is purely destructive. We've created life in our own image.\nStephen Hawking"

func TestBase64MimeEncoderShort( t *testing.T) {
	// encode the message
	a := goEncoderBase64.Base64MimeEncoder(msg_short)

	b, err := base64.StdEncoding.DecodeString(a)
	if err != nil {
		t.Error(err)
	}
	// compare decoded message
	if msg_short != string(b) {
		t.Error("Different messages!")
	} else {
		t.Log("Short Message encoding passed!")
	}
}

func TestBase64MimeEncoderLong( t *testing.T) {
	// encode the message
	a := goEncoderBase64.Base64MimeEncoder(msg_long)

	b, err := base64.StdEncoding.DecodeString(a)
	if err != nil {
		t.Error(err)
	}
	// compare decoded message
	if msg_long != string(b) {
		t.Error("Different messages!")
	} else {
		t.Log("Long Message encoding passed!")
	}
}
