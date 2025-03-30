//go:build !darwin
// +build !darwin

package darwin

// This variable declaration will fail to compile on non-darwin systems
// because the constant 'error_compiling_this_package_requires_goos_darwin' is not defined for the
// !darwin build constraint within this package.
// The error message itself signals the requirement.
var _ = error_compiling_this_package_requires_goos_darwin
