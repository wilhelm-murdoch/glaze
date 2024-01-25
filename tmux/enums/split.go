package enums

type Full int

func (f Full) String() string {
	switch f {
	case FullHeight:
		return FullHeightString
	case FullWidth:
		return FullWidthString
	}

	return FullUnknownString
}

const (
	FullHeight Full = iota
	FullWidth
	FullUnknown
)

const (
	FullHeightString  = "height"
	FullWidthString   = "width"
	FullUnknownString = "unknown"
)

var FullList = []string{
	FullHeightString,
	FullWidthString,
}

func FullFromString(s string) Full {
	switch s {
	case FullHeightString:
		return FullHeight
	case FullWidthString:
		return FullWidth
	}

	return FullUnknown
}
