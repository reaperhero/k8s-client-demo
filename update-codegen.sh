#!/bin/bash
# gcl git@git.zhlh6.cn:kubernetes/code-generator.git
$GOPATH/src/github.com/kubernetes/code-generator/generate-groups.sh \
  all github.com/reaperhero/k8s-client-demo/pkg/client \
  github.com/reaperhero/k8s-client-demo/pkg/apis \
  nginx_controller:v1
