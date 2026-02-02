BUILD_DIR = ./copy-contents-to-xcode

clean:
	rm -rf $(BUILD_DIR)

ios-hw: clean
	mkdir -p $(BUILD_DIR)
	CGO_ENABLED=1 \
	GOOS=darwin \
	GOARCH=arm64 \
	GOARM64=v8.3,crypto \
	SDK=iphoneos \
	CC=$(PWD)/clangwrap.sh \
	CGO_CFLAGS="-fembed-bitcode" \
	go build -buildmode=c-archive -tags ios -o $(BUILD_DIR)/cmutton-ios-hw.a ..
	echo '#import "types.h"\n#import "cmutton.h"' > $(BUILD_DIR)/cmutton-Bridging-Header-ios-hw.h
	cp ../types.h $(BUILD_DIR)/types.h

ios-sim: clean
	mkdir -p $(BUILD_DIR)
	CGO_ENABLED=1 \
	GOOS=darwin \
	GOARCH=arm64 \
	GOARM64=v8.4,crypto \
	SDK=iphonesimulator \
	CC=$(PWD)/clangwrap.sh \
	go build -buildmode=c-archive -tags ios -o $(BUILD_DIR)/cmutton-ios-sim.a ..
	echo '#import "types.h"\n#import "cmutton.h"' > $(BUILD_DIR)/cmutton-Bridging-Header-ios-sim.h
	cp ../types.h $(BUILD_DIR)/types.h

macos: clean
	mkdir -p $(BUILD_DIR)
	CGO_ENABLED=1 \
	GOOS=darwin \
	GOARCH=arm64 \
	GOARM64=v8.4,crypto \
	SDK=macosx \
	go build -buildmode=c-archive -o $(BUILD_DIR)/cmutton-macos.a ..
	echo '#import "types.h"\n#import "cmutton.h"' > $(BUILD_DIR)/cmutton-Bridging-Header-macos.h
	cp ../types.h $(BUILD_DIR)/types.h

apple: clean ios-hw ios-sim macos
