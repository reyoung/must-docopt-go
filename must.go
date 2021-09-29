package must_docopt_go

import (
	"github.com/docopt/docopt-go"
	"os"
)

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Opts struct {
	docopt.Opts
}

func MustParse(doc, version string) Opts {
	opts, err := docopt.ParseArgs(doc, os.Args, version)
	panicErr(err)
	return Opts{opts}
}

func (o Opts) MustString(key string) string {
	v, err := o.String(key)
	panicErr(err)
	return v
}

func (o Opts) MustFloat64(key string) float64 {
	v, err := o.Float64(key)
	panicErr(err)
	return v
}

func (o Opts) MustInt(key string) int {
	v, err := o.Int(key)
	panicErr(err)
	return v
}

func (o Opts) MustBool(key string) bool {
	v, err := o.Bool(key)
	panicErr(err)
	return v
}

func (o Opts) MustBind(v interface{}) {
	panicErr(o.Bind(v))
}
