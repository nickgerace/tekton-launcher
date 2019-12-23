# Tekton Launcher

**author: [Nick Gerace](https://nickgerace.dev)**

A simple Tekton launcher for Kubernetes.

## Requirements

Currently, the exact version strings required have not been confirmed.
Moreover, the requirements are constantly changing.
The end goal is for the requirements to be as minimal as possible.

*Permanent Requirements*

- Kubernetes cluster
- Tekton Pipelines (installation make target included)
- Tekton Dashboard (optional, not included)

*Temporary Requirements*

- Go 1.13+
- Kubectl
- Make

## License

- MIT License, Copyright (c) Nick Gerace
- See 'LICENSE' file for more information

## Notes

The client-go dependency version was changed to "0.17.0" since the library switched to matching Kubernetes version strings.
