package semver

import (
	"strconv"
	"strings"
)

type preReleaseSort []string

func (p preReleaseSort) Len() int {
	return len(p)
}

func (p preReleaseSort) Swap(x, y int) {
	p[x], p[y] = p[y], p[x]
}

func (p preReleaseSort) Less(x, y int) bool {
	xSplit := strings.Split(p[x], ".")
	ySplit := strings.Split(p[y], ".")

	if len(xSplit) > 0 && len(ySplit) > 0 {
		for i := 0; i < len(xSplit); i++ {
			if i >= len(ySplit) { // alpha.beta > alpha
				return false
			}
			if xSplit[i] == ySplit[i] { // alpha == alpha
				continue
			}
			xVal, xErr := strconv.ParseInt(xSplit[i], 10, 32)
			yVal, yErr := strconv.ParseInt(ySplit[i], 10, 32)
			if xErr != nil && yErr != nil { // both strings
				return xSplit[i] < ySplit[i] // alpha < beta
			} else if xErr == nil && yErr == nil { // both ints
				if xVal == yVal { // beta.2 == beta.2
					continue
				}
				return xVal < yVal // beta.2 < beta.11
			} else if xErr == nil && yErr != nil { // xVal is int, lower precedence than yVal
				return true // alpha.1 < alpha.beta
			} else if xErr != nil && yErr == nil { // xVal is string, higher precedence than yVal
				return false // alpha.beta > alpha.1
			}
		}
		return true // alpha < alpha.beta
	} else if len(xSplit) > 0 { // y has no pre-release, it is higher precedence than x
		return true // 1.0.0-rc1 < 1.0.0
	}
	return false // 1.0.0 > 1.0.0-rc1
}
