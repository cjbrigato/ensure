//go:build !windows
// +build !windows

package windows

// This variable declaration will fail to compile on non-windows systems
// because the constant 'error_compiling_this_package_requires_goos_windows' is not defined for the
// !windows build constraint within this package.
// The error message itself signals the requirement.
var _ = error_compiling_this_package_requires_goos_windows
