//go:build !android
// +build !android

package android

// This variable declaration will fail to compile on non-android systems
// because the constant 'error_compiling_this_package_requires_goos_android' is not defined for the
// !android build constraint within this package.
// The error message itself signals the requirement.
var _ = error_compiling_this_package_requires_goos_android
