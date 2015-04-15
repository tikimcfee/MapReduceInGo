@echo off
echo ----------------------------------
echo NUM_PROCS	INPUT_COUNT	RUNTIME
start /b /wait go run MapReduce_slow.go -input_directory=32 -processor_count=8
start /b /wait go run MapReduce_slow.go -input_directory=64 -processor_count=8
start /b /wait go run MapReduce_slow.go -input_directory=128 -processor_count=8
start /b /wait go run MapReduce_slow.go -input_directory=256 -processor_count=8
echo ----------------------------------
echo Trial 1 Complete!