//go:build !plan9
// +build !plan9

package plan9

// This variable declaration will fail to compile on non-plan9 systems
// because the constant 'error_compiling_this_package_requires_goos_plan9' is not defined for the
// !plan9 build constraint within this package.
// The error message itself signals the requirement.
var _ = error_compiling_this_package_requires_goos_plan9
