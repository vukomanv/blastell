package rsepparser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type TokenType int

const (
	STRING TokenType = iota
	ARRAY
	INT
	FLOAT
	NIL
	BOOL
)

type Token struct {
	TokenType TokenType
	Value     interface{}
}

func Parse(input []byte) ([]Token, error) {
	splitInput := strings.Split(string(input), "\\r\\n")
	if len(splitInput) < 2 {
		return []Token{}, errors.New("falty input provided")
	}

	splitInput = splitInput[:len(splitInput)-1]

	fmt.Println(splitInput)
	tokenArr := []Token{}

	v := splitInput[0]
	if v[0] != '*' {
		return []Token{}, errors.New("error: input is invalid, number of array elements if wrong")
	}
	v = splitInput[1]
	if v[0] != '$' {
		return []Token{}, errors.New("error: input is invalid, no command given")
	}

	for i := 0; i < len(splitInput); i++ {
		v := splitInput[i]
		switch v[0] {
		case ':':
			val, err := strconv.Atoi(v[1:])
			if err != nil {
				return []Token{}, errors.New("error: " + v + " is formatted incorrectly")
			}
			tokenArr = append(tokenArr, Token{INT, val})
		case ',':
			val, err := strconv.ParseFloat(v[1:], 32)
			if err != nil {
				return []Token{}, errors.New("error: " + v + " is formatted incorrectly")
			}
			tokenArr = append(tokenArr, Token{FLOAT, val})
		case '#':
			inputVal := v[1:]
			if inputVal != "f" && inputVal != "t" {
				return []Token{}, errors.New("error: " + v + " is formatted incorrectly")
			}
			val := true
			if inputVal == "f" {
				val = false
			}
			tokenArr = append(tokenArr, Token{BOOL, val})
		case '$':
			strLen, err := strconv.Atoi(v[1:])
			if err != nil {
				return []Token{}, errors.New("error: " + v + " is formatted incorrectly")
			}

			str := splitInput[i+1]
			i++
			if len(str) != int(strLen) {
				return []Token{}, errors.New("error: length of '" + str + "' doesn't match the defined length " + v[1:])
			}
			tokenArr = append(tokenArr, Token{STRING, str})
		case '*':
			arrLen, err := strconv.Atoi(v[1:])
			if err != nil {
				return []Token{}, errors.New("error: " + v + " is formatted incorrectly")
			}

			tokenArr = append(tokenArr, Token{ARRAY, arrLen})
		default:
			return []Token{}, errors.New("error: " + v + " is formatted incorrectly")
		}
	}

	if tokenArr[0].Value != len(tokenArr)-1 {
		return []Token{}, fmt.Errorf("error: defined array length %v doesn't match the actual array length %v", tokenArr[0].Value, len(tokenArr))
	}

	fmt.Println(tokenArr)
	// check if each arr token has the correct number of tokens "after" it

	return tokenArr, nil
}

// +OK\r\n
// -ErrorMessage\r\n
// :-50\r\n
// ,2.5\r\n
// #t\r\n
// $5\r\nHello\r\n
// *3\r\n:12\r\n$7\r\nWelcome\r\n (*0\r\n)
// _\r\n (nil)

// split by \r\n -> []{type, value}
// example
// *3\r\n*2\r\n:5\r\n$3\r\nHey\r\n*1\r\n,2.3\r\n*0\r\n
// ARRAY 3, ARRAY 2, INT 5, STRING Hey, ARRAY 1, DOUBLE 2.3, ARRAY 0
//
// +PING\r\n
// SIMPLE_STRING PING
// if slice len == 1

// bytes -> tokens -> array | simple str | err
