# PART 1 OPTIMIZED 1

## Key Improvements
- Static [][]bool Grid for Visited Tracking:
  - Eliminates the overhead of dynamic map allocation.
- Early Exit Optimization:
  - Although not explicitly implemented here, a hash-based or path-tracking mechanism could be used to detect cycles in the guard's path.
- Lookup Table for Direction Changes:
  - Avoids modulo operations by using a map for direction transitions.
- Inline Updates:
  - Updates to positions and directions are performed directly without intermediate function calls.  

## Performance Gains  
- Memory Efficiency: Fixed-size grid for visited tracking reduces memory fragmentation and dynamic allocation overhead.  
- Speed: Inline computations and lookup tables minimize unnecessary computation and function call overhead.  

## Potential Next Steps  
- Cycle Detection: Identify if the guard repeats a previously visited state (position + direction) to terminate early in cyclic cases.  
- Concurrency: Parallelize the visited grid updates if the input grid is large.  


# PART 2 OPTIMIZED 1

To optimize the obstruction simulation further, we can focus on these aspects:

- Minimize State Representation Overhead:
  - Use a compact representation for the guard's state (position + direction) to reduce memory usage and lookup times.
- Avoid Unnecessary Simulations:
  - Pre-filter potential obstruction positions to exclude those where the guard's path is unlikely to change significantly.
- Optimize Obstruction Checks:
  - Simulate the patrol once and store the reachable positions. Only test obstructions at reachable positions for efficiency.
- Parallel Processing:
  - Split the grid into chunks and evaluate obstruction positions in parallel.

# Part 2 OPTIMIZED 2

The primary optimization in Code 2 compared to Code 1 is the significant reduction of overhead in detecting visited states and handling the grid structure. Here are the key points that make Code 2 more efficient:  

- Use of a Fixed-Size Visited Array Instead of a Map:
  - In Code 1, the simulation uses a map keyed by a struct (containing position and direction) to track visited states. Maps involve hashing and can be relatively slow and memory-intensive for frequent lookups and insertions, especially in tight loops.
  - In Code 2, the code replaces this map with a three-dimensional boolean array visited[x][y][dir]. This array provides constant-time lookups and updates, removing the hashing overhead and memory allocation costs associated with maps. This results in far fewer CPU cycles spent on state checks.

- Ensuring a Rectangular Grid for Simpler Indexing:
  - Code 2 pre-processes the input to ensure a perfectly rectangular grid by padding shorter lines with '.' characters. This guarantees that all rows have the same length, allowing direct indexing without conditional checks. The uniform grid dimension simplifies boundary checks and ensures consistent, cache-friendly memory access patterns.

- Elimination of Redundant Grid Copies:
  - Code 1 uses a helper function copyGrid to duplicate the grid for each simulation with an obstruction. Code 2 does not copy the grid for each test; instead, it simulates the obstruction by considering the target cell blocked only during the specific run. This removes the overhead of repeated memory allocations and copies.

- Explicit CPU Utilization (GOMAXPROCS):
  - While both versions parallelize work, Code 2 explicitly sets runtime.GOMAXPROCS(runtime.NumCPU()) to utilize all available CPU cores fully. Although Code 1 also uses concurrency, Code 2 makes it more explicit that all logical processors should be used efficiently.  

In summary, the main optimization in Code 2 is switching from a state-checking method that relies on a high-overhead map to a direct, array-based state-checking mechanism, combined with ensuring rectangular grid dimensions and reducing unnecessary copying. These changes significantly cut down on overhead and increase the performance of the simulation.  
