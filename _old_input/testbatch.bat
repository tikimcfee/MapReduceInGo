@echo off
echo ----------------------------------
echo Starting new_MapReduce.go tests... [input32]
for %%i in (1, 2, 3) do start /b /wait go run new_MapReduce.go [input32]
echo ----------------------------------
echo ----------------------------------
echo Starting wordcount.go tests...		[input32]
for %%i in (1, 2, 3) do start /b /wait go run .\wordcount\wordcount.go [input32]
echo ----------------------------------
echo Complete!


echo ----------------------------------
echo Starting new_MapReduce.go tests... [input64]
for %%i in (1, 2, 3) do start /b /wait go run new_MapReduce.go [input64]
echo ----------------------------------
echo ----------------------------------
echo Starting wordcount.go tests...		[input64]
for %%i in (1, 2, 3) do start /b /wait go run .\wordcount\wordcount.go [input64]
echo ----------------------------------
echo Complete!


echo ----------------------------------
echo Starting new_MapReduce.go tests... [input128]
for %%i in (1, 2, 3) do start /b /wait go run new_MapReduce.go [input128]
echo ----------------------------------
echo ----------------------------------
echo Starting wordcount.go tests...		[input128]
for %%i in (1, 2, 3) do start /b /wait go run .\wordcount\wordcount.go [input128]
echo ----------------------------------
echo Complete!

echo ----------------------------------
echo Starting new_MapReduce.go tests... [input256]
for %%i in (1, 2, 3) do start /b /wait go run new_MapReduce.go [input256]
echo ----------------------------------
echo ----------------------------------
echo Starting wordcount.go tests...		[input256]
for %%i in (1, 2, 3) do start /b /wait go run .\wordcount\wordcount.go [input256]
echo ---------------------------------- 
echo Complete!


echo ----------------------------------
echo Starting new_MapReduce.go tests... [input512]
for %%i in (1, 2, 3) do start /b /wait go run new_MapReduce.go [input512]
echo ----------------------------------
echo ----------------------------------
echo Starting wordcount.go tests...		[input512]
for %%i in (1, 2, 3) do start /b /wait go run .\wordcount\wordcount.go [input512]
echo ----------------------------------
echo Complete!

echo ----------------------------------
echo Starting new_MapReduce.go tests... [input1024]
for %%i in (1, 2, 3) do start /b /wait go run new_MapReduce.go [input1024]
echo ----------------------------------
echo ----------------------------------
echo Starting wordcount.go tests...		[input1024]
for %%i in (1, 2, 3) do start /b /wait go run .\wordcount\wordcount.go [input1024]
echo ----------------------------------
echo Complete!

echo ----------------------------------
echo Starting new_MapReduce.go tests... [input2048]
for %%i in (1, 2, 3) do start /b /wait go run new_MapReduce.go [input2048]
echo ----------------------------------
echo ----------------------------------
echo Starting wordcount.go tests...		[input2048]
for %%i in (1, 2, 3) do start /b /wait go run .\wordcount\wordcount.go [input2048]
echo ----------------------------------
echo Complete!

echo ----------------------------------
echo Starting new_MapReduce.go tests... [input4096]
for %%i in (1, 2, 3) do start /b /wait go run new_MapReduce.go [input4096]
echo ----------------------------------
echo ----------------------------------
echo Starting wordcount.go tests...		[input4096]
for %%i in (1, 2, 3) do start /b /wait go run .\wordcount\wordcount.go [input4096]
echo ----------------------------------
echo Complete!