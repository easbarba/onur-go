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

FROM golang:1.22

MAINTAINER EAS Barbosa <easbarba@outlook.com>
LABEL version=${ONUR_VERSION}
LABEL description="Easily manage multiple FLOSS repositories."

ENV USER_NAME easbarba
ENV XDG_CONFIG_HOME /home/$USER_NAME/.config
ENV APP_HOME /home/$USER_NAME/app

RUN groupadd -r $USER_NAME &&\
       useradd -r -g $USER_NAME -m -s /bin/bash $USER_NAME &&\
       chown -R $USER_NAME /home/$USER_NAME &&\
       usermod -aG sudo $USER_NAME

WORKDIR $APP_HOME
COPY examples/ $XDG_CONFIG_HOME/onur

RUN chown -R $USER_NAME:$USER_NAME /home/$USER_NAME
USER $USER_NAME

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download

COPY . .
RUN ./prepare.bash

CMD [ "go", "test", "-v", "./..." ]
