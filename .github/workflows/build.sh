#!/usr/bin/env bash
#
# This script builds the application from source for multiple platforms.

# Get the parent directory of where this script is.
SOURCE="/home/runner/work/power-beta-go-sdk/power-beta-go-sdk/go/src/github.com/terraform-provider-ibm/.github"
echo "Source: $SOURCE"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
echo "DIR: $DIR"
# Change into that directory
cd "/home/runner/work/power-beta-go-sdk/power-beta-go-sdk/go/src/github.com/terraform-provider-ibm"

# Package which has the version information, required to set the Version, GitCommit info
VERSION_PACKAGE="github.com/terraform-providers/terraform-provider-ibm/version"

# Get the git commit
GIT_COMMIT=$(git rev-parse HEAD)
GIT_DIRTY=$(test -n "`git status --porcelain`" && echo "+CHANGES" || true)

# Determine the arch/os combos we're building for
XC_ARCH=${XC_ARCH:-"amd64" "arm64" "arm"}
XC_OS=${XC_OS:-linux darwin windows}
# XC_ARCH=${XC_ARCH:-"amd64" "arm64" }
# XC_OS=${XC_OS:-linux }
XC_EXCLUDE_OSARCH="!darwin/386 !windows/arm64 !windows/arm !darwin/arm"

# Delete the old dir
echo "==> Removing old directory..."
rm -f bin/*
rm -rf pkg/*
mkdir -p bin/
mkdir -p pkg/

# If its dev mode, only build for ourself
if [ "${TF_DEV}x" != "x" ]; then
    XC_OS=$(go env GOOS)
    XC_ARCH=$(go env GOARCH)
fi

# if ! which gox > /dev/null; then
#     echo "==> Installing gox..."
#     go get  github.com/mitchellh/gox@latest
#     go get  golang.org/x/sys/unix@latest
# fi
#  echo "===> check gox download...`which gox`" 
# instruct gox to build statically linked binaries
export CGO_ENABLED=0

# Allow LD_FLAGS to be appended during development compilations
LD_FLAGS="-X ${VERSION_PACKAGE}.GitCommit=${GIT_COMMIT}${GIT_DIRTY} $LD_FLAGS"

# In release mode we don't want debug information in the binary
if [[ -n "${TF_RELEASE}" ]]; then
    LD_FLAGS="-X ${VERSION_PACKAGE}.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X ${VERSION_PACKAGE}.VersionPrerelease= -s -w"
fi

# Build!
echo "==> Building..."
gox \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -osarch="${XC_EXCLUDE_OSARCH}" \
    -ldflags "${LD_FLAGS}" \
    -output "pkg/{{.OS}}_{{.Arch}}/terraform-provider-ibm" \
    .

# # Move all the compiled things to the $GOPATH/bin
# GOPATH=${GOPATH:-$(go env GOPATH)}
# case $(uname) in
#     CYGWIN*)
#         GOPATH="$(cygpath $GOPATH)"
#         ;;
# esac
# OLDIFS=$IFS
# IFS=: MAIN_GOPATH=($GOPATH)
# IFS=$OLDIFS

# # Create GOPATH/bin if it's doesn't exists
# if [ ! -d $MAIN_GOPATH/bin ]; then
#     echo "==> Creating GOPATH/bin directory..."
#     mkdir -p $MAIN_GOPATH/bin
# fi

# # Copy our OS/Arch to the bin/ directory
# DEV_PLATFORM="./pkg/$(go env GOOS)_$(go env GOARCH)"
# if [[ -d "${DEV_PLATFORM}" ]]; then
#     for F in $(find ${DEV_PLATFORM} -mindepth 1 -maxdepth 1 -type f); do
#         cp ${F} bin/
#         cp ${F} ${MAIN_GOPATH}/bin/
#     done
# fi

if [ "${TF_DEV}x" = "x" ]; then
    # Zip and copy to the dist dir
    echo "==> Packaging..."
    for PLATFORM in $(find ./pkg -mindepth 1 -maxdepth 1 -type d); do
        OSARCH=$(basename ${PLATFORM})
        echo "--> ${OSARCH}"

        pushd $PLATFORM >/dev/null 2>&1
        zip ../${OSARCH}.zip ./*
        echo "==> Deleiting dir: ${PLATFORM} ... or ${OSARCH}"
        rm -rf ${OSARCH}
        popd >/dev/null 2>&1
    done
fi

# Done!
echo
echo "==> Results:"
echo "==> bin/:..."
ls -hl bin/
echo "==> pkg/:..."
ls -hl pkg/
# zip_files=""
# while IFS= read -r -d $'\0' file; do
#     zip_files+="$(printf "%q" "$file")\n"
# done < <(find ./pkg -type f -name "*.zip" -print0)

# echo -e "Zip files:\n$zip_files"
