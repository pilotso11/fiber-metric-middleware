# Fiber Middleware for Metric
Golang Fiber Middleware to Expose zserge/metric

[Metric](https://github.com/zserge/metric) middleware for [Fiber](https://github.com/gofiber/fiber) 
that servers metric created HTML pages via Fiber's HTTP server.  Metric exposes
application metrics with a lightweight graphical UI and integrates with expvar.  It is exposed
by default on `/debug/metrics`.

See also [expvar middleware](https://docs.gofiber.io/api/middleware/expvar) and [pprof middleware](https://docs.gofiber.io/api/middleware/pprof).

# Usage
```shell
go get github.com/pilotso11/metricmware
```
```go
import "github.com/pilotso11/metricmware"

app.Use(metricmware.New())
```

# Configuration
`metricmware.Config` exposes one option, which is a prefix before "/debug".

# Examples
Don't forget to checkout your stats at http://localhost:8000/debug/metrics

```go
import (
    "expvar"
    "github.com/gofiber/fiber/v2"
    expvarmw "github.com/gofiber/fiber/v2/middleware/expvar"
    "github.com/zserge/metric"
    "github.com/vade-mecum/mixfi-fix/middleware/metric"
)

app := fiber.New()
app.Use(expvarmw.New())   // Recommended if you use metric with expvar
app.Use(metricmware.New())

expvar.Publish("mycounter", metric.NewCounter("5m1s", "15m30s", "1h1m"))
expvar.Publish("mystat", metric.NewGauge("30m1m", "5h5m"))
expvar.Publish("mylatency", metric.NewHistogram("5m1s", "15m30s", "1h1m"))

_ = app.Listen("127.0.0.1:8000")
```
```go
// In my handlers
expvar.Get("mycounter").(metric.Metric).Add(1)   // To increase the counter
expvar.Get("mystat").(metric.Metric).Add(delta)  // Update the gauge 

starTime := time.Now()
// Do some work
expvar.Get("mylatency").(metric.Metric).Add(time.Since(startTime).Seconds())
```

# Thanks


