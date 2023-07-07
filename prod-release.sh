#!/bin/bash

# $1=<new release tag>

# parse to get tags needed
IFS='.'; read -r v major minor  num<<< "$1"

# if new tag is v1.100.0, it is a fresh sprint release, look for previous major version and keep minor version
# if new tag is v1.100.1, it is a hotfix, keep major version and look for previous minor version
if [[ "$minor" -eq 0 ]]; then
    RELEASE_TITLE="V$major.$minor.$num"
    PREV_MAJOR=$((major-1))
    PREV_MINOR=$minor
else
    echo "creating hotfix"
    RELEASE_TITLE="V$major.$minor.$num"
    PREV_MAJOR=$major
    PREV_MINOR=$minor
fi

echo "creating draft release: $1; with title: $RELEASE_TITLE"

unset IFS
# create the draft release
echo "$GITHUB_TOKEN" | gh auth login -p ssh -h "github.com" --with-token

echo "did you sync up with upstream? (enter to continue...)"
read -r

git log --oneline --no-decorate "$v.$PREV_MAJOR.$PREV_MINOR..." | gh release create "$1"  --generate-notes -R 'michaelkad/power-beta-go-sdk' -t "$RELEASE_TITLE"  