// Package ensure provides subpackages to enforce compile-time checks
// for the target operating system (GOOS).
//
// This root package itself contains no functionality. Its purpose is to
// act as a container and documentation hub for the OS-specific subpackages
// like `linux`, `windows`, `darwin`, etc.
//
// To use the compile-time check, import the desired subpackage anonymously.
// For example, to ensure your program only compiles for Linux:
//
//	import _ "github.com/cjbrigato/ensure/linux"
//
// Attempting to compile code with such an import for a different GOOS
// will result in a compile-time error indicating the required OS.
package ensure
