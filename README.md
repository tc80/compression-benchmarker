# Benchmark Tool for Loseless Compression Algorithms

A simple program to generate the resulting size in bytes for a number of [cdnjs](https://cdnjs.com/) files after compressing with an algorithm.

It was developed with Go version 1.14.4 and macOS Catalina 10.15.5.

## Input
The application reads from `input.txt`, which contains the list of target cdnjs files for compression. Each entry must begin on a new line and prefixed with `/ajax/libs/`, corresponding to entries in the [cdnjs GitHub repo](https://github.com/cdnjs/cdnjs).

For example:
`/ajax/libs/font-awesome/4.7.0/fonts/fontawesome-webfont.woff2`

More examples can be found in this repo's [input.txt](input.txt).

## Environment
The application relies on an environment variable `CDNJS_PATH`. This variable must contain the path to the local [cdnjs GitHub repo](https://github.com/cdnjs/cdnjs) after cloning.

For example:
`export CDNJS_PATH="/tmp/cdnjs/cdnjs"`

Here, `/tmp/cdnjs/cdnjs/` will contain the `ajax/` directory.

## Output
In order for this tool to work, the algorithm *must* output its result to `STDOUT`. This is because this tool focuses on the size of the resulting output, not writing the output itself to a file.

## Build

To build, use the `Makefile` provided.

Simply run `make` to build and `make clean` to remove the executable/output directory.

To run the program for an algorithm with `n` arguments:

`./compress <shell command for algorithm> <algorithm argument 1> ... <algorithm argument n>`

For example, to benchmark `zopfli -i1000 --gzip` on all libaries in `input.txt`, run:

`./compress zopfli -c -i1000 --gzip`

Note that we added `-c` here to ensure the output is written to `STDOUT`.

## Built-In Algorithms

The provided `Makefile` contains a number of built-in options.

- [x] uncompressed
- [x] [zopfli](https://github.com/google/zopfli)
    - [x] gzip
    - [x] zlib
    - [x] deflate
- [x] [brotli](https://github.com/google/brotli)
    - [x] q5
    - [x] q11

To run all of them at once:

`make run`

This will generate a number of csv files in the `output/` directory.

