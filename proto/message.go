package proto

import "time"

type ErrorMessage struct {
	Error string
}

func (e *ErrorMessage) Unmarshal(values []string) {
	if len(values) > 1 {
		e.Error = values[1]
	}
}

type SystemMessage struct {
	Type    string
	Message []string
}

func (s *SystemMessage) Unmarshal(values []string) {
	if len(values) > 1 {
		s.Type = values[1]
	}
	s.Message = values
}

type TimestampMessage struct {
	Timestamp time.Time
}

func (t *TimestampMessage) Unmarshal(values []string) {
	if len(values) > 1 {
		t.Timestamp = ParseDateTime(values[1])
	}
}

type SymbolNotFoundMessage struct {
	Symbol string
}

func (s *SymbolNotFoundMessage) Unmarshal(values []string) {
	if len(values) > 1 {
		s.Symbol = values[1]
	}
}
