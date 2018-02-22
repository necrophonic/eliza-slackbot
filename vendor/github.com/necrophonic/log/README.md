# go-simple-logger

```
  import "github.com/necrophonic/log"
```

Very simple logger for golang. Wrap default go [log](https://golang.org/pkg/log/) package to add granular functions.

```go

import "github.com/necrophonic/log"

// LevelTrace > LevelDebug > LevelWarn > LevelInfo > LevelError > LevelNone
log.Init(log.LevelInfo)

log.Trace("Trace this out")

log.Info("This is more informational")

log.Infof("This %s is nice", "thing")

```

Functions
---------

  - ``Fatal(v ...interface{})``
  - ``Fatalf(format string, v ...interface{})``
  - ``Trace(v ...interface{})``
  - ``Tracef(format string, v ...interface{})``
  - ``Debug(v ...interface{})``
  - ``Debugf(format string, v ...interface{})``
  - ``Info(v ...interface{})``
  - ``Infof(format string, v ...interface{})``
  - ``Warn(v ...interface{})``
  - ``Warnf(format string, v ...interface{})``
  - ``Error(v ...interface{})``
  - ``Errorf(format string, v ...interface{})``
