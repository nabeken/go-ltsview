package ltsview

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func DoTestLTSView(t *testing.T, f func(v *LTSView)) {
	exampleLTSV, err := os.Open("example.ltsv")
	if err != nil {
		t.Fatal(err)
	}
	defer exampleLTSV.Close()
	v := &LTSView{
		Reader: exampleLTSV,
	}
	f(v)
}

func TestLTSView(t *testing.T) {
	expected := `---
host: 192.168.0.3
req: GET / HTTP/1.1
status: 200
time: 17/Aug/2014:06:27:16 +0000
ua: Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)
---
host: 192.168.0.4
req: POST / HTTP/1.1
status: 302
time: 18/Aug/2014:06:27:16 +0000
ua: Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)
`

	actual := &bytes.Buffer{}
	DoTestLTSView(t, func(v *LTSView) {
		v.Writer = actual
		v.Start()
	})
	assert.Equal(t, expected, actual.String())
}

func TestLTSView_Keys(t *testing.T) {
	expected := `---
host: 192.168.0.3
status: 200
---
host: 192.168.0.4
status: 302
`

	actual := &bytes.Buffer{}
	DoTestLTSView(t, func(v *LTSView) {
		v.Writer = actual
		v.Keys = ParseKeysByFlag("host,status")
		v.Start()
	})
	assert.Equal(t, expected, actual.String())
}

func TestLTSView_Ikeys(t *testing.T) {
	expected := `---
host: 192.168.0.3
req: GET / HTTP/1.1
status: 200
---
host: 192.168.0.4
req: POST / HTTP/1.1
status: 302
`

	actual := &bytes.Buffer{}
	DoTestLTSView(t, func(v *LTSView) {
		v.Writer = actual
		v.Ikeys = ParseKeysByFlag("ua,time")
		v.Start()
	})
	assert.Equal(t, expected, actual.String())
}
