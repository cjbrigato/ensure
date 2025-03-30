//go:build !solaris
// +build !solaris

package solaris

// This variable declaration will fail to compile on non-solaris systems
// because the constant 'error_compiling_this_package_requires_goos_solaris' is not defined for the
// !solaris build constraint within this package.
// The error message itself signals the requirement.
var _ = error_compiling_this_package_requires_goos_solaris
