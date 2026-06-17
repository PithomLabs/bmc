package nullrun

import (
	"fmt"
)

// RunNullModels returns the null runs for a given profile.
func RunNullModels(profile string) ([]NullModelRun, error) {
	if profile != "bmc0a-nullrun" {
		return nil, fmt.Errorf("unknown profile: %s", profile)
	}

	valMinR := 1.2e-3
	valMaxQ := 4.5
	valMaxPG := 8.9

	runs := []NullModelRun{
		{
			NullModelID:          "constant_phase_control",
			RunStatus:            RunStatusDiagnosticsGenerated,
			DiagnosticProvenance: ProvenanceDeterministicFixture,
			Seed: &NullRunSeed{
				NullModelID: "constant_phase_control",
				Seed:        42,
				RngKind:     "none",
			},
			DiagnosticStatus: DiagStatusFinite,
			Diagnostics: NullDiagnostics{
				NodeContactFree:          true,
				TrajectoryFiniteness:     "finite",
				NumValidTrajectoryPoints: 100,
				NumClockSegments:         1,
				NumTurningPoints:         0,
				PhiClockGlobalStatus:     "monotonic",
				LocalBranchStatus:        "stable",
				MinAmplitudeR:            &valMinR,
				MaxAbsQAwayFromNodes:     &valMaxQ,
				MaxPhaseGradient:         &valMaxPG,
				DiagnosticWarnings:       []string{},
			},
			Notes: "Phase is held constant across coordinates; acts as simplest static control.",
		},
		{
			NullModelID:          "randomized_phase_control",
			RunStatus:            RunStatusDiagnosticsGenerated,
			DiagnosticProvenance: ProvenanceDeterministicFixture,
			Seed: &NullRunSeed{
				NullModelID: "randomized_phase_control",
				Seed:        12345,
				RngKind:     "xorshift64",
			},
			DiagnosticStatus: DiagStatusFinite,
			Diagnostics: NullDiagnostics{
				NodeContactFree:          true,
				TrajectoryFiniteness:     "finite",
				NumValidTrajectoryPoints: 100,
				NumClockSegments:         2,
				NumTurningPoints:         1,
				PhiClockGlobalStatus:     "non_monotonic",
				LocalBranchStatus:        "fragile",
				MinAmplitudeR:            &valMinR,
				MaxAbsQAwayFromNodes:     &valMaxQ,
				MaxPhaseGradient:         &valMaxPG,
				DiagnosticWarnings:       []string{"Phase fluctuations degrade relational monotonicity."},
			},
			Notes: "Phase values randomized at coordinates with recorded seed.",
		},
		{
			NullModelID:          "matched_amplitude_randomized_phase_control",
			RunStatus:            RunStatusDiagnosticsGenerated,
			DiagnosticProvenance: ProvenanceDeterministicFixture,
			Seed: &NullRunSeed{
				NullModelID: "matched_amplitude_randomized_phase_control",
				Seed:        54321,
				RngKind:     "xorshift64",
			},
			DiagnosticStatus: DiagStatusFinite,
			Diagnostics: NullDiagnostics{
				NodeContactFree:          true,
				TrajectoryFiniteness:     "finite",
				NumValidTrajectoryPoints: 100,
				NumClockSegments:         2,
				NumTurningPoints:         1,
				PhiClockGlobalStatus:     "non_monotonic",
				LocalBranchStatus:        "fragile",
				MinAmplitudeR:            &valMinR,
				MaxAbsQAwayFromNodes:     &valMaxQ,
				MaxPhaseGradient:         &valMaxPG,
				DiagnosticWarnings:       []string{"Phase fluctuations degrade relational monotonicity."},
			},
			Notes: "Amplitude is matched to target superposition wave while phase is randomized.",
		},
		{
			NullModelID:          "classical_frw_reference_trajectory",
			RunStatus:            RunStatusDiagnosticsGenerated,
			DiagnosticProvenance: ProvenanceDeterministicFixture,
			DiagnosticStatus:     DiagStatusFinite,
			Diagnostics: NullDiagnostics{
				NodeContactFree:          true,
				TrajectoryFiniteness:     "finite",
				NumValidTrajectoryPoints: 100,
				NumClockSegments:         1,
				NumTurningPoints:         0,
				PhiClockGlobalStatus:     "monotonic",
				LocalBranchStatus:        "stable",
				MinAmplitudeR:            nil,
				MaxAbsQAwayFromNodes:     nil,
				MaxPhaseGradient:         nil,
				DiagnosticWarnings:       []string{},
			},
			Notes: "reference comparator only; no residual or recovery interpretation",
		},
		{
			NullModelID:          "same_branch_segmentation_under_null_wavefunctions",
			RunStatus:            RunStatusDeferred,
			DiagnosticProvenance: ProvenanceDeferred,
			DiagnosticStatus:     DiagStatusNotAvailable,
			Diagnostics: NullDiagnostics{
				NodeContactFree:          false,
				TrajectoryFiniteness:     "not_available",
				NumValidTrajectoryPoints: 0,
				NumClockSegments:         0,
				NumTurningPoints:         0,
				PhiClockGlobalStatus:     "not_available",
				LocalBranchStatus:        "not_available",
				MinAmplitudeR:            nil,
				MaxAbsQAwayFromNodes:     nil,
				MaxPhaseGradient:         nil,
				DiagnosticWarnings:       []string{},
			},
			Notes: "Deferred until full relational clock segmentation is active on null sets.",
		},
		{
			NullModelID:          "node_neighborhood_stress_case",
			RunStatus:            RunStatusBlocked,
			DiagnosticProvenance: ProvenanceBlocked,
			BlockedReason:        "Blocked near nodes due to infinite quantum potential and division by zero risks.",
			DiagnosticStatus:     DiagStatusNodeBlocked,
			Diagnostics: NullDiagnostics{
				NodeContactFree:          false,
				TrajectoryFiniteness:     "node_blocked",
				NumValidTrajectoryPoints: 0,
				NumClockSegments:         0,
				NumTurningPoints:         0,
				PhiClockGlobalStatus:     "not_available",
				LocalBranchStatus:        "not_available",
				MinAmplitudeR:            nil,
				MaxAbsQAwayFromNodes:     nil,
				MaxPhaseGradient:         nil,
				DiagnosticWarnings:       []string{},
			},
			Notes: "Unable to generate trajectories inside node neighborhoods.",
		},
		{
			NullModelID:          "clock_choice_alternative_branch_diagnostic",
			RunStatus:            RunStatusDeferred,
			DiagnosticProvenance: ProvenanceDeferred,
			DiagnosticStatus:     DiagStatusNotAvailable,
			Diagnostics: NullDiagnostics{
				NodeContactFree:          false,
				TrajectoryFiniteness:     "not_available",
				NumValidTrajectoryPoints: 0,
				NumClockSegments:         0,
				NumTurningPoints:         0,
				PhiClockGlobalStatus:     "not_available",
				LocalBranchStatus:        "not_available",
				MinAmplitudeR:            nil,
				MaxAbsQAwayFromNodes:     nil,
				MaxPhaseGradient:         nil,
				DiagnosticWarnings:       []string{},
			},
			Notes: "Deferred until alternative relational clocks (e.g. scale factor a) are integrated in nullspec.",
		},
	}

	return runs, nil
}
