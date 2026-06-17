The patch keeps residual/recovery/superiority claims blocked and the tests/CLI/Lean build pass,
  but the validator and generated report still allow several schema and accounting violations. These
  issues should be repaired before accepting Sprint 9 as an honest null-model runner artifact.

  Full review comments:

  - [P2] Reject decorative comparison records — /home/chaschel/Documents/go/bmc/internal/bmc/
    nullrun/validate.go:212-218
    When target_null_comparison_computed is true, this loop only requires comparison_id,
    comparison_computed, and an allowed interpretation status; it does not require metrics_compared
    to be nonempty or null_model_ids to reference null runs with generated diagnostics. A report can
    therefore validate with a comparison record that compares no metrics or points at no comparable
    null, which makes the target/null comparison gate decorative rather than an honest diagnostic
    comparison.

  - [P2] Reject sentinel float metrics — /home/chaschel/Documents/go/bmc/internal/bmc/nullrun/
    validate.go:176-186
    This check rejects nonfinite floats but still accepts finite sentinel values such as -1 for
    min_amplitude_r, max_abs_q_away_from_nodes, or max_phase_gradient. For unavailable diagnostics
    those fields are required to be JSON null rather than sentinel numbers, so a report can
    currently validate while encoding missing float diagnostics as -1.

  - [P2] Preserve unavailable metrics as explicit nulls — /home/chaschel/Documents/go/bmc/internal/
    bmc/nullrun/report.go:24-26
    Because these pointer fields use omitempty, unavailable metrics are omitted from the generated
    JSON instead of being emitted as null values; the current classical reference run already drops
    all three fields. The Sprint 9 schema requires unavailable numeric diagnostics to be represented
    as JSON null, so downstream validators/readers cannot distinguish an intentionally unavailable
    metric from a missing field.

  - [P2] Use accepted EBP debt status values — /home/chaschel/Documents/go/bmc/internal/bmc/nullrun/
    report.go:289-299
    Several generated EBP debt fields are set to active, but the accepted debt vocabulary for this
    review is limited to unpaid, partial, retired, contested, overclaimed, and absent. Since the
    validator only checks containsFinalTruthClaim and needFaithfulnessReview, the artifact can
    validate while carrying out-of-vocabulary debt classifications.

  - [P2] Account for deferred null models in the summary — /home/chaschel/Documents/go/bmc/cmd/ptw-
    bmc/main.go:437-438
    The summary prints registered, diagnostics-generated, and blocked counts, but the default report
    has 7 runs with only 4 diagnostics-generated and 1 blocked, leaving the 2 deferred runs
    unreported. This makes the run accounting opaque for the advertised Sprint 9 summary and should
    explicitly show the deferred count or otherwise reconcile all seven models.
