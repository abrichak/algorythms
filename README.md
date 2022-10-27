# Overview
The aim of the project is to implement and compare different algorithms including the algorithms of sorting

## Usage
1. To see the particular algorithm working, set the necessary value of the `N` constant (the number of elements in array), and uncomment the line with the name of necessary algorithm in the `main()` function. Then run the project. The array of `N` pseudo-random elements will be created and sorted. The initial array and result will be printed to console. 

2. The project contains a set of benchmarks to estimate the time of an algorithm for different values of `N`. To get the result of an algorithm, uncomment the lines with its name in all the `Benchmark...` functions of the `main_sorting_test.go` file. Then, run the command: 

- `go test -bench=.`

    The set of benchmarks allows to compare the time necessary to the algorithm for the array sizes of `N`, `10N`, `100N`, `1000N`, and also for different cases of initial array:
    
    - the 'best case' - when the initial array is already sorted;
    - an 'average case' - non-sorted initial array consisting of pseudo-random elements;
    - the 'worst case' - when the initial array is sorted in the backward direction.
  
3. The `golang.org/x/perf/cmd/benchstat` util allows to compare the benchmarks for different algorithms. To do the comparison, we need to save the benchmark results for every algorithm to a separate file. First, run the set of benchmark for an algorithm_1:
- `go test -bench=. | tee algorithm_1.txt`
  
  Then, change the calls in the benchmarks code to an `algorithm_2` and run the set of benchmarks directing the results to another file:
- `go test -bench=. | tee algorithm_2.txt`

  Install and run the `benchstat` utility:
- `go install golang.org/x/perf/cmd/benchstat@latest`
- `benchstat algorithm_1.txt algorithm_2.txt`

  The benchmarks for different algorithms and their comparison are located in the `benchmarks` directory.

