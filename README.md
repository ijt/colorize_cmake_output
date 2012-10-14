colorize_cmake_output
=====================

Highlight errors and warnings in the output of CMake.

Installing
----------

    go build
    sudo mv colorize_cmake_output /usr/local/bin
    echo '

    function cm() {
        cmake "$@" 2>&1 | colorize_cmake_output
    }

    ' >> ~/.bashrc
    source ~/.bashrc

Usage
-----
    
    cd build
    cm ../src

