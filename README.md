# supcolor

[![Go Reference](https://pkg.go.dev/badge/gopkg.ilharper.com/x/supcolor.svg)](https://pkg.go.dev/gopkg.ilharper.com/x/supcolor)
[![Go Report Card](https://goreportcard.com/badge/github.com/ifrstr/supcolor)](https://goreportcard.com/report/github.com/ifrstr/supcolor)

Determine whether a terminal supports color.

## Install

```sh
go get gopkg.ilharper.com/x/supcolor
```

## Usage

```go
import "gopkg.ilharper.com/x/supcolor"

// Variables (int8)
// https://github.com/chalk/supports-color#api
supcolor.Stdout // supcolor.SupColor(os.Stdout)
supcolor.Stderr // supcolor.SupColor(os.Stderr)
```

[![Go Reference](https://pkg.go.dev/badge/gopkg.ilharper.com/x/supcolor.svg)](https://pkg.go.dev/gopkg.ilharper.com/x/supcolor)

## LICENSE

[MIT](https://github.com/ifrstr/supcolor/blob/master/LICENSE)
