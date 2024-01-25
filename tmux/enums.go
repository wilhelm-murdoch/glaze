package tmux

const (
	PlacementAbove = "above"
	PlacementLeft  = "left"

	FullHeight = "height"
	FullWidth  = "width"

	SplitHorizontal = "horizontal"
	SplitVertical   = "vertical"

	LayoutEventHorizontal = "even-horizontal"
	LayoutEvenVertical    = "even-vertical"
	LayoutMainHorizontal  = "main-horizontal"
	LayoutMainVertical    = "main-vertical"
	LayoutTiled           = "tiled"
	LayoutUnknown         = "unknown"
)

var (
	PlacementList = []string{
		PlacementAbove,
		PlacementLeft,
	}

	FullList = []string{
		FullHeight,
		FullWidth,
	}

	SplitList = []string{
		SplitHorizontal,
		SplitVertical,
	}

	LayoutList = []string{
		LayoutEventHorizontal,
		LayoutEvenVertical,
		LayoutMainHorizontal,
		LayoutMainVertical,
		LayoutTiled,
	}
)

func Contains(list []string, value string) bool {
	for _, v := range list {
		if string(v) == value {
			return true
		}
	}

	return false
}
