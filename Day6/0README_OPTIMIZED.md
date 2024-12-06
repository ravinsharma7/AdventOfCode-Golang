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

