#!/bin/sh

MIN_VERSION=26.2

SDK_PATH=`xcrun --sdk $SDK --show-sdk-path`
CLANG=`xcrun --sdk $SDK --find clang`

if [ "$SDK" = "iphoneos" ]; then
  export TARGET="-target $CARCH-apple-ios$MIN_VERSION"
elif [ "$SDK" = "iphonesimulator" ]; then
  export TARGET="-target $CARCH-apple-ios$MIN_VERSION-simulator"
fi

if [ "$GOARCH" == "amd64" ]; then
    CARCH="x86_64"
elif [ "$GOARCH" == "arm64" ]; then
    CARCH="arm64"
fi

exec $CLANG -arch $CARCH $TARGET -isysroot $SDK_PATH -mios-version-min=$MIN_VERSION "$@"
