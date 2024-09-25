package enums

import (
	"slices"
	"strconv"
	"strings"
)

type OptionsWindow int

const (
	OptionsWindowAggressiveResize = iota + 1
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

func OptionsWindowFromString(s string) OptionsWindow {
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

var OptionsWindowValidators = map[string]func(v string) (bool, []string){
	OptionsWindowAggressiveResizeString: func(v string) (bool, []string) {
		choices := []string{"on", "off"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsWindowAutomaticRenameString: func(v string) (bool, []string) {
		choices := []string{"on", "off"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsWindowAutomaticRenameFormatString: func(v string) (bool, []string) {
		// Any string here will do for now. Tmux will report the error and it will
		// bubble up to the user. May add additional sub-validators in the future.
		return true, nil
	},
	OptionsWindowClockModeColourString: func(v string) (bool, []string) { return true, nil },
	OptionsWindowClockModeStyleString: func(v string) (bool, []string) {
		choices := []string{"12", "24"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsWindowFillCharacterString: func(v string) (bool, []string) {
		if len(strings.TrimSpace(v)) != 1 {
			return false, nil
		}

		return true, nil
	},
	OptionsWindowMainPaneHeightString:            func(v string) (bool, []string) { return true, nil },
	OptionsWindowMainPaneWidthString:             func(v string) (bool, []string) { return true, nil },
	OptionsWindowCopyModeMatchStyleString:        func(v string) (bool, []string) { return true, nil },
	OptionsWindowCopyModeMarkStyleString:         func(v string) (bool, []string) { return true, nil },
	OptionsWindowCopyModeCurrentMatchStyleString: func(v string) (bool, []string) { return true, nil },
	OptionsWindowModeKeysString: func(v string) (bool, []string) {
		choices := []string{"vi", "emacs"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsWindowModeStyleString: func(v string) (bool, []string) { return true, nil },
	OptionsWindowMonitorActivityString: func(v string) (bool, []string) {
		choices := []string{"on", "off"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsWindowMonitorBellString: func(v string) (bool, []string) {
		choices := []string{"on", "off"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsWindowMonitorSilenceString: func(v string) (bool, []string) {
		if _, err := strconv.Atoi(v); err != nil {
			return false, nil
		}

		return true, nil
	},
	OptionsWindowOtherPaneHeightString:       func(v string) (bool, []string) { return true, nil },
	OptionsWindowOtherPaneWidthString:        func(v string) (bool, []string) { return true, nil },
	OptionsWindowPaneActiveBorderStyleString: func(v string) (bool, []string) { return true, nil },
	OptionsWindowPaneBaseIndexString: func(v string) (bool, []string) {
		if _, err := strconv.Atoi(v); err != nil {
			return false, nil
		}

		return true, nil
	},
	OptionsWindowPaneBorderFormatString: func(v string) (bool, []string) { return true, nil },
	OptionsWindowPaneBorderIndicatorsString: func(v string) (bool, []string) {
		choices := []string{"off", "colour", "arrows", "both"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsWindowPaneBorderLinesString: func(v string) (bool, []string) {
		choices := []string{"single", "double", "heavy", "simple", "number"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsWindowPaneBorderStatusString: func(v string) (bool, []string) {
		choices := []string{"off", "top", "bottom"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsWindowPaneBorderStyleString:  func(v string) (bool, []string) { return true, nil },
	OptionsWindowPopupStyleString:       func(v string) (bool, []string) { return true, nil },
	OptionsWindowPopupBorderStyleString: func(v string) (bool, []string) { return true, nil },
	OptionsWindowPopupBorderLinesString: func(v string) (bool, []string) {
		choices := []string{"single", "rounded", "double", "heavy", "simple", "padded", "none"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsWindowWindowStatusActivityStyleString: func(v string) (bool, []string) { return true, nil },
	OptionsWindowWindowStatusBellStyleString:     func(v string) (bool, []string) { return true, nil },
	OptionsWindowWindowStatusCurrentFormatString: func(v string) (bool, []string) { return true, nil },
	OptionsWindowWindowStatusCurrentStyleString:  func(v string) (bool, []string) { return true, nil },
	OptionsWindowWindowStatusFormatString:        func(v string) (bool, []string) { return true, nil },
	OptionsWindowWindowStatusLastStyleString:     func(v string) (bool, []string) { return true, nil },
	OptionsWindowWindowStatusSeparatorString:     func(v string) (bool, []string) { return true, nil },
	OptionsWindowWindowStatusStyleString:         func(v string) (bool, []string) { return true, nil },
	OptionsWindowWindowSizeString: func(v string) (bool, []string) {
		choices := []string{"largest", "smallest", "manual", "latest"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsWindowWrapSearchString: func(v string) (bool, []string) {
		choices := []string{"on", "off"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
}
