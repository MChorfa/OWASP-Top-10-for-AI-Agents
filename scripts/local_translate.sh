#!/bin/bash

SRC_LANG="en"
TARGET_LANG=$1
BASE_DIR="v1/translations"

if [ -z "$TARGET_LANG" ]; then
    echo "Usage: $0 <target-language-code>"
    exit 1
fi

# Create target language directory
mkdir -p "$BASE_DIR/$TARGET_LANG"

if [ "$TARGET_LANG" == "ar" ]; then
    # Handle RTL language specifics if necessary
    echo "Setting up RTL support for Arabic translations"
    # ...additional configurations...
fi

# Copy metadata template
cp "$BASE_DIR/$SRC_LANG/metadata.yaml" "$BASE_DIR/$TARGET_LANG/"

# Copy markdown files for translation
for file in "$BASE_DIR/$SRC_LANG"/*.md; do
    basename=$(basename "$file")
    cp "$file" "$BASE_DIR/$TARGET_LANG/$basename"
done

echo "Translation structure created for $TARGET_LANG"
echo "Files are ready for translation in $BASE_DIR/$TARGET_LANG/"
