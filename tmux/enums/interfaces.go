package enums

type OptionTyper[OT OptionsPane | OptionsWindow | OptionsSession] interface {
	OptionsPane | OptionsWindow | OptionsSession
	FromString(string) OT
	IsKnown(string) bool
	GetValidator(string) (ValidatorFunc, bool)
}
