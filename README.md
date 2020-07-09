# Tekton Launcher

**author: [Nick Gerace](https://nickgerace.dev)**

![GitHub Actions](https://github.com/nickgerace/tekton-launcher/workflows/Go/badge.svg)

A simple Tekton launcher for Kubernetes.

## Requirements

Even though the dependencies and version strings changing (pre-alpha), the end goal is for the requirements to be as minimal as possible.

*Permanent Requirements*

- Kubernetes cluster v?
- [Tekton Pipelines](https://github.com/tektoncd/pipeline) v? (make target included)
- [Tekton Dashboard](https://github.com/tektoncd/dashboard) v? (optional, make target not included)

*Temporary Requirements*

- Go 1.13+
- Kubectl v?
- Make v?

## License

- MIT License, Copyright (c) Nick Gerace
- See 'LICENSE' file for more information

## Notes

The client-go dependency version was manually changed (in *go.mod*) to "0.17.0" since the library switched to matching Kubernetes version strings.
