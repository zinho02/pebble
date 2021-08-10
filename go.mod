module github.com/zinho02/pebble

replace gopkg.in/square/go-jose.v2 => ../go-jose

require (
	github.com/letsencrypt/challtestsrv v1.2.0
	github.com/letsencrypt/pebble v1.0.1
	github.com/miekg/dns v1.1.15
	github.com/zinho02/go-jose v2.6.0+incompatible // indirect
	gopkg.in/square/go-jose.v2 v2.6.0
)

go 1.16
