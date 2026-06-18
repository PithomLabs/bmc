# BMC-REVIEW-0003 Implementation Plan: Literature Acquisition and Source-Provenance Intake

Create a documentation-only source intake and provenance ledger for literature required to review the BMC-0B minisuperspace massive-scalar Wheeler-DeWitt operator.

## User Review Required

> [!IMPORTANT]
> This ticket is strictly documentation-only. All mathematical physical debts (such as `OperatorFormDebt`, `MetricSignatureDebt`, etc.) will remain unpaid as "documented_unpaid" or "dependent_unpaid" due to missing external literature source texts in the repository.

## Open Questions

None. The 21 required sections and intake templates are fully specified.

## Proposed Changes

### Documentation Component

#### [NEW] [bmc_review_0003_literature_acquisition_source_provenance_intake.md](file:///home/chaschel/Documents/go/bmc/docs/reviews/bmc_review_0003_literature_acquisition_source_provenance_intake.md)
Create the primary review document containing the 21 required sections:
1. Review purpose and non-goals
2. Current accepted boundary from REVIEW-0001 and REVIEW-0002
3. Source-intake methodology
4. Repository source inventory
5. Existing citation/reference inventory
6. Missing source acquisition ledger
7. Source relevance criteria
8. Equation-intake template
9. Claim-intake template
10. Candidate source categories
11. Minisuperspace WdW operator source requirements
12. Massive scalar potential source requirements
13. Factor-ordering source requirements
14. Units/constants/convention source requirements
15. Boundary-condition source requirements
16. Classical-recovery target source requirements
17. Source-to-project notation mapping requirements
18. Human literature-review checklist
19. Debt status after source intake
20. No-promotion audit
21. Next allowed ticket recommendation

## Verification Plan

### Automated Tests
- Run Go package regression tests:
  ```bash
  GOCACHE=/tmp/go-build-cache go test ./... -count=1
  ```
- Run Lake build for Lean component:
  ```bash
  cd BMC && /home/chaschel/.elan/bin/lake build
  ```

### Manual Verification
- Perform the forbidden phrase scan to verify that any matches only appear in the blocked/negated/no-promotion sections:
  ```bash
  grep -R "validated\|proved\|recovered\|ready\|successful\|physics_success\|bmc_validated\|friedmann_recovered\|quantum_gravity_progress\|full bmc unblocked\|bmc beats nulls" docs/reviews docs/gates docs/postmortem || true
  ```
