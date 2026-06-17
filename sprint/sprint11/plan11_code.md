 Implemented Sprint 11 as the residual/null comparison audit layer.

  What changed:

  - Added internal/bmc/residualaudit/ with report types, audit generation, stability diagnostics,
    strict validation, deterministic JSON I/O, gates, constants, and tests.

  - Added CLI routing:
      - ptw-bmc audit-residuals --profile bmc0a-residual-audit --out ...
      - ptw-bmc validate --report ...
      - ptw-bmc summarize --report ...

  - Used the tightened enum: comparison_integrity_structurally_honest; comparison_integrity_passed
    is not present.

  - Added policy-only Lean file BMC/BMC/ResidualAudit.lean:1 and imported it from BMC/BMC.lean:1.
  - Kept the audit candidate-only: no recovery, novelty, null failure, superiority, or full-BMC
    promotion claims.

  Verification passed:

  - GOCACHE=/tmp/go-build-cache go test ./internal/bmc/residualaudit -count=1
  - GOCACHE=/tmp/go-build-cache go test ./...
  - GOCACHE=/tmp/go-build-cache go build -buildvcs=false -o /tmp/ptw-bmc-s11 ./cmd/ptw-bmc
  - /tmp/ptw-bmc-s11 audit-residuals --profile bmc0a-residual-audit --out /tmp/
    bmc0a_residual_audit_s11.json

  - /tmp/ptw-bmc-s11 validate --report /tmp/bmc0a_residual_audit_s11.json
  - /tmp/ptw-bmc-s11 summarize --report /tmp/bmc0a_residual_audit_s11.json
  - lake build

  Note: the worktree already contains other uncommitted/generated files, and lake build updated Lean
  build outputs. I left unrelated existing changes untouched.

