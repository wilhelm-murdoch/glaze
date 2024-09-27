package enums

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

var OptionsSessionValidators = map[string]ValidatorFunc{
	OptionsSessionDisplayPanesActiveColourString: validatorColour,
	OptionsSessionDisplayPanesColourString:       validatorColour,
	OptionsSessionDisplayPanesTimeString:         validatorIsNumber,
	OptionsSessionDisplayTimeString:              validatorIsNumber,
	OptionsSessionHistoryLimitString:             validatorIsNumber,
	OptionsSessionKeyTableString:                 validatorDefault,
	OptionsSessionLockAfterTimeString:            validatorIsNumber,
	OptionsSessionLockCommandString:              validatorDefault,
	OptionsSessionMouseString:                    validatorToggle,
	OptionsSessionPrefixString:                   validatorDefault,
	OptionsSessionPrefix2String:                  validatorDefault,
	OptionsSessionRenumberWindowsString:          validatorToggle,
	OptionsSessionRepeatTimeString:               validatorIsNumber,
	OptionsSessionSetTitlesString:                validatorToggle,
	OptionsSessionStatusIntervalString:           validatorIsNumber,
	OptionsSessionStatusLeftString:               validatorDefault,
	OptionsSessionStatusLeftLengthString:         validatorIsNumber,
	OptionsSessionStatusRightString:              validatorDefault,
	OptionsSessionStatusRightLengthString:        validatorIsNumber,
	OptionsSessionVisualActivityString:           validatorToggle,
	OptionsSessionVisualBellString:               validatorToggle,
	OptionsSessionWordSeparatorsString:           validatorDefault,

	OptionsSessionMenuBorderLinesString: func(v string) (bool, []string) {
		return validatorContains(v, "single", "rounded", "double", "heavy", "simple", "padded", "none")
	},
	OptionsSessionActivityActionString: func(v string) (bool, []string) {
		return validatorContains(v, "any", "none", "current", "other")
	},
	OptionsSessionDetachOnDestroyString: func(v string) (bool, []string) {
		return validatorContains(v, "off", "on", "no-detatched", "previous", "next")
	},
	OptionsSessionMessageLineString: func(v string) (bool, []string) {
		return validatorContains(v, "0", "1", "2", "3", "4")
	},
	OptionsSessionSilenceActionString: func(v string) (bool, []string) {
		return validatorContains(v, "any", "none", "current", "other")
	},
	OptionsSessionStatusString: func(v string) (bool, []string) {
		return validatorContains(v, "off", "on", "2", "3", "4", "5")
	},
	OptionsSessionStatusJustifyString: func(v string) (bool, []string) {
		return validatorContains(v, "left", "centre", "right", "absolute-centre")
	},
	OptionsSessionStatusKeysString: func(v string) (bool, []string) {
		return validatorContains(v, "vi", "emacs")
	},
	OptionsSessionStatusPositionString: func(v string) (bool, []string) {
		return validatorContains(v, "top", "bottom")
	},
	OptionsSessionVisualSilenceString: func(v string) (bool, []string) {
		return validatorContains(v, "on", "off", "both")
	},

	// STYLE options are supported, but not yet validated properly:
	OptionsSessionMessageCommandStyleString: validatorDefault,
	OptionsSessionMessageStyleString:        validatorDefault,
	OptionsSessionStatusLeftStyleString:     validatorDefault,
	OptionsSessionStatusRightStyleString:    validatorDefault,
	OptionsSessionStatusStyleString:         validatorDefault,
	OptionsSessionMenuStyleString:           validatorDefault,
	OptionsSessionMenuSelectedStyleString:   validatorDefault,
	OptionsSessionMenuBorderStyleString:     validatorDefault,
}
