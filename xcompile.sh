#!/bin/sh

for OS in windows linux darwin; do
  mkdir -p ${OS}
  GOOS=${OS} GOARCH=amd64 go build -o ${OS}/getpod
  if [ ${OS} == "windows" ]; then
    mv windows/getpod windows/getpod.exe
  fi
  zip getpod-${OS}-amd64.zip -r ${OS}/
  rm -rf ${OS}/
done
