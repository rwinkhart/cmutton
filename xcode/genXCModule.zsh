#!/usr/bin/env zsh
# TODO Support iOS simulator

# setup working dirs w/modulemap
rm -rf ./build
mkdir -p ./build/{ios,macos}
echo "module cmutton {
    header "cmutton.h"
    link "cmutton"
    export *
}" > ./build/{ios,macos}/module.modulemap

# cd to Go source
cd ..

# build static archive for all platforms
## iOS
CGO_ENABLED=1 GOOS=ios SDK=iphoneos go build -buildmode=c-archive
mv ./{cmutton.a,cmutton.h} ./xcode/build/ios/
cp ./types.h ./xcode/build/ios/
## macOS
CGO_ENABLED=1 GOOS=darwin go build -buildmode=c-archive
mv ./{cmutton.a,cmutton.h} ./xcode/build/macos/
cp ./types.h ./xcode/build/macos/

cd ./xcode
