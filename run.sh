#!/bin/bash
set -e
make build
exec ./out/epictectus start_http
