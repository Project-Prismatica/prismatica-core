# prismatica-core

# Go usage:

Use the ```prismatica_grpc``` package:

```go
import (
	"github.com/Project-Prismatica/prismatica-core/go/prismatica_grpc"
)
```

## Compiling go bindings

1. Add the ```protoc``` plugin for [go](https://github.com/golang/protobuf).

2. run:
```bash
$ protoc -I $PROTOBUF_STDLIB -I ./grpc \
    --go_out=grpc,import_path=prismatica_grpc:./gp \
    ./grpc/*
```

Where ```$PROTOBUF_STDLIB``` is the path to the root of the protobug standard
library which should contain, for example,
```google/protobuf/timestamp.proto```.
