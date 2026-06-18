# BMC-REVIEW-0004 Implementation Plan: Source Text Acquisition and Equation Intake

Create a documentation-only source-text acquisition and equation-intake review mapping acquired equations and claims against project conventions.

## User Review Required

> [!IMPORTANT]
> This review is strictly documentation-only. All mathematical physical debts will remain unpaid (e.g. "documented_unpaid" or "dependent_unpaid") since no peer-reviewed external publications are available in the repository. Candidate equations extracted from Case 1 project notes are unreviewed.

## Open Questions

None. The 23 required sections and templates are fully specified.

## Proposed Changes

### Documentation Component

#### [NEW] [bmc_review_0004_source_text_acquisition_equation_intake.md](file:///home/chaschel/Documents/go/bmc/docs/reviews/bmc_review_0004_source_text_acquisition_equation_intake.md)
Create the primary review document containing the 23 required sections:
1. Review purpose and non-goals
2. Current accepted boundary from REVIEW-0001 through REVIEW-0003
3. Source acquisition status update
4. Case 2 and Case 3 resolution status
5. Newly available source-text inventory
6. Sources still missing
7. Equation-search methodology
8. Equation-intake ledger
9. Notation-mapping ledger
10. Convention-mapping ledger
11. Operator-form candidate intake
12. Metric/signature intake
13. Massive scalar potential intake
14. Factor-ordering intake
15. Units/constants intake
16. Boundary/domain intake
17. Bohmian guidance and Q-potential intake
18. Classical-recovery target intake
19. Source-to-project fixture comparison
20. Debt status after equation intake
21. Human physics-review checklist
22. No-promotion audit
23. Next allowed ticket recommendation

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
