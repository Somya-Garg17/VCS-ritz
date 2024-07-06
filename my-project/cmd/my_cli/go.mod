module my-project/cmd/my_cli

go 1.22.3

replace my-project/pkg/ritz => ../../pkg/ritz

replace my-project/pkg/utility => ../../pkg/utility

require my-project/pkg/ritz v0.0.0

require (
	github.com/sergi/go-diff v1.3.1 // indirect
	my-project/pkg/utility v0.0.0 // indirect
)
