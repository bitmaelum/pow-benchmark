Benchmarking the proof-of-work aka how many bits do we need to work on?

To compile:

    go build main.go
    $ ./main   (or main.exe)
    
    
Output should be something similar like:


    Hi there! This program benchmarks your proof-of-work capabilities. This is basically a test to see
    how fast your system is in hashing SHA256 numbers. With this information, I can find a balance
    between effort (time) it takes for doing proof-of-work in BitMaelum, and workability. My aim is to
    find a bit-size that would result in a proof-of-work between 2-3 minutes on an average machine.
    
    Hit CTRL-C to stop. The more data, the better though. Please share the results on slack
    (DM to @jaytaph in PHPNL slack)
    ---------- START CPU INFO -------------
    (Don't send this part if you do not feel comfortable)
    
    OS     : windows
    Arch   : amd64
    CPUs   : 12 (cores: 6)
    Vendor : GenuineIntel
    Model  : Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
    ---------- END CPU INFO -------------
    ---------- START WORK INFO -------------
    Bits: 01   Cnt: 1000   Avg: 2.998µs         Total: 2.9982ms
    Bits: 02   Cnt: 1000   Avg: 999ns           Total: 999.9µs
    Bits: 03   Cnt: 1000   Avg: 6.003µs         Total: 6.0036ms
    Bits: 04   Cnt: 1000   Avg: 5µs             Total: 5.0001ms
    Bits: 05   Cnt: 1000   Avg: 9.999µs         Total: 9.999ms
    Bits: 06   Cnt: 1000   Avg: 16.001µs        Total: 16.001ms
    Bits: 07   Cnt: 1000   Avg: 143.999µs       Total: 143.9998ms
    Bits: 08   Cnt: 1000   Avg: 107µs           Total: 107.0005ms
    Bits: 09   Cnt: 1000   Avg: 294.583µs       Total: 294.5831ms
    Bits: 10   Cnt: 1000   Avg: 133.032µs       Total: 133.0322ms
    Bits: 11   Cnt: 1000   Avg: 424.488µs       Total: 424.4888ms
    Bits: 12   Cnt: 1000   Avg: 3.220716ms      Total: 3.2207165s
    Bits: 13   Cnt: 0571   Avg: 8.756981ms      Total: 5.0002362s
