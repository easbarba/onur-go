// Onur is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Onur is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Onur. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"github.com/alecthomas/kong"
)

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	Backup BackupCmd `cmd help:"compress selected projects"`
	Grab   GrabCmd   `cmd help:"grab all projects"`
	Config ConfigCmd `cmd help:"Manage configurations"`
}

type Context struct {
	Debug bool
}

type BackupCmd struct {
	Verbose  bool   `help:"Provide more information."`
	Packages string `arg name:"package" help:"get package dependencies."`
}

func (r *BackupCmd) Run(ctx *Context) error {
	Backup(&r.Packages, &r.Verbose)
	return nil
}

type GrabCmd struct {
	Verbose bool   `help:"Provide more information."`
	Topic   string `arg:"" optional:"" help:"Name and Topic of configuration."`
}

func (r *GrabCmd) Run(ctx *Context) error {
	Grab(&r.Topic, &r.Verbose)
	return nil
}

type ConfigCmd struct {
	Verbose bool `help:"Provide more information."`

	Topic  string `arg:"" help:"Name and Topic of configuration."`
	Name   string `arg:"" optional:"" help:"Name of configuration."`
	Url    string `arg:"" optional:"" help:"Url of configuration."`
	Branch string `arg:"" optional:"" help:"Branch of configuration."`
}

func (r *ConfigCmd) Run(ctx *Context) error {
	Config(&r.Topic, &r.Name, &r.Url, &r.Branch, &r.Verbose)
	return nil
}
