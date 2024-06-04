#!/bin/bash

set -e
set -x

function BuildChessEngineContainer () {
  APP_VERSION=$(grep "LABEL APP_VERSION" docker/Dockerfile | grep -o '"[^"]*"' | sed 's/"//g')
  docker build . -t chess-engine:"${APP_VERSION}" -f docker/Dockerfile --no-cache
}

BuildChessEngineContainer
