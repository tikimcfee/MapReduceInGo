@echo off
echo ----------------------------------
echo NUM_PROCS	INPUT_COUNT	RUNTIME
start /b /wait go run MapReduce_slow.go -input_directory=C:\GO_SRC\_old_input\32 -processor_count=8
start /b /wait go run MapReduce_slow.go -input_directory=C:\GO_SRC\_old_input\64 -processor_count=8
start /b /wait go run MapReduce_slow.go -input_directory=C:\GO_SRC\_old_input\128 -processor_count=8
start /b /wait go run MapReduce_slow.go -input_directory=C:\GO_SRC\_old_input\256 -processor_count=8
echo ----------------------------------
echo Trial 1 Complete!