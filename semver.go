package semver

import (
	"strconv"
	"strings"
)

type SemVersion struct {
	Major int64
	Minor int64
	Patch int64
}

// IsNewer returns true if compare is > version
func IsNewer(version, compare string) (bool, error) {
	first, err := GetAsSemversion(version)
	if err != nil {
		return false, nil
	}
	second, err := GetAsSemversion(compare)
	if err != nil {
		return false, nil
	}

	if second.Major > first.Major {
		return true, nil
	} else if second.Major == first.Major && second.Minor > first.Minor {
		return true, nil
	} else if second.Major == first.Major && second.Minor == first.Minor && second.Patch > first.Patch {
		return true, nil
	}

	return false, nil
}

// GetAsSemversion takes a string and retuns a Semversion type
func GetAsSemversion(version string) (SemVersion, error) {
	semversion := SemVersion{}
	var err error
	parts := strings.Split(version, ".")
	partsLen := len(parts)

	majorInt, err := getAsInt(parts[0])
	if err != nil {
		return semversion, err
	}
	semversion.Major = majorInt

	if partsLen >= 2 {
		minorInt, err := getAsInt(parts[1])
		if err != nil {
			return semversion, err
		}
		semversion.Minor = minorInt
	}

	if partsLen >= 3 {
		patchInt, err := getAsInt(parts[2])
		if err != nil {
			return semversion, err
		}
		semversion.Patch = patchInt
	}

	return semversion, err
}

func getAsInt(text string) (int64, error) {
	if text == "" {
		return 0, nil
	}

	return strconv.ParseInt(text, 10, 16)
}
