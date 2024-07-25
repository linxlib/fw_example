module github.com/linxlib/fw_example

go 1.22.0

replace (
	github.com/linxlib/astp => ../astp
	github.com/linxlib/fw => ../../repos/fw
)

require github.com/linxlib/fw v0.0.0-00010101000000-000000000000

require (
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/fasthttp/router v1.5.2 // indirect
	github.com/fasthttp/websocket v1.5.10 // indirect
	github.com/gookit/color v1.5.4 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/linxlib/astp v0.0.0-00010101000000-000000000000 // indirect
	github.com/linxlib/config v0.1.1 // indirect
	github.com/linxlib/conv v0.0.0-20200419055849-46faf16ac98f // indirect
	github.com/linxlib/inject v0.1.3 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/savsgio/gotils v0.0.0-20240704082632-aef3928b8a38 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.55.0 // indirect
	github.com/xo/terminfo v0.0.0-20210125001918-ca9a967f8778 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
