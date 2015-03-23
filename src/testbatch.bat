@echo off
echo ----------------------------------
echo Starting new_MapReduce.go tests...
for %%i in (1, 2, 3) do start /b /wait go run new_MapReduce.go
echo ----------------------------------
echo ----------------------------------
echo Starting wordcount.go tests...
for %%i in (1, 2, 3) do start /b /wait go run .\wordcount\wordcount.go
echo ----------------------------------
echo Complete!