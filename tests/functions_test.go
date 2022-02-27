package tests

import (
	"cli_converter/functions"
	"testing"
)

func TestConvertion(t *testing.T) {
	t.Run("Decimal To Binary", func(t *testing.T) {
		var a, res = "23", "10111"

		realRes, err := functions.DecimalToBinary(a)
		if err != nil {
			t.Error("Error returned:", err)
		}

		if res != realRes {
			t.Errorf("Expected result: %s != %s real result", res, realRes)
		}

		t.Run("Decimal to Binary, error test", func(t *testing.T) {
			_, err := functions.DecimalToBinary("should return an error")
			if err == nil {
				t.Error("Didn't return an error")
			}
		})
	})

	t.Run("Binary to Decimal", func(t *testing.T) {
		var a, res = "10111", "23"

		realRes, err := functions.BinaryToDecimal(a)
		if err != nil {
			t.Error("Error returned:", err)
		}

		if res != realRes {
			t.Errorf("Expected result: %s != %s real result", res, realRes)
		}

		t.Run("Binary to Decimal, error test", func(t *testing.T) {
			_, err := functions.BinaryToDecimal("should return an error")
			if err == nil {
				t.Error("Didn't return an error 1")
			}
		})
	})

	t.Run("Decimal to octal", func(t *testing.T) {
		var a, res = "23", "27"

		realRes, err := functions.DecimalToOctal(a)
		if err != nil {
			t.Error("Error returned", err)
		}

		if res != realRes {
			t.Errorf("Expected result: %s != %s real result", res, realRes)
		}

		t.Run("Decimal to octal, error test", func(t *testing.T) {
			_, err := functions.DecimalToOctal("should return an error")
			if err == nil {
				t.Error("Didn't return an error")
			}
		})
	})

	t.Run("Octal to decimal", func(t *testing.T) {
		var a, res = "27", "23"

		realRes, err := functions.OctalToDecimal(a)
		if err != nil {
			t.Error("Error returned", err)
		}

		if res != realRes {
			t.Errorf("Expected result: %s != %s real result", res, realRes)
		}

		t.Run("Octal to decimal, error test", func(t *testing.T) {
			_, err := functions.OctalToDecimal("should return an error")
			if err == nil {
				t.Error("Didn't return an error")
			}
		})
	})

	t.Run("Decimal to hex", func(t *testing.T) {
		var a, res = "100", "64"

		realRes, err := functions.DecimalToHex(a)
		if err != nil {
			t.Error("Error returned", err)
		}

		if res != realRes {
			t.Errorf("Expected result: %s != %s hex result", res, realRes)
		}

		t.Run("Decimal to hex, error test", func(t *testing.T) {
			_, err := functions.DecimalToHex("should return an error")
			if err == nil {
				t.Error("Didn't return an error")
			}
		})
	})

	t.Run("Hex to decimal", func(t *testing.T) {
		var a, res = "64", "100"

		realRes, err := functions.HexToDecimal(a)
		if err != nil {
			t.Error("Error returned", err)
		}

		if res != realRes {
			t.Errorf("Expected result: %s != %s hex result", res, realRes)
		}
	})
}
