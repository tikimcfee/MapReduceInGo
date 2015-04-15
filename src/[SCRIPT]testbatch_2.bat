@echo off
echo ----------------------------------
echo NUM_PROCS	INPUT_COUNT	RUNTIME
start /b /wait go run MapReduce_prettyPrint.go -input_directory=32 -processor_count=1
start /b /wait go run MapReduce_prettyPrint.go -input_directory=32 -processor_count=2
start /b /wait go run MapReduce_prettyPrint.go -input_directory=32 -processor_count=4
start /b /wait go run MapReduce_prettyPrint.go -input_directory=32 -processor_count=8
start /b /wait go run MapReduce_prettyPrint.go -input_directory=64 -processor_count=1
start /b /wait go run MapReduce_prettyPrint.go -input_directory=64 -processor_count=2
start /b /wait go run MapReduce_prettyPrint.go -input_directory=64 -processor_count=4
start /b /wait go run MapReduce_prettyPrint.go -input_directory=64 -processor_count=8
start /b /wait go run MapReduce_prettyPrint.go -input_directory=128 -processor_count=1
start /b /wait go run MapReduce_prettyPrint.go -input_directory=128 -processor_count=2
start /b /wait go run MapReduce_prettyPrint.go -input_directory=128 -processor_count=4
start /b /wait go run MapReduce_prettyPrint.go -input_directory=128 -processor_count=8
start /b /wait go run MapReduce_prettyPrint.go -input_directory=256 -processor_count=1
start /b /wait go run MapReduce_prettyPrint.go -input_directory=256 -processor_count=2
start /b /wait go run MapReduce_prettyPrint.go -input_directory=256 -processor_count=4
start /b /wait go run MapReduce_prettyPrint.go -input_directory=256 -processor_count=8
echo ----------------------------------
echo Complete!