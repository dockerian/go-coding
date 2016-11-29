// +build skip demo pipe stream test

package demo

// see https://gist.github.com/ImJasonH/da090817f7d513441d09

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"testing"
)

const dataSize = 1000000

func BenchmarkBuffer(b *testing.B) {
	m := makeData()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		_ = json.NewEncoder(&buf).Encode(m)
		_, _ = io.Copy(ioutil.Discard, &buf)
	}
}

func BenchmarkPipe(b *testing.B) {
	m := makeData()
	for i := 0; i < b.N; i++ {
		pr, pw := io.Pipe()

		go func() {
			pw.CloseWithError(json.NewEncoder(pw).Encode(m))
		}()
		_, _ = io.Copy(ioutil.Discard, pr)
	}
}

func TestPipeStream(t *testing.T) {
	t.Logf("Preparing test data %v ...\n", dataSize)
	data := makeData()
	t.Log("Testing pipeStream ...")
	pipeStream(data)
}

func makeData() map[string]string {
	m := map[string]string{}
	for i := 0; i < dataSize; i++ {
		m[fmt.Sprintf("key-%d", i)] = fmt.Sprintf("val-%d", i)
	}
	return m
}
