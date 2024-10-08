package enums

type OptionsWindow int

const (
	OptionsWindowAggressiveResize OptionsWindow = iota + 1
	OptionsWindowAutomaticRename
	OptionsWindowAutomaticRenameFormat
	OptionsWindowClockModeColour
	OptionsWindowClockModeStyle
	OptionsWindowFillCharacter
	OptionsWindowMainPaneHeight
	OptionsWindowMainPaneWidth
	OptionsWindowCopyModeMatchStyle
	OptionsWindowCopyModeMarkStyle
	OptionsWindowCopyModeCurrentMatchStyle
	OptionsWindowModeKeys
	OptionsWindowModeStyle
	OptionsWindowMonitorActivity
	OptionsWindowMonitorBell
	OptionsWindowMonitorSilence
	OptionsWindowOtherPaneHeight
	OptionsWindowOtherPaneWidth
	OptionsWindowPaneActiveBorderStyle
	OptionsWindowPaneBaseIndex
	OptionsWindowPaneBorderFormat
	OptionsWindowPaneBorderIndicators
	OptionsWindowPaneBorderLines
	OptionsWindowPaneBorderStatus
	OptionsWindowPaneBorderStyle
	OptionsWindowPopupStyle
	OptionsWindowPopupBorderStyle
	OptionsWindowPopupBorderLines
	OptionsWindowWindowStatusActivityStyle
	OptionsWindowWindowStatusBellStyle
	OptionsWindowWindowStatusCurrentFormat
	OptionsWindowWindowStatusCurrentStyle
	OptionsWindowWindowStatusFormat
	OptionsWindowWindowStatusLastStyle
	OptionsWindowWindowStatusSeparator
	OptionsWindowWindowStatusStyle
	OptionsWindowWindowSize
	OptionsWindowWrapSearch
	OptionsWindowUnknown
)

const (
	OptionsWindowAggressiveResizeString          = "aggressive-resize"
	OptionsWindowAutomaticRenameString           = "automatic-rename"
	OptionsWindowAutomaticRenameFormatString     = "automatic-rename-format"
	OptionsWindowClockModeColourString           = "clock-mode-colour"
	OptionsWindowClockModeStyleString            = "clock-mode-style"
	OptionsWindowFillCharacterString             = "fill-character"
	OptionsWindowMainPaneHeightString            = "main-pane-height"
	OptionsWindowMainPaneWidthString             = "main-pane-width"
	OptionsWindowCopyModeMatchStyleString        = "copy-mode-match-style"
	OptionsWindowCopyModeMarkStyleString         = "copy-mode-mark-style"
	OptionsWindowCopyModeCurrentMatchStyleString = "copy-mode-current-match-style"
	OptionsWindowModeKeysString                  = "mode-keys"
	OptionsWindowModeStyleString                 = "mode-style"
	OptionsWindowMonitorActivityString           = "monitor-activity"
	OptionsWindowMonitorBellString               = "monitor-bell"
	OptionsWindowMonitorSilenceString            = "monitor-silence"
	OptionsWindowOtherPaneHeightString           = "other-pane-height"
	OptionsWindowOtherPaneWidthString            = "other-pane-width"
	OptionsWindowPaneActiveBorderStyleString     = "pane-active-border-style"
	OptionsWindowPaneBaseIndexString             = "pane-base-index"
	OptionsWindowPaneBorderFormatString          = "pane-border-format"
	OptionsWindowPaneBorderIndicatorsString      = "pane-border-indicators"
	OptionsWindowPaneBorderLinesString           = "pane-border-lines"
	OptionsWindowPaneBorderStatusString          = "pane-border-status"
	OptionsWindowPaneBorderStyleString           = "pane-border-style"
	OptionsWindowPopupStyleString                = "popup-style"
	OptionsWindowPopupBorderStyleString          = "popup-border-style"
	OptionsWindowPopupBorderLinesString          = "popup-border-lines"
	OptionsWindowWindowStatusActivityStyleString = "window-status-activity-style"
	OptionsWindowWindowStatusBellStyleString     = "window-status-bell-style"
	OptionsWindowWindowStatusCurrentFormatString = "window-status-current-format"
	OptionsWindowWindowStatusCurrentStyleString  = "window-status-current-style"
	OptionsWindowWindowStatusFormatString        = "window-status-format"
	OptionsWindowWindowStatusLastStyleString     = "window-status-last-style"
	OptionsWindowWindowStatusSeparatorString     = "window-status-separator"
	OptionsWindowWindowStatusStyleString         = "window-status-style"
	OptionsWindowWindowSizeString                = "window-size"
	OptionsWindowWrapSearchString                = "wrap-search"
	OptionsWindowUnknownString                   = "unknown"
)

var OptionsWindowList = []string{
	OptionsWindowAggressiveResizeString,
	OptionsWindowAutomaticRenameString,
	OptionsWindowAutomaticRenameFormatString,
	OptionsWindowClockModeColourString,
	OptionsWindowClockModeStyleString,
	OptionsWindowFillCharacterString,
	OptionsWindowMainPaneHeightString,
	OptionsWindowMainPaneWidthString,
	OptionsWindowCopyModeMatchStyleString,
	OptionsWindowCopyModeMarkStyleString,
	OptionsWindowCopyModeCurrentMatchStyleString,
	OptionsWindowModeKeysString,
	OptionsWindowModeStyleString,
	OptionsWindowMonitorActivityString,
	OptionsWindowMonitorBellString,
	OptionsWindowMonitorSilenceString,
	OptionsWindowOtherPaneHeightString,
	OptionsWindowOtherPaneWidthString,
	OptionsWindowPaneActiveBorderStyleString,
	OptionsWindowPaneBaseIndexString,
	OptionsWindowPaneBorderFormatString,
	OptionsWindowPaneBorderIndicatorsString,
	OptionsWindowPaneBorderLinesString,
	OptionsWindowPaneBorderStatusString,
	OptionsWindowPaneBorderStyleString,
	OptionsWindowPopupStyleString,
	OptionsWindowPopupBorderStyleString,
	OptionsWindowPopupBorderLinesString,
	OptionsWindowWindowStatusActivityStyleString,
	OptionsWindowWindowStatusBellStyleString,
	OptionsWindowWindowStatusCurrentFormatString,
	OptionsWindowWindowStatusCurrentStyleString,
	OptionsWindowWindowStatusFormatString,
	OptionsWindowWindowStatusLastStyleString,
	OptionsWindowWindowStatusSeparatorString,
	OptionsWindowWindowStatusStyleString,
	OptionsWindowWindowSizeString,
	OptionsWindowWrapSearchString,
	OptionsWindowUnknownString,
}

func (o OptionsWindow) String() string {
	switch o {
	case OptionsWindowAggressiveResize:
		return OptionsWindowAggressiveResizeString
	case OptionsWindowAutomaticRename:
		return OptionsWindowAutomaticRenameString
	case OptionsWindowAutomaticRenameFormat:
		return OptionsWindowAutomaticRenameFormatString
	case OptionsWindowClockModeColour:
		return OptionsWindowClockModeColourString
	case OptionsWindowClockModeStyle:
		return OptionsWindowClockModeStyleString
	case OptionsWindowFillCharacter:
		return OptionsWindowFillCharacterString
	case OptionsWindowMainPaneHeight:
		return OptionsWindowMainPaneHeightString
	case OptionsWindowMainPaneWidth:
		return OptionsWindowMainPaneWidthString
	case OptionsWindowCopyModeMatchStyle:
		return OptionsWindowCopyModeMatchStyleString
	case OptionsWindowCopyModeMarkStyle:
		return OptionsWindowCopyModeMarkStyleString
	case OptionsWindowCopyModeCurrentMatchStyle:
		return OptionsWindowCopyModeCurrentMatchStyleString
	case OptionsWindowModeKeys:
		return OptionsWindowModeKeysString
	case OptionsWindowModeStyle:
		return OptionsWindowModeStyleString
	case OptionsWindowMonitorActivity:
		return OptionsWindowMonitorActivityString
	case OptionsWindowMonitorBell:
		return OptionsWindowMonitorBellString
	case OptionsWindowMonitorSilence:
		return OptionsWindowMonitorSilenceString
	case OptionsWindowOtherPaneHeight:
		return OptionsWindowOtherPaneHeightString
	case OptionsWindowOtherPaneWidth:
		return OptionsWindowOtherPaneWidthString
	case OptionsWindowPaneActiveBorderStyle:
		return OptionsWindowPaneActiveBorderStyleString
	case OptionsWindowPaneBaseIndex:
		return OptionsWindowPaneBaseIndexString
	case OptionsWindowPaneBorderFormat:
		return OptionsWindowPaneBorderFormatString
	case OptionsWindowPaneBorderIndicators:
		return OptionsWindowPaneBorderIndicatorsString
	case OptionsWindowPaneBorderLines:
		return OptionsWindowPaneBorderLinesString
	case OptionsWindowPaneBorderStatus:
		return OptionsWindowPaneBorderStatusString
	case OptionsWindowPaneBorderStyle:
		return OptionsWindowPaneBorderStyleString
	case OptionsWindowPopupStyle:
		return OptionsWindowPopupStyleString
	case OptionsWindowPopupBorderStyle:
		return OptionsWindowPopupBorderStyleString
	case OptionsWindowPopupBorderLines:
		return OptionsWindowPopupBorderLinesString
	case OptionsWindowWindowStatusActivityStyle:
		return OptionsWindowWindowStatusActivityStyleString
	case OptionsWindowWindowStatusBellStyle:
		return OptionsWindowWindowStatusBellStyleString
	case OptionsWindowWindowStatusCurrentFormat:
		return OptionsWindowWindowStatusCurrentFormatString
	case OptionsWindowWindowStatusCurrentStyle:
		return OptionsWindowWindowStatusCurrentStyleString
	case OptionsWindowWindowStatusFormat:
		return OptionsWindowWindowStatusFormatString
	case OptionsWindowWindowStatusLastStyle:
		return OptionsWindowWindowStatusLastStyleString
	case OptionsWindowWindowStatusSeparator:
		return OptionsWindowWindowStatusSeparatorString
	case OptionsWindowWindowStatusStyle:
		return OptionsWindowWindowStatusStyleString
	case OptionsWindowWindowSize:
		return OptionsWindowWindowSizeString
	case OptionsWindowWrapSearch:
		return OptionsWindowWrapSearchString
	}

	return OptionsWindowUnknownString
}

func (o OptionsWindow) IsKnown(s string) bool {
	return o.FromString(s) != OptionsWindowUnknown
}

func (o OptionsWindow) FromString(s string) OptionsWindow {
	switch s {
	case OptionsWindowAggressiveResizeString:
		return OptionsWindowAggressiveResize
	case OptionsWindowAutomaticRenameString:
		return OptionsWindowAutomaticRename
	case OptionsWindowAutomaticRenameFormatString:
		return OptionsWindowAutomaticRenameFormat
	case OptionsWindowClockModeColourString:
		return OptionsWindowClockModeColour
	case OptionsWindowClockModeStyleString:
		return OptionsWindowClockModeStyle
	case OptionsWindowFillCharacterString:
		return OptionsWindowFillCharacter
	case OptionsWindowMainPaneHeightString:
		return OptionsWindowMainPaneHeight
	case OptionsWindowMainPaneWidthString:
		return OptionsWindowMainPaneWidth
	case OptionsWindowCopyModeMatchStyleString:
		return OptionsWindowCopyModeMatchStyle
	case OptionsWindowCopyModeMarkStyleString:
		return OptionsWindowCopyModeMarkStyle
	case OptionsWindowCopyModeCurrentMatchStyleString:
		return OptionsWindowCopyModeCurrentMatchStyle
	case OptionsWindowModeKeysString:
		return OptionsWindowModeKeys
	case OptionsWindowModeStyleString:
		return OptionsWindowModeStyle
	case OptionsWindowMonitorActivityString:
		return OptionsWindowMonitorActivity
	case OptionsWindowMonitorBellString:
		return OptionsWindowMonitorBell
	case OptionsWindowMonitorSilenceString:
		return OptionsWindowMonitorSilence
	case OptionsWindowOtherPaneHeightString:
		return OptionsWindowOtherPaneHeight
	case OptionsWindowOtherPaneWidthString:
		return OptionsWindowOtherPaneWidth
	case OptionsWindowPaneActiveBorderStyleString:
		return OptionsWindowPaneActiveBorderStyle
	case OptionsWindowPaneBaseIndexString:
		return OptionsWindowPaneBaseIndex
	case OptionsWindowPaneBorderFormatString:
		return OptionsWindowPaneBorderFormat
	case OptionsWindowPaneBorderIndicatorsString:
		return OptionsWindowPaneBorderIndicators
	case OptionsWindowPaneBorderLinesString:
		return OptionsWindowPaneBorderLines
	case OptionsWindowPaneBorderStatusString:
		return OptionsWindowPaneBorderStatus
	case OptionsWindowPaneBorderStyleString:
		return OptionsWindowPaneBorderStyle
	case OptionsWindowPopupStyleString:
		return OptionsWindowPopupStyle
	case OptionsWindowPopupBorderStyleString:
		return OptionsWindowPopupBorderStyle
	case OptionsWindowPopupBorderLinesString:
		return OptionsWindowPopupBorderLines
	case OptionsWindowWindowStatusActivityStyleString:
		return OptionsWindowWindowStatusActivityStyle
	case OptionsWindowWindowStatusBellStyleString:
		return OptionsWindowWindowStatusBellStyle
	case OptionsWindowWindowStatusCurrentFormatString:
		return OptionsWindowWindowStatusCurrentFormat
	case OptionsWindowWindowStatusCurrentStyleString:
		return OptionsWindowWindowStatusCurrentStyle
	case OptionsWindowWindowStatusFormatString:
		return OptionsWindowWindowStatusFormat
	case OptionsWindowWindowStatusLastStyleString:
		return OptionsWindowWindowStatusLastStyle
	case OptionsWindowWindowStatusSeparatorString:
		return OptionsWindowWindowStatusSeparator
	case OptionsWindowWindowStatusStyleString:
		return OptionsWindowWindowStatusStyle
	case OptionsWindowWindowSizeString:
		return OptionsWindowWindowSize
	case OptionsWindowWrapSearchString:
		return OptionsWindowWrapSearch
	}

	return OptionsWindowUnknown
}

func (o OptionsWindow) GetValidator(name string) (ValidatorFunc, bool) {
	if out, ok := OptionsWindowValidators[name]; ok {
		return out, true
	}

	return nil, false
}

var OptionsWindowValidators = map[string]ValidatorFunc{
	OptionsWindowAggressiveResizeString:      validatorToggle,
	OptionsWindowAutomaticRenameString:       validatorToggle,
	OptionsWindowClockModeColourString:       validatorColour,
	OptionsWindowFillCharacterString:         validatorNonEmpty,
	OptionsWindowMainPaneHeightString:        validatorDimension,
	OptionsWindowMainPaneWidthString:         validatorDimension,
	OptionsWindowMonitorActivityString:       validatorToggle,
	OptionsWindowMonitorBellString:           validatorToggle,
	OptionsWindowMonitorSilenceString:        validatorIsNumber,
	OptionsWindowOtherPaneHeightString:       validatorDimension,
	OptionsWindowOtherPaneWidthString:        validatorDimension,
	OptionsWindowPaneBaseIndexString:         validatorIsNumber,
	OptionsWindowWrapSearchString:            validatorToggle,
	OptionsWindowWindowStatusSeparatorString: validatorDefault,
	OptionsWindowClockModeStyleString:        validatorContains("12", "24"),
	OptionsWindowModeKeysString:              validatorContains("vi", "emacs"),
	OptionsWindowPaneBorderIndicatorsString:  validatorContains("off", "colour", "arrows", "both"),
	OptionsWindowPaneBorderLinesString:       validatorContains("single", "double", "heavy", "simple", "number"),
	OptionsWindowPaneBorderStatusString:      validatorContains("off", "top", "bottom"),
	OptionsWindowPopupBorderLinesString:      validatorContains("single", "rounded", "double", "heavy", "simple", "padded", "none"),
	OptionsWindowWindowSizeString:            validatorContains("largest", "smallest", "manual", "latest"),

	// FORMAT options are supported, but not yet validated properly:
	OptionsWindowWindowStatusFormatString:        validatorDefault,
	OptionsWindowWindowStatusCurrentFormatString: validatorDefault,
	OptionsWindowPaneBorderFormatString:          validatorDefault,
	OptionsWindowAutomaticRenameFormatString:     validatorDefault,

	// STYLE options are supported, but not yet validated properly:
	OptionsWindowWindowStatusLastStyleString:     validatorDefault,
	OptionsWindowWindowStatusStyleString:         validatorDefault,
	OptionsWindowWindowStatusCurrentStyleString:  validatorDefault,
	OptionsWindowWindowStatusActivityStyleString: validatorDefault,
	OptionsWindowWindowStatusBellStyleString:     validatorDefault,
	OptionsWindowPaneActiveBorderStyleString:     validatorDefault,
	OptionsWindowModeStyleString:                 validatorDefault,
	OptionsWindowCopyModeCurrentMatchStyleString: validatorDefault,
	OptionsWindowCopyModeMatchStyleString:        validatorDefault,
	OptionsWindowCopyModeMarkStyleString:         validatorDefault,
	OptionsWindowPaneBorderStyleString:           validatorDefault,
	OptionsWindowPopupStyleString:                validatorDefault,
	OptionsWindowPopupBorderStyleString:          validatorDefault,
}
