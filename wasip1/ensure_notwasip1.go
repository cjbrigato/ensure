//go:build !wasip1
// +build !wasip1

package wasip1

// This variable declaration will fail to compile on non-wasip1 systems
// because the constant 'error_compiling_this_package_requires_goos_wasip1' is not defined for the
// !wasip1 build constraint within this package.
// The error message itself signals the requirement.
var _ = error_compiling_this_package_requires_goos_wasip1
