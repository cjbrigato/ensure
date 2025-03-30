//go:build aix
// +build aix

package aix

// This constant provides the definition needed by ensure_notaix.go,
// but only when compiling for aix.
const error_compiling_this_package_requires_goos_aix = true
