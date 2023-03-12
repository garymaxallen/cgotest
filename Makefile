ios-arm64:
	$(info $$IOS_OUT is [$(IOS_OUT)])
	CGO_ENABLED=1 \
	GOOS=darwin \
	GOARCH=arm64 \
	SDK=iphoneos \
	CC=/Users/pcl/Documents/tmp/gotest/gotest/clangwrap.sh \
	CGO_CFLAGS="-fembed-bitcode" \
	go build -buildmode=c-archive -tags ios -o $(IOS_OUT)/arm64.a ./cmd/libfoo

ios-x86_64:
	CGO_ENABLED=1 \
	GOOS=darwin \
	GOARCH=amd64 \
	SDK=iphonesimulator \
	CC=/Users/pcl/Documents/tmp/gotest/gotest/clangwrap.sh \
	go build -buildmode=c-archive -tags ios -o $(IOS_OUT)/x86_64.a ./cmd/libfoo

ios: ios-arm64 ios-x86_64
	lipo $(IOS_OUT)/x86_64.a $(IOS_OUT)/arm64.a -create -output $(IOS_OUT)/foo.a
	cp $(IOS_OUT)/arm64.h $(IOS_OUT)/foo.h
