module rfgodatamongo

go 1.18

replace rfgocore => E:/trabajo/repos/go/rfgocore

replace rfgodata => E:/trabajo/repos/go/rfgodata

require (
	gopkg.in/guregu/null.v4 v4.0.0
	rfgocore v0.0.1
	rfgodata v0.0.1
)

require (
	go.mongodb.org/mongo-driver v1.9.0
	golang.org/x/text v0.3.7 // indirect
)

require (
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.15.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
)
