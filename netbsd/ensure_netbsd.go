//go:build netbsd
// +build netbsd

package netbsd

// This constant provides the definition needed by ensure_notnetbsd.go,
// but only when compiling for netbsd.
const error_compiling_this_package_requires_goos_netbsd = true
