package imports

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Resolver is the interface that must be implemented by all resolvers.
type Resolver interface {
	// Resolve resolves a package by its name and version.
	Resolve(pkg *ExactPackage) error
}

var _ Resolver = (*UnimplementedResolver)(nil)

// UnimplementedResolver is a resolver that is not implemented.
type UnimplementedResolver struct{}

// Resolve resolves a package by its name and version.
func (r *UnimplementedResolver) Resolve(_ *ExactPackage) error {
	return nil
}

// Require is a struct that represents a package with its name and version.
type Require struct {
	// File is the file path of the package. This is the local path not the full url in the CDN.
	File string
	// Raw is the raw file path of the package.
	Raw string
	// As is the alias for the package.
	As string
}

// ExactPackage is a struct that represents a package with its name, version, and files.
type ExactPackage struct {
	// Name is the name of the package.
	Name string
	// Version is the version of the package.
	Version string
	// Files is a list of files associated with the package.
	Files []isFileType
}

// Hash is a struct that represents a hash of a file.
func (e *ExactPackage) Hash() string {
	hasher := sha256.New()
	fmt.Fprintf(hasher, "%s@%s", e.Name, e.Version)
	return hex.EncodeToString(hasher.Sum(nil))
}

type isFileType interface {
	isFileType()
}

// FileJS is a struct that represents a JavaScript file.
type FileJS struct {
	Path      string
	LocalPath string
	Integrity string
}

func (*FileJS) isFileType() {}

// FileCSS is a struct that represents a CSS file.
type FileCSS struct {
	Path      string
	LocalPath string
	Integrity string
}

func (*FileCSS) isFileType() {}

// FileUnkown is a struct that represents an unknown file type.
type FileUnkown struct {
	Path      string
	LocalPath string
	Integrity string
}

func (*FileUnkown) isFileType() {}
