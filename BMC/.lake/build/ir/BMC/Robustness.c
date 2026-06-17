// Lean compiler output
// Module: BMC.Robustness
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
lean_object* lean_nat_to_int(lean_object*);
lean_object* l_Bool_repr___redArg(uint8_t);
lean_object* lp_bmc_instReprCheckStatus_repr(uint8_t, lean_object*);
lean_object* lean_string_length(lean_object*);
uint8_t lp_bmc_checkPassed(uint8_t);
uint8_t lp_bmc_checkDeferred(uint8_t);
static const lean_string_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__0_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 3, .m_capacity = 3, .m_length = 2, .m_data = "{ "};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__0 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__0_value;
static const lean_string_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__1_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 16, .m_capacity = 16, .m_length = 15, .m_data = "toyAnalysisOnly"};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__1 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__1_value;
static const lean_ctor_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__2_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__1_value)}};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__2 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__2_value;
static const lean_ctor_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__3_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*2 + 0, .m_other = 2, .m_tag = 5}, .m_objs = {((lean_object*)(((size_t)(0) << 1) | 1)),((lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__2_value)}};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__3 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__3_value;
static const lean_string_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__4_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 5, .m_capacity = 5, .m_length = 4, .m_data = " := "};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__4 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__4_value;
static const lean_ctor_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__5_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__4_value)}};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__5 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__5_value;
static const lean_ctor_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__6_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*2 + 0, .m_other = 2, .m_tag = 5}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__3_value),((lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__5_value)}};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__6 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__6_value;
static lean_once_cell_t lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__7_once = LEAN_ONCE_CELL_INITIALIZER;
static lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__7;
static const lean_string_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__8_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 2, .m_capacity = 2, .m_length = 1, .m_data = ","};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__8 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__8_value;
static const lean_ctor_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__9_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__8_value)}};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__9 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__9_value;
static const lean_string_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__10_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 16, .m_capacity = 16, .m_length = 15, .m_data = "finalTruthClaim"};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__10 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__10_value;
static const lean_ctor_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__11_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__10_value)}};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__11 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__11_value;
static const lean_string_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__12_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 20, .m_capacity = 20, .m_length = 19, .m_data = "technicalGateStatus"};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__12 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__12_value;
static const lean_ctor_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__13_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__12_value)}};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__13 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__13_value;
static lean_once_cell_t lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__14_once = LEAN_ONCE_CELL_INITIALIZER;
static lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__14;
static const lean_string_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__15_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 20, .m_capacity = 20, .m_length = 19, .m_data = "promotionGateStatus"};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__15 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__15_value;
static const lean_ctor_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__16_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__15_value)}};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__16 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__16_value;
static const lean_string_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__17_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 18, .m_capacity = 18, .m_length = 17, .m_data = "friedmannResidual"};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__17 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__17_value;
static const lean_ctor_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__18_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__17_value)}};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__18 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__18_value;
static lean_once_cell_t lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__19_once = LEAN_ONCE_CELL_INITIALIZER;
static lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__19;
static const lean_string_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__20_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 13, .m_capacity = 13, .m_length = 12, .m_data = "faithfulness"};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__20 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__20_value;
static const lean_ctor_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__21_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__20_value)}};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__21 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__21_value;
static lean_once_cell_t lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__22_once = LEAN_ONCE_CELL_INITIALIZER;
static lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__22;
static const lean_string_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__23_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 3, .m_capacity = 3, .m_length = 2, .m_data = " }"};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__23 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__23_value;
static lean_once_cell_t lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__24_once = LEAN_ONCE_CELL_INITIALIZER;
static lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__24;
static lean_once_cell_t lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__25_once = LEAN_ONCE_CELL_INITIALIZER;
static lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__25;
static const lean_ctor_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__26_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__0_value)}};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__26 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__26_value;
static const lean_ctor_object lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__27_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__23_value)}};
static const lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__27 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__27_value;
LEAN_EXPORT lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___boxed(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_instReprBMCRobustnessReport_repr(lean_object*, lean_object*);
LEAN_EXPORT lean_object* lp_bmc_instReprBMCRobustnessReport_repr___boxed(lean_object*, lean_object*);
static const lean_closure_object lp_bmc_instReprBMCRobustnessReport___closed__0_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_closure_object) + sizeof(void*)*0, .m_other = 0, .m_tag = 245}, .m_fun = (void*)lp_bmc_instReprBMCRobustnessReport_repr___boxed, .m_arity = 2, .m_num_fixed = 0, .m_objs = {} };
static const lean_object* lp_bmc_instReprBMCRobustnessReport___closed__0 = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport___closed__0_value;
LEAN_EXPORT const lean_object* lp_bmc_instReprBMCRobustnessReport = (const lean_object*)&lp_bmc_instReprBMCRobustnessReport___closed__0_value;
LEAN_EXPORT uint8_t lp_bmc_reportPassesBMC0ARobustnessAuditGate(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_reportPassesBMC0ARobustnessAuditGate___boxed(lean_object*);
LEAN_EXPORT uint8_t lp_bmc_reportPassesFullBMCForRobustness(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_reportPassesFullBMCForRobustness___boxed(lean_object*);
LEAN_EXPORT lean_object* lp_bmc___private_BMC_Robustness_0__checkDeferred_match__1_splitter___redArg(uint8_t, lean_object*, lean_object*);
LEAN_EXPORT lean_object* lp_bmc___private_BMC_Robustness_0__checkDeferred_match__1_splitter___redArg___boxed(lean_object*, lean_object*, lean_object*);
LEAN_EXPORT lean_object* lp_bmc___private_BMC_Robustness_0__checkDeferred_match__1_splitter(lean_object*, uint8_t, lean_object*, lean_object*);
LEAN_EXPORT lean_object* lp_bmc___private_BMC_Robustness_0__checkDeferred_match__1_splitter___boxed(lean_object*, lean_object*, lean_object*, lean_object*);
LEAN_EXPORT lean_object* lp_bmc___private_BMC_Robustness_0__checkPassed_match__1_splitter___redArg(uint8_t, lean_object*, lean_object*);
LEAN_EXPORT lean_object* lp_bmc___private_BMC_Robustness_0__checkPassed_match__1_splitter___redArg___boxed(lean_object*, lean_object*, lean_object*);
LEAN_EXPORT lean_object* lp_bmc___private_BMC_Robustness_0__checkPassed_match__1_splitter(lean_object*, uint8_t, lean_object*, lean_object*);
LEAN_EXPORT lean_object* lp_bmc___private_BMC_Robustness_0__checkPassed_match__1_splitter___boxed(lean_object*, lean_object*, lean_object*, lean_object*);
static const lean_ctor_object lp_bmc_sprint3RobustnessWitness___closed__0_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*0 + 8, .m_other = 0, .m_tag = 0}, .m_objs = {LEAN_SCALAR_PTR_LITERAL(1, 0, 0, 3, 2, 3, 0, 0)}};
static const lean_object* lp_bmc_sprint3RobustnessWitness___closed__0 = (const lean_object*)&lp_bmc_sprint3RobustnessWitness___closed__0_value;
LEAN_EXPORT const lean_object* lp_bmc_sprint3RobustnessWitness = (const lean_object*)&lp_bmc_sprint3RobustnessWitness___closed__0_value;
static lean_object* _init_lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__7(void){
_start:
{
lean_object* v___x_14_; lean_object* v___x_15_; 
v___x_14_ = lean_unsigned_to_nat(19u);
v___x_15_ = lean_nat_to_int(v___x_14_);
return v___x_15_;
}
}
static lean_object* _init_lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__14(void){
_start:
{
lean_object* v___x_25_; lean_object* v___x_26_; 
v___x_25_ = lean_unsigned_to_nat(23u);
v___x_26_ = lean_nat_to_int(v___x_25_);
return v___x_26_;
}
}
static lean_object* _init_lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__19(void){
_start:
{
lean_object* v___x_33_; lean_object* v___x_34_; 
v___x_33_ = lean_unsigned_to_nat(21u);
v___x_34_ = lean_nat_to_int(v___x_33_);
return v___x_34_;
}
}
static lean_object* _init_lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__22(void){
_start:
{
lean_object* v___x_38_; lean_object* v___x_39_; 
v___x_38_ = lean_unsigned_to_nat(16u);
v___x_39_ = lean_nat_to_int(v___x_38_);
return v___x_39_;
}
}
static lean_object* _init_lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__24(void){
_start:
{
lean_object* v___x_41_; lean_object* v___x_42_; 
v___x_41_ = ((lean_object*)(lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__0));
v___x_42_ = lean_string_length(v___x_41_);
return v___x_42_;
}
}
static lean_object* _init_lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__25(void){
_start:
{
lean_object* v___x_43_; lean_object* v___x_44_; 
v___x_43_ = lean_obj_once(&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__24, &lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__24_once, _init_lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__24);
v___x_44_ = lean_nat_to_int(v___x_43_);
return v___x_44_;
}
}
LEAN_EXPORT lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg(lean_object* v_x_49_){
_start:
{
uint8_t v_toyAnalysisOnly_50_; uint8_t v_finalTruthClaim_51_; uint8_t v_technicalGateStatus_52_; uint8_t v_promotionGateStatus_53_; uint8_t v_friedmannResidual_54_; uint8_t v_faithfulness_55_; lean_object* v___x_56_; lean_object* v___x_57_; lean_object* v___x_58_; lean_object* v___x_59_; lean_object* v___x_60_; lean_object* v___x_61_; uint8_t v___x_62_; lean_object* v___x_63_; lean_object* v___x_64_; lean_object* v___x_65_; lean_object* v___x_66_; lean_object* v___x_67_; lean_object* v___x_68_; lean_object* v___x_69_; lean_object* v___x_70_; lean_object* v___x_71_; lean_object* v___x_72_; lean_object* v___x_73_; lean_object* v___x_74_; lean_object* v___x_75_; lean_object* v___x_76_; lean_object* v___x_77_; lean_object* v___x_78_; lean_object* v___x_79_; lean_object* v___x_80_; lean_object* v___x_81_; lean_object* v___x_82_; lean_object* v___x_83_; lean_object* v___x_84_; lean_object* v___x_85_; lean_object* v___x_86_; lean_object* v___x_87_; lean_object* v___x_88_; lean_object* v___x_89_; lean_object* v___x_90_; lean_object* v___x_91_; lean_object* v___x_92_; lean_object* v___x_93_; lean_object* v___x_94_; lean_object* v___x_95_; lean_object* v___x_96_; lean_object* v___x_97_; lean_object* v___x_98_; lean_object* v___x_99_; lean_object* v___x_100_; lean_object* v___x_101_; lean_object* v___x_102_; lean_object* v___x_103_; lean_object* v___x_104_; lean_object* v___x_105_; lean_object* v___x_106_; lean_object* v___x_107_; lean_object* v___x_108_; lean_object* v___x_109_; lean_object* v___x_110_; lean_object* v___x_111_; lean_object* v___x_112_; lean_object* v___x_113_; lean_object* v___x_114_; lean_object* v___x_115_; lean_object* v___x_116_; lean_object* v___x_117_; lean_object* v___x_118_; lean_object* v___x_119_; lean_object* v___x_120_; lean_object* v___x_121_; 
v_toyAnalysisOnly_50_ = lean_ctor_get_uint8(v_x_49_, 0);
v_finalTruthClaim_51_ = lean_ctor_get_uint8(v_x_49_, 1);
v_technicalGateStatus_52_ = lean_ctor_get_uint8(v_x_49_, 2);
v_promotionGateStatus_53_ = lean_ctor_get_uint8(v_x_49_, 3);
v_friedmannResidual_54_ = lean_ctor_get_uint8(v_x_49_, 4);
v_faithfulness_55_ = lean_ctor_get_uint8(v_x_49_, 5);
v___x_56_ = ((lean_object*)(lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__5));
v___x_57_ = ((lean_object*)(lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__6));
v___x_58_ = lean_obj_once(&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__7, &lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__7_once, _init_lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__7);
v___x_59_ = lean_unsigned_to_nat(0u);
v___x_60_ = l_Bool_repr___redArg(v_toyAnalysisOnly_50_);
v___x_61_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_61_, 0, v___x_58_);
lean_ctor_set(v___x_61_, 1, v___x_60_);
v___x_62_ = 0;
v___x_63_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_63_, 0, v___x_61_);
lean_ctor_set_uint8(v___x_63_, sizeof(void*)*1, v___x_62_);
v___x_64_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_64_, 0, v___x_57_);
lean_ctor_set(v___x_64_, 1, v___x_63_);
v___x_65_ = ((lean_object*)(lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__9));
v___x_66_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_66_, 0, v___x_64_);
lean_ctor_set(v___x_66_, 1, v___x_65_);
v___x_67_ = lean_box(1);
v___x_68_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_68_, 0, v___x_66_);
lean_ctor_set(v___x_68_, 1, v___x_67_);
v___x_69_ = ((lean_object*)(lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__11));
v___x_70_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_70_, 0, v___x_68_);
lean_ctor_set(v___x_70_, 1, v___x_69_);
v___x_71_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_71_, 0, v___x_70_);
lean_ctor_set(v___x_71_, 1, v___x_56_);
v___x_72_ = l_Bool_repr___redArg(v_finalTruthClaim_51_);
v___x_73_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_73_, 0, v___x_58_);
lean_ctor_set(v___x_73_, 1, v___x_72_);
v___x_74_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_74_, 0, v___x_73_);
lean_ctor_set_uint8(v___x_74_, sizeof(void*)*1, v___x_62_);
v___x_75_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_75_, 0, v___x_71_);
lean_ctor_set(v___x_75_, 1, v___x_74_);
v___x_76_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_76_, 0, v___x_75_);
lean_ctor_set(v___x_76_, 1, v___x_65_);
v___x_77_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_77_, 0, v___x_76_);
lean_ctor_set(v___x_77_, 1, v___x_67_);
v___x_78_ = ((lean_object*)(lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__13));
v___x_79_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_79_, 0, v___x_77_);
lean_ctor_set(v___x_79_, 1, v___x_78_);
v___x_80_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_80_, 0, v___x_79_);
lean_ctor_set(v___x_80_, 1, v___x_56_);
v___x_81_ = lean_obj_once(&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__14, &lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__14_once, _init_lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__14);
v___x_82_ = lp_bmc_instReprCheckStatus_repr(v_technicalGateStatus_52_, v___x_59_);
v___x_83_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_83_, 0, v___x_81_);
lean_ctor_set(v___x_83_, 1, v___x_82_);
v___x_84_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_84_, 0, v___x_83_);
lean_ctor_set_uint8(v___x_84_, sizeof(void*)*1, v___x_62_);
v___x_85_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_85_, 0, v___x_80_);
lean_ctor_set(v___x_85_, 1, v___x_84_);
v___x_86_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_86_, 0, v___x_85_);
lean_ctor_set(v___x_86_, 1, v___x_65_);
v___x_87_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_87_, 0, v___x_86_);
lean_ctor_set(v___x_87_, 1, v___x_67_);
v___x_88_ = ((lean_object*)(lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__16));
v___x_89_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_89_, 0, v___x_87_);
lean_ctor_set(v___x_89_, 1, v___x_88_);
v___x_90_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_90_, 0, v___x_89_);
lean_ctor_set(v___x_90_, 1, v___x_56_);
v___x_91_ = lp_bmc_instReprCheckStatus_repr(v_promotionGateStatus_53_, v___x_59_);
v___x_92_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_92_, 0, v___x_81_);
lean_ctor_set(v___x_92_, 1, v___x_91_);
v___x_93_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_93_, 0, v___x_92_);
lean_ctor_set_uint8(v___x_93_, sizeof(void*)*1, v___x_62_);
v___x_94_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_94_, 0, v___x_90_);
lean_ctor_set(v___x_94_, 1, v___x_93_);
v___x_95_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_95_, 0, v___x_94_);
lean_ctor_set(v___x_95_, 1, v___x_65_);
v___x_96_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_96_, 0, v___x_95_);
lean_ctor_set(v___x_96_, 1, v___x_67_);
v___x_97_ = ((lean_object*)(lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__18));
v___x_98_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_98_, 0, v___x_96_);
lean_ctor_set(v___x_98_, 1, v___x_97_);
v___x_99_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_99_, 0, v___x_98_);
lean_ctor_set(v___x_99_, 1, v___x_56_);
v___x_100_ = lean_obj_once(&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__19, &lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__19_once, _init_lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__19);
v___x_101_ = lp_bmc_instReprCheckStatus_repr(v_friedmannResidual_54_, v___x_59_);
v___x_102_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_102_, 0, v___x_100_);
lean_ctor_set(v___x_102_, 1, v___x_101_);
v___x_103_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_103_, 0, v___x_102_);
lean_ctor_set_uint8(v___x_103_, sizeof(void*)*1, v___x_62_);
v___x_104_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_104_, 0, v___x_99_);
lean_ctor_set(v___x_104_, 1, v___x_103_);
v___x_105_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_105_, 0, v___x_104_);
lean_ctor_set(v___x_105_, 1, v___x_65_);
v___x_106_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_106_, 0, v___x_105_);
lean_ctor_set(v___x_106_, 1, v___x_67_);
v___x_107_ = ((lean_object*)(lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__21));
v___x_108_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_108_, 0, v___x_106_);
lean_ctor_set(v___x_108_, 1, v___x_107_);
v___x_109_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_109_, 0, v___x_108_);
lean_ctor_set(v___x_109_, 1, v___x_56_);
v___x_110_ = lean_obj_once(&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__22, &lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__22_once, _init_lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__22);
v___x_111_ = lp_bmc_instReprCheckStatus_repr(v_faithfulness_55_, v___x_59_);
v___x_112_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_112_, 0, v___x_110_);
lean_ctor_set(v___x_112_, 1, v___x_111_);
v___x_113_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_113_, 0, v___x_112_);
lean_ctor_set_uint8(v___x_113_, sizeof(void*)*1, v___x_62_);
v___x_114_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_114_, 0, v___x_109_);
lean_ctor_set(v___x_114_, 1, v___x_113_);
v___x_115_ = lean_obj_once(&lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__25, &lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__25_once, _init_lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__25);
v___x_116_ = ((lean_object*)(lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__26));
v___x_117_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_117_, 0, v___x_116_);
lean_ctor_set(v___x_117_, 1, v___x_114_);
v___x_118_ = ((lean_object*)(lp_bmc_instReprBMCRobustnessReport_repr___redArg___closed__27));
v___x_119_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_119_, 0, v___x_117_);
lean_ctor_set(v___x_119_, 1, v___x_118_);
v___x_120_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_120_, 0, v___x_115_);
lean_ctor_set(v___x_120_, 1, v___x_119_);
v___x_121_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_121_, 0, v___x_120_);
lean_ctor_set_uint8(v___x_121_, sizeof(void*)*1, v___x_62_);
return v___x_121_;
}
}
LEAN_EXPORT lean_object* lp_bmc_instReprBMCRobustnessReport_repr___redArg___boxed(lean_object* v_x_122_){
_start:
{
lean_object* v_res_123_; 
v_res_123_ = lp_bmc_instReprBMCRobustnessReport_repr___redArg(v_x_122_);
lean_dec_ref(v_x_122_);
return v_res_123_;
}
}
LEAN_EXPORT lean_object* lp_bmc_instReprBMCRobustnessReport_repr(lean_object* v_x_124_, lean_object* v_prec_125_){
_start:
{
lean_object* v___x_126_; 
v___x_126_ = lp_bmc_instReprBMCRobustnessReport_repr___redArg(v_x_124_);
return v___x_126_;
}
}
LEAN_EXPORT lean_object* lp_bmc_instReprBMCRobustnessReport_repr___boxed(lean_object* v_x_127_, lean_object* v_prec_128_){
_start:
{
lean_object* v_res_129_; 
v_res_129_ = lp_bmc_instReprBMCRobustnessReport_repr(v_x_127_, v_prec_128_);
lean_dec(v_prec_128_);
lean_dec_ref(v_x_127_);
return v_res_129_;
}
}
LEAN_EXPORT uint8_t lp_bmc_reportPassesBMC0ARobustnessAuditGate(lean_object* v_r_132_){
_start:
{
uint8_t v_toyAnalysisOnly_133_; 
v_toyAnalysisOnly_133_ = lean_ctor_get_uint8(v_r_132_, 0);
if (v_toyAnalysisOnly_133_ == 0)
{
return v_toyAnalysisOnly_133_;
}
else
{
uint8_t v_finalTruthClaim_134_; 
v_finalTruthClaim_134_ = lean_ctor_get_uint8(v_r_132_, 1);
if (v_finalTruthClaim_134_ == 0)
{
uint8_t v_technicalGateStatus_135_; uint8_t v_friedmannResidual_136_; uint8_t v___x_137_; 
v_technicalGateStatus_135_ = lean_ctor_get_uint8(v_r_132_, 2);
v_friedmannResidual_136_ = lean_ctor_get_uint8(v_r_132_, 4);
v___x_137_ = lp_bmc_checkPassed(v_technicalGateStatus_135_);
if (v___x_137_ == 0)
{
return v___x_137_;
}
else
{
uint8_t v___x_138_; 
v___x_138_ = lp_bmc_checkDeferred(v_friedmannResidual_136_);
return v___x_138_;
}
}
else
{
uint8_t v___x_139_; 
v___x_139_ = 0;
return v___x_139_;
}
}
}
}
LEAN_EXPORT lean_object* lp_bmc_reportPassesBMC0ARobustnessAuditGate___boxed(lean_object* v_r_140_){
_start:
{
uint8_t v_res_141_; lean_object* v_r_142_; 
v_res_141_ = lp_bmc_reportPassesBMC0ARobustnessAuditGate(v_r_140_);
lean_dec_ref(v_r_140_);
v_r_142_ = lean_box(v_res_141_);
return v_r_142_;
}
}
LEAN_EXPORT uint8_t lp_bmc_reportPassesFullBMCForRobustness(lean_object* v_r_143_){
_start:
{
uint8_t v_toyAnalysisOnly_144_; 
v_toyAnalysisOnly_144_ = lean_ctor_get_uint8(v_r_143_, 0);
if (v_toyAnalysisOnly_144_ == 0)
{
return v_toyAnalysisOnly_144_;
}
else
{
uint8_t v_finalTruthClaim_145_; 
v_finalTruthClaim_145_ = lean_ctor_get_uint8(v_r_143_, 1);
if (v_finalTruthClaim_145_ == 0)
{
uint8_t v_technicalGateStatus_146_; uint8_t v_friedmannResidual_147_; uint8_t v_faithfulness_148_; uint8_t v___x_149_; 
v_technicalGateStatus_146_ = lean_ctor_get_uint8(v_r_143_, 2);
v_friedmannResidual_147_ = lean_ctor_get_uint8(v_r_143_, 4);
v_faithfulness_148_ = lean_ctor_get_uint8(v_r_143_, 5);
v___x_149_ = lp_bmc_checkPassed(v_technicalGateStatus_146_);
if (v___x_149_ == 0)
{
return v___x_149_;
}
else
{
uint8_t v___x_150_; 
v___x_150_ = lp_bmc_checkPassed(v_friedmannResidual_147_);
if (v___x_150_ == 0)
{
return v___x_150_;
}
else
{
uint8_t v___x_151_; 
v___x_151_ = lp_bmc_checkPassed(v_faithfulness_148_);
return v___x_151_;
}
}
}
else
{
uint8_t v___x_152_; 
v___x_152_ = 0;
return v___x_152_;
}
}
}
}
LEAN_EXPORT lean_object* lp_bmc_reportPassesFullBMCForRobustness___boxed(lean_object* v_r_153_){
_start:
{
uint8_t v_res_154_; lean_object* v_r_155_; 
v_res_154_ = lp_bmc_reportPassesFullBMCForRobustness(v_r_153_);
lean_dec_ref(v_r_153_);
v_r_155_ = lean_box(v_res_154_);
return v_r_155_;
}
}
LEAN_EXPORT lean_object* lp_bmc___private_BMC_Robustness_0__checkDeferred_match__1_splitter___redArg(uint8_t v_x_156_, lean_object* v_h__1_157_, lean_object* v_h__2_158_){
_start:
{
if (v_x_156_ == 2)
{
lean_object* v___x_159_; lean_object* v___x_160_; 
lean_dec(v_h__2_158_);
v___x_159_ = lean_box(0);
v___x_160_ = lean_apply_1(v_h__1_157_, v___x_159_);
return v___x_160_;
}
else
{
lean_object* v___x_161_; lean_object* v___x_162_; 
lean_dec(v_h__1_157_);
v___x_161_ = lean_box(v_x_156_);
v___x_162_ = lean_apply_2(v_h__2_158_, v___x_161_, lean_box(0));
return v___x_162_;
}
}
}
LEAN_EXPORT lean_object* lp_bmc___private_BMC_Robustness_0__checkDeferred_match__1_splitter___redArg___boxed(lean_object* v_x_163_, lean_object* v_h__1_164_, lean_object* v_h__2_165_){
_start:
{
uint8_t v_x_17__boxed_166_; lean_object* v_res_167_; 
v_x_17__boxed_166_ = lean_unbox(v_x_163_);
v_res_167_ = lp_bmc___private_BMC_Robustness_0__checkDeferred_match__1_splitter___redArg(v_x_17__boxed_166_, v_h__1_164_, v_h__2_165_);
return v_res_167_;
}
}
LEAN_EXPORT lean_object* lp_bmc___private_BMC_Robustness_0__checkDeferred_match__1_splitter(lean_object* v_motive_168_, uint8_t v_x_169_, lean_object* v_h__1_170_, lean_object* v_h__2_171_){
_start:
{
if (v_x_169_ == 2)
{
lean_object* v___x_172_; lean_object* v___x_173_; 
lean_dec(v_h__2_171_);
v___x_172_ = lean_box(0);
v___x_173_ = lean_apply_1(v_h__1_170_, v___x_172_);
return v___x_173_;
}
else
{
lean_object* v___x_174_; lean_object* v___x_175_; 
lean_dec(v_h__1_170_);
v___x_174_ = lean_box(v_x_169_);
v___x_175_ = lean_apply_2(v_h__2_171_, v___x_174_, lean_box(0));
return v___x_175_;
}
}
}
LEAN_EXPORT lean_object* lp_bmc___private_BMC_Robustness_0__checkDeferred_match__1_splitter___boxed(lean_object* v_motive_176_, lean_object* v_x_177_, lean_object* v_h__1_178_, lean_object* v_h__2_179_){
_start:
{
uint8_t v_x_28__boxed_180_; lean_object* v_res_181_; 
v_x_28__boxed_180_ = lean_unbox(v_x_177_);
v_res_181_ = lp_bmc___private_BMC_Robustness_0__checkDeferred_match__1_splitter(v_motive_176_, v_x_28__boxed_180_, v_h__1_178_, v_h__2_179_);
return v_res_181_;
}
}
LEAN_EXPORT lean_object* lp_bmc___private_BMC_Robustness_0__checkPassed_match__1_splitter___redArg(uint8_t v_x_182_, lean_object* v_h__1_183_, lean_object* v_h__2_184_){
_start:
{
if (v_x_182_ == 0)
{
lean_object* v___x_185_; lean_object* v___x_186_; 
lean_dec(v_h__2_184_);
v___x_185_ = lean_box(0);
v___x_186_ = lean_apply_1(v_h__1_183_, v___x_185_);
return v___x_186_;
}
else
{
lean_object* v___x_187_; lean_object* v___x_188_; 
lean_dec(v_h__1_183_);
v___x_187_ = lean_box(v_x_182_);
v___x_188_ = lean_apply_2(v_h__2_184_, v___x_187_, lean_box(0));
return v___x_188_;
}
}
}
LEAN_EXPORT lean_object* lp_bmc___private_BMC_Robustness_0__checkPassed_match__1_splitter___redArg___boxed(lean_object* v_x_189_, lean_object* v_h__1_190_, lean_object* v_h__2_191_){
_start:
{
uint8_t v_x_17__boxed_192_; lean_object* v_res_193_; 
v_x_17__boxed_192_ = lean_unbox(v_x_189_);
v_res_193_ = lp_bmc___private_BMC_Robustness_0__checkPassed_match__1_splitter___redArg(v_x_17__boxed_192_, v_h__1_190_, v_h__2_191_);
return v_res_193_;
}
}
LEAN_EXPORT lean_object* lp_bmc___private_BMC_Robustness_0__checkPassed_match__1_splitter(lean_object* v_motive_194_, uint8_t v_x_195_, lean_object* v_h__1_196_, lean_object* v_h__2_197_){
_start:
{
if (v_x_195_ == 0)
{
lean_object* v___x_198_; lean_object* v___x_199_; 
lean_dec(v_h__2_197_);
v___x_198_ = lean_box(0);
v___x_199_ = lean_apply_1(v_h__1_196_, v___x_198_);
return v___x_199_;
}
else
{
lean_object* v___x_200_; lean_object* v___x_201_; 
lean_dec(v_h__1_196_);
v___x_200_ = lean_box(v_x_195_);
v___x_201_ = lean_apply_2(v_h__2_197_, v___x_200_, lean_box(0));
return v___x_201_;
}
}
}
LEAN_EXPORT lean_object* lp_bmc___private_BMC_Robustness_0__checkPassed_match__1_splitter___boxed(lean_object* v_motive_202_, lean_object* v_x_203_, lean_object* v_h__1_204_, lean_object* v_h__2_205_){
_start:
{
uint8_t v_x_28__boxed_206_; lean_object* v_res_207_; 
v_x_28__boxed_206_ = lean_unbox(v_x_203_);
v_res_207_ = lp_bmc___private_BMC_Robustness_0__checkPassed_match__1_splitter(v_motive_202_, v_x_28__boxed_206_, v_h__1_204_, v_h__2_205_);
return v_res_207_;
}
}
lean_object* initialize_Init(uint8_t builtin);
lean_object* initialize_Init(uint8_t builtin);
lean_object* initialize_bmc_BMC_ToyReport(uint8_t builtin);
static bool _G_initialized = false;
LEAN_EXPORT lean_object* initialize_bmc_BMC_Robustness(uint8_t builtin) {
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
