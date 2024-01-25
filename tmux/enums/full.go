package enums

type Split int

func (s Split) String() string {
	switch s {
	case SplitHorizontal:
		return SplitHorizontalString
	case SplitVertical:
		return SplitVerticalString
	}

	return SplitUnknownString
}

const (
	SplitHorizontal Split = iota
	SplitVertical
	SplitUnknown
)

const (
	SplitHorizontalString = "horizontal"
	SplitVerticalString   = "vertical"
	SplitUnknownString    = "unknown"
)

var SplitList = []string{
	SplitHorizontalString,
	SplitVerticalString,
}

func SplitFromString(s string) Split {
	switch s {
	case SplitHorizontalString:
		return SplitHorizontal
	case SplitVerticalString:
		return SplitVertical
	}

	return SplitUnknown
}
