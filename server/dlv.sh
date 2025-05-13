#!/bin/sh
cd /build
dlv --listen=:40000 --headless=true --api-version=2 --accept-multiclient exec /cmd/web/bin/golang-react-todo-app
