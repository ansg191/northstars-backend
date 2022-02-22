module github.com/ansg191/northstars-backend/users

go 1.15

require (
	github.com/ansg191/northstars-backend/cookie-stealer v0.0.0-20220221022916-799a541a8c4a
	github.com/ansg191/northstars-backend/database v0.0.0-20220222051536-7092d4ebdafd
	github.com/golang/protobuf v1.5.2
	github.com/micro/micro/v3 v3.9.0
	github.com/nyaruka/phonenumbers v1.0.74
	google.golang.org/protobuf v1.27.1
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
//replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
