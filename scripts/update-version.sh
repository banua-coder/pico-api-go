#!/bin/bash
# Version Update Script
# Reads .version-config.yml and updates version in specified files

set -e

VERSION="$1"
CONFIG_FILE="${2:-.version-config.yml}"

if [ -z "$VERSION" ]; then
    echo "Usage: $0 <version> [config-file]"
    echo "Example: $0 1.2.3"
    exit 1
fi

# Remove 'v' prefix if present
CLEAN_VERSION=$(echo "$VERSION" | sed 's/^v//')
MAJOR_VERSION=$(echo "$CLEAN_VERSION" | cut -d. -f1)

echo "🔄 Updating version to $CLEAN_VERSION using config: $CONFIG_FILE"

# Function to process file updates
process_file_update() {
    if [ -z "$CURRENT_FILE" ] || [ -z "$CURRENT_PATTERN" ] || [ -z "$CURRENT_REPLACEMENT" ]; then
        return
    fi

    # Check if file exists
    if [ ! -f "$CURRENT_FILE" ]; then
        echo "⚠️  File $CURRENT_FILE not found, skipping"
        return
    fi

    # Check 'when' condition
    if [ -n "$CURRENT_WHEN" ]; then
        if [[ "$CURRENT_WHEN" == "major_version_only" ]]; then
            # Only update for major version changes
            # This is a simple check - in real implementation you'd compare with previous version
            echo "ℹ️  Skipping $CURRENT_FILE (major version only)"
            return
        fi
    fi

    # Prepare replacement string
    REPLACEMENT="$CURRENT_REPLACEMENT"
    REPLACEMENT="${REPLACEMENT//\{version\}/$CLEAN_VERSION}"
    REPLACEMENT="${REPLACEMENT//\{major\}/$MAJOR_VERSION}"

    echo "🔄 Updating $CURRENT_FILE..."
    echo "   Pattern: $CURRENT_PATTERN"
    echo "   Replacement: $REPLACEMENT"

    # Use perl for more reliable regex replacement
    if command -v perl >/dev/null 2>&1; then
        perl -i -pe "s|$CURRENT_PATTERN|$REPLACEMENT|g" "$CURRENT_FILE"
    else
        # Fallback to sed (less reliable for complex patterns)
        sed -i "s|$CURRENT_PATTERN|$REPLACEMENT|g" "$CURRENT_FILE"
    fi

    echo "✅ Updated $CURRENT_FILE"
}

if [ ! -f "$CONFIG_FILE" ]; then
    echo "⚠️  Config file $CONFIG_FILE not found, using default file updates"

    # Fallback to hardcoded updates if config doesn't exist
    if [ -f "cmd/main.go" ]; then
        sed -i "s/@version\s\+[0-9]\+\.[0-9]\+\.[0-9]\+/@version\t\t$CLEAN_VERSION/" cmd/main.go
        echo "✅ Updated cmd/main.go"
    fi

    if [ -f "internal/handler/covid_handler.go" ]; then
        sed -i "s/\"version\":\s*\"[^\"]*\"/\"version\": \"$CLEAN_VERSION\"/" internal/handler/covid_handler.go
        echo "✅ Updated internal/handler/covid_handler.go"
    fi

    exit 0
fi

# Read version_files from YAML config
# This is a simple YAML parser for the version_files section
IN_VERSION_FILES=false
CURRENT_FILE=""
CURRENT_PATTERN=""
CURRENT_REPLACEMENT=""
CURRENT_WHEN=""

while IFS= read -r line; do
    # Check if we're entering the version_files section
    if [[ "$line" =~ ^version_files: ]]; then
        IN_VERSION_FILES=true
        continue
    fi

    # Check if we're leaving the version_files section
    if [[ "$IN_VERSION_FILES" == true && "$line" =~ ^[a-zA-Z] ]]; then
        IN_VERSION_FILES=false
        break
    fi

    if [[ "$IN_VERSION_FILES" == true ]]; then
        # Parse YAML entries
        if [[ "$line" =~ ^[[:space:]]*-[[:space:]]*path:[[:space:]]*\"(.*)\" ]]; then
            # Process previous file if we have one
            if [ -n "$CURRENT_FILE" ]; then
                process_file_update
            fi

            CURRENT_FILE="${BASH_REMATCH[1]}"
            CURRENT_PATTERN=""
            CURRENT_REPLACEMENT=""
            CURRENT_WHEN=""
        elif [[ "$line" =~ ^[[:space:]]*pattern:[[:space:]]*\"(.*)\" ]] || [[ "$line" =~ ^[[:space:]]*pattern:[[:space:]]*\'(.*)\' ]]; then
            CURRENT_PATTERN="${BASH_REMATCH[1]}"
        elif [[ "$line" =~ ^[[:space:]]*replacement:[[:space:]]*\"(.*)\" ]] || [[ "$line" =~ ^[[:space:]]*replacement:[[:space:]]*\'(.*)\' ]]; then
            CURRENT_REPLACEMENT="${BASH_REMATCH[1]}"
        elif [[ "$line" =~ ^[[:space:]]*when:[[:space:]]*\"(.*)\" ]] || [[ "$line" =~ ^[[:space:]]*when:[[:space:]]*\'(.*)\' ]]; then
            CURRENT_WHEN="${BASH_REMATCH[1]}"
        fi
    fi
done < "$CONFIG_FILE"

# Process the last file
if [ -n "$CURRENT_FILE" ]; then
    process_file_update
fi

echo "✅ Version update completed!"