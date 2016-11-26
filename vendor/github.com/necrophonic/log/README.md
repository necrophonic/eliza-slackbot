# go-simple-logger

[![](https://godoc.org/github.com/necrophonic/log?status.svg)](http://godoc.org/github.com/necrophonic/log) [![Go Report Card](https://goreportcard.com/badge/github.com/necrophonic/log)](https://goreportcard.com/report/github.com/necrophonic/log)

```
  import "github.com/necrophonic/log"
```

Very simple logger for golang. Similar to Go's core `log`
but with granular log levels (trace, debug, warn, error, info, fatal) and always outputs unbuffered to `STDOUT` and `STDERR` as appropriate - it's up to client to redirect this to somewhere more appropriate if desired.

```go
import "github.com/necrophonic/log"

// LevelTrace > LevelDebug > LevelWarn > LevelInfo > LevelError > LevelNone
log.Init(log.LevelInfo)

// Log at TRACE level - would be suppressed in example
// due to level being set at INFO.
log.Trace("Trace this out")

// Log at INFO level - equivalent operation to fmt.Print
log.Info("This is more informational")

// Log at INFO level with formatting - equivalent operation to fmt.Printf
log.Infof("This %s is nice", "thing")
```

Functions
---------

Provides functions for each level of granularity. Each has both "regular"
(equivalent to ``fmt.Print``) and "formatted" (equivalent to ``fmt.Printf``) versions.

See [go doc](http://godoc.org/github.com/necrophonic/log) for details

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


Log levels
----------

If you've ever used something like [log4j](http://logging.apache.org/log4j/) then you should be familiar with levels.

Essentially you set a level for output and only logging statements that match or exceed that level will be output.

The hierarchy of levels is thus (high to low):

  - None (no log output apart from Fatal)
  - Error
  - Info
  - Warn
  - Debug
  - Trace (everything)


Licence
-------

This code is licenced under the permissive MIT licence. Basically, you can do whatever you want as long as you include the original copyright and license notice in any copy of the software/source. [Read this](https://tldrlegal.com/license/mit-license) for a summary.
