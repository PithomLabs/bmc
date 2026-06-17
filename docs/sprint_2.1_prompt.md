Adversarial review result accepted. Do not start Sprint 3 yet.

Sprint 2 is provisionally successful but still `accept_with_repairs`.

Required repairs before Sprint 2 acceptance:

1. Add explicit safe-profile node-avoidance sampling test.
   - Name suggestion:
     TestSafeSuperpositionAmplitudeStaysAboveNodeThreshold
   - Generate or integrate the bmc0a-superposition-safe trajectory.
   - For every trajectory point, compute R = |Ψ(α, φ)|.
   - Assert R >= NodeThresh for all sampled points.
   - This must independently verify safe node avoidance, not only rely on node_contact_free pass.

2. Add exact-name node-probe semantic test required by Sprint 2 plan.
   - Name suggestion:
     TestNodeProbeShortCircuitAndValidationGate
   - Use bmc0a-superposition-node-probe.
   - Assert:
     node_detection = pass
     node_contact_free = fail
     node obstruction severity = blocker
     trajectory is empty or explicitly short-circuited/blocked
     safe superposition gate does not pass
     technical gate = node_detection_validation_gate
     technical gate status = pass
     full BMC toy gate remains blocked

3. Replace weak RK4 test with a non-constant-velocity test.
   - Current RK4 test is too weak if it uses constant velocity.
   - Use a simple known non-constant ODE, for example:
       dα/dλ = α
       dφ/dλ = -φ
     with initial α=1, φ=1.
   - Exact solution:
       α(λ)=exp(λ)
       φ(λ)=exp(-λ)
   - Compare one or multiple steps and assert RK4 error < Euler error.
   - Keep this as an integrator correctness test only; do not imply physics.

After repairs, rerun:

go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc run --profile bmc0a-superposition-safe --out out/bmc0a_superposition_safe.json
./ptw-bmc validate --report out/bmc0a_superposition_safe.json
./ptw-bmc summarize --report out/bmc0a_superposition_safe.json
./ptw-bmc run --profile bmc0a-superposition-node-probe --out out/bmc0a_superposition_node_probe.json
./ptw-bmc validate --report out/bmc0a_superposition_node_probe.json
./ptw-bmc summarize --report out/bmc0a_superposition_node_probe.json
cd BMC && lake build

Return a concise repair report with:
- files changed
- tests added/renamed
- command results
- updated Sprint 2 promotion status

Do not implement optional repairs unless trivial and scope-safe.
Do not implement massive scalar, Friedmann residual, LQC, Page-Wootters, perturbations, black holes, fermions, gauge fields, Lorentz recovery, or full quantum gravity.