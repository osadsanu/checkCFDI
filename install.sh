#! /bin/bash
rm -r bin/*
rm /usr/local/bin/cfdi
echo clean
go build -o bin/cfdi
echo compiled
mv bin/cfdi /usr/local/bin
echo installed
