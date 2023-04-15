// Qas is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Qas is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Qas. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"github.com/alecthomas/kong"
	"github.com/easbarba/qas/internal/actions"
)

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	Backup BackupCmd `cmd help:"clean system residual packages dependencies"`
	Grab   GrabCmd   `cmd help:"install dependencies packages of package"`
}

type Context struct {
	Debug bool
}

type BackupCmd struct {
	Verbose  bool   `help:"Provide more information."`
	Packages string `arg name:"package" help:"get package dependencies."`
}

func (r *BackupCmd) Run(ctx *Context) error {
	actions.Backup(&r.Packages, &r.Verbose)
	return nil
}

type GrabCmd struct {
	Verbose bool `help:"Provide more information."`
}

func (r *GrabCmd) Run(ctx *Context) error {
	actions.Grab(&r.Verbose)
	return nil
}
