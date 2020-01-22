package semver

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/rs/zerolog/log"
)

// Diff returns a map of the version differences
func (s *SemVer) Diff(otherVer *SemVer) map[string]interface{} {
	result := make(map[string]interface{})

	majorDiff := int64(otherVer.Major) - int64(s.Major)
	result["major"] = majorDiff
	switch {
	case majorDiff == 0: // 1.x.x == 1.x.x
		minorDiff := int64(otherVer.Minor) - int64(s.Minor)
		result["minor"] = minorDiff
		switch {
		case minorDiff == 0: // 1.1.x == 1.1.x
			patchDiff := int64(otherVer.Patch) - int64(s.Patch)
			result["patch"] = patchDiff
			switch {
			case patchDiff == 0: // 1.1.1 == 1.1.1
				if otherVer.PreRelease == s.PreRelease {
					result["prerelease"] = "0"
				} else {
					prerelease := []string{otherVer.PreRelease, s.PreRelease}
					sort.Strings(prerelease)
					result["prerelease"] = prerelease[len(prerelease)-1]
				}
				if otherVer.Build == s.Build {
					result["build"] = "0"
				} else {
					result["build"] = otherVer.Build
				}
			case patchDiff < 0: // 1.1.2 > 1.1.1
				result["prerelease"] = "0"
				result["build"] = "0"
			case patchDiff > 0: // 1.1.1 < 1.1.2
				if otherVer.PreRelease != "" {
					result["prerelease"] = otherVer.PreRelease
				} else {
					result["prerelease"] = "0"
				}
				if otherVer.Build != "" {
					result["build"] = otherVer.Build
				} else {
					result["build"] = "0"
				}
			}
		case minorDiff < 0: // 1.2.x > 1.1.x
			result["patch"] = 0
			result["prerelease"] = "0"
			result["build"] = "0"
		case minorDiff > 0: // 1.1.x < 1.2.x
			result["patch"] = otherVer.Patch
			result["prerelease"] = otherVer.PreRelease
			result["build"] = otherVer.Build
		}
	case majorDiff < 0: // 2.x.x > 1.x.x
		result["minor"] = 0
		result["patch"] = 0
		result["prerelease"] = "0"
		result["build"] = "0"
	case majorDiff > 0:
		result["minor"] = otherVer.Minor
		result["patch"] = otherVer.Patch
		result["prerelease"] = otherVer.PreRelease
		result["build"] = otherVer.Build
	}

	return result
}

// OutputDiff sends formatted output to stdout
func OutputDiff(results map[string]interface{}, outputType string, suppressHeaders bool) {
	switch outputType {
	case "table":
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)
		if suppressHeaders == false {
			fmt.Fprintln(w, "MAJOR\tMINOR\tPATCH\tPRERELEASE\tBUILD\t")
		}
		fmt.Fprintf(w, "%d\t%d\t%d\t%s\t%s\t\n", results["major"], results["minor"], results["patch"], results["prerelease"], results["build"])
		w.Flush()
	case "json":
		jsonString, err := json.Marshal(results)
		if err != nil {
			log.Fatal().Msgf("Error while outputting JSON: %v", err)
		}
		fmt.Printf("%s", jsonString)
	}
}
