module github.com/ansg191/northstars-backend/database

go 1.15

require (
	github.com/golang/protobuf v1.5.2
	github.com/gotidy/ptr v1.3.0
	github.com/micro/micro/v3 v3.9.0
	github.com/nyaruka/phonenumbers v1.0.74
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/postgres v1.3.1
	gorm.io/gorm v1.23.1
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
//replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
