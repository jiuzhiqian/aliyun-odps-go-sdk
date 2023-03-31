module odpsexample

go 1.17

replace github.com/jiuzhiqian/aliyun-odps-go-sdk/arrow => ../arrow

replace github.com/jiuzhiqian/aliyun-odps-go-sdk => ../

require (
	github.com/jiuzhiqian/aliyun-odps-go-sdk v0.0.1
	github.com/jiuzhiqian/aliyun-odps-go-sdk/arrow v0.0.1
	github.com/pkg/errors v0.9.1
)

require (
	github.com/golang/snappy v0.0.3 // indirect
	github.com/google/flatbuffers v2.0.0+incompatible // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/pierrec/lz4/v4 v4.1.11 // indirect
	golang.org/x/exp v0.0.0-20211123021643-48cbe7f80d7c // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
)
