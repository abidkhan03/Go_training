#!/bin/sh

# if a user write -b or --build, then build the project
if [ "$1" = "-b" ] || [ "$1" = "--build" ]; then
    echo "Building project..."
    ./scripts/build.sh
fi

./bin/server
