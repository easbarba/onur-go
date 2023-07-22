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

# PROTOCOL

## Config

- cache latest bundled configs as ONUR_CONFIG/.cached
- bundle all files with syntax as

```json
"lang": "py",
"projects": {"name": "python", "branch": "dev", "https://github.com/python/python"}
```

- branch field defaults to Git, that is `master`
- check if empty
- check if config has properly syntax
- `url` to URI
- `name` as lowercase and \_ only

### Api

- parse all files

## Local

- parse json files and bundle as api sintaxe
- url to uri

### Single

- bundle file as config syntax
- consumes piped input

## Front

- consume array properly processed
- run action on array
- log actions and it's output
