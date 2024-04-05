#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[@]}")/..

CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

bash "${CODEGEN_PKG}"/generate-internal-groups.sh \
  "deepcopy,conversion,defaulter" \
  github.com/mokaz111/mokazSche/pkg/generated \
  github.com/mokaz111/mokazSche/pkg/plugins/apis \
  github.com/mokaz111/mokazSche/pkg/plugins/apis \
  "config:v1beta2,v1beta3" \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt
