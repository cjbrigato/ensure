# Go Ensure GOOS Package - Guarantee target OS for Go builds via compile-time checks

[![GoDoc](https://godoc.org/github.com/cjbrigato/ensure?status.svg)](https://godoc.org/github.com/cjbrigato/ensure)

This Go package provides a simple, zero-runtime-cost mechanism to enforce **compile-time** checks for the target operating system (`GOOS`).

If your Go program or library has strict operating system requirements (e.g., it uses Linux-specific syscalls or Windows API calls), you can use this package to make the build fail *immediately* if someone tries to compile it for an unsupported OS, providing a clear error message.

## How it Works

The package leverages Go's build constraints (`//go:build`). Each OS-specific subpackage (e.g., `ensure/linux`) contains two files:

1.  One file that compiles *only* on the target OS (e.g., `//go:build linux`) and defines a specific constant.
2.  Another file that compiles *only* when *not* on the target OS (e.g., `//go:build !linux`) and attempts to use that constant.

If you compile for the wrong OS, the constant will be undefined, causing a compile-time error with a descriptive name indicating the required OS.

## Usage

To enforce a specific operating system, simply import the corresponding subpackage using an **anonymous import** (`_`) somewhere in your main package or any package that is part of your final build artifact. You don't need to call any functions.

**Ensure Linux:**
```go
package main

import (
    _ "github.com/cjbrigato/ensure/linux" // Fails compilation if GOOS != linux
    "fmt"
)

func main() {
    fmt.Println("This program is guaranteed to be compiled for Linux.")
    // ... Linux-specific code
}  
```
**Ensure Windows:**
```go
package main

import (
    _ "github.com/cjbrigato/ensure/windows" // Fails compilation if GOOS != windows
    "fmt"
)

func main() {
    fmt.Println("This program is guaranteed to be compiled for Windows.")
    // ... Windows-specific code
}
```

**Ensure macOS (Darwin):**
```go
package main

import (
    _ "github.com/cjbrigato/ensure/darwin" // Fails compilation if GOOS != darwin
    "fmt"
)

func main() {
    fmt.Println("This program is guaranteed to be compiled for macOS.")
    // ... macOS-specific code
}
```

## Subpackage Generation

The OS-specific subpackages (like `linux`, `windows`, `darwin`, etc.) in this repository were automatically generated to ensure consistency and reduce manual effort.

This was done using the following process:

1.  A list of currently known target operating systems (`GOOS`) was obtained using the standard Go toolchain command:
    ```bash
    go tool dist list | cut -d / -f 1 | sort -u
    ```
2.  A helper script (`./generate_ensure_subpackage.sh`) was created. This script takes a single `GOOS` value as input and generates the corresponding subdirectory (e.g., `./freebsd/`) containing:
    *   A `doc.go` file for package documentation.
    *   An `ensure_$GOOS.go` file containing the constant definition, guarded by `//go:build $GOOS`.
    *   An `ensure_not$GOOS.go` file attempting to use the constant, guarded by `//go:build !$GOOS`.
3.  A master script (`./generate_all.sh`) iterated through the list obtained in step 1 and called the helper script for each `GOOS` value.

This ensures that all subpackages follow the same pattern and that support for new target operating systems recognized by the Go toolchain can be added easily by re-running the generation scripts. You can inspect these scripts in the root of the repository.

## Available Subpackages
If ever a GOOS was missed or you need something custom, just use `./generate_ensure_subpackage.sh CUSTOM_GOOS`
   
## Installation
`go get github.com/cjbrigato/ensure`
    
Then, add the specific import to your code.

## License

This package is licensed under the [MIT License](LICENSE).
