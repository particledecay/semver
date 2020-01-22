package semver

import (
	"fmt"
	"regexp"
	"strconv"
)

// SemVer is a parsed Semantic Version string
type SemVer struct {
	Major      uint64
	Minor      uint64
	Patch      uint64
	PreRelease string
	Build      string
}

// Parse takes a version string and returns a parsed SemVer
func Parse(version string) (*SemVer, error) {
	sv := &SemVer{}

	rParse := regexp.MustCompile(`^[vV]?(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
	match := rParse.FindStringSubmatch(version)
	for i, name := range rParse.SubexpNames() {
		if i == 0 {
			continue
		}
		switch name {
		case "major":
			ver, err := strconv.ParseUint(match[i], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("%v is not a valid major version", match[i])
			}
			sv.Major = ver
		case "minor":
			ver, err := strconv.ParseUint(match[i], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("%v is not a valid minor version", match[i])
			}
			sv.Minor = ver
		case "patch":
			ver, err := strconv.ParseUint(match[i], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("%v is not a valid patch version", match[i])
			}
			sv.Patch = ver
		case "prerelease":
			if match[i] != "" {
				sv.PreRelease = match[i]
			}
		case "buildmetadata":
			if match[i] != "" {
				sv.Build = match[i]
			}
		}
	}

	return sv, nil
}
