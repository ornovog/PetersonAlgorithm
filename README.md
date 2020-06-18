# PetersonAlgorithm
in this project we can see a common issue when writing aync code.
the compiler doing instuctions reordering which can effect the logic corectness of the code.

in this example I implement Peterson algorithm for mutual exclusion that allows two or more processes to
share a single-use resource without conflict, using only shared memory for communication.

the first implementation use ordinary instructions, which we can see the compiler reorder with running
"go build -gcflags -S petersonAlgorithm.go" in terminal, and his test fails (it can fail because deadlock or race condition).

the second implementation use ordinary instructions, which prevent those issues.
