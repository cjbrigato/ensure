#!/bin/bash

# Script to generate an OS-specific subpackage for github.com/cjbrigato/ensure
# Creates a directory named <goos_name> with the required Go files.

# --- Configuration ---
# You might want to change this if your constant naming convention differs.
# The GOOS name will be appended to this prefix.
ERROR_VAR_PREFIX="error_compiling_this_package_requires_goos_"
# Change this if your base import path is different
BASE_IMPORT_PATH="github.com/cjbrigato/ensure"

# --- Input Validation ---
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <goos_name>"
    echo "  <goos_name>: A valid GOOS value (e.g., linux, windows, darwin, js, plan9)."
    echo "Example: $0 freebsd"
    # You could enhance this by fetching valid GOOS values:
    # echo "Valid GOOS values:"
    # go tool dist list | cut -d / -f 1 | sort -u
    exit 1
fi

# Ensure lowercase and store the GOOS name
goos_name=$(echo "$1" | tr '[:upper:]' '[:lower:]')
error_var_name="${ERROR_VAR_PREFIX}${goos_name}"

# Basic check for typical GOOS format (letters/numbers) - enhance if needed
if [[ ! "$goos_name" =~ ^[a-z0-9]+$ ]]; then
   echo "Error: Invalid GOOS name format '$goos_name'. Should be simple lowercase letters/numbers."
   exit 1
fi

echo "--> Generating ensure subpackage for GOOS: '$goos_name'"

# --- Directory Creation ---
if [ -d "$goos_name" ]; then
    echo "Warning: Directory '$goos_name' already exists. Files may be overwritten."
else
    mkdir "$goos_name"
    if [ $? -ne 0 ]; then
        echo "Error: Failed to create directory '$goos_name'."
        exit 1
    fi
    echo "Created directory: $goos_name/"
fi

# --- File Generation ---

# 1. doc.go
doc_file="$goos_name/doc.go"
echo "Generating $doc_file ..."
cat << EOF > "$doc_file"
// Package $goos_name forces a compile-time error if the target GOOS is not $goos_name.
//
// Import this package anonymously to ensure your code is only built for $goos_name:
//
//	import _ "${BASE_IMPORT_PATH}/$goos_name"
package $goos_name
EOF

# 2. ensure_<goos_name>.go (compiles ON the target OS)
#    Filename uses the specific GOOS name.
ensure_file="$goos_name/ensure_${goos_name}.go"
echo "Generating $ensure_file ..."
cat << EOF > "$ensure_file"
//go:build $goos_name
// +build $goos_name

package $goos_name

// This constant provides the definition needed by ensure_not${goos_name}.go,
// but only when compiling for $goos_name.
const $error_var_name = true
EOF

# 3. ensure_not<goos_name>.go (compiles OFF the target OS)
#    Filename explicitly indicates it's for the 'not' case.
ensure_not_file="$goos_name/ensure_not${goos_name}.go"
echo "Generating $ensure_not_file ..."
cat << EOF > "$ensure_not_file"
//go:build !$goos_name
// +build !$goos_name

package $goos_name

// This variable declaration will fail to compile on non-$goos_name systems
// because the constant '$error_var_name' is not defined for the
// !$goos_name build constraint within this package.
// The error message itself signals the requirement.
var _ = $error_var_name
EOF

echo "--> Successfully generated subpackage files in '$goos_name/'"
exit 0