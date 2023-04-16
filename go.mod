module github.com/cyb3rko/signal-back-2

go 1.12

require (
	github.com/cyb3rko/signal-back-2/signal v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.1.0
	github.com/h2non/filetype v1.1.3
	github.com/pkg/errors v0.8.1
	github.com/urfave/cli v1.20.0
	golang.org/x/crypto v0.0.0-20180808211826-de0752318171
	golang.org/x/sys v0.0.0-20181122145206-62eef0e2fa9b // indirect
)

replace github.com/cyb3rko/signal-back-2/signal => ./signal
