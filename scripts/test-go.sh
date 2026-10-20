#!/usr/bin/env bash
set -euo pipefail

if [[ -n "${RUBIO_TEST_GATE_COUNT_OUTPUT:-}" ]]; then
  rm -f "$RUBIO_TEST_GATE_COUNT_OUTPUT"
fi

results=$(mktemp)
trap 'rm -f "$results"' EXIT
mkdir -p reports/unit
gotestsum --junitfile=reports/unit/junit.xml --jsonfile="$results" -- -count=1 -p 1 ./...
count=$(jq -s '[.[] | select(.Action == "pass" and .Test != null)] | length' "$results")
if [[ "$count" -le 0 ]]; then
  echo "Canonical Go suite executed zero passing tests" >&2
  exit 1
fi

bash scripts/picker_snapshot_tmux.sh

if [[ -n "${RUBIO_TEST_GATE_COUNT_OUTPUT:-}" ]]; then
  printf '%s\n' "$count" >"$RUBIO_TEST_GATE_COUNT_OUTPUT"
fi
