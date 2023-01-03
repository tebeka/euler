#!/bin/bash

case $1 in
    -h | --help ) echo "usage: $(basename $0) file"; exit;;
esac

if [ $# -ne 1 ]; then
    echo "error: wrong number of arguments" >&2
    exit 1
fi

set -e
rustc -o /tmp/eulers $1
/tmp/eulers
