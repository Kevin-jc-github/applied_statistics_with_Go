# Bootstrap Median Estimator in Go

## Overview
This project demonstrates the implementation of a **Bootstrap method** to estimate the **standard error of the median** using the Go programming language. The bootstrap method is a resampling technique that is widely used in statistics to measure the accuracy of sample estimates.

In this project, we refactor the commonly used bootstrap sampling in R to Go, exploring the benefits of Go's concurrency capabilities and its efficient use of system resources.

## Selected Statistical Method
We have selected **Bootstrap Sampling** as our statistical method. Bootstrap sampling provides a robust, distribution-free approach to estimating the standard error of the median for different types of distributions (e.g., positively skewed, symmetric, and negatively skewed).

### R Implementation
The implementation in R can be done using the `boot` package or the `bootstrap` package available in CRAN:
- [boot package](https://cran.r-project.org/web/packages/boot/index.html)
- [bootstrap package](https://cran.r-project.org/web/packages/bootstrap/index.html)

These packages offer functions for conducting bootstrap analyses in R, which we have used for reference and comparison.

## Go Implementation
In the Go implementation, we:
1. Developed a function to perform **bootstrap resampling** on input data.
2. Implemented the logic to **estimate the median** for each bootstrap sample.
3. Calculated the **standard error** of the median across bootstrap samples.
4. **Employed Go testing, benchmarking, software profiling, and logging** to enhance performance and ensure correctness.

### Code Features
- **Concurrency:** The bootstrap sampling process can be optimized using Go's goroutines to fully utilize multiple CPU cores.
- **Testing and Profiling:** The implementation includes unit tests for critical functions, performance profiling using Go's `testing` and `pprof` packages, and logging to monitor key steps and any potential errors.

### How to Run the Code
1. **Clone the repository** using the command:
   ```sh
   git clone <repository-url>.git
   ```
2. **Navigate** to the project directory:
   ```sh
   cd bootstrap-median-go
   ```
3. **Run the code**:
   ```sh
   go run main.go
   ```
   The output will include bootstrap medians and the calculated standard error of the median, similar to the sample output shown below:
   ```
   Bootstrap Medians: [1.2 1.2 7.3 5.8 ...]
   Standard Error of Median: 7.1405
   ```
4. **Run tests**:
   ```sh
   go test
   ```
   Unit tests validate the correctness of the core functions and ensure expected behavior.

## Comparison with R
We compared the Go and R implementations using the same input data:
- **Memory Requirements:** The Go implementation had significantly lower memory overhead compared to R.
- **Processing Time:** By leveraging Go's concurrency, the bootstrap method runs faster, especially on multi-core machines.

### Improvements Made to Meet Assignment Requirements
1. **Go Testing and Benchmarking**: We added unit tests (`bootstrap_test.go`) to verify the correctness of the key functions, especially the median calculation and bootstrap sampling.
2. **Software Profiling**: Using Go's `pprof` package, we profiled the program to identify bottlenecks and optimize performance. This was crucial to compare processing times effectively against the R implementation.
3. **Logging**: Implemented logging to track key operations and to ensure traceability when running larger datasets.
4. **Refactoring**: The code was refactored to make functions modular and reusable, which aligns with best practices in Go.
5. **Memory Monitoring**: Memory usage during the execution was monitored using Go's profiling tools to compare resource usage with the R implementation, highlighting Go's efficiency.

## Cloud Computing Cost Savings
We also evaluated the potential cost savings when running these implementations on cloud services. By using Go, the consultancy can reduce cloud computing costs, especially when handling large datasets, due to:
- **Lower memory usage**.
- **Better utilization of multi-core processors**.

### Cloud Service Provider
We recommend using **Google Cloud Platform (GCP)** for running the Go implementation, as it offers flexible pricing for virtual machines (VMs) with multiple CPUs. Initial estimates show a potential **30-40% reduction** in cloud computing costs compared to using R.

## Recommendations
- **Use Go** for computationally intensive methods like bootstrapping when working with large datasets. The performance benefits and cost savings are substantial when leveraging Go's concurrent processing capabilities.
- R can still be preferred for rapid prototyping and for users more comfortable with its extensive library of statistical methods.
