@echo off
echo ----------------------------------
echo NUM_PROCS		INPUT_COUNT		RUNTIME
start /b /wait go run MapReduce_prettyPrint.go -input_directory=32 -processor_count=8
echo ----------------------------------
echo Complete!