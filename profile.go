package glaze

import "os"

// profile management and searching functions go here

// ProfileManager is a struct that contains a list of Profile objects
// Profile objects contain valid paths to glaze profiles
// ProfileManager searches for and selects the most appropriate profile
// Will return an error if no profiles are found

type Profile struct {
	Path string
}

func (p *Profile) Open() ([]byte, error) {
	return os.ReadFile(p.Path)
}

type ProfileManager struct {
	Profiles []Profile
}

func (pm *ProfileManager) Find() error {
	return nil
}
