package enums

type Placement int

func (p Placement) String() string {
	switch p {
	case PlacementAbove:
		return PlacementAboveString
	case PlacementLeft:
		return PlacementLeftString
	}

	return PlacementUnknownString
}

const (
	PlacementAbove Placement = iota
	PlacementLeft
	PlacementUnknown
)

const (
	PlacementAboveString   = "above"
	PlacementLeftString    = "left"
	PlacementUnknownString = "unknown"
)

var PlacementList = []string{
	PlacementAboveString,
	PlacementLeftString,
}

func PlacementFromString(s string) Placement {
	switch s {
	case PlacementAboveString:
		return PlacementAbove
	case PlacementLeftString:
		return PlacementLeft
	}

	return PlacementUnknown
}
