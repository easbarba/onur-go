<!--
Qas is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Qas is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Qas. If not, see <https://www.gnu.org/licenses/>.
-->

# Qas

[![building and testing](https://github.com/easbarba/qas_go/workflows/ci/badge.svg)](https://github.com/easbarba/qas_go/actions)

Easily manage multiple FLOSS repositories

[php](https://github.com/easbarba/qas_php) | [api](https://github.com/easbarba/qas_api) | [ruby](https://github.com/easbarba/qas_rb)

## Usage

`qas` consumes configuration in the following manners:

By default it looks for configuration files at `$XDG_CONFIG/qas` or in the
directory set in the `$QAS_CONFIG_HOME` environment variable.

```shell
qas grab
qas archive nuxt,awesomewm,gitignore
```

<!-- Of course, a `JSON` configuration file can provide projects; -->

<!-- ```shell -->
<!-- qas grab --json ~/Downloads/misc.json -->
<!-- ``` -->

<!-- or it consumes even a REST API `JSON` resource providing all the projects. -->

<!-- ```shell -->
<!-- qas grab --api localhost:5000/configs -->
<!-- ``` -->

<!-- PS: an API example is at: https://github.com/easbarba/qas_api. -->

## Configuration file

A `qas` single configuration file:

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

## Settings file

You may too set the behavior of `qas` with the `settings.toml`

```
singlebranch = true
depth = 2
```

More examples of configuration files are at `docs/examples`.

## Options

Consult `qas --help` for more options.

# Installation

`go install github.com/easbarba/qas-go@latest`

## GNU Guix

In a system with GNU Guix binary installed, its even easier to grab all
dependencies: `guix shell`.

## TODO

Check the `TODO.md` for more information.

## LICENSE

[GPL-v3](https://www.gnu.org/licenses/gpl-3.0.en.html)
