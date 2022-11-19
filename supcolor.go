// Package supcolor determines whether a terminal supports color.
//
// # Usage
//
//	import "gopkg.ilharper.com/x/supcolor"
//
//	// Variables (int8)
//	// https://github.com/chalk/supports-color#api
//	supcolor.Stdout // supcolor.SupColor(os.Stdout)
//	supcolor.Stderr // supcolor.SupColor(os.Stderr)
package supcolor

import "os"

var (
	Stdout int8
	Stderr int8
)

func init() {
	Stdout = SupColor(os.Stdout)
	Stderr = SupColor(os.Stderr)
}
