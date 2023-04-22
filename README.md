


### Directory Structure

- `docs/`: Application + Infrastructure documentation and diagrams.
- `infra/tf`: Terraform modules specific to application infrastructure (e.g database, assets, vms, lb)
- `infra/kube`: Kubernetes manifests.

## Project Overview 

### Requirements

fictional company lifespan: 5years

Key generation algorithm
- needs to be short (4-8char)
- always must be unique
- base62 encoding

Throughput:
- 10M Daily active users
- 250K Daily users that create a new link
- 2.8935185185 Write/s
- 115.740740741 Read/s

2x Read Replica
1x Master, 1x Failover

2x pgbench nodes

Data storage:
- 250,000*365*5 =  456,250,000 Required URLs
- 62^5 = 916,132,832 Possible URLs 
- Assuming average long URL length of 30char, total storage per url would be:
36+5(user id)+10 = 51bytes
51*250,000 bytes = 12.75mb/day written (387MB/month, 4,653.6MB/Year)
0.5GB Read/day
#### Design considerat

#### Functional

- Given a URL, our service should generate a *shorter* and *unique* alias for it
- Users should be redirected to the original URL when they visit the short link
- Links should expire after a default timespan

#### Non-functional

- High availability
- Scalable, efficient
- Simple as reasonably possible
- Prevent abuse of service
- Record logs 

### Tech stack

#### Frontend
- React + Typescript
- S3 Static Site

#### Backend

- Go REST API
- 