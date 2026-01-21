#!/usr/bin/env bash

# Get all packages in libmutton
mapfile -s 1 -t packages < <(go list github.com/rwinkhart/libmutton/... 2>/dev/null | sed 's|github.com/rwinkhart/libmutton/||')

# If no packages found, exit
if [ ${#packages[@]} -eq 0 ]; then
    echo "No packages found in libmutton"
    exit 1
fi

# Iterate through each package
for package in "${packages[@]}"; do
    echo "- [ ] $package"
    mapfile -t lines < <(go doc -short "github.com/rwinkhart/libmutton/$package" 2>/dev/null | grep "func ")
    for line in "${lines[@]}"; do
        line="${line#"${line%%[![:space:]]*}"}"
        echo "    - [ ] ${line:5}"
    done
done
