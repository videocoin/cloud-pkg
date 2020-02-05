package project

import (
	"regexp"
	"strings"

	cstr "github.com/videocoin/cloud-pkg/strings"

	"github.com/videocoin/cloud-pkg/api/resources"
)

const (
	// CollectionID is the identifier of the resource that contains a list of
	// projects.
	CollectionID = "projects"
)

var (
	// ErrInvalidName indicates that the project name is invalid.
	ErrInvalidName = resources.PatternError(NamePattern.String())
	// ErrInvalidID indicates that the project identifier is invalid.
	ErrInvalidID = resources.PatternError(IDPattern.String())
)

var (
	// IDPattern is the project identifier pattern.
	IDPattern = regexp.MustCompile(`^[a-z0-9][-a-z0-9]{3,48}[a-z0-9]$`)
	// NamePattern is the project name pattern.
	NamePattern = regexp.MustCompile(`^projects/[a-z0-9][-a-z0-9]{3,48}[a-z0-9]$`)
)

// Name returns the project name given a project identifier.
func Name(ID string) (string, error) {
	if ok := IsValidID(ID); !ok {
		return "", ErrInvalidID
	}
	return cstr.JoinWithSeparator(resources.NameSeparator, CollectionID, ID), nil
}

// IDFromName derives the project identifier from its name.
func IDFromName(name string) (string, error) {
	if ok := IsValidName(name); !ok {
		return "", ErrInvalidName
	}
	return strings.SplitN(name, resources.NameSeparator, 2)[1], nil
}

// IsValidID reports whether a project identifier is valid.
func IsValidID(ID string) bool {
	return IDPattern.MatchString(ID)
}

// IsValidName reports whether a project name is valid.
func IsValidName(name string) bool {
	return NamePattern.MatchString(name)
}
