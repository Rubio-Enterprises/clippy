#!/usr/bin/env bash
set -euo pipefail

mkdir -p reports/test
gotestsum --junitfile=reports/test/junit.xml -- -count=1 -p 1 ./...

bash scripts/picker_snapshot_tmux.sh
