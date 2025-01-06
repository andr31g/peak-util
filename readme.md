## Peak Detector Utility

### NB: Currently this is more of a test driver, than an actually useful utility.

### Checkout the repository

```sh
❯ git clone https://github.com/andr31g/peak-util.git
❯ 
```

### Build the `go` executable

It was tested with `go` version `1.21` and up.

```sh
❯ go build
❯ ls -l
...
-rwxr-xr-x 1 ...  peakdetect
...
❯
```

### Run tests

To run tests, invoke the command with the `test` or `t` sub-command.

```sh
❯ ./peakdetect test
...
3 3 3 3 3 3 3 2]
[0 1 2 3 4 5 6]
[3 3 3 3 3 3 3 3]
[]
Total: 65536
❯
```

### Process arbitrary input

To process arbitrary input, invoke the command with the `detect`
or `d` sub-command, and provide a list of comma-separated integer
samples via the `--samples` or `-s` flag. An optional `--iterations`
or `-i` flag can be used to specify the number of iterations of
the algorithm over the input data.

```sh
❯ ./peakdetect d -s 1,8,3,5,7,2,3,6,9,0,1,3,12,1,4,8,2,4,1,9
    Original Samples  [1 8 3 5 7 2 3 6 9 0 1 3 12 1 4 8 2 4 1 9]
         Peak Values  [0 8 0 0 7 0 0 0 9 0 0 0 12 0 0 8 0 4 0 9]
Aligned Peak Indices  [0 1 0 0 4 0 0 0 8 0 0 0 12 0 0 15 0 17 0 19]
        Peak Indices  [1 4 8 12 15 17 19]
❯
```
To specify a custom number of iterations, issue:

```sh
❯ ./peakdetect d -s 1,8,3,5,7,2,3,6,9,0,1,3,12,1,4,8,2,4,1,9 -i 2
Running 2 iterations
    Original Samples  [1 8 3 5 7 2 3 6 9 0 1 3 12 1 4 8 2 4 1 9]
         Peak Values  [0 8 0 0 0 0 0 0 0 0 0 0 12 0 0 0 0 0 0 9]
Aligned Peak Indices  [0 1 0 0 0 0 0 0 0 0 0 0 12 0 0 0 0 0 0 19]
        Peak Indices  [1 12 19]
❯
```


