#!/bin/bash

##################################################
# This command is executed when calling either
# 'ash make:exec main' or simply 'ash make'
#
# This command builds the project and runs it
##################################################
Command__main() {
    # Install deps if vendor file doesn't exist
    if [[ ! -d vendor ]]; then
        Command__install_deps
        local result="$?"
        if [[ "$result" -ne 0 ]]; then
            return "$result"
        fi
    fi

    # Generate Docs - TODO

    # Build
    Command__build
    local result="$?"
    if [[ "$result" -ne 0 ]]; then
        Logger__error ""
        return "$result"
    fi

    # Run the server
    Command__run
    return "$?"
}

##################################################
# This command is executed when calling
# 'ash make:exec install_deps'
#
# This command installs all of the dependencies
# as specified in the Gomfile
##################################################
Command__install_deps() {
    # Log
    Logger__warning "Installing dependencies..."

    # Checking if we have gom
    local gom_location="$(which gom)"
    if [[ "$gom_location" = "" ]]; then
        Logger__error "gom is not installed, which is required to install dependencies"
        Logger__alert "To install gom run: 'go get github.com/mattn/gom' and try again"
        return 1
    fi

    # Installing dependencies
    gom install

    # Verify Success
    local result=$?
    if [[ "$result" -ne 0 ]]; then
        Logger__error "Failed to install dependencies"
        Logger__error "Something is likely wrong with your Gomfile"
    fi

    # Pass along success
    return "$result"
}

##################################################
# This command is executed when calling
# 'ash make:exec build'
#
# This command builds the project
##################################################
Command__build() {
    # Log
    Logger__warning "Building..."

    # Build
    gom build
    return $?
}

##################################################
# This command is executed when calling
# 'ash make:exec run'
#
# Runs the server
##################################################
Command__run(){
    # Log
    Logger__warning "Starting the server..."

    if [[ ! -f "${PWD##*/}" ]]; then
        Logger__error "You need to build the project before running it"
        Logger__alert "Consider using 'ash make' as opposed to running this command directly"
        return 1
    fi

    ./${PWD##*/} development
    return $?
}
