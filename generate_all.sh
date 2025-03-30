#!/bin/bash

# Master script to generate all GOOS subpackages for github.com/cjbrigato/ensure
# using generate_ensure_subpackage.sh.

# --- Configuration ---
GENERATOR_SCRIPT="./generate_ensure_subpackage.sh"

# --- Sanity Checks ---
if [ ! -f "$GENERATOR_SCRIPT" ]; then
    echo "Error: Generator script '$GENERATOR_SCRIPT' not found."
    echo "Please ensure it exists in the current directory."
    exit 1
fi

if [ ! -x "$GENERATOR_SCRIPT" ]; then
    echo "Error: Generator script '$GENERATOR_SCRIPT' is not executable."
    echo "Please run: chmod +x $GENERATOR_SCRIPT"
    exit 1
fi

# --- Get Unique GOOS List ---
echo "Fetching unique GOOS values from 'go tool dist list'..."
# Assign to an array (works in bash 4+)
readarray -t ALL_GOOS < <(go tool dist list | cut -d / -f 1 | sort -u)
# Alternatively, for older bash/sh:
# ALL_GOOS=$(go tool dist list | cut -d / -f 1 | sort -u)

if [ ${#ALL_GOOS[@]} -eq 0 ]; then
     echo "Error: Failed to get GOOS list from 'go tool dist list'."
     exit 1
fi
echo "Found ${#ALL_GOOS[@]} unique GOOS values."


# --- Generate Subpackages ---
echo "----------------------------------------"
echo "Generating subpackages..."
echo "----------------------------------------"
generation_failed=0
success_count=0
for goos in "${ALL_GOOS[@]}"; do
    echo "--- Generating for GOOS: $goos ---"
    # Execute the generator script, capturing its output only on error maybe?
    # Or just let it print its usual messages.
    if "$GENERATOR_SCRIPT" "$goos"; then
        echo "Successfully generated for '$goos'."
        ((success_count++))
    else
        # The generator script should already print errors
        echo "ERROR: Failed to generate package for '$goos'. Check logs above."
        generation_failed=1
        # Decide whether to stop or continue
        # exit 1 # Option: Stop on first error
    fi
    echo # Add a newline for readability
done

echo "----------------------------------------"
echo "Generation complete."
echo "Successfully generated: $success_count"
if [ $generation_failed -ne 0 ]; then
    failed_count=$(( ${#ALL_GOOS[@]} - success_count ))
    echo "Failed to generate: $failed_count. Please review the logs above."
    exit 1 # Exit with error if any generation failed
fi
echo "----------------------------------------"
exit 0