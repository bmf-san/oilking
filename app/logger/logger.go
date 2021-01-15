package logger

import (
	"encoding/json"
	"io"
	"os"
	"reflect"
	"time"
)

const (
	// LevelInfo is something notable infomation.
	LevelInfo int = iota
	// LevelWarn is warning.
	LevelWarn
	// LevelError is unexpected runtime error.
	LevelError
	// LevelFatal is an abend error.
	LevelFatal
	// LevelTextInfo is the text for info.
	LevelTextInfo = "info"
	// LevelTextWarn is the text for warn.
	LevelTextWarn = "warn"
	// LevelTextError is the text for error.
	LevelTextError = "error"
	// LevelTextFatal is the text for fatal.
	LevelTextFatal = "fatal"
	// LabelBalance is a label for balance.
	LabelBalance = "[Balance]"
	// LabelBalanceHistory is a label for balance history.
	LabelBalanceHistory = "[Balance History]"
	// LabelChat is a label for chat.
	LabelChat = "[Chat]"
	// LabelChildOrder is a label for child order.
	LabelChildOrder = "[ChildOrder]"
	// LabelCollateral is a label for collateral.
	LabelCollateral = "[Collateral]"
	// LabelExecution is a label for execution.
	LabelExecution = "[Execution]"
	// LabelPosition is a label for position.
	LabelPosition = "[Position]"
	// LabelTicker is a label for Ticker.
	LabelTicker = "[Ticker]"
	// LabelTradingCommission is a label for trading commission.
	LabelTradingCommission = "[Trading Commission]"
)

var levelText = map[int]string{
	LevelInfo:  LevelTextInfo,
	LevelWarn:  LevelTextWarn,
	LevelError: LevelTextError,
	LevelFatal: LevelTextFatal,
}

// Fields is a fields
type Fields map[string]interface{}

// LevelText returns the text of level.
func LevelText(level int) string {
	return levelText[level]
}

// A Logger represents a logger.
type Logger struct {
	out       io.Writer
	threshold int
	location  *time.Location
	entry     json.RawMessage
}

// An Entry represents an entry for basic log.
type Entry struct {
	Level   string    `json:"level"`
	Time    time.Time `json:"time"` // UTC
	Message string    `json:"message"`
}

// An AccessLogEntry represents an entry for access log.
type AccessLogEntry struct {
	Level  string    `json:"level"`
	Time   time.Time `json:"time"` // UTC
	Method string    `json:"method"`
	URL    string    `json:"url"`
	Body   string    `json:"body"`
}

// A TradingLogEntry represents an entry for trading log.
type TradingLogEntry struct {
	Level  string    `json:"level"`
	Time   time.Time `json:"time"` // UTC
	Label  string    `json:"label"`
	Action string    `json:"action"`
}

// NewLogger creates a logger.
func NewLogger(threshold int, location *time.Location) *Logger {
	return &Logger{
		out:       io.Writer(os.Stdout),
		threshold: threshold,
		location:  location,
	}
}

// Info outputs a info level log.
func (l *Logger) Info(entry interface{}) {
	if LevelInfo >= l.threshold {
		l.OutputJSON(LevelText(LevelInfo), entry)
	}
}

// Warn outputs a warn level log.
func (l *Logger) Warn(entry interface{}) {
	if LevelWarn >= l.threshold {
		l.OutputJSON(LevelText(LevelWarn), entry)
	}
}

// Error outputs a error level log.
func (l *Logger) Error(entry interface{}) {
	if LevelError >= l.threshold {
		l.OutputJSON(LevelText(LevelError), entry)
	}
}

// Fatal outputs a fatal level log.
func (l *Logger) Fatal(entry interface{}) {
	if LevelFatal >= l.threshold {
		l.OutputJSON(LevelText(LevelFatal), entry)
	}
}

// OutputJSON outputs logs.
func (l *Logger) OutputJSON(level string, entry interface{}) error {
	switch reflect.TypeOf(entry).Name() {
	case "Entry":
		if e, ok := entry.(Entry); ok {
			e.Level = level
			e.Time = time.Now().UTC().In(l.location)
			ebytes, err := json.Marshal(e)
			if err != nil {
				return err
			}
			l.entry = ebytes
		}
	case "AccessLogEntry":
		if e, ok := entry.(AccessLogEntry); ok {
			e.Level = level
			e.Time = time.Now().UTC().In(l.location)
			ebytes, err := json.Marshal(e)
			if err != nil {
				return err
			}
			l.entry = ebytes
		}
	case "TradingLogEntry":
		if e, ok := entry.(TradingLogEntry); ok {
			e.Level = level
			e.Time = time.Now().UTC().In(l.location)
			ebytes, err := json.Marshal(e)
			if err != nil {
				return err
			}
			l.entry = ebytes
		}
	}

	bytes, err := json.Marshal(l.entry)
	if err != nil {
		return err
	}
	bytes = append(bytes, '\n')
	_, err = l.out.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
