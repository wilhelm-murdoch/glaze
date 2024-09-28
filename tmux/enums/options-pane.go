package enums

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

var OptionsPaneValidators = map[string]ValidatorFunc{
	OptionsPaneAllowRenameString:        validatorToggle,
	OptionsPaneAllowSetTitleString:      validatorToggle,
	OptionsPaneAlternateScreenString:    validatorToggle,
	OptionsPaneCursorColourString:       validatorColour,
	OptionsPaneRemainOnExitFormatString: validatorDefault,
	OptionsPaneScrollOnClearString:      validatorToggle,
	OptionsPaneSynchronizePanesString:   validatorToggle,
	OptionsPaneRemainOnExitString:       validatorContains("on", "off", "failed"),
	OptionsPaneAllowPassthroughString:   validatorContains("on", "off", "all"),

	// STYLE options are supported, but not yet validated properly:
	OptionsPaneWindowActiveStyleString: validatorDefault,
	OptionsPaneWindowStyleString:       validatorDefault,
	OptionsPaneCursorStyleString:       validatorDefault,
}
