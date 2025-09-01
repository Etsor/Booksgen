#!/bin/bash

RESET='\033[0m'

FG_RED='\033[31m'
FG_GREEN='\033[32m'
FG_WHITE='\033[37m'

go build -o ./build/booksgen ./cmd/

if [ $? -ne 0 ]; then
    echo -e "${FG_RED}Error while building${RESET}"
    exit 1
else 
    echo -e "${FG_GREEN}Build successful${RESET}"
fi