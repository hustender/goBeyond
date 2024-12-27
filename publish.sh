#!/bin/bash

if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <commit_message> <version>"
    exit 1
fi

COMMIT_MESSAGE=$1
VERSION=$2

echo "The following files are available to be committed:"
git status --short

read -p "Do you want to stage all changes? (y/n): " STAGE_ALL

if [ "$STAGE_ALL" == "y" ]; then
    git add .
else
    echo "No files were staged. Please stage files manually if needed."
    echo "You can use 'git add <file>' to add specific files."
    exit 1
fi

git commit -m "${COMMIT_MESSAGE}@${VERSION}"

git push origin master

git tag "${VERSION}"

git push origin "${VERSION}"

GOPROXY=proxy.golang.org go list -m github.com/hustender/goBeyond@"${VERSION}"

if [ $? -eq 0 ]; then
    echo "Version ${VERSION} tagged and validated successfully."
else
    echo "An error occurred during the validation of version ${VERSION}."
fi
