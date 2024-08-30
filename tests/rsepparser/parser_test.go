package rsepparser_test

import (
	"math"
	"testing"

	"github.com/vukomanv/blastell/internal/rsepparser"
)

func TestParse(t *testing.T) {
	cases := []struct {
		name     string
		input    []byte
		expected []rsepparser.Token
	}{
		{
			"strings pass",
			[]byte{42, 52, 92, 114, 92, 110, 36, 53, 92, 114, 92, 110, 112, 114, 105, 110, 116, 92, 114, 92, 110, 36, 52, 92, 114, 92, 110, 82, 65, 78, 68, 92, 114, 92, 110, 36, 51, 92, 114, 92, 110, 100, 79, 71, 92, 114, 92, 110, 36, 48, 92, 114, 92, 110, 92, 114, 92, 110},
			[]rsepparser.Token{
				{TokenType: rsepparser.ARRAY, Value: 4},
				{TokenType: rsepparser.STRING, Value: "print"},
				{TokenType: rsepparser.STRING, Value: "RAND"},
				{TokenType: rsepparser.STRING, Value: "dOG"},
				{TokenType: rsepparser.STRING, Value: ""},
			},
		},
		{
			"integers pass",
			[]byte{42, 52, 92, 114, 92, 110, 36, 53, 92, 114, 92, 110, 112, 114, 105, 110, 116, 92, 114, 92, 110, 58, 52, 51, 92, 114, 92, 110, 58, 45, 52, 51, 50, 92, 114, 92, 110, 58, 48, 92, 114, 92, 110},
			[]rsepparser.Token{
				{TokenType: rsepparser.ARRAY, Value: 4},
				{TokenType: rsepparser.STRING, Value: "print"},
				{TokenType: rsepparser.INT, Value: 43},
				{TokenType: rsepparser.INT, Value: -432},
				{TokenType: rsepparser.INT, Value: 0},
			},
		},
		{
			"floats pass",
			[]byte{42, 53, 92, 114, 92, 110, 36, 53, 92, 114, 92, 110, 112, 114, 105, 110, 116, 92, 114, 92, 110, 44, 53, 50, 49, 46, 48, 50, 51, 92, 114, 92, 110, 44, 45, 48, 46, 52, 57, 92, 114, 92, 110, 44, 48, 92, 114, 92, 110, 44, 49, 49, 92, 114, 92, 110},
			[]rsepparser.Token{
				{TokenType: rsepparser.ARRAY, Value: 5},
				{TokenType: rsepparser.STRING, Value: "print"},
				{TokenType: rsepparser.FLOAT, Value: 521.023},
				{TokenType: rsepparser.FLOAT, Value: -0.49},
				{TokenType: rsepparser.FLOAT, Value: 11},
				{TokenType: rsepparser.FLOAT, Value: 0},
			},
		},
	}

	for _, tc := range cases {
		expectedOutput := tc.expected
		testName := tc.name

		output, err := rsepparser.Parse(tc.input)
		if err != nil {
			t.Fatalf("method: Parse, case: %v, %s", testName, err.Error())
		}

		if len(expectedOutput) != len(output) {
			t.Fatalf("method: Parse, case: %v, test error: %v token array length doesn't match the expected %v length", testName, len(output), len(expectedOutput))
		}

		for i := 0; i < len(expectedOutput); i++ {
			p := output[i]
			e := expectedOutput[i]
			if p.TokenType != e.TokenType {
				t.Fatalf("method: Parse, case: %v, test error: %v token type doesn't match the expected token type %v", testName, p.TokenType, e.TokenType)
			}

			if p.TokenType == rsepparser.FLOAT && compareFloats(e.Value, p.Value, t) {
				continue
			}
			if p.Value != e.Value {
				t.Fatalf("method: Parse, case: %v, test error: %v token value doesn't match the expected token value %v", testName, p.Value, e.Value)
			}
		}
	}
}

func compareFloats(expectedValue interface{}, actualValue interface{}, t *testing.T) bool {
	ev, ok := expectedValue.(float64)
	if ok == false {
		t.Fatalf("value %v isn't a float64 and it should be", expectedValue)
	}
	av, ok := actualValue.(float64)
	if ok == false {
		t.Fatalf("value %v isn't a float64 and it should be", actualValue)
	}
	return math.Abs(ev-av) < 0.1
}
