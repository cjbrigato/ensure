//go:build !freebsd
// +build !freebsd

package freebsd

// This variable declaration will fail to compile on non-freebsd systems
// because the constant 'error_compiling_this_package_requires_goos_freebsd' is not defined for the
// !freebsd build constraint within this package.
// The error message itself signals the requirement.
var _ = error_compiling_this_package_requires_goos_freebsd
