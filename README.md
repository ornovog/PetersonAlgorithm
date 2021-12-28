# PetersonAlgorithm
in this project we can see a common issue when writing multi-threaded code.
the compiler do instructions reordering which can affect the logic and correctness of the code.

in this example I implement Peterson algorithm for mutual exclusion that allows two or more processes to
share a single-use resource without conflict, using only shared memory for communication.

the first implementation uses ordinary instructions, which we can see the compiler reorder, with running
"go build -gcflags -S petersonAlgorithm.go" in terminal. also we can see that his test fails (it can fail because deadlock or
race condition in multi-cores computer).   

the second implementation uses ordinary instructions, which prevent those issues.
