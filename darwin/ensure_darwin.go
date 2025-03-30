//go:build darwin
// +build darwin

package darwin

// This constant provides the definition needed by ensure_notdarwin.go,
// but only when compiling for darwin.
const error_compiling_this_package_requires_goos_darwin = true
