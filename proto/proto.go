package proto

import "strconv"

const (
	ProtocolVersion                = "6.0"
	ProtocolTerminatingCharacters  = "\r\n"
	ProtocolEndOfMessageCharacters = "!ENDMSG!"
	ProtocolNoDataCharacters       = "!NO_DATA!"
	ProtocolSyntaxErrorCharacters  = "!SYNTAX_ERROR!"
	ProtocolLineFeedCharacter      = '\n'
	ProtocolDelimiterCharacter     = ','
	PrototolErrorCharacter         = 'E'
)

func Split(s string, expected []string) []string {
	mark := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ProtocolDelimiterCharacter:
			expected = append(expected, s[mark:i])
			mark = i + 1
		}
	}

	if mark == len(s) {
		expected = append(expected, "")
	} else {
		expected = append(expected, s[mark:])
	}
	return expected
}

func CharAt(s string, index int) byte {
	if index < 0 || index >= len(s) {
		return 0
	}
	return s[index]
}

func SetClientName(b []byte, name string) []byte {
	b = append(b, "S,SET CLIENT NAME,"...)
	b = append(b, name...)
	b = append(b, ProtocolTerminatingCharacters...)
	return b
}

func SetProtocol(b []byte, version string) []byte {
	b = append(b, "S,SET PROTOCOL,"...)
	b = append(b, version...)
	b = append(b, ProtocolTerminatingCharacters...)
	return b
}

func ParseInt(str string) int64 {
	v, _ := strconv.ParseInt(str, 10, 64)
	return v
}

func ParseFloat(str string) float64 {
	v, _ := strconv.ParseFloat(str, 64)
	return v
}

func Value(values []string, index int) string {
	if index < 0 || index >= len(values) {
		return ""
	}
	return values[index]
}

func IsNumeral(b byte) bool {
	switch b {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}

func NumeralToWord(b byte) string {
	switch b {
	case '0':
		return "Zero"
	case '1':
		return "One"
	case '2':
		return "Two"
	case '3':
		return "Three"
	case '4':
		return "Four"
	case '5':
		return "Five"
	case '6':
		return "Six"
	case '7':
		return "Seven"
	case '8':
		return "Eight"
	case '9':
		return "Nine"
	}
	return ""
}
