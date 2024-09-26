package enums

import (
	"slices"
	"strconv"
)

type OptionsSession int

const (
	OptionsSessionActivityAction = iota + 1
	OptionsSessionAssumePasteTime
	OptionsSessionBaseIndex
	OptionsSessionBellAction
	OptionsSessionDefaultCommand
	OptionsSessionDefaultShell
	OptionsSessionDefaultSize
	OptionsSessionDestroyUnattached
	OptionsSessionDetachOnDestroy
	OptionsSessionDisplayPanesActiveColour
	OptionsSessionDisplayPanesColour
	OptionsSessionDisplayPanesTime
	OptionsSessionDisplayTime
	OptionsSessionHistoryLimit
	OptionsSessionKeyTable
	OptionsSessionLockAfterTime
	OptionsSessionLockCommand
	OptionsSessionMenuStyle
	OptionsSessionMenuSelectedStyle
	OptionsSessionMenuBorderStyle
	OptionsSessionMenuBorderLines
	OptionsSessionMessageCommandStyle
	OptionsSessionMessageLine
	OptionsSessionMessageStyle
	OptionsSessionMouse
	OptionsSessionPrefix
	OptionsSessionPrefix2
	OptionsSessionRenumberWindows
	OptionsSessionRepeatTime
	OptionsSessionSetTitles
	OptionsSessionSilenceAction
	OptionsSessionStatus
	OptionsSessionStatusInterval
	OptionsSessionStatusJustify
	OptionsSessionStatusKeys
	OptionsSessionStatusLeft
	OptionsSessionStatusLeftLength
	OptionsSessionStatusLeftStyle
	OptionsSessionStatusPosition
	OptionsSessionStatusRight
	OptionsSessionStatusRightLength
	OptionsSessionStatusRightStyle
	OptionsSessionStatusStyle
	OptionsSessionVisualActivity
	OptionsSessionVisualBell
	OptionsSessionVisualSilence
	OptionsSessionWordSeparators
	OptionsSessionUnknown
)

const (
	OptionsSessionActivityActionString           = "activity-action"
	OptionsSessionAssumePasteTimeString          = "assume-paste-time"
	OptionsSessionBaseIndexString                = "base-index"
	OptionsSessionBellActionString               = "bell-action"
	OptionsSessionDefaultCommandString           = "default-command"
	OptionsSessionDefaultShellString             = "default-shell"
	OptionsSessionDefaultSizeString              = "default-size"
	OptionsSessionDestroyUnattachedString        = "destroy-unattached"
	OptionsSessionDetachOnDestroyString          = "detach-on-destroy"
	OptionsSessionDisplayPanesActiveColourString = "display-panes-active-colour"
	OptionsSessionDisplayPanesColourString       = "display-panes-colour"
	OptionsSessionDisplayPanesTimeString         = "display-panes-time"
	OptionsSessionDisplayTimeString              = "display-time"
	OptionsSessionHistoryLimitString             = "history-limit"
	OptionsSessionKeyTableString                 = "key-table"
	OptionsSessionLockAfterTimeString            = "lock-after-time"
	OptionsSessionLockCommandString              = "lock-command"
	OptionsSessionMenuStyleString                = "menu-style"
	OptionsSessionMenuSelectedStyleString        = "menu-selected-style"
	OptionsSessionMenuBorderStyleString          = "menu-border-style"
	OptionsSessionMenuBorderLinesString          = "menu-border-lines"
	OptionsSessionMessageCommandStyleString      = "message-command-style"
	OptionsSessionMessageLineString              = "message-line"
	OptionsSessionMessageStyleString             = "message-style"
	OptionsSessionMouseString                    = "mouse"
	OptionsSessionPrefixString                   = "prefix"
	OptionsSessionPrefix2String                  = "prefix2"
	OptionsSessionRenumberWindowsString          = "renumber-windows"
	OptionsSessionRepeatTimeString               = "repeat-time"
	OptionsSessionSetTitlesString                = "set-titles"
	OptionsSessionSilenceActionString            = "silence-action"
	OptionsSessionStatusString                   = "status"
	OptionsSessionStatusIntervalString           = "status-interval"
	OptionsSessionStatusJustifyString            = "status-justify"
	OptionsSessionStatusKeysString               = "status-keys"
	OptionsSessionStatusLeftString               = "status-left"
	OptionsSessionStatusLeftLengthString         = "status-left-length"
	OptionsSessionStatusLeftStyleString          = "status-left-style"
	OptionsSessionStatusPositionString           = "status-position"
	OptionsSessionStatusRightString              = "status-right"
	OptionsSessionStatusRightLengthString        = "status-right-length"
	OptionsSessionStatusRightStyleString         = "status-right-style"
	OptionsSessionStatusStyleString              = "status-style"
	OptionsSessionVisualActivityString           = "visual-activity"
	OptionsSessionVisualBellString               = "visual-bell"
	OptionsSessionVisualSilenceString            = "visual-silence"
	OptionsSessionWordSeparatorsString           = "word-separators"
	OptionsSessionUnknownString                  = "unknown"
)

var OptionsSessionList = []string{
	OptionsSessionActivityActionString,
	OptionsSessionAssumePasteTimeString,
	OptionsSessionBaseIndexString,
	OptionsSessionBellActionString,
	OptionsSessionDefaultCommandString,
	OptionsSessionDefaultShellString,
	OptionsSessionDefaultSizeString,
	OptionsSessionDestroyUnattachedString,
	OptionsSessionDetachOnDestroyString,
	OptionsSessionDisplayPanesActiveColourString,
	OptionsSessionDisplayPanesColourString,
	OptionsSessionDisplayPanesTimeString,
	OptionsSessionDisplayTimeString,
	OptionsSessionHistoryLimitString,
	OptionsSessionKeyTableString,
	OptionsSessionLockAfterTimeString,
	OptionsSessionLockCommandString,
	OptionsSessionMenuStyleString,
	OptionsSessionMenuSelectedStyleString,
	OptionsSessionMenuBorderStyleString,
	OptionsSessionMenuBorderLinesString,
	OptionsSessionMessageCommandStyleString,
	OptionsSessionMessageLineString,
	OptionsSessionMessageStyleString,
	OptionsSessionMouseString,
	OptionsSessionPrefixString,
	OptionsSessionPrefix2String,
	OptionsSessionRenumberWindowsString,
	OptionsSessionRepeatTimeString,
	OptionsSessionSetTitlesString,
	OptionsSessionSilenceActionString,
	OptionsSessionStatusString,
	OptionsSessionStatusIntervalString,
	OptionsSessionStatusJustifyString,
	OptionsSessionStatusKeysString,
	OptionsSessionStatusLeftString,
	OptionsSessionStatusLeftLengthString,
	OptionsSessionStatusLeftStyleString,
	OptionsSessionStatusPositionString,
	OptionsSessionStatusRightString,
	OptionsSessionStatusRightLengthString,
	OptionsSessionStatusRightStyleString,
	OptionsSessionStatusStyleString,
	OptionsSessionVisualActivityString,
	OptionsSessionVisualBellString,
	OptionsSessionVisualSilenceString,
	OptionsSessionWordSeparatorsString,
}

func (o OptionsSession) String() string {
	switch o {
	case OptionsSessionActivityAction:
		return OptionsSessionActivityActionString
	case OptionsSessionAssumePasteTime:
		return OptionsSessionAssumePasteTimeString
	case OptionsSessionBaseIndex:
		return OptionsSessionBaseIndexString
	case OptionsSessionBellAction:
		return OptionsSessionBellActionString
	case OptionsSessionDefaultCommand:
		return OptionsSessionDefaultCommandString
	case OptionsSessionDefaultShell:
		return OptionsSessionDefaultShellString
	case OptionsSessionDefaultSize:
		return OptionsSessionDefaultSizeString
	case OptionsSessionDestroyUnattached:
		return OptionsSessionDestroyUnattachedString
	case OptionsSessionDetachOnDestroy:
		return OptionsSessionDetachOnDestroyString
	case OptionsSessionDisplayPanesActiveColour:
		return OptionsSessionDisplayPanesActiveColourString
	case OptionsSessionDisplayPanesColour:
		return OptionsSessionDisplayPanesColourString
	case OptionsSessionDisplayPanesTime:
		return OptionsSessionDisplayPanesTimeString
	case OptionsSessionDisplayTime:
		return OptionsSessionDisplayTimeString
	case OptionsSessionHistoryLimit:
		return OptionsSessionHistoryLimitString
	case OptionsSessionKeyTable:
		return OptionsSessionKeyTableString
	case OptionsSessionLockAfterTime:
		return OptionsSessionLockAfterTimeString
	case OptionsSessionLockCommand:
		return OptionsSessionLockCommandString
	case OptionsSessionMenuStyle:
		return OptionsSessionMenuStyleString
	case OptionsSessionMenuSelectedStyle:
		return OptionsSessionMenuSelectedStyleString
	case OptionsSessionMenuBorderStyle:
		return OptionsSessionMenuBorderStyleString
	case OptionsSessionMenuBorderLines:
		return OptionsSessionMenuBorderLinesString
	case OptionsSessionMessageCommandStyle:
		return OptionsSessionMessageCommandStyleString
	case OptionsSessionMessageLine:
		return OptionsSessionMessageLineString
	case OptionsSessionMessageStyle:
		return OptionsSessionMessageStyleString
	case OptionsSessionMouse:
		return OptionsSessionMouseString
	case OptionsSessionPrefix:
		return OptionsSessionPrefixString
	case OptionsSessionPrefix2:
		return OptionsSessionPrefix2String
	case OptionsSessionRenumberWindows:
		return OptionsSessionRenumberWindowsString
	case OptionsSessionRepeatTime:
		return OptionsSessionRepeatTimeString
	case OptionsSessionSetTitles:
		return OptionsSessionSetTitlesString
	case OptionsSessionSilenceAction:
		return OptionsSessionSilenceActionString
	case OptionsSessionStatus:
		return OptionsSessionStatusString
	case OptionsSessionStatusInterval:
		return OptionsSessionStatusIntervalString
	case OptionsSessionStatusJustify:
		return OptionsSessionStatusJustifyString
	case OptionsSessionStatusKeys:
		return OptionsSessionStatusKeysString
	case OptionsSessionStatusLeft:
		return OptionsSessionStatusLeftString
	case OptionsSessionStatusLeftLength:
		return OptionsSessionStatusLeftLengthString
	case OptionsSessionStatusLeftStyle:
		return OptionsSessionStatusLeftStyleString
	case OptionsSessionStatusPosition:
		return OptionsSessionStatusPositionString
	case OptionsSessionStatusRight:
		return OptionsSessionStatusRightString
	case OptionsSessionStatusRightLength:
		return OptionsSessionStatusRightLengthString
	case OptionsSessionStatusRightStyle:
		return OptionsSessionStatusRightStyleString
	case OptionsSessionStatusStyle:
		return OptionsSessionStatusStyleString
	case OptionsSessionVisualActivity:
		return OptionsSessionVisualActivityString
	case OptionsSessionVisualBell:
		return OptionsSessionVisualBellString
	case OptionsSessionVisualSilence:
		return OptionsSessionVisualSilenceString
	case OptionsSessionWordSeparators:
		return OptionsSessionWordSeparatorsString
	case OptionsSessionUnknown:
		return OptionsSessionUnknownString
	}

	return OptionsSessionUnknownString
}

func OptionsSessionFromString(s string) OptionsSession {
	switch s {
	case OptionsSessionActivityActionString:
		return OptionsSessionActivityAction
	case OptionsSessionAssumePasteTimeString:
		return OptionsSessionAssumePasteTime
	case OptionsSessionBaseIndexString:
		return OptionsSessionBaseIndex
	case OptionsSessionBellActionString:
		return OptionsSessionBellAction
	case OptionsSessionDefaultCommandString:
		return OptionsSessionDefaultCommand
	case OptionsSessionDefaultShellString:
		return OptionsSessionDefaultShell
	case OptionsSessionDefaultSizeString:
		return OptionsSessionDefaultSize
	case OptionsSessionDestroyUnattachedString:
		return OptionsSessionDestroyUnattached
	case OptionsSessionDetachOnDestroyString:
		return OptionsSessionDetachOnDestroy
	case OptionsSessionDisplayPanesActiveColourString:
		return OptionsSessionDisplayPanesActiveColour
	case OptionsSessionDisplayPanesColourString:
		return OptionsSessionDisplayPanesColour
	case OptionsSessionDisplayPanesTimeString:
		return OptionsSessionDisplayPanesTime
	case OptionsSessionDisplayTimeString:
		return OptionsSessionDisplayTime
	case OptionsSessionHistoryLimitString:
		return OptionsSessionHistoryLimit
	case OptionsSessionKeyTableString:
		return OptionsSessionKeyTable
	case OptionsSessionLockAfterTimeString:
		return OptionsSessionLockAfterTime
	case OptionsSessionLockCommandString:
		return OptionsSessionLockCommand
	case OptionsSessionMenuStyleString:
		return OptionsSessionMenuStyle
	case OptionsSessionMenuSelectedStyleString:
		return OptionsSessionMenuSelectedStyle
	case OptionsSessionMenuBorderStyleString:
		return OptionsSessionMenuBorderStyle
	case OptionsSessionMenuBorderLinesString:
		return OptionsSessionMenuBorderLines
	case OptionsSessionMessageCommandStyleString:
		return OptionsSessionMessageCommandStyle
	case OptionsSessionMessageLineString:
		return OptionsSessionMessageLine
	case OptionsSessionMessageStyleString:
		return OptionsSessionMessageStyle
	case OptionsSessionMouseString:
		return OptionsSessionMouse
	case OptionsSessionPrefixString:
		return OptionsSessionPrefix
	case OptionsSessionPrefix2String:
		return OptionsSessionPrefix2
	case OptionsSessionRenumberWindowsString:
		return OptionsSessionRenumberWindows
	case OptionsSessionRepeatTimeString:
		return OptionsSessionRepeatTime
	case OptionsSessionSetTitlesString:
		return OptionsSessionSetTitles
	case OptionsSessionSilenceActionString:
		return OptionsSessionSilenceAction
	case OptionsSessionStatusString:
		return OptionsSessionStatus
	case OptionsSessionStatusIntervalString:
		return OptionsSessionStatusInterval
	case OptionsSessionStatusJustifyString:
		return OptionsSessionStatusJustify
	case OptionsSessionStatusKeysString:
		return OptionsSessionStatusKeys
	case OptionsSessionStatusLeftString:
		return OptionsSessionStatusLeft
	case OptionsSessionStatusLeftLengthString:
		return OptionsSessionStatusLeftLength
	case OptionsSessionStatusLeftStyleString:
		return OptionsSessionStatusLeftStyle
	case OptionsSessionStatusPositionString:
		return OptionsSessionStatusPosition
	case OptionsSessionStatusRightString:
		return OptionsSessionStatusRight
	case OptionsSessionStatusRightLengthString:
		return OptionsSessionStatusRightLength
	case OptionsSessionStatusRightStyleString:
		return OptionsSessionStatusRightStyle
	case OptionsSessionStatusStyleString:
		return OptionsSessionStatusStyle
	case OptionsSessionVisualActivityString:
		return OptionsSessionVisualActivity
	case OptionsSessionVisualBellString:
		return OptionsSessionVisualBell
	case OptionsSessionVisualSilenceString:
		return OptionsSessionVisualSilence
	case OptionsSessionWordSeparatorsString:
		return OptionsSessionWordSeparators
	case OptionsSessionUnknownString:
		return OptionsSessionUnknown
	}

	return OptionsSessionUnknown
}

var OptionsSessionValidators = map[string]func(v string) (bool, []string){
	OptionsSessionActivityActionString: func(v string) (bool, []string) {
		return slices.Contains([]string{"any", "none", "current", "other"}, v), nil
	},
	OptionsSessionDetachOnDestroyString: func(v string) (bool, []string) {
		return slices.Contains([]string{"off", "on", "no-detatched", "previous", "next"}, v), nil
	},
	OptionsSessionDisplayPanesActiveColourString: validateColour,
	OptionsSessionDisplayPanesColourString:       validateColour,
	OptionsSessionDisplayPanesTimeString: func(v string) (bool, []string) {
		integer, err := strconv.Atoi(v)
		if err != nil {
			return false, nil
		}

		if integer < 0 {
			return false, nil
		}
		return true, nil
	},
	OptionsSessionDisplayTimeString: func(v string) (bool, []string) {
		integer, err := strconv.Atoi(v)
		if err != nil {
			return false, nil
		}

		if integer < 0 {
			return false, nil
		}
		return true, nil
	},
	OptionsSessionHistoryLimitString: func(v string) (bool, []string) {
		integer, err := strconv.Atoi(v)
		if err != nil {
			return false, nil
		}

		if integer < 0 {
			return false, nil
		}
		return true, nil
	},
	OptionsSessionKeyTableString: func(v string) (bool, []string) { return true, nil },
	OptionsSessionLockAfterTimeString: func(v string) (bool, []string) {
		integer, err := strconv.Atoi(v)
		if err != nil {
			return false, nil
		}

		if integer < 0 {
			return false, nil
		}
		return true, nil
	},
	OptionsSessionLockCommandString: func(v string) (bool, []string) { return true, nil },
	OptionsSessionMenuBorderLinesString: func(v string) (bool, []string) {
		choices := []string{"single", "rounded", "double", "heavy", "simple", "padded", "none"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsSessionMessageLineString: func(v string) (bool, []string) {
		choices := []string{"0", "1", "2", "3", "4"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsSessionMouseString: func(v string) (bool, []string) {
		choices := []string{"on", "off"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsSessionPrefixString:  func(v string) (bool, []string) { return true, nil },
	OptionsSessionPrefix2String: func(v string) (bool, []string) { return true, nil },
	OptionsSessionRenumberWindowsString: func(v string) (bool, []string) {
		choices := []string{"on", "off"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsSessionRepeatTimeString: func(v string) (bool, []string) {
		integer, err := strconv.Atoi(v)
		if err != nil {
			return false, nil
		}

		if integer < 0 {
			return false, nil
		}
		return true, nil
	},
	OptionsSessionSetTitlesString: func(v string) (bool, []string) {
		choices := []string{"on", "off"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsSessionSilenceActionString: func(v string) (bool, []string) {
		choices := []string{"any", "none", "current", "other"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsSessionStatusString: func(v string) (bool, []string) {
		choices := []string{"off", "on", "2", "3", "4", "5"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsSessionStatusIntervalString: func(v string) (bool, []string) {
		integer, err := strconv.Atoi(v)
		if err != nil {
			return false, nil
		}

		if integer < 0 {
			return false, nil
		}
		return true, nil
	},
	OptionsSessionStatusJustifyString: func(v string) (bool, []string) {
		choices := []string{"left", "centre", "right", "absolute-centre"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsSessionStatusKeysString: func(v string) (bool, []string) {
		choices := []string{"vi", "emacs"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsSessionStatusLeftString: func(v string) (bool, []string) { return true, nil },
	OptionsSessionStatusLeftLengthString: func(v string) (bool, []string) {
		integer, err := strconv.Atoi(v)
		if err != nil {
			return false, nil
		}

		if integer < 0 {
			return false, nil
		}
		return true, nil
	},
	OptionsSessionStatusPositionString: func(v string) (bool, []string) {
		choices := []string{"top", "bottom"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsSessionStatusRightString: func(v string) (bool, []string) { return true, nil },
	OptionsSessionStatusRightLengthString: func(v string) (bool, []string) {
		integer, err := strconv.Atoi(v)
		if err != nil {
			return false, nil
		}

		if integer < 0 {
			return false, nil
		}
		return true, nil
	},
	OptionsSessionVisualActivityString: func(v string) (bool, []string) {
		choices := []string{"on", "off"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsSessionVisualBellString: func(v string) (bool, []string) {
		choices := []string{"on", "off"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsSessionVisualSilenceString: func(v string) (bool, []string) {
		choices := []string{"on", "off", "both"}

		if found := slices.Contains(choices, v); !found {
			return false, choices
		}

		return true, nil
	},
	OptionsSessionWordSeparatorsString: func(v string) (bool, []string) { return true, nil },

	// STYLE options are supported, but not yet validated properly:
	OptionsSessionMessageCommandStyleString: func(v string) (bool, []string) { return true, nil },
	OptionsSessionMessageStyleString:        func(v string) (bool, []string) { return true, nil },
	OptionsSessionStatusLeftStyleString:     func(v string) (bool, []string) { return true, nil },
	OptionsSessionStatusRightStyleString:    func(v string) (bool, []string) { return true, nil },
	OptionsSessionStatusStyleString:         func(v string) (bool, []string) { return true, nil },
	OptionsSessionMenuStyleString:           func(v string) (bool, []string) { return true, nil },
	OptionsSessionMenuSelectedStyleString:   func(v string) (bool, []string) { return true, nil },
	OptionsSessionMenuBorderStyleString:     func(v string) (bool, []string) { return true, nil },
}
