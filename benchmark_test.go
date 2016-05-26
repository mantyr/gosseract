package gosseract_test

import (
	"testing"

	"github.com/mantyr/gosseract"
)

func BenchmarkMust(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gosseract.Must(gosseract.Params{
			Src:       "./.samples/png/sample000.png",
			Languages: "eng",
		})
	}
}
