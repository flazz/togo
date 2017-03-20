#! /bin/sh -e

pkg=the
name=ThePdf
input=$pkg/The.pdf

go run main.go -pkg $pkg -name $name -input $input
go test ./$pkg
