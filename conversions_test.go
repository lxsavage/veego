package veego

import "testing"

func TestRGB(t *testing.T) {
	expected := colorRGB(0xFEDFED)
	var r, g, b uint8 = 0xFE, 0xDF, 0xED

	of := RGB(r, g, b)

	if of != expected {
		t.Fatalf("RGB(%d, %d, %d) = %x, expected %x", r, g, b, of, expected)
	}
}

func TestRGBComponents(t *testing.T) {
	var expectedR, expectedG, expectedB uint8 = 0xFE, 0xDF, 0xED
	of := RGB(expectedR, expectedG, expectedB)

	ofR, ofG, ofB := of.Components()

	if ofR != expectedR || ofG != expectedG || ofB != expectedB {
		t.Fatalf(
			"RGB(%d, %d, %d) = %d, %d, %d, expected %d, %d, %d",
			expectedR,
			expectedG,
			expectedB,
			ofR,
			ofG,
			ofB,
			expectedR,
			expectedG,
			expectedB,
		)
	}
}
