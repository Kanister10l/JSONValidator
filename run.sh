#!/bin/bash

if go build -o bin/JSONValidator; then
    ./bin/JSONValidator
else
    echo "Build Failed"
fi
