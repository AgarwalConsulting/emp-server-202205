#!/usr/bin/env bash

docker run -it --rm -p 8000:9000 -e PORT=9000 emp-server
