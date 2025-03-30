//go:build dragonfly
// +build dragonfly

package dragonfly

// This constant provides the definition needed by ensure_notdragonfly.go,
// but only when compiling for dragonfly.
const error_compiling_this_package_requires_goos_dragonfly = true
