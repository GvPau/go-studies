Considerations

1. Scalability:

- The in-memory approach works well for single-instance applications.
  However, in a distributed setup with multiple instances, each instance would have its own rate limit map, which can lead to inconsistent rate limiting.
  For distributed applications, a centralized store like Redis is preferred to ensure consistent rate limiting across all instances.

2. Memory Usage:

- In-memory rate limiting can consume significant memory if there are many unique clients.
  You may need to implement cleanup logic to remove old entries from the map.
