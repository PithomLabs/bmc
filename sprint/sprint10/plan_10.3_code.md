Implemented Sprint 10.3.

  Key changes:

  - Added explicit non-omitempty Lambda *float64 \json:"lambda"`` on internal/bmc/residualrun/
    residual.go:14.

  - Updated residual generation to emit one audited residual input point per finite-difference
    interval, with num_evaluation_points == len(residual_input_points) in internal/bmc/residualrun/
    report.go:499.

  - Replaced hard-coded comparison target candidate_residual_branch_0 with dynamically collected
    computed residual IDs in internal/bmc/residualrun/report.go:601.

  - Tightened validator checks for point count, branch/index alignment, finite lambda/alpha/phi/LHS/
    RHS, local strictly increasing lambda, allowed provenance, and phrase-safe forbidden formula IDs
    in internal/bmc/residualrun/validate.go:368.

  - Added regression tests for generated comparison targets, residual input point integrity, count
    mismatch, invalid provenance, local lambda monotonicity, and forbidden formula IDs.

  Verification passed:

  - GOCACHE=/tmp/go-build-cache go test ./internal/bmc/residualrun -count=1
  - GOCACHE=/tmp/go-build-cache go test ./...
  - CLI run-residuals, validate, and summarize against /tmp/bmc_s10_3_residual.json
  - lake build

  Note: the worktree already contains broader uncommitted/untracked Sprint 10.x and generated Lean/
  build/output artifacts; I did not revert or clean those.
