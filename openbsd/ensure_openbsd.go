//go:build openbsd
// +build openbsd

package openbsd

// This constant provides the definition needed by ensure_notopenbsd.go,
// but only when compiling for openbsd.
const error_compiling_this_package_requires_goos_openbsd = true
