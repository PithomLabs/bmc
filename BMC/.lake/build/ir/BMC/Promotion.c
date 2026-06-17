// Lean compiler output
// Module: BMC.Promotion
// Imports: public import Init public meta import Init public import BMC.ToyReport
#include <lean/lean.h>
#if defined(__clang__)
#pragma clang diagnostic ignored "-Wunused-parameter"
#pragma clang diagnostic ignored "-Wunused-label"
#elif defined(__GNUC__) && !defined(__CLANG__)
#pragma GCC diagnostic ignored "-Wunused-parameter"
#pragma GCC diagnostic ignored "-Wunused-label"
#pragma GCC diagnostic ignored "-Wunused-but-set-variable"
#endif
#ifdef __cplusplus
extern "C" {
#endif
uint8_t lp_bmc_checkPassed(uint8_t);
uint8_t lp_bmc_checkDeferred(uint8_t);
LEAN_EXPORT uint8_t lp_bmc_reportPassesBMC0AControlGate(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_reportPassesBMC0AControlGate___boxed(lean_object*);
LEAN_EXPORT uint8_t lp_bmc_reportPassesFullBMCToyGate(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_reportPassesFullBMCToyGate___boxed(lean_object*);
LEAN_EXPORT uint8_t lp_bmc_reportPassesBMC0ASuperpositionSafeGate(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_reportPassesBMC0ASuperpositionSafeGate___boxed(lean_object*);
LEAN_EXPORT uint8_t lp_bmc_reportPassesBMC0ANodeDetectionGate(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_reportPassesBMC0ANodeDetectionGate___boxed(lean_object*);
static const lean_ctor_object lp_bmc_sprint2SafeWitness___closed__0_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*0 + 16, .m_other = 0, .m_tag = 0}, .m_objs = {LEAN_SCALAR_PTR_LITERAL(1, 0, 0, 0, 0, 0, 0, 0),LEAN_SCALAR_PTR_LITERAL(0, 0, 2, 3, 0, 0, 0, 0)}};
static const lean_object* lp_bmc_sprint2SafeWitness___closed__0 = (const lean_object*)&lp_bmc_sprint2SafeWitness___closed__0_value;
LEAN_EXPORT const lean_object* lp_bmc_sprint2SafeWitness = (const lean_object*)&lp_bmc_sprint2SafeWitness___closed__0_value;
static const lean_ctor_object lp_bmc_sprint2NodeProbeWitness___closed__0_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*0 + 16, .m_other = 0, .m_tag = 0}, .m_objs = {LEAN_SCALAR_PTR_LITERAL(1, 0, 0, 1, 1, 0, 1, 3),LEAN_SCALAR_PTR_LITERAL(3, 1, 2, 3, 0, 0, 0, 0)}};
static const lean_object* lp_bmc_sprint2NodeProbeWitness___closed__0 = (const lean_object*)&lp_bmc_sprint2NodeProbeWitness___closed__0_value;
LEAN_EXPORT const lean_object* lp_bmc_sprint2NodeProbeWitness = (const lean_object*)&lp_bmc_sprint2NodeProbeWitness___closed__0_value;
LEAN_EXPORT uint8_t lp_bmc_reportPassesBMC0AControlGate(lean_object* v_r_1_){
_start:
{
uint8_t v_toyAnalysisOnly_2_; 
v_toyAnalysisOnly_2_ = lean_ctor_get_uint8(v_r_1_, 0);
if (v_toyAnalysisOnly_2_ == 0)
{
return v_toyAnalysisOnly_2_;
}
else
{
uint8_t v_finalTruthClaim_3_; 
v_finalTruthClaim_3_ = lean_ctor_get_uint8(v_r_1_, 1);
if (v_finalTruthClaim_3_ == 0)
{
uint8_t v_wdwResidual_4_; uint8_t v_trajectoryFinite_5_; uint8_t v_clockMonotonic_6_; uint8_t v_nodeDetection_7_; uint8_t v_nodeContactFree_8_; uint8_t v_qFiniteAwayFromNodes_9_; uint8_t v_phaseGradientFinite_10_; uint8_t v_classicalLimit_11_; uint8_t v_friedmannResidual_12_; uint8_t v___x_13_; 
v_wdwResidual_4_ = lean_ctor_get_uint8(v_r_1_, 2);
v_trajectoryFinite_5_ = lean_ctor_get_uint8(v_r_1_, 3);
v_clockMonotonic_6_ = lean_ctor_get_uint8(v_r_1_, 4);
v_nodeDetection_7_ = lean_ctor_get_uint8(v_r_1_, 5);
v_nodeContactFree_8_ = lean_ctor_get_uint8(v_r_1_, 6);
v_qFiniteAwayFromNodes_9_ = lean_ctor_get_uint8(v_r_1_, 7);
v_phaseGradientFinite_10_ = lean_ctor_get_uint8(v_r_1_, 8);
v_classicalLimit_11_ = lean_ctor_get_uint8(v_r_1_, 9);
v_friedmannResidual_12_ = lean_ctor_get_uint8(v_r_1_, 10);
v___x_13_ = lp_bmc_checkPassed(v_wdwResidual_4_);
if (v___x_13_ == 0)
{
return v___x_13_;
}
else
{
uint8_t v___x_14_; 
v___x_14_ = lp_bmc_checkPassed(v_trajectoryFinite_5_);
if (v___x_14_ == 0)
{
return v___x_14_;
}
else
{
uint8_t v___x_15_; 
v___x_15_ = lp_bmc_checkPassed(v_clockMonotonic_6_);
if (v___x_15_ == 0)
{
return v___x_15_;
}
else
{
uint8_t v___x_16_; 
v___x_16_ = lp_bmc_checkPassed(v_nodeDetection_7_);
if (v___x_16_ == 0)
{
return v___x_16_;
}
else
{
uint8_t v___x_17_; 
v___x_17_ = lp_bmc_checkPassed(v_nodeContactFree_8_);
if (v___x_17_ == 0)
{
return v___x_17_;
}
else
{
uint8_t v___x_18_; 
v___x_18_ = lp_bmc_checkPassed(v_qFiniteAwayFromNodes_9_);
if (v___x_18_ == 0)
{
return v___x_18_;
}
else
{
uint8_t v___x_19_; 
v___x_19_ = lp_bmc_checkPassed(v_phaseGradientFinite_10_);
if (v___x_19_ == 0)
{
return v___x_19_;
}
else
{
uint8_t v___x_20_; 
v___x_20_ = lp_bmc_checkPassed(v_classicalLimit_11_);
if (v___x_20_ == 0)
{
return v___x_20_;
}
else
{
uint8_t v___x_21_; 
v___x_21_ = lp_bmc_checkDeferred(v_friedmannResidual_12_);
return v___x_21_;
}
}
}
}
}
}
}
}
}
else
{
uint8_t v___x_22_; 
v___x_22_ = 0;
return v___x_22_;
}
}
}
}
LEAN_EXPORT lean_object* lp_bmc_reportPassesBMC0AControlGate___boxed(lean_object* v_r_23_){
_start:
{
uint8_t v_res_24_; lean_object* v_r_25_; 
v_res_24_ = lp_bmc_reportPassesBMC0AControlGate(v_r_23_);
lean_dec_ref(v_r_23_);
v_r_25_ = lean_box(v_res_24_);
return v_r_25_;
}
}
LEAN_EXPORT uint8_t lp_bmc_reportPassesFullBMCToyGate(lean_object* v_r_26_){
_start:
{
uint8_t v_toyAnalysisOnly_27_; 
v_toyAnalysisOnly_27_ = lean_ctor_get_uint8(v_r_26_, 0);
if (v_toyAnalysisOnly_27_ == 0)
{
return v_toyAnalysisOnly_27_;
}
else
{
uint8_t v_finalTruthClaim_28_; 
v_finalTruthClaim_28_ = lean_ctor_get_uint8(v_r_26_, 1);
if (v_finalTruthClaim_28_ == 0)
{
uint8_t v_wdwResidual_29_; uint8_t v_trajectoryFinite_30_; uint8_t v_clockMonotonic_31_; uint8_t v_nodeDetection_32_; uint8_t v_nodeContactFree_33_; uint8_t v_qFiniteAwayFromNodes_34_; uint8_t v_phaseGradientFinite_35_; uint8_t v_classicalLimit_36_; uint8_t v_friedmannResidual_37_; uint8_t v_faithfulness_38_; uint8_t v___x_39_; 
v_wdwResidual_29_ = lean_ctor_get_uint8(v_r_26_, 2);
v_trajectoryFinite_30_ = lean_ctor_get_uint8(v_r_26_, 3);
v_clockMonotonic_31_ = lean_ctor_get_uint8(v_r_26_, 4);
v_nodeDetection_32_ = lean_ctor_get_uint8(v_r_26_, 5);
v_nodeContactFree_33_ = lean_ctor_get_uint8(v_r_26_, 6);
v_qFiniteAwayFromNodes_34_ = lean_ctor_get_uint8(v_r_26_, 7);
v_phaseGradientFinite_35_ = lean_ctor_get_uint8(v_r_26_, 8);
v_classicalLimit_36_ = lean_ctor_get_uint8(v_r_26_, 9);
v_friedmannResidual_37_ = lean_ctor_get_uint8(v_r_26_, 10);
v_faithfulness_38_ = lean_ctor_get_uint8(v_r_26_, 11);
v___x_39_ = lp_bmc_checkPassed(v_wdwResidual_29_);
if (v___x_39_ == 0)
{
return v___x_39_;
}
else
{
uint8_t v___x_40_; 
v___x_40_ = lp_bmc_checkPassed(v_trajectoryFinite_30_);
if (v___x_40_ == 0)
{
return v___x_40_;
}
else
{
uint8_t v___x_41_; 
v___x_41_ = lp_bmc_checkPassed(v_clockMonotonic_31_);
if (v___x_41_ == 0)
{
return v___x_41_;
}
else
{
uint8_t v___x_42_; 
v___x_42_ = lp_bmc_checkPassed(v_nodeDetection_32_);
if (v___x_42_ == 0)
{
return v___x_42_;
}
else
{
uint8_t v___x_43_; 
v___x_43_ = lp_bmc_checkPassed(v_nodeContactFree_33_);
if (v___x_43_ == 0)
{
return v___x_43_;
}
else
{
uint8_t v___x_44_; 
v___x_44_ = lp_bmc_checkPassed(v_qFiniteAwayFromNodes_34_);
if (v___x_44_ == 0)
{
return v___x_44_;
}
else
{
uint8_t v___x_45_; 
v___x_45_ = lp_bmc_checkPassed(v_phaseGradientFinite_35_);
if (v___x_45_ == 0)
{
return v___x_45_;
}
else
{
uint8_t v___x_46_; 
v___x_46_ = lp_bmc_checkPassed(v_classicalLimit_36_);
if (v___x_46_ == 0)
{
return v___x_46_;
}
else
{
uint8_t v___x_47_; 
v___x_47_ = lp_bmc_checkPassed(v_friedmannResidual_37_);
if (v___x_47_ == 0)
{
return v___x_47_;
}
else
{
uint8_t v___x_48_; 
v___x_48_ = lp_bmc_checkPassed(v_faithfulness_38_);
return v___x_48_;
}
}
}
}
}
}
}
}
}
}
else
{
uint8_t v___x_49_; 
v___x_49_ = 0;
return v___x_49_;
}
}
}
}
LEAN_EXPORT lean_object* lp_bmc_reportPassesFullBMCToyGate___boxed(lean_object* v_r_50_){
_start:
{
uint8_t v_res_51_; lean_object* v_r_52_; 
v_res_51_ = lp_bmc_reportPassesFullBMCToyGate(v_r_50_);
lean_dec_ref(v_r_50_);
v_r_52_ = lean_box(v_res_51_);
return v_r_52_;
}
}
LEAN_EXPORT uint8_t lp_bmc_reportPassesBMC0ASuperpositionSafeGate(lean_object* v_r_53_){
_start:
{
uint8_t v_toyAnalysisOnly_54_; 
v_toyAnalysisOnly_54_ = lean_ctor_get_uint8(v_r_53_, 0);
if (v_toyAnalysisOnly_54_ == 0)
{
return v_toyAnalysisOnly_54_;
}
else
{
uint8_t v_finalTruthClaim_55_; 
v_finalTruthClaim_55_ = lean_ctor_get_uint8(v_r_53_, 1);
if (v_finalTruthClaim_55_ == 0)
{
uint8_t v_wdwResidual_56_; uint8_t v_trajectoryFinite_57_; uint8_t v_clockMonotonic_58_; uint8_t v_nodeDetection_59_; uint8_t v_nodeContactFree_60_; uint8_t v_qFiniteAwayFromNodes_61_; uint8_t v_phaseGradientFinite_62_; uint8_t v_classicalLimit_63_; uint8_t v_friedmannResidual_64_; uint8_t v___x_65_; 
v_wdwResidual_56_ = lean_ctor_get_uint8(v_r_53_, 2);
v_trajectoryFinite_57_ = lean_ctor_get_uint8(v_r_53_, 3);
v_clockMonotonic_58_ = lean_ctor_get_uint8(v_r_53_, 4);
v_nodeDetection_59_ = lean_ctor_get_uint8(v_r_53_, 5);
v_nodeContactFree_60_ = lean_ctor_get_uint8(v_r_53_, 6);
v_qFiniteAwayFromNodes_61_ = lean_ctor_get_uint8(v_r_53_, 7);
v_phaseGradientFinite_62_ = lean_ctor_get_uint8(v_r_53_, 8);
v_classicalLimit_63_ = lean_ctor_get_uint8(v_r_53_, 9);
v_friedmannResidual_64_ = lean_ctor_get_uint8(v_r_53_, 10);
v___x_65_ = lp_bmc_checkPassed(v_wdwResidual_56_);
if (v___x_65_ == 0)
{
return v___x_65_;
}
else
{
uint8_t v___x_66_; 
v___x_66_ = lp_bmc_checkPassed(v_trajectoryFinite_57_);
if (v___x_66_ == 0)
{
return v___x_66_;
}
else
{
uint8_t v___x_67_; 
v___x_67_ = lp_bmc_checkPassed(v_clockMonotonic_58_);
if (v___x_67_ == 0)
{
return v___x_67_;
}
else
{
uint8_t v___x_68_; 
v___x_68_ = lp_bmc_checkPassed(v_nodeDetection_59_);
if (v___x_68_ == 0)
{
return v___x_68_;
}
else
{
uint8_t v___x_69_; 
v___x_69_ = lp_bmc_checkPassed(v_nodeContactFree_60_);
if (v___x_69_ == 0)
{
return v___x_69_;
}
else
{
uint8_t v___x_70_; 
v___x_70_ = lp_bmc_checkPassed(v_qFiniteAwayFromNodes_61_);
if (v___x_70_ == 0)
{
return v___x_70_;
}
else
{
uint8_t v___x_71_; 
v___x_71_ = lp_bmc_checkPassed(v_phaseGradientFinite_62_);
if (v___x_71_ == 0)
{
return v___x_71_;
}
else
{
uint8_t v___x_72_; 
v___x_72_ = lp_bmc_checkPassed(v_classicalLimit_63_);
if (v___x_72_ == 0)
{
return v___x_72_;
}
else
{
uint8_t v___x_73_; 
v___x_73_ = lp_bmc_checkDeferred(v_friedmannResidual_64_);
return v___x_73_;
}
}
}
}
}
}
}
}
}
else
{
uint8_t v___x_74_; 
v___x_74_ = 0;
return v___x_74_;
}
}
}
}
LEAN_EXPORT lean_object* lp_bmc_reportPassesBMC0ASuperpositionSafeGate___boxed(lean_object* v_r_75_){
_start:
{
uint8_t v_res_76_; lean_object* v_r_77_; 
v_res_76_ = lp_bmc_reportPassesBMC0ASuperpositionSafeGate(v_r_75_);
lean_dec_ref(v_r_75_);
v_r_77_ = lean_box(v_res_76_);
return v_r_77_;
}
}
LEAN_EXPORT uint8_t lp_bmc_reportPassesBMC0ANodeDetectionGate(lean_object* v_r_78_){
_start:
{
uint8_t v_toyAnalysisOnly_79_; 
v_toyAnalysisOnly_79_ = lean_ctor_get_uint8(v_r_78_, 0);
if (v_toyAnalysisOnly_79_ == 0)
{
return v_toyAnalysisOnly_79_;
}
else
{
uint8_t v_finalTruthClaim_80_; 
v_finalTruthClaim_80_ = lean_ctor_get_uint8(v_r_78_, 1);
if (v_finalTruthClaim_80_ == 0)
{
uint8_t v_wdwResidual_81_; uint8_t v_nodeDetection_82_; uint8_t v_friedmannResidual_83_; uint8_t v___x_84_; 
v_wdwResidual_81_ = lean_ctor_get_uint8(v_r_78_, 2);
v_nodeDetection_82_ = lean_ctor_get_uint8(v_r_78_, 5);
v_friedmannResidual_83_ = lean_ctor_get_uint8(v_r_78_, 10);
v___x_84_ = lp_bmc_checkPassed(v_nodeDetection_82_);
if (v___x_84_ == 0)
{
return v___x_84_;
}
else
{
uint8_t v___x_85_; 
v___x_85_ = lp_bmc_checkPassed(v_wdwResidual_81_);
if (v___x_85_ == 0)
{
return v___x_85_;
}
else
{
uint8_t v___x_86_; 
v___x_86_ = lp_bmc_checkDeferred(v_friedmannResidual_83_);
return v___x_86_;
}
}
}
else
{
uint8_t v___x_87_; 
v___x_87_ = 0;
return v___x_87_;
}
}
}
}
LEAN_EXPORT lean_object* lp_bmc_reportPassesBMC0ANodeDetectionGate___boxed(lean_object* v_r_88_){
_start:
{
uint8_t v_res_89_; lean_object* v_r_90_; 
v_res_89_ = lp_bmc_reportPassesBMC0ANodeDetectionGate(v_r_88_);
lean_dec_ref(v_r_88_);
v_r_90_ = lean_box(v_res_89_);
return v_r_90_;
}
}
lean_object* initialize_Init(uint8_t builtin);
lean_object* initialize_Init(uint8_t builtin);
lean_object* initialize_bmc_BMC_ToyReport(uint8_t builtin);
static bool _G_initialized = false;
LEAN_EXPORT lean_object* initialize_bmc_BMC_Promotion(uint8_t builtin) {
lean_object * res;
if (_G_initialized) return lean_io_result_mk_ok(lean_box(0));
_G_initialized = true;
res = initialize_Init(builtin);
if (lean_io_result_is_error(res)) return res;
lean_dec_ref(res);
res = initialize_Init(builtin);
if (lean_io_result_is_error(res)) return res;
lean_dec_ref(res);
res = initialize_bmc_BMC_ToyReport(builtin);
if (lean_io_result_is_error(res)) return res;
lean_dec_ref(res);
return lean_io_result_mk_ok(lean_box(0));
}
#ifdef __cplusplus
}
#endif
