//go:build !linux
// +build !linux

package linux

// This variable declaration will fail to compile on non-linux systems
// because the constant 'error_compiling_this_package_requires_goos_linux' is not defined for the
// !linux build constraint within this package.
// The error message itself signals the requirement.
var _ = error_compiling_this_package_requires_goos_linux
