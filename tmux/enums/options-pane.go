package enums

import "slices"

type OptionsPane int

const (
	OptionsPaneAllowPassthrough = iota + 1
	OptionsPaneAllowRename
	OptionsPaneAllowSetTitle
	OptionsPaneAlternateScreen
	OptionsPaneCursorColour
	OptionsPanePaneColours
	OptionsPaneCursorStyle
	OptionsPaneRemainOnExit
	OptionsPaneRemainOnExitFormat
	OptionsPaneScrollOnClear
	OptionsPaneSynchronizePanes
	OptionsPaneWindowActiveStyle
	OptionsPaneWindowStyle
	OptionsPaneUnknown
)

const (
	OptionsPaneAllowPassthroughString   = "allow-passthrough"
	OptionsPaneAllowRenameString        = "allow-rename"
	OptionsPaneAllowSetTitleString      = "allow-set-title"
	OptionsPaneAlternateScreenString    = "alternate-screen"
	OptionsPaneCursorColourString       = "cursor-colour"
	OptionsPanePaneColoursString        = "pane-colours"
	OptionsPaneCursorStyleString        = "cursor-style"
	OptionsPaneRemainOnExitString       = "remain-on-exit"
	OptionsPaneRemainOnExitFormatString = "remain-on-exit-format"
	OptionsPaneScrollOnClearString      = "scroll-on-clear"
	OptionsPaneSynchronizePanesString   = "synchronize-panes"
	OptionsPaneWindowActiveStyleString  = "window-active-style"
	OptionsPaneWindowStyleString        = "window-style"
	OptionsPaneUnknownString            = "unknown"
)

var OptionsPaneList = []string{
	OptionsPaneAllowPassthroughString,
	OptionsPaneAllowRenameString,
	OptionsPaneAllowSetTitleString,
	OptionsPaneAlternateScreenString,
	OptionsPaneCursorColourString,
	OptionsPanePaneColoursString,
	OptionsPaneCursorStyleString,
	OptionsPaneRemainOnExitString,
	OptionsPaneRemainOnExitFormatString,
	OptionsPaneScrollOnClearString,
	OptionsPaneSynchronizePanesString,
	OptionsPaneWindowActiveStyleString,
	OptionsPaneWindowStyleString,
	OptionsPaneUnknownString,
}

func (o OptionsPane) String() string {
	switch o {
	case OptionsPaneAllowPassthrough:
		return OptionsPaneAllowPassthroughString
	case OptionsPaneAllowRename:
		return OptionsPaneAllowRenameString
	case OptionsPaneAllowSetTitle:
		return OptionsPaneAllowSetTitleString
	case OptionsPaneAlternateScreen:
		return OptionsPaneAlternateScreenString
	case OptionsPaneCursorColour:
		return OptionsPaneCursorColourString
	case OptionsPanePaneColours:
		return OptionsPanePaneColoursString
	case OptionsPaneCursorStyle:
		return OptionsPaneCursorStyleString
	case OptionsPaneRemainOnExit:
		return OptionsPaneRemainOnExitString
	case OptionsPaneRemainOnExitFormat:
		return OptionsPaneRemainOnExitFormatString
	case OptionsPaneScrollOnClear:
		return OptionsPaneScrollOnClearString
	case OptionsPaneSynchronizePanes:
		return OptionsPaneSynchronizePanesString
	case OptionsPaneWindowActiveStyle:
		return OptionsPaneWindowActiveStyleString
	case OptionsPaneWindowStyle:
		return OptionsPaneWindowStyleString
	}

	return OptionsPaneUnknownString
}

func OptionsPaneFromString(s string) OptionsPane {
	switch s {
	case OptionsPaneAllowPassthroughString:
		return OptionsPaneAllowPassthrough
	case OptionsPaneAllowRenameString:
		return OptionsPaneAllowRename
	case OptionsPaneAllowSetTitleString:
		return OptionsPaneAllowSetTitle
	case OptionsPaneAlternateScreenString:
		return OptionsPaneAlternateScreen
	case OptionsPaneCursorColourString:
		return OptionsPaneCursorColour
	case OptionsPanePaneColoursString:
		return OptionsPanePaneColours
	case OptionsPaneCursorStyleString:
		return OptionsPaneCursorStyle
	case OptionsPaneRemainOnExitString:
		return OptionsPaneRemainOnExit
	case OptionsPaneRemainOnExitFormatString:
		return OptionsPaneRemainOnExitFormat
	case OptionsPaneScrollOnClearString:
		return OptionsPaneScrollOnClear
	case OptionsPaneSynchronizePanesString:
		return OptionsPaneSynchronizePanes
	case OptionsPaneWindowActiveStyleString:
		return OptionsPaneWindowActiveStyle
	case OptionsPaneWindowStyleString:
		return OptionsPaneWindowStyle
	}

	return OptionsPaneUnknown
}

var OptionsPaneValidators = map[string]func(v string) bool{
	OptionsPaneAllowPassthroughString: func(v string) bool {
		return slices.Contains([]string{"on", "off", "all"}, v)
	},
	OptionsPaneAllowRenameString:        func(v string) bool { return true },
	OptionsPaneAllowSetTitleString:      func(v string) bool { return true },
	OptionsPaneAlternateScreenString:    func(v string) bool { return true },
	OptionsPaneCursorColourString:       func(v string) bool { return true },
	OptionsPanePaneColoursString:        func(v string) bool { return true },
	OptionsPaneCursorStyleString:        func(v string) bool { return true },
	OptionsPaneRemainOnExitString:       func(v string) bool { return true },
	OptionsPaneRemainOnExitFormatString: func(v string) bool { return true },
	OptionsPaneScrollOnClearString:      func(v string) bool { return true },
	OptionsPaneSynchronizePanesString:   func(v string) bool { return true },
	OptionsPaneWindowActiveStyleString:  func(v string) bool { return true },
	OptionsPaneWindowStyleString:        func(v string) bool { return true },
}
