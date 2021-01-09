package gslice

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/slice
func Slice(s []string, beginIndex int, endIndex ...int) []string {
	endIdx := len(s)

	if len(endIndex) > 0 {
		endIdx = endIndex[0]
	}

	if beginIndex < 0 {
		beginIndex = len(s) + beginIndex
	}

	if beginIndex > len(s) {
		beginIndex = len(s)
	}

	if endIdx < 0 {
		endIdx = len(s) + endIdx
	}

	if endIdx > len(s) {
		endIdx = len(s)
	}

	copyElems := make([]string, endIdx - beginIndex)
	copy(copyElems, s[beginIndex: endIdx])
	return copyElems
}
