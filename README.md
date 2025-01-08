# Monte Carlo Simulation with Concurrent Programming

**Course:** Concurrent Programming (ICP-361)  
**Institution:** Federal University of Rio de Janeiro (UFRJ)  

---

## Problem Description

The Monte Carlo algorithm estimates the area of a geometric figure using random points generated within a known area, such as a square. For each point, a condition determines whether it falls within the target figure, incrementing a counter accordingly. By calculating the ratio of points inside the target figure to the total points in the known area, the algorithm estimates the area of the target figure.

The algorithm becomes more accurate as the number of points (`N`) increases. In this project, the input is the number of points (`N`), and the output is an estimation of π.

Parallelizing the Monte Carlo simulation improves performance significantly due to the independence of point generation and verification tasks. Benefits of parallelization include reduced execution time, scalability with increased threads, and efficiency when handling large datasets.

---

## Solution Design

### Concurrent Strategies

Two strategies were considered for implementing concurrency:

1. **Task Pool**  
   A shared counter tracks the total points generated, and threads increment it as they process points.  
   - **Advantages:** Simplicity and dynamic task distribution prevent idle threads.  
   - **Disadvantages:** Possible contention on the shared counter in highly concurrent scenarios.

2. **Partitioned Workload**  
   The total number of points is evenly divided among threads.  
   - **Advantages:** Balanced workload among threads ensures consistent performance.  
   - **Disadvantages:** Increased implementation complexity and potential coordination overhead.

**Chosen Strategy:** The task pool approach was selected due to its ability to dynamically balance workloads and reduce overhead associated with task coordination.

---

## Testing and Validation

### Correctness Tests

The algorithm's correctness is evaluated based on the precision and consistency of the π estimation. Three test cases are used:

1. **Few Points:** Tests with a small number of points; expected results show low variation around 3.14.  
2. **Moderate Points:** Tests with a moderate number of points, yielding results closer to 3.14.  
3. **Many Points:** Tests with a large number of points, achieving high precision near 3.1415.

Validation involves comparing the computed π value with Python's `math.pi` to ensure accuracy.

### Performance Tests

Performance is assessed by varying the number of points and threads:

1. **Few Points, Multiple Threads (1, 2, 4, 8):**  
   - Slight reduction in execution time as threads increase.  
2. **Moderate Points, Multiple Threads (1, 2, 4, 8):**  
   - Noticeable reduction in execution time compared to fewer points.  
3. **Many Points, Multiple Threads (1, 2, 4, 8):**  
   - Significant reduction in execution time, demonstrating good scalability.

### Scalability Expectations

- **Speedup:** Expected to increase with more threads up to a certain limit.  
- **Efficiency:** Initially improves with additional threads but may decline due to thread management overhead beyond optimal concurrency levels.

---

## References

- Amazon Web Services. [What is Monte Carlo simulation?](https://aws.amazon.com/pt/what-is/monte-carlo-simulation/)  
- IBM. [Monte Carlo simulation](https://www.ibm.com/topics/monte-carlo-simulation)  
- MathWorks. [Improving Performance of Monte Carlo Simulation with Parallel Computing](https://www.mathworks.com/help/finance/improving-performance-of-monte-carlo-simulation-with-parallel-computing.html)  
- Kristia, A. (2016). [Parallel Monte Carlo](https://agustinus.kristia.de/techblog/2016/06/13/parallel-monte-carlo/)  

---

This project demonstrates the practical application of parallel computing principles to enhance the performance and scalability of Monte Carlo simulations.
