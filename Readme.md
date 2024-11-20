# Simple Microservice in Go

### Plan
- Build a gateway that will accept HTTP REST API calls.
- Build services that will have specific tasks to manage.
- TCP protocol for ISC (Inter service communication)
    - Custom HTTP Impl with Path, method, payload


### Services
- Auth 
    - Design | done 
    - Implementation | done
- Math
    - Design | pending 
    - Implementation | pending
- Physics
    - Design | pending 
    - Implementation | pending

### Request Trace
- Request Entry - Request Exit
    - Entry -- When the request arrives at our doorstep (gateway)
    - Exit -- When the request is completed (response returned to the issuer/client, http request session is closed)

- Trace - lifecycle
    - Span - smaller segments within the lifecycle

### Improvements
- Use environment or config files to load service ports
    - environment is preferred because it can be changed easily when launching the service
    - config files are conveninent when the configuration is sharable (e.g app/service names, prefixes, configuration options without secrets, etc.)
