//go:build !js
// +build !js

package js

// This variable declaration will fail to compile on non-js systems
// because the constant 'error_compiling_this_package_requires_goos_js' is not defined for the
// !js build constraint within this package.
// The error message itself signals the requirement.
var _ = error_compiling_this_package_requires_goos_js
