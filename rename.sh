#!/bin/bash

# Usage: ./rename.sh github.com/oldname/project github.com/newname/project

if [ "$#" -ne 2 ]; then
  echo "Usage: $0 <old-module-path> <new-module-path>"
  exit 1
fi

OLD=$1
NEW=$2

echo "🔄 Replacing all occurrences of:"
echo "  $OLD → $NEW"
echo

# Replace in all .go, .mod, .sum, and templ files
grep -rl "$OLD" . --include \*.go --include go.mod --include go.sum --include \*.templ | while read -r file; do
  echo "↪ Updating $file"
  sed -i '' "s|$OLD|$NEW|g" "$file"
done

echo "✅ Done."
