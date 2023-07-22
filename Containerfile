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

FROM golang:1.20 AS build
MAINTAINER EAS Barbosa <easbarba@outlook.com>
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY examples /root/.config/onur
RUN touch ~/.config/onur/emptyfile.json && ln -sf ~/nonexistentfile ~/.config/onur/baz.json
COPY . .
CMD [ "go", "test", "-v", "./..." ]
