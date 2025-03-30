//go:build android
// +build android

package android

// This constant provides the definition needed by ensure_notandroid.go,
// but only when compiling for android.
const error_compiling_this_package_requires_goos_android = true
