//go:build !dragonfly
// +build !dragonfly

package dragonfly

// This variable declaration will fail to compile on non-dragonfly systems
// because the constant 'error_compiling_this_package_requires_goos_dragonfly' is not defined for the
// !dragonfly build constraint within this package.
// The error message itself signals the requirement.
var _ = error_compiling_this_package_requires_goos_dragonfly
