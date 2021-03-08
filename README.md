# Openshift Knative Eventing Kafka

This repository holds Openshift's fork of
[`knative-sandbox/eventing-kafka`](https://github.com/knative-sandbox/eventing-kafka) with additions and
fixes needed only for the OpenShift side of things.

## List of releases

- [release-v0.19.1](https://github.com/openshift-knative/eventing-kafka/tree/release-v0.19.1)

## How this repository works ?

The `main` branch holds up-to-date specific [openshift files](./openshift)
that are necessary for CI setups and maintaining it. This includes:

- Scripts to create a new release branch from `upstream`
- CI setup files
  - operator configuration (for Openshift's CI setup)
  - tests scripts
- Operator's base configurations

Each release branch holds the upstream code for that release and our
openshift's specific files.

## CI Setup

For the CI setup, two repositories are of importance:

- This repository
- [openshift/release](https://github.com/openshift/release) which
  contains the configuration of CI jobs that are run on this
  repository
