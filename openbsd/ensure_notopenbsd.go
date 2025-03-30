//go:build !openbsd
// +build !openbsd

package openbsd

// This variable declaration will fail to compile on non-openbsd systems
// because the constant 'error_compiling_this_package_requires_goos_openbsd' is not defined for the
// !openbsd build constraint within this package.
// The error message itself signals the requirement.
var _ = error_compiling_this_package_requires_goos_openbsd
