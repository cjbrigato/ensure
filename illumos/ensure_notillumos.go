//go:build !illumos
// +build !illumos

package illumos

// This variable declaration will fail to compile on non-illumos systems
// because the constant 'error_compiling_this_package_requires_goos_illumos' is not defined for the
// !illumos build constraint within this package.
// The error message itself signals the requirement.
var _ = error_compiling_this_package_requires_goos_illumos
