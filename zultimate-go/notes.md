# 1. Design Guidelines
- Think about the cost of the design, it is always give & take relationship. There is no good or bad design, it is all about the trade-off.
- Think about what we are doing, what is the problem we are trying to solve.
- Performance vs Productivity, we need to balance both.
- Go models are about machine, as it interacts directly with the hardware.
- Do optimization for correctness, not for performance. Do benchmarking and profiling.
- Optimization can be done only when something is working, it can not be done prematurely.
- Integrity is important, It's about the reliability of the system.
- Every Read/Write/Memory management should be consistent.
- Error handling can not be exception based, it should be explicit. It should be a part of code.
- Code should be readable, it should be easy to understand.
- Engineering is about understanding the cost of the design/solution.
- Mental model of the code should be clear.
- Avoid debugger to solve a bug as in production, it is not possible to debug.
- Your code should not hide cost of the design.

# 2. Language Syntax
