#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

HACK_DIR=$(dirname "${BASH_SOURCE[0]}")
REPO_ROOT="${HACK_DIR}/.."

"${REPO_ROOT}/vendor/k8s.io/code-generator/generate-groups.sh" all \
  github.com/cuijxin/postgres-operator-atom/pkg/generated github.com/cuijxin/postgres-operator-atom/pkg/apis \
  "acid.zalan.do:v1" \
  --go-header-file "${REPO_ROOT}"/hack/custom-boilerplate.go.txt