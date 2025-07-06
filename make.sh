#! /bin/sh

GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build -ldflags="-w -s" -o bin/jiggler-MacOSARM64
GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -ldflags="-w -s" -o bin/jiggler-MacOSAMD64
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" go build -ldflags="-w -s" -o bin/jiggler.exe #  -H windowsgui -r img2icons.rc" -o bin/jiggler.exe
# GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -ldflags="-w -s" -o bin/jiggler-LinuxAMD64
