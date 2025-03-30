//go:build !ios
// +build !ios

package ios

// This variable declaration will fail to compile on non-ios systems
// because the constant 'error_compiling_this_package_requires_goos_ios' is not defined for the
// !ios build constraint within this package.
// The error message itself signals the requirement.
var _ = error_compiling_this_package_requires_goos_ios
