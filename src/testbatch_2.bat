@echo off
echo ----------------------------------
echo Starting new_MapReduce.go tests... [input1024]
for %%i in (1, 2, 3) do start /b /wait go run new_MapReduce.go [input1024]
echo ----------------------------------
echo ----------------------------------
echo Starting wordcount.go tests...		[input1024]
for %%i in (1, 2, 3) do start /b /wait go run .\wordcount\wordcount.go [input1024]
echo ----------------------------------
echo Complete!