// Package rcflags implements command line flags to set up the remote control
package rcflags

import (
	"github.com/ncw/rclone/cmd/serve/httplib/httpflags"
	"github.com/ncw/rclone/fs/config/flags"
	"github.com/ncw/rclone/fs/rc"
	"github.com/spf13/pflag"
)

// Options set by command line flags
var (
	Opt = rc.DefaultOpt
)

// AddFlags adds the remote control flags to the flagSet
func AddFlags(flagSet *pflag.FlagSet) {
	rc.AddOption("rc", &Opt)
	flags.BoolVarP(flagSet, &Opt.Enabled, "rc", "", false, "Enable the remote control server.")
	flags.StringVarP(flagSet, &Opt.Files, "rc-files", "", "", "Serve these files on the HTTP server.")
	httpflags.AddFlagsPrefix(flagSet, "rc-", &Opt.HTTPOptions)
}
