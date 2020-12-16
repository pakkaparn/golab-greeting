package greeting

import (
	"testing"
)

func TestHellloToReturnString(t *testing.T) {
	message, err := Hello("Pan")

	if "Hi, Pan. Welcome!" != message && err == nil {
		t.Error("Hello of Pan should return 'Hi, Pan. Welcome!' but return ", message)
	}
}

func TestHellloToReturnError(t *testing.T) {
	message, err := Hello("")

	if len(message) == 0 && err == nil {
		t.Error("Hello should Error but return ", message)
	}
}
