package axum

import (
	"regexp"
	"strings"
)

const (
	varSegmentPattern   = `^\{[a-zA-Z]+[a-zA-Z0-9_]*\}$`
	constSegmentPattern = `^[a-zA-Z]+[a-zA-Z0-9_]*$`
)

const (
	segmentTypeVar   = 1
	segmentTypeConst = 2
	segmentTypeRegex = 3
)

type pathSegmentMatcher interface {
	Match(string) bool
	Type() int
}

type varMatcher struct{}

func (m *varMatcher) Match(_ string) bool {
	return true
}

func (m *varMatcher) Type() int {
	return segmentTypeVar
}

func newVarMatch() *varMatcher {
	return &varMatcher{}
}

type constMatcher struct {
	value string
}

func (m *constMatcher) Match(s string) bool {
	return m.value == s
}

func (m *constMatcher) Type() int {
	return segmentTypeConst
}

func newConstMatcher(segment string) *constMatcher {
	return &constMatcher{
		value: segment,
	}
}

type regexMatcher struct {
	regex *regexp.Regexp
}

func (m *regexMatcher) Match(s string) bool {
	return m.regex.MatchString(s)

}

func (m *regexMatcher) Type() int {
	return segmentTypeRegex
}

func newRegexMatcher(segment string) *regexMatcher {
	return &regexMatcher{
		regex: regexp.MustCompile(segment),
	}
}

type pathSegment struct {
	value   string
	matcher pathSegmentMatcher
	arg     string
}

func (seg *pathSegment) Match(segment string) bool {
	result := seg.matcher.Match(segment)

	if seg.matcher.Type() == segmentTypeVar {
		key := seg.value[1 : len(seg.value)-1]
		seg.arg = key + "=" + segment
	}

	return result
}

func newPathSegment(v string) *pathSegment {
	path := &pathSegment{
		value: v,
	}

	varRegex := regexp.MustCompile(varSegmentPattern)
	constRegex := regexp.MustCompile(constSegmentPattern)

	if varRegex.MatchString(v) {
		path.matcher = newVarMatch()
	} else if constRegex.MatchString(v) {
		path.matcher = newConstMatcher(v)
	} else {
		path.matcher = newRegexMatcher(v)
	}

	return path
}

type pathSegments []*pathSegment

func (ps pathSegments) Match(path string) bool {
	path = strings.Trim(path, "/")
	segs := strings.Split(path, "/")

	if len(segs) != len(ps) {
		return false
	}

	for i, p := range ps {
		if !p.Match(segs[i]) {
			return false
		}
	}

	return true
}
