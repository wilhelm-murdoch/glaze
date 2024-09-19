package enums

type Layout int

func (l Layout) String() string {
	switch l {
	case LayoutEvenHorizontal:
		return LayoutEvenHorizontalString
	case LayoutEvenVertical:
		return LayoutEvenVerticalString
	case LayoutMainHorizontal:
		return LayoutMainHorizontalString
	case LayoutMainVertical:
		return LayoutMainVerticalString
	case LayoutTiled:
		return LayoutTiledString
	}

	return LayoutUnknownString
}

const (
	LayoutEvenHorizontal Layout = iota + 1
	LayoutEvenVertical
	LayoutMainHorizontal
	LayoutMainVertical
	LayoutTiled
	LayoutUnknown
)

const (
	LayoutEvenHorizontalString = "even-horizontal"
	LayoutEvenVerticalString   = "even-vertical"
	LayoutMainHorizontalString = "main-horizontal"
	LayoutMainVerticalString   = "main-vertical"
	LayoutTiledString          = "tiled"
	LayoutUnknownString        = "unknown"
)

var LayoutList = []string{
	LayoutEvenHorizontalString,
	LayoutEvenVerticalString,
	LayoutMainHorizontalString,
	LayoutMainVerticalString,
	LayoutTiledString,
}

func LayoutFromString(s string) Layout {
	switch s {
	case LayoutEvenHorizontalString:
		return LayoutEvenHorizontal
	case LayoutEvenVerticalString:
		return LayoutEvenVertical
	case LayoutMainHorizontalString:
		return LayoutMainHorizontal
	case LayoutMainVerticalString:
		return LayoutMainVertical
	case LayoutTiledString:
		return LayoutTiled
	}

	return LayoutUnknown
}
