import (
	"github.com/fuzzbuzz/platform/sfz/fuzz"
	uuid "github.com/google/uuid"
)

func FuzzParseBytes(f *fuzz.F) {
	b := f.Bytes("uuid")
	uuid.ParseBytes(b.Get())
}

