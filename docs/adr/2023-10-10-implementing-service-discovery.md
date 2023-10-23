# Implementing Service Discovery

## Status

ACCEPTED

## Context

Service discovery has been one of our non-functional requirements
since we have already decided to use microservice architecture.
Although we want to automated daunt tasks such as static IP configuration,
over-engineering must be avoided as much as possible.

Moreover, service discovery service should not create additional task to the team members.

Our candidates are Consul, Etcd, and Kubernetes Ingress (K8S Ingress).

Etcd is the least preferred since no member in our team has experience on it.

Consul is a powerful service discovery and service mesh tools.
Its documentation is straightforward and well-written.
However, both Consul and Etcd require deployment of sidecar container
and extra implementation.

K8S Ingress seems to be the most suitable
since it requires only images from other services
which has already been implemented and most of our members
have experience with K8S.

## Decision

We chose K8S Ingress by unanimity.

## Consequences

Better dev experience since static configuration is not needed.
