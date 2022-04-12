module rfgodatamongo

go 1.17

replace rfgocore => E:/trabajo/repos/go/rfgocore

replace rfgodata => E:/trabajo/repos/go/rfgodata

require (
	go.mongodb.org/mongo-driver v1.9.0
	gopkg.in/guregu/null.v4 v4.0.0
	rfgocore v0.0.1
	rfgodata v0.0.1
)
