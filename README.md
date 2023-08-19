<!--
Onur is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Onur is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Onur. If not, see <https://www.gnu.org/licenses/>.
-->

# Onur

Easily manage multiple FLOSS repositories

[rust](https://github.com/easbarba/onur-rust) | [java](https://github.com/easbarba/onur-java) | [python](https://github.com/easbarba/onur-python) | [php](https://github.com/easbarba/onur-php) | [ruby](https://github.com/easbarba/onur-ruby)

## Usage

`onur` consumes configuration in the following manners:

By default it looks for configuration files at `$XDG_CONFIG/onur` or in the
directory set in the `$ONUR_CONFIG_HOME` environment variable.

```shell
onur grab
onur archive nuxt,awesomewm,gitignore
```

<!-- Of course, a `JSON` configuration file can provide projects; -->

<!-- ```shell -->
<!-- onur grab --json ~/Downloads/misc.json -->
<!-- ``` -->

<!-- or it consumes even a REST API `JSON` resource providing all the projects. -->

<!-- ```shell -->
<!-- onur grab --api localhost:5000/configs -->
<!-- ``` -->

<!-- PS: an API example is at: https://github.com/easbarba/onur_api. -->

## Configuration file

A `onur` single configuration file:

```json
[
  {
    "name": "awesomewm",
    "branch": "master",
    "url": "https://github.com/awesomeWM/awesome"
  },
  {
    "name": "nuxt",
    "branch": "main",
    "url": "https://github.com/nuxt/framework"
  }
]
```

## Settings

A TOML settings file may define the behavior of `onur`:

```toml
single-branch = true
quiet = true
depth = 1
```

More examples of configuration files are at `docs/examples`.

## Options

Consult `onur --help` for more options.

# Installation

`go install github.com/easbarba/onur-go@latest`

## GNU Guix

In a system with GNU Guix binary installed, its even easier to grab all
dependencies: `guix shell`.

## TODO

Check the `TODO.md` for more information.

## LICENSE

[GPL-v3](https://www.gnu.org/licenses/gpl-3.0.en.html)
