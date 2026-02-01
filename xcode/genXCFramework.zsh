#!/usr/bin/env zsh
# TODO Support iOS simulator

# setup build dirs
rm -rf ./build
mkdir -p ./build/{ios,macos}

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

# cd to xcode build dir and build XCFramework
cd ./xcode/build
xcodebuild -create-framework \
    -library ./ios/cmutton.a \
    -headers ./ios/cmutton.h \
    -library ./macos/cmutton.a \
    -headers ./macos/cmutton.h \
    -output cmutton.xcframework
