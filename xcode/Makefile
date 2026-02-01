clean:
	rm -rf ./cmutton*

ios-hw:
	CGO_ENABLED=1 \
	GOOS=darwin \
	GOARCH=arm64 \
	SDK=iphoneos \
	CC=$(PWD)/clangwrap.sh \
	CGO_CFLAGS="-fembed-bitcode" \
	go build -buildmode=c-archive -tags ios -o ./cmutton-ios-hw.a ..

ios-sim:
	CGO_ENABLED=1 \
	GOOS=darwin \
	GOARCH=arm64 \
	SDK=iphonesimulator \
	CC=$(PWD)/clangwrap.sh \
	go build -buildmode=c-archive -tags ios -o ./cmutton-ios-sim.a ..

macos:
	CGO_ENABLED=1 \
	GOOS=darwin \
	GOARCH=arm64 \
	SDK=macosx \
	go build -buildmode=c-archive -o ./cmutton-macos-arm64.a ..

ios: ios-hw ios-sim
	xcodebuild -create-xcframework \
	-library ./cmutton-ios-hw.a \
	-headers ./cmutton-ios-hw.h \
	-library ./cmutton-ios-sim.a \
	-headers ./cmutton-ios-sim.h \
	-output ./cmutton.xcframework

apple: ios-hw ios-sim macos
	xcodebuild -create-xcframework \
	-library ./cmutton-ios-hw.a \
	-headers ./cmutton-ios-hw.h \
	-library ./cmutton-ios-sim.a \
	-headers ./cmutton-ios-sim.h \
	-library ./cmutton-macos-arm64.a \
	-headers ./cmutton-macos-arm64.h \
	-output ./cmutton.xcframework
