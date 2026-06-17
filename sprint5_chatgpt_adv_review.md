Sprint 5 adversarial review received. Do not start Sprint 6.

Current status: `accept_with_repairs`.

Required repairs before Sprint 5 acceptance:

1. Fix JSON field-name consistency.

   * Ensure the generated JSON field is exactly:
     `friedmann_readiness`
   * Ensure Go struct tags, decoder, validator, summarizer, and tests all use `json:"friedmann_readiness"`.
   * Remove or fix any accidental `"friedmann readiness"` key with a space.

2. Verify Sprint 5 readiness report contains all 12 fragile-configuration branch-audit runs.
   Required configs:

   * c2=0.50, k2=2.1, omega2=-2.1
   * c2=0.55, k2=1.9, omega2=-1.9
   * c2=0.55, k2=2.0, omega2=-2.0
   * c2=0.55, k2=2.1, omega2=-2.1

   Required step sizes:

   * dt=0.05, steps=200
   * dt=0.025, steps=400
   * dt=0.0125, steps=800

   Total required:

   * 4 configs × 3 dt values = 12 branch-audit runs.

3. Update the readiness scope string.

   * If all 12 runs are present, state that exactly.
   * If fewer than 12 are present for a justified reason, state exactly what is missing and why.
   * Do not use vague wording like “additional step sizes under validation.”

4. Complete tests.
   Add or update tests for:

   * all 4 fragile configs × 3 step sizes,
   * deterministic ordering of the 12 branch-audit runs,
   * flat segment edge case,
   * single-point final segment edge case,
   * unsorted turning-point rejection,
   * finite clock-independent diagnostics,
   * `DisallowUnknownFields` rejecting unexpected keys,
   * `local_only_candidate` not unblocking full BMC.

5. Fix EBP debt status.

   * `clock_choice_debt` must not be `retired`.
   * Use `active` in the report and `contested` only where human review is explicitly required.
   * Full BMC toy gate remains blocked.
   * Friedmann residual remains deferred.
   * `local_only_candidate` must not mean ready for Friedmann recovery.

6. Lean policy repair.

   * Do not add undefined EVP/batchV2/PCM-lite obligations unless those are already approved specs.
   * If the reviewer’s “human review gap” is addressed, do it with a clear policy-only field/theorem, for example:

     * `faithfulnessContested : Bool`
     * `humanFaithfulnessReviewRequired : Bool`
     * theorem: `clock_readiness_requires_faithfulness_contested_or_human_review`
   * Do not prove clock physics or Friedmann readiness in Lean.

7. Rerun verification:

   ```bash
   go test ./...
   go build -buildvcs=false ./cmd/ptw-bmc
   ./ptw-bmc segment-clock --profile bmc0a-clock-readiness --out out/bmc0a_clock_readiness.json
   ./ptw-bmc validate --report out/bmc0a_clock_readiness.json
   ./ptw-bmc summarize --report out/bmc0a_clock_readiness.json
   cd BMC && /home/chaschel/.elan/bin/lake build
   ```

Return a concise repair report with:

* files changed,
* confirmation of exact `friedmann_readiness` JSON key,
* exact branch-audit run count,
* test names added,
* command results,
* updated Sprint 5 promotion status.

Do not implement Friedmann residual recovery, massive scalar, LQC/Page-Wootters comparison, full BMC promotion, or any full quantum-gravity claim.
