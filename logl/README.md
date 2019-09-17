# This is

A simple logging library with simple writes buffering.

# Usage

Import dependency

```go
import (
	"github.com/shved/advanced_go_exercises/logl"
)
```

Provide logger options struct
```go
var opts LoggerOptions = LoggerOptions{
	Dest: os.Open("some/file"), // any io.Writer is supported; default to stdout
	Source: "my-app",
	Separator: " // ", // default to \\t
	Level: logl.Info,
	BufferLength: 1024 // default to 4096, min 1024
}
```

Initialize logger
```go
logger := logl.NewLogger(opts)
```

And log messages
```go
logger.Log("my super duper event fired!", time.Now(), logl.Info) // time will be casted to UTC
```
