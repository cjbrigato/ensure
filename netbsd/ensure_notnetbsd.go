//go:build !netbsd
// +build !netbsd

package netbsd

// This variable declaration will fail to compile on non-netbsd systems
// because the constant 'error_compiling_this_package_requires_goos_netbsd' is not defined for the
// !netbsd build constraint within this package.
// The error message itself signals the requirement.
var _ = error_compiling_this_package_requires_goos_netbsd
