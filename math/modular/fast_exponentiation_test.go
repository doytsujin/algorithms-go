package modular

import "testing"

func TestFastExponentiation(t *testing.T) {
	t.Run("Testing base = 2 and exponent = 16", func(b *testing.T) {
		result := FastExponentiation(2, 16)
		if result != 65536 {
			t.Fatalf("Expected: 65536, Received: %d", result)
		}
	})
}
