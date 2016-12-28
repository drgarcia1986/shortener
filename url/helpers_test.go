package url

import "testing"

func TestGenerateShortWithLenght(t *testing.T) {
	expectedLength := 4
	short := GenerateShort(expectedLength)
	if len(short) != expectedLength {
		t.Errorf("Expected %d, got %d", expectedLength, len(short))
	}
}

func TestGenerateShortUnique(t *testing.T) {
	length := 4
	short := GenerateShort(length)

	for i := 0; i < 5; i++ {
		if short == GenerateShort(length) {
			t.Error("Generate same short")
			break
		}
	}
}
