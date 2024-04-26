#!/usr/bin/env bash

# Onur is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# Onur is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with Onur. If not, see <https://www.gnu.org/licenses/>.

mkdir -p ~/.config/onur

cp ./examples/* ~/.config/onur

touch ~/.config/onur/emptyfile.json

ln -sf ~/nonexistentfile ~/.config/onur/danglingsymlink.json

rm ~/.config/onur/lua.json && cp ./examples/lua.json ~/ && ln -sf ~/lua.json ~/.config/onur/lualink.json
