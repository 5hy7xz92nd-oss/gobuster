package libgobuster

import (
	"errors"
	"time"
)

// OutputFormat defines the output format for results
type OutputFormat int

const (
	// OutputFormatText is the default human-readable output
	OutputFormatText OutputFormat = iota
	// OutputFormatJSON outputs one JSON object per result
	OutputFormatJSON
)

// ParseOutputFormat converts a string to an OutputFormat
func ParseOutputFormat(s string) (OutputFormat, error) {
	switch s {
	case "text", "":
		return OutputFormatText, nil
	case "json":
		return OutputFormatJSON, nil
	default:
		return OutputFormatText, errors.New("output-format must be one of: text, json")
	}
}

// Options holds all options that can be passed to libgobuster
type Options struct {
	Threads             int
	Debug               bool
	Wordlist            string
	WordlistOffset      int
	PatternFile         string
	DiscoverPatternFile string
	Patterns            []string
	DiscoverPatterns    []string
	OutputFilename      string
	OutputFormat        OutputFormat
	OutputAppend        bool
	NoProgress          bool
	NoError             bool
	Quiet               bool
	Delay               time.Duration
	RateLimit           int
	NoDuplicates        bool
}
