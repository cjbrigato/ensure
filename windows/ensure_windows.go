//go:build windows
// +build windows

package windows

// This constant provides the definition needed by ensure_notwindows.go,
// but only when compiling for windows.
const error_compiling_this_package_requires_goos_windows = true
