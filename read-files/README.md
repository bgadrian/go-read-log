I built this fake readers just to see what is the fastest way
to read trough a file and get its content.

The file is a random apache access log.

Because I used only buffered techniques, the file size should not matter, right?! But it seems it does.

If you know any other method or I did a mistake please let me know or make a PR.

The results are [in the 200kb log](benchmark_200kb.log) and [in the 1MBlog](benchmark_1MB.log).