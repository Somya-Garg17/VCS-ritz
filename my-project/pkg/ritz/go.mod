module my-project/pkg/ritz

go 1.22.3

replace my-project/pkg/utility => ../../pkg/utility

require my-project/pkg/utility v0.0.0

require github.com/sergi/go-diff v1.3.1 // indirect
