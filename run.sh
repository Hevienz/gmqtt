#!/usr/bin/env sh

set -eux
set -o pipefail

gmqttd mergeconfig -o gmqttd.yml

gmqttd start -c gmqttd.yml
