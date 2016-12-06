# Logger

A tiny little, over opinionated, logging package.

Most of the time the standard log package is fine. Sometimes you find
you need levels, or named loggers. This package aims to satisfy those
use cases.

By default logging is written to `os.Stdout`, and all levels are logged.

## Usage

See the [examples](./examples/) directory.

```
package main

import "github.com/scottjbarr/logger"

func main() {
	log := logger.New("SomeProcess")

	log.Debug("Hello %v, nice to see you %v", 42, "today")
	log.Info("Yar %s", "hello")
	log.Warn("Yar %s", "hello")
	log.Error("Yar %s", "hello")
	log.Info("Boring")
}
```

## Log Level

You can set the Level of the Logger.

Example.

    // this will only log WARN level and higher
    log.Level = logger.LevelWarn

## Log to an io.Writer?

You can also choose to log to another location other than Stdout by
setting `log.Out` to any `io.Writer`.

## License

The MIT License (MIT)

Copyright (c) 2016 Scott Barr

See [LICENSE.md](LICENSE.md)
