//go:build !aix
// +build !aix

package aix

// This variable declaration will fail to compile on non-aix systems
// because the constant 'error_compiling_this_package_requires_goos_aix' is not defined for the
// !aix build constraint within this package.
// The error message itself signals the requirement.
var _ = error_compiling_this_package_requires_goos_aix
