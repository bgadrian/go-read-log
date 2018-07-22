#Golang fastest method to read a text file, line by line
I built this fake readers just to see what is the fastest way
to read trough a file and get its content.

The file is a random nginx access log.

If you know any other method or I did a mistake please let me know or make a PR.

The results are saved in the "benchmark_*.log" files. 
The actual logs that are read are "log_*.log".

Algorithms in [seq_readers](seq_readers.go) are the default "line" parsers from bufio. 

Algorithms in [seq_custom](seq_custom.go) uses ReaderAt implementations and does a custom line parsing. 

Algorithms in [seq_parallel](seq_parallel.go) read the file using multiple goroutines, splitting large files in chunks, hopefully doing it in parallel.