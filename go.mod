module github.com/hootuu/nineorai

go 1.23

replace github.com/hootuu/gelato => ../gelato

require github.com/hootuu/gelato v0.0.0-20241211102738-092d538d30c2

require (
	filippo.io/edwards25519 v1.1.0
	github.com/mr-tron/base58 v1.2.0
)

require github.com/rs/xid v1.6.0 // indirect
