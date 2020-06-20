# Makefile for building and running the loseless compression benchmarking tool.
# @author tc80, June 2020
#
# Note, project was built with:
# 	Go version: go1.14.4 darwin/amd64
#	macOS Catalina 10.15.5

EXEC=compress
OUTPUT_DIR=output

# build the project
all: 
	go build -o $(EXEC)
	mkdir -p $(OUTPUT_DIR)

# runs all provided examples and writes results to file
run: uncompressed zopfli-gzip zopfli-zlib zopfli-deflate brotli-q5 brotli-q11

# view uncompressed sizes
uncompressed:
	./$(EXEC) cat > $(OUTPUT_DIR)/uncompressed.csv

# view zopfli gzip with 1000 iterations
zopfli-gzip:
	./$(EXEC) zopfli -c -i1000 --gzip > $(OUTPUT_DIR)/zopfli-gzip.csv

# view zopfli zlib with 1000 iterations
zopfli-zlib:
	./$(EXEC) zopfli -c -i1000 --zlib > $(OUTPUT_DIR)/zopfli-zlib.csv

# view zopfli deflate with 1000 iterations
zopfli-deflate:
	./$(EXEC) zopfli -c -i1000 --deflate > $(OUTPUT_DIR)/zopfli-deflate.csv

# view brotli q5
brotli-q5:
	./$(EXEC) brotli -c -q 5 > $(OUTPUT_DIR)/brotli-q5.csv

# view brotli q11
brotli-q11:
	./$(EXEC) brotli -c -q 11 > $(OUTPUT_DIR)/brotli-q11.csv

# cleans the executable and output dir
clean: 
	rm $(EXEC)
	rm -rf $(OUTPUT_DIR)