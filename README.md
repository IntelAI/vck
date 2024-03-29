# DISCONTINUATION OF PROJECT #
This project will no longer be maintained by Intel.
Intel has ceased development and contributions including, but not limited to, maintenance, bug fixes, new releases, or updates, to this project.
Intel no longer accepts patches to this project.
# Volume Controller for Kubernetes (VCK)

[![CircleCI](https://circleci.com/gh/IntelAI/vck.svg?style=svg)](https://circleci.com/gh/IntelAI/vck)
## Overview

This project provides basic volume and data management in Kubernetes v1.9+
using [custom resource definitions][crd] (CRDs), custom controllers,
[volumes][vols] and [volume sources][volsources]. It also 
establishes a relationship between volumes and data and provides a way to
abstract the details away from the user. When using VCK, users 
are expected to only interact with [custom resources][cr] (CRs).

![VCK overview figure](./docs/images/vck.png)

## Further Reading

- [Architecture][arch-doc]
- [Developer Manual][dev-doc]
- [Operator Manual][ops-doc]
- [User Manual][user-doc]
- [Best Practices][best-practices-doc]
- [FAQs][faqs-doc]

[arch-doc]: docs/arch.md
[dev-doc]: docs/dev.md
[ops-doc]: docs/ops.md
[user-doc]: docs/user.md
[faqs-doc]: docs/faq.md
[best-practices-doc]: docs/best-practices.md
[crd]: https://kubernetes.io/docs/tasks/access-kubernetes-api/extend-api-custom-resource-definitions/
[cr]: https://kubernetes.io/docs/concepts/api-extension/custom-resources/
[vols]: https://kubernetes.io/docs/concepts/storage/volumes/
[volsources]: https://github.com/kubernetes/api/blob/master/core/v1/types.go#L250
