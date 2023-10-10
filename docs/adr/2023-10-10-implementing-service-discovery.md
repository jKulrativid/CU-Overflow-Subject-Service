# Implementing Service Discovery

## Status

waiting for team discussion and approval

## Context

Service discovery has been one of our non-functional requirements
since we have already decided to use microservice architecture.

While data-plane service discovery tools such as Nginx and Envoy
might work well for a simple microservice architecture, complex services, for example,
elastically scalable services require more sophisticated tools like service mesh.

## Decision

Most service mesh tools have supported static service registration,
thus moving to service mesh tools may be a better choice.

According to [DevOps Cube](https://devopscube.com/open-source-service-discovery/),
three potential candidates have been proposed, which are Etcd, Consul, and Apache Zookeeper.

The zookeeper should be remove from consideration since it does not support HTTP/JSON API.

The docs of Etcd is quite sucks, on the contrary, Consul's official toturial
provides elaborated walkthrough guides with detailed examples.

Furthermore, Consul has support for load-balancer integration, fostering service scalability
which is the most significant non-functional property of our project.

## Consequences

Not decided yet!