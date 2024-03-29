package simplemath

import "fmt"

type SemanticVersion struct {
	major, minor, patch int
}

func NewSemanticVersion(major, minor, patch int) SemanticVersion {

	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

func GetSemanticVersion(sv SemanticVersion) string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}

func (sv SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}

func (sv SemanticVersion) IncrementMajor() {
	sv.major += 1
}

func (sv SemanticVersion) IncMajor() SemanticVersion {
	sv.major += 1
	return sv
}

func (sv *SemanticVersion) IncrementMajorVer() {
	sv.major += 1
}

func (sv SemanticVersion) IncrementMinor() {
	sv.minor += 1
}
func (sv SemanticVersion) IncrementPatch() {
	sv.patch += 1
}
