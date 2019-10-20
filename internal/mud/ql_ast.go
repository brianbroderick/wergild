package mud

import (
	"errors"
	"math"
	"strings"
	"time"
)

// Statement represents a single command
type Statement interface {
	String() string // mostly used for testing
	execute()
	setPlayer(*Player)
	// RequiredPrivileges() (ExecutionPrivileges, error)
}

// DataType represents the primitive data types available in InfluxQL.
type DataType int

const (
	// Unknown primitive data type.
	Unknown DataType = 0
	// Float means the data type is a float.
	Float DataType = 1
	// Integer means the data type is an integer.
	Integer DataType = 2
	// String means the data type is a string of text.
	String DataType = 3
	// Boolean means the data type is a boolean.
	Boolean DataType = 4
	// Time means the data type is a time.
	Time DataType = 5
	// Duration means the data type is a duration of time.
	Duration DataType = 6
	// Tag means the data type is a tag.
	Tag DataType = 7
	// AnyField means the data type is any field.
	AnyField DataType = 8
	// Unsigned means the data type is an unsigned integer.
	Unsigned DataType = 9
)

const (
	// MinTime is the minumum time that can be represented.
	//
	// 1677-09-21 00:12:43.145224194 +0000 UTC
	//
	// The two lowest minimum integers are used as sentinel values.  The
	// minimum value needs to be used as a value lower than any other value for
	// comparisons and another separate value is needed to act as a sentinel
	// default value that is unusable by the user, but usable internally.
	// Because these two values need to be used for a special purpose, we do
	// not allow users to write points at these two times.
	MinTime = int64(math.MinInt64) + 2

	// MaxTime is the maximum time that can be represented.
	//
	// 2262-04-11 23:47:16.854775806 +0000 UTC
	//
	// The highest time represented by a nanosecond needs to be used for an
	// exclusive range in the shard group, so the maximum time needs to be one
	// less than the possible maximum number of nanoseconds representable by an
	// int64 so that we don't lose a point at that one time.
	MaxTime = int64(math.MaxInt64) - 1
)

var (
	// ErrInvalidTime is returned when the timestamp string used to
	// compare against time field is invalid.
	ErrInvalidTime = errors.New("invalid timestamp string")
)

// InspectDataType returns the data type of a given value.
func InspectDataType(v interface{}) DataType {
	switch v.(type) {
	case float64:
		return Float
	case int64, int32, int:
		return Integer
	case string:
		return String
	case bool:
		return Boolean
	case uint64:
		return Unsigned
	case time.Time:
		return Time
	case time.Duration:
		return Duration
	default:
		return Unknown
	}
}

// DataTypeFromString returns a data type given the string representation of that
// data type.
func DataTypeFromString(s string) DataType {
	switch s {
	case "float":
		return Float
	case "integer":
		return Integer
	case "unsigned":
		return Unsigned
	case "string":
		return String
	case "boolean":
		return Boolean
	case "time":
		return Time
	case "duration":
		return Duration
	case "tag":
		return Tag
	case "field":
		return AnyField
	default:
		return Unknown
	}
}

// LessThan returns true if the other DataType has greater precedence than the
// current data type. Unknown has the lowest precedence.
//
// NOTE: This is not the same as using the `<` or `>` operator because the
// integers used decrease with higher precedence, but Unknown is the lowest
// precedence at the zero value.
func (d DataType) LessThan(other DataType) bool {
	if d == Unknown {
		return true
	} else if d == Unsigned {
		return other != Unknown && other <= Integer
	} else if other == Unsigned {
		return d >= String
	}
	return other != Unknown && other < d
}

var (
	zeroFloat64  interface{} = float64(0)
	zeroInt64    interface{} = int64(0)
	zeroUint64   interface{} = uint64(0)
	zeroString   interface{} = ""
	zeroBoolean  interface{} = false
	zeroTime     interface{} = time.Time{}
	zeroDuration interface{} = time.Duration(0)
)

// Zero returns the zero value for the DataType.
// The return value of this method, when sent back to InspectDataType,
// may not produce the same value.
func (d DataType) Zero() interface{} {
	switch d {
	case Float:
		return zeroFloat64
	case Integer:
		return zeroInt64
	case Unsigned:
		return zeroUint64
	case String, Tag:
		return zeroString
	case Boolean:
		return zeroBoolean
	case Time:
		return zeroTime
	case Duration:
		return zeroDuration
	}
	return nil
}

// String returns the human-readable string representation of the DataType.
func (d DataType) String() string {
	switch d {
	case Float:
		return "float"
	case Integer:
		return "integer"
	case Unsigned:
		return "unsigned"
	case String:
		return "string"
	case Boolean:
		return "boolean"
	case Time:
		return "time"
	case Duration:
		return "duration"
	case Tag:
		return "tag"
	case AnyField:
		return "field"
	}
	return "unknown"
}

// Statements represents a list of statements.
type Statements []Statement

// String returns a string representation of the statements.
func (a Statements) String() string {
	var str []string
	for _, stmt := range a {
		str = append(str, stmt.String())
	}
	return strings.Join(str, ";\n")
}
