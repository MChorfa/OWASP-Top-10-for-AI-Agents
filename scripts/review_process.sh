#!/bin/bash

# Translation review process script
TARGET_LANG=$1
FILE_PATH=$2

if [ -z "$TARGET_LANG" ] || [ -z "$FILE_PATH" ]; then
    echo "Usage: $0 <language-code> <file-path>"
    exit 1
fi

# Include 'ar' in supported languages
SUPPORTED_LANGS=("en" "fr" "es" "ar")

# Create review metadata
echo "Creating review entry for $FILE_PATH"
REVIEW_DATE=$(date +%Y-%m-%d)
REVIEW_ID="${TARGET_LANG}_${REVIEW_DATE}"

# Add review metadata to file
cat << EOF >> "$FILE_PATH.review"
review_id: $REVIEW_ID
language: $TARGET_LANG
status: pending
created: $REVIEW_DATE
reviewers_required: 2
technical_review: pending
language_review: pending
EOF

echo "Review process initiated for $FILE_PATH"
