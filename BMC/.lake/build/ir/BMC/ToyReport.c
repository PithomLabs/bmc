// Lean compiler output
// Module: BMC.ToyReport
// Imports: public import Init public meta import Init
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
lean_object* l_Repr_addAppParen(lean_object*, lean_object*);
uint8_t lean_nat_dec_le(lean_object*, lean_object*);
lean_object* lean_nat_to_int(lean_object*);
lean_object* l_Bool_repr___redArg(uint8_t);
lean_object* lean_string_length(lean_object*);
uint8_t lean_nat_dec_le(lean_object*, lean_object*);
uint8_t lean_nat_dec_eq(lean_object*, lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_ctorIdx(uint8_t);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_ctorIdx___boxed(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_toCtorIdx(uint8_t);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_toCtorIdx___boxed(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_ctorElim___redArg(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_ctorElim___redArg___boxed(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_ctorElim(lean_object*, lean_object*, uint8_t, lean_object*, lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_ctorElim___boxed(lean_object*, lean_object*, lean_object*, lean_object*, lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_pass_elim___redArg(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_pass_elim___redArg___boxed(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_pass_elim(lean_object*, uint8_t, lean_object*, lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_pass_elim___boxed(lean_object*, lean_object*, lean_object*, lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_fail_elim___redArg(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_fail_elim___redArg___boxed(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_fail_elim(lean_object*, uint8_t, lean_object*, lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_fail_elim___boxed(lean_object*, lean_object*, lean_object*, lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_deferred_elim___redArg(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_deferred_elim___redArg___boxed(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_deferred_elim(lean_object*, uint8_t, lean_object*, lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_deferred_elim___boxed(lean_object*, lean_object*, lean_object*, lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_contested_elim___redArg(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_contested_elim___redArg___boxed(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_contested_elim(lean_object*, uint8_t, lean_object*, lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_contested_elim___boxed(lean_object*, lean_object*, lean_object*, lean_object*);
LEAN_EXPORT uint8_t lp_bmc_CheckStatus_ofNat(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_ofNat___boxed(lean_object*);
LEAN_EXPORT uint8_t lp_bmc_instDecidableEqCheckStatus(uint8_t, uint8_t);
LEAN_EXPORT lean_object* lp_bmc_instDecidableEqCheckStatus___boxed(lean_object*, lean_object*);
static const lean_string_object lp_bmc_instReprCheckStatus_repr___closed__0_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 17, .m_capacity = 17, .m_length = 16, .m_data = "CheckStatus.pass"};
static const lean_object* lp_bmc_instReprCheckStatus_repr___closed__0 = (const lean_object*)&lp_bmc_instReprCheckStatus_repr___closed__0_value;
static const lean_ctor_object lp_bmc_instReprCheckStatus_repr___closed__1_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprCheckStatus_repr___closed__0_value)}};
static const lean_object* lp_bmc_instReprCheckStatus_repr___closed__1 = (const lean_object*)&lp_bmc_instReprCheckStatus_repr___closed__1_value;
static const lean_string_object lp_bmc_instReprCheckStatus_repr___closed__2_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 17, .m_capacity = 17, .m_length = 16, .m_data = "CheckStatus.fail"};
static const lean_object* lp_bmc_instReprCheckStatus_repr___closed__2 = (const lean_object*)&lp_bmc_instReprCheckStatus_repr___closed__2_value;
static const lean_ctor_object lp_bmc_instReprCheckStatus_repr___closed__3_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprCheckStatus_repr___closed__2_value)}};
static const lean_object* lp_bmc_instReprCheckStatus_repr___closed__3 = (const lean_object*)&lp_bmc_instReprCheckStatus_repr___closed__3_value;
static const lean_string_object lp_bmc_instReprCheckStatus_repr___closed__4_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 21, .m_capacity = 21, .m_length = 20, .m_data = "CheckStatus.deferred"};
static const lean_object* lp_bmc_instReprCheckStatus_repr___closed__4 = (const lean_object*)&lp_bmc_instReprCheckStatus_repr___closed__4_value;
static const lean_ctor_object lp_bmc_instReprCheckStatus_repr___closed__5_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprCheckStatus_repr___closed__4_value)}};
static const lean_object* lp_bmc_instReprCheckStatus_repr___closed__5 = (const lean_object*)&lp_bmc_instReprCheckStatus_repr___closed__5_value;
static const lean_string_object lp_bmc_instReprCheckStatus_repr___closed__6_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 22, .m_capacity = 22, .m_length = 21, .m_data = "CheckStatus.contested"};
static const lean_object* lp_bmc_instReprCheckStatus_repr___closed__6 = (const lean_object*)&lp_bmc_instReprCheckStatus_repr___closed__6_value;
static const lean_ctor_object lp_bmc_instReprCheckStatus_repr___closed__7_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprCheckStatus_repr___closed__6_value)}};
static const lean_object* lp_bmc_instReprCheckStatus_repr___closed__7 = (const lean_object*)&lp_bmc_instReprCheckStatus_repr___closed__7_value;
static lean_once_cell_t lp_bmc_instReprCheckStatus_repr___closed__8_once = LEAN_ONCE_CELL_INITIALIZER;
static lean_object* lp_bmc_instReprCheckStatus_repr___closed__8;
static lean_once_cell_t lp_bmc_instReprCheckStatus_repr___closed__9_once = LEAN_ONCE_CELL_INITIALIZER;
static lean_object* lp_bmc_instReprCheckStatus_repr___closed__9;
LEAN_EXPORT lean_object* lp_bmc_instReprCheckStatus_repr(uint8_t, lean_object*);
LEAN_EXPORT lean_object* lp_bmc_instReprCheckStatus_repr___boxed(lean_object*, lean_object*);
static const lean_closure_object lp_bmc_instReprCheckStatus___closed__0_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_closure_object) + sizeof(void*)*0, .m_other = 0, .m_tag = 245}, .m_fun = (void*)lp_bmc_instReprCheckStatus_repr___boxed, .m_arity = 2, .m_num_fixed = 0, .m_objs = {} };
static const lean_object* lp_bmc_instReprCheckStatus___closed__0 = (const lean_object*)&lp_bmc_instReprCheckStatus___closed__0_value;
LEAN_EXPORT const lean_object* lp_bmc_instReprCheckStatus = (const lean_object*)&lp_bmc_instReprCheckStatus___closed__0_value;
LEAN_EXPORT uint8_t lp_bmc_checkPassed(uint8_t);
LEAN_EXPORT lean_object* lp_bmc_checkPassed___boxed(lean_object*);
LEAN_EXPORT uint8_t lp_bmc_checkDeferred(uint8_t);
LEAN_EXPORT lean_object* lp_bmc_checkDeferred___boxed(lean_object*);
static const lean_string_object lp_bmc_instReprBMCReport_repr___redArg___closed__0_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 3, .m_capacity = 3, .m_length = 2, .m_data = "{ "};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__0 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__0_value;
static const lean_string_object lp_bmc_instReprBMCReport_repr___redArg___closed__1_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 16, .m_capacity = 16, .m_length = 15, .m_data = "toyAnalysisOnly"};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__1 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__1_value;
static const lean_ctor_object lp_bmc_instReprBMCReport_repr___redArg___closed__2_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__1_value)}};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__2 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__2_value;
static const lean_ctor_object lp_bmc_instReprBMCReport_repr___redArg___closed__3_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*2 + 0, .m_other = 2, .m_tag = 5}, .m_objs = {((lean_object*)(((size_t)(0) << 1) | 1)),((lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__2_value)}};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__3 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__3_value;
static const lean_string_object lp_bmc_instReprBMCReport_repr___redArg___closed__4_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 5, .m_capacity = 5, .m_length = 4, .m_data = " := "};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__4 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__4_value;
static const lean_ctor_object lp_bmc_instReprBMCReport_repr___redArg___closed__5_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__4_value)}};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__5 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__5_value;
static const lean_ctor_object lp_bmc_instReprBMCReport_repr___redArg___closed__6_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*2 + 0, .m_other = 2, .m_tag = 5}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__3_value),((lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__5_value)}};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__6 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__6_value;
static lean_once_cell_t lp_bmc_instReprBMCReport_repr___redArg___closed__7_once = LEAN_ONCE_CELL_INITIALIZER;
static lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__7;
static const lean_string_object lp_bmc_instReprBMCReport_repr___redArg___closed__8_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 2, .m_capacity = 2, .m_length = 1, .m_data = ","};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__8 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__8_value;
static const lean_ctor_object lp_bmc_instReprBMCReport_repr___redArg___closed__9_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__8_value)}};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__9 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__9_value;
static const lean_string_object lp_bmc_instReprBMCReport_repr___redArg___closed__10_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 16, .m_capacity = 16, .m_length = 15, .m_data = "finalTruthClaim"};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__10 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__10_value;
static const lean_ctor_object lp_bmc_instReprBMCReport_repr___redArg___closed__11_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__10_value)}};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__11 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__11_value;
static const lean_string_object lp_bmc_instReprBMCReport_repr___redArg___closed__12_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 12, .m_capacity = 12, .m_length = 11, .m_data = "wdwResidual"};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__12 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__12_value;
static const lean_ctor_object lp_bmc_instReprBMCReport_repr___redArg___closed__13_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__12_value)}};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__13 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__13_value;
static lean_once_cell_t lp_bmc_instReprBMCReport_repr___redArg___closed__14_once = LEAN_ONCE_CELL_INITIALIZER;
static lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__14;
static const lean_string_object lp_bmc_instReprBMCReport_repr___redArg___closed__15_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 17, .m_capacity = 17, .m_length = 16, .m_data = "trajectoryFinite"};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__15 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__15_value;
static const lean_ctor_object lp_bmc_instReprBMCReport_repr___redArg___closed__16_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__15_value)}};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__16 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__16_value;
static lean_once_cell_t lp_bmc_instReprBMCReport_repr___redArg___closed__17_once = LEAN_ONCE_CELL_INITIALIZER;
static lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__17;
static const lean_string_object lp_bmc_instReprBMCReport_repr___redArg___closed__18_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 15, .m_capacity = 15, .m_length = 14, .m_data = "clockMonotonic"};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__18 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__18_value;
static const lean_ctor_object lp_bmc_instReprBMCReport_repr___redArg___closed__19_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__18_value)}};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__19 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__19_value;
static lean_once_cell_t lp_bmc_instReprBMCReport_repr___redArg___closed__20_once = LEAN_ONCE_CELL_INITIALIZER;
static lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__20;
static const lean_string_object lp_bmc_instReprBMCReport_repr___redArg___closed__21_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 14, .m_capacity = 14, .m_length = 13, .m_data = "nodeDetection"};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__21 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__21_value;
static const lean_ctor_object lp_bmc_instReprBMCReport_repr___redArg___closed__22_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__21_value)}};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__22 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__22_value;
static lean_once_cell_t lp_bmc_instReprBMCReport_repr___redArg___closed__23_once = LEAN_ONCE_CELL_INITIALIZER;
static lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__23;
static const lean_string_object lp_bmc_instReprBMCReport_repr___redArg___closed__24_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 16, .m_capacity = 16, .m_length = 15, .m_data = "nodeContactFree"};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__24 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__24_value;
static const lean_ctor_object lp_bmc_instReprBMCReport_repr___redArg___closed__25_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__24_value)}};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__25 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__25_value;
static const lean_string_object lp_bmc_instReprBMCReport_repr___redArg___closed__26_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 21, .m_capacity = 21, .m_length = 20, .m_data = "qFiniteAwayFromNodes"};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__26 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__26_value;
static const lean_ctor_object lp_bmc_instReprBMCReport_repr___redArg___closed__27_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__26_value)}};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__27 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__27_value;
static lean_once_cell_t lp_bmc_instReprBMCReport_repr___redArg___closed__28_once = LEAN_ONCE_CELL_INITIALIZER;
static lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__28;
static const lean_string_object lp_bmc_instReprBMCReport_repr___redArg___closed__29_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 20, .m_capacity = 20, .m_length = 19, .m_data = "phaseGradientFinite"};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__29 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__29_value;
static const lean_ctor_object lp_bmc_instReprBMCReport_repr___redArg___closed__30_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__29_value)}};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__30 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__30_value;
static lean_once_cell_t lp_bmc_instReprBMCReport_repr___redArg___closed__31_once = LEAN_ONCE_CELL_INITIALIZER;
static lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__31;
static const lean_string_object lp_bmc_instReprBMCReport_repr___redArg___closed__32_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 15, .m_capacity = 15, .m_length = 14, .m_data = "classicalLimit"};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__32 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__32_value;
static const lean_ctor_object lp_bmc_instReprBMCReport_repr___redArg___closed__33_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__32_value)}};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__33 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__33_value;
static const lean_string_object lp_bmc_instReprBMCReport_repr___redArg___closed__34_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 18, .m_capacity = 18, .m_length = 17, .m_data = "friedmannResidual"};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__34 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__34_value;
static const lean_ctor_object lp_bmc_instReprBMCReport_repr___redArg___closed__35_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__34_value)}};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__35 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__35_value;
static lean_once_cell_t lp_bmc_instReprBMCReport_repr___redArg___closed__36_once = LEAN_ONCE_CELL_INITIALIZER;
static lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__36;
static const lean_string_object lp_bmc_instReprBMCReport_repr___redArg___closed__37_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 13, .m_capacity = 13, .m_length = 12, .m_data = "faithfulness"};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__37 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__37_value;
static const lean_ctor_object lp_bmc_instReprBMCReport_repr___redArg___closed__38_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__37_value)}};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__38 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__38_value;
static lean_once_cell_t lp_bmc_instReprBMCReport_repr___redArg___closed__39_once = LEAN_ONCE_CELL_INITIALIZER;
static lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__39;
static const lean_string_object lp_bmc_instReprBMCReport_repr___redArg___closed__40_value = {.m_header = {.m_rc = 0, .m_cs_sz = 0, .m_other = 0, .m_tag = 249}, .m_size = 3, .m_capacity = 3, .m_length = 2, .m_data = " }"};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__40 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__40_value;
static lean_once_cell_t lp_bmc_instReprBMCReport_repr___redArg___closed__41_once = LEAN_ONCE_CELL_INITIALIZER;
static lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__41;
static lean_once_cell_t lp_bmc_instReprBMCReport_repr___redArg___closed__42_once = LEAN_ONCE_CELL_INITIALIZER;
static lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__42;
static const lean_ctor_object lp_bmc_instReprBMCReport_repr___redArg___closed__43_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__0_value)}};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__43 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__43_value;
static const lean_ctor_object lp_bmc_instReprBMCReport_repr___redArg___closed__44_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_ctor_object) + sizeof(void*)*1 + 0, .m_other = 1, .m_tag = 3}, .m_objs = {((lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__40_value)}};
static const lean_object* lp_bmc_instReprBMCReport_repr___redArg___closed__44 = (const lean_object*)&lp_bmc_instReprBMCReport_repr___redArg___closed__44_value;
LEAN_EXPORT lean_object* lp_bmc_instReprBMCReport_repr___redArg(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_instReprBMCReport_repr___redArg___boxed(lean_object*);
LEAN_EXPORT lean_object* lp_bmc_instReprBMCReport_repr(lean_object*, lean_object*);
LEAN_EXPORT lean_object* lp_bmc_instReprBMCReport_repr___boxed(lean_object*, lean_object*);
static const lean_closure_object lp_bmc_instReprBMCReport___closed__0_value = {.m_header = {.m_rc = 0, .m_cs_sz = sizeof(lean_closure_object) + sizeof(void*)*0, .m_other = 0, .m_tag = 245}, .m_fun = (void*)lp_bmc_instReprBMCReport_repr___boxed, .m_arity = 2, .m_num_fixed = 0, .m_objs = {} };
static const lean_object* lp_bmc_instReprBMCReport___closed__0 = (const lean_object*)&lp_bmc_instReprBMCReport___closed__0_value;
LEAN_EXPORT const lean_object* lp_bmc_instReprBMCReport = (const lean_object*)&lp_bmc_instReprBMCReport___closed__0_value;
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_ctorIdx(uint8_t v_x_1_){
_start:
{
switch(v_x_1_)
{
case 0:
{
lean_object* v___x_2_; 
v___x_2_ = lean_unsigned_to_nat(0u);
return v___x_2_;
}
case 1:
{
lean_object* v___x_3_; 
v___x_3_ = lean_unsigned_to_nat(1u);
return v___x_3_;
}
case 2:
{
lean_object* v___x_4_; 
v___x_4_ = lean_unsigned_to_nat(2u);
return v___x_4_;
}
default: 
{
lean_object* v___x_5_; 
v___x_5_ = lean_unsigned_to_nat(3u);
return v___x_5_;
}
}
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_ctorIdx___boxed(lean_object* v_x_6_){
_start:
{
uint8_t v_x_boxed_7_; lean_object* v_res_8_; 
v_x_boxed_7_ = lean_unbox(v_x_6_);
v_res_8_ = lp_bmc_CheckStatus_ctorIdx(v_x_boxed_7_);
return v_res_8_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_toCtorIdx(uint8_t v_x_9_){
_start:
{
lean_object* v___x_10_; 
v___x_10_ = lp_bmc_CheckStatus_ctorIdx(v_x_9_);
return v___x_10_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_toCtorIdx___boxed(lean_object* v_x_11_){
_start:
{
uint8_t v_x_4__boxed_12_; lean_object* v_res_13_; 
v_x_4__boxed_12_ = lean_unbox(v_x_11_);
v_res_13_ = lp_bmc_CheckStatus_toCtorIdx(v_x_4__boxed_12_);
return v_res_13_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_ctorElim___redArg(lean_object* v_k_14_){
_start:
{
lean_inc(v_k_14_);
return v_k_14_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_ctorElim___redArg___boxed(lean_object* v_k_15_){
_start:
{
lean_object* v_res_16_; 
v_res_16_ = lp_bmc_CheckStatus_ctorElim___redArg(v_k_15_);
lean_dec(v_k_15_);
return v_res_16_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_ctorElim(lean_object* v_motive_17_, lean_object* v_ctorIdx_18_, uint8_t v_t_19_, lean_object* v_h_20_, lean_object* v_k_21_){
_start:
{
lean_inc(v_k_21_);
return v_k_21_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_ctorElim___boxed(lean_object* v_motive_22_, lean_object* v_ctorIdx_23_, lean_object* v_t_24_, lean_object* v_h_25_, lean_object* v_k_26_){
_start:
{
uint8_t v_t_boxed_27_; lean_object* v_res_28_; 
v_t_boxed_27_ = lean_unbox(v_t_24_);
v_res_28_ = lp_bmc_CheckStatus_ctorElim(v_motive_22_, v_ctorIdx_23_, v_t_boxed_27_, v_h_25_, v_k_26_);
lean_dec(v_k_26_);
lean_dec(v_ctorIdx_23_);
return v_res_28_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_pass_elim___redArg(lean_object* v_pass_29_){
_start:
{
lean_inc(v_pass_29_);
return v_pass_29_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_pass_elim___redArg___boxed(lean_object* v_pass_30_){
_start:
{
lean_object* v_res_31_; 
v_res_31_ = lp_bmc_CheckStatus_pass_elim___redArg(v_pass_30_);
lean_dec(v_pass_30_);
return v_res_31_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_pass_elim(lean_object* v_motive_32_, uint8_t v_t_33_, lean_object* v_h_34_, lean_object* v_pass_35_){
_start:
{
lean_inc(v_pass_35_);
return v_pass_35_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_pass_elim___boxed(lean_object* v_motive_36_, lean_object* v_t_37_, lean_object* v_h_38_, lean_object* v_pass_39_){
_start:
{
uint8_t v_t_boxed_40_; lean_object* v_res_41_; 
v_t_boxed_40_ = lean_unbox(v_t_37_);
v_res_41_ = lp_bmc_CheckStatus_pass_elim(v_motive_36_, v_t_boxed_40_, v_h_38_, v_pass_39_);
lean_dec(v_pass_39_);
return v_res_41_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_fail_elim___redArg(lean_object* v_fail_42_){
_start:
{
lean_inc(v_fail_42_);
return v_fail_42_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_fail_elim___redArg___boxed(lean_object* v_fail_43_){
_start:
{
lean_object* v_res_44_; 
v_res_44_ = lp_bmc_CheckStatus_fail_elim___redArg(v_fail_43_);
lean_dec(v_fail_43_);
return v_res_44_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_fail_elim(lean_object* v_motive_45_, uint8_t v_t_46_, lean_object* v_h_47_, lean_object* v_fail_48_){
_start:
{
lean_inc(v_fail_48_);
return v_fail_48_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_fail_elim___boxed(lean_object* v_motive_49_, lean_object* v_t_50_, lean_object* v_h_51_, lean_object* v_fail_52_){
_start:
{
uint8_t v_t_boxed_53_; lean_object* v_res_54_; 
v_t_boxed_53_ = lean_unbox(v_t_50_);
v_res_54_ = lp_bmc_CheckStatus_fail_elim(v_motive_49_, v_t_boxed_53_, v_h_51_, v_fail_52_);
lean_dec(v_fail_52_);
return v_res_54_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_deferred_elim___redArg(lean_object* v_deferred_55_){
_start:
{
lean_inc(v_deferred_55_);
return v_deferred_55_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_deferred_elim___redArg___boxed(lean_object* v_deferred_56_){
_start:
{
lean_object* v_res_57_; 
v_res_57_ = lp_bmc_CheckStatus_deferred_elim___redArg(v_deferred_56_);
lean_dec(v_deferred_56_);
return v_res_57_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_deferred_elim(lean_object* v_motive_58_, uint8_t v_t_59_, lean_object* v_h_60_, lean_object* v_deferred_61_){
_start:
{
lean_inc(v_deferred_61_);
return v_deferred_61_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_deferred_elim___boxed(lean_object* v_motive_62_, lean_object* v_t_63_, lean_object* v_h_64_, lean_object* v_deferred_65_){
_start:
{
uint8_t v_t_boxed_66_; lean_object* v_res_67_; 
v_t_boxed_66_ = lean_unbox(v_t_63_);
v_res_67_ = lp_bmc_CheckStatus_deferred_elim(v_motive_62_, v_t_boxed_66_, v_h_64_, v_deferred_65_);
lean_dec(v_deferred_65_);
return v_res_67_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_contested_elim___redArg(lean_object* v_contested_68_){
_start:
{
lean_inc(v_contested_68_);
return v_contested_68_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_contested_elim___redArg___boxed(lean_object* v_contested_69_){
_start:
{
lean_object* v_res_70_; 
v_res_70_ = lp_bmc_CheckStatus_contested_elim___redArg(v_contested_69_);
lean_dec(v_contested_69_);
return v_res_70_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_contested_elim(lean_object* v_motive_71_, uint8_t v_t_72_, lean_object* v_h_73_, lean_object* v_contested_74_){
_start:
{
lean_inc(v_contested_74_);
return v_contested_74_;
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_contested_elim___boxed(lean_object* v_motive_75_, lean_object* v_t_76_, lean_object* v_h_77_, lean_object* v_contested_78_){
_start:
{
uint8_t v_t_boxed_79_; lean_object* v_res_80_; 
v_t_boxed_79_ = lean_unbox(v_t_76_);
v_res_80_ = lp_bmc_CheckStatus_contested_elim(v_motive_75_, v_t_boxed_79_, v_h_77_, v_contested_78_);
lean_dec(v_contested_78_);
return v_res_80_;
}
}
LEAN_EXPORT uint8_t lp_bmc_CheckStatus_ofNat(lean_object* v_n_81_){
_start:
{
lean_object* v___x_82_; uint8_t v___x_83_; 
v___x_82_ = lean_unsigned_to_nat(1u);
v___x_83_ = lean_nat_dec_le(v_n_81_, v___x_82_);
if (v___x_83_ == 0)
{
lean_object* v___x_84_; uint8_t v___x_85_; 
v___x_84_ = lean_unsigned_to_nat(2u);
v___x_85_ = lean_nat_dec_le(v_n_81_, v___x_84_);
if (v___x_85_ == 0)
{
uint8_t v___x_86_; 
v___x_86_ = 3;
return v___x_86_;
}
else
{
uint8_t v___x_87_; 
v___x_87_ = 2;
return v___x_87_;
}
}
else
{
lean_object* v___x_88_; uint8_t v___x_89_; 
v___x_88_ = lean_unsigned_to_nat(0u);
v___x_89_ = lean_nat_dec_le(v_n_81_, v___x_88_);
if (v___x_89_ == 0)
{
uint8_t v___x_90_; 
v___x_90_ = 1;
return v___x_90_;
}
else
{
uint8_t v___x_91_; 
v___x_91_ = 0;
return v___x_91_;
}
}
}
}
LEAN_EXPORT lean_object* lp_bmc_CheckStatus_ofNat___boxed(lean_object* v_n_92_){
_start:
{
uint8_t v_res_93_; lean_object* v_r_94_; 
v_res_93_ = lp_bmc_CheckStatus_ofNat(v_n_92_);
lean_dec(v_n_92_);
v_r_94_ = lean_box(v_res_93_);
return v_r_94_;
}
}
LEAN_EXPORT uint8_t lp_bmc_instDecidableEqCheckStatus(uint8_t v_x_95_, uint8_t v_y_96_){
_start:
{
lean_object* v___x_97_; lean_object* v___x_98_; uint8_t v___x_99_; 
v___x_97_ = lp_bmc_CheckStatus_ctorIdx(v_x_95_);
v___x_98_ = lp_bmc_CheckStatus_ctorIdx(v_y_96_);
v___x_99_ = lean_nat_dec_eq(v___x_97_, v___x_98_);
lean_dec(v___x_98_);
lean_dec(v___x_97_);
return v___x_99_;
}
}
LEAN_EXPORT lean_object* lp_bmc_instDecidableEqCheckStatus___boxed(lean_object* v_x_100_, lean_object* v_y_101_){
_start:
{
uint8_t v_x_13__boxed_102_; uint8_t v_y_14__boxed_103_; uint8_t v_res_104_; lean_object* v_r_105_; 
v_x_13__boxed_102_ = lean_unbox(v_x_100_);
v_y_14__boxed_103_ = lean_unbox(v_y_101_);
v_res_104_ = lp_bmc_instDecidableEqCheckStatus(v_x_13__boxed_102_, v_y_14__boxed_103_);
v_r_105_ = lean_box(v_res_104_);
return v_r_105_;
}
}
static lean_object* _init_lp_bmc_instReprCheckStatus_repr___closed__8(void){
_start:
{
lean_object* v___x_118_; lean_object* v___x_119_; 
v___x_118_ = lean_unsigned_to_nat(2u);
v___x_119_ = lean_nat_to_int(v___x_118_);
return v___x_119_;
}
}
static lean_object* _init_lp_bmc_instReprCheckStatus_repr___closed__9(void){
_start:
{
lean_object* v___x_120_; lean_object* v___x_121_; 
v___x_120_ = lean_unsigned_to_nat(1u);
v___x_121_ = lean_nat_to_int(v___x_120_);
return v___x_121_;
}
}
LEAN_EXPORT lean_object* lp_bmc_instReprCheckStatus_repr(uint8_t v_x_122_, lean_object* v_prec_123_){
_start:
{
lean_object* v___y_125_; lean_object* v___y_132_; lean_object* v___y_139_; lean_object* v___y_146_; 
switch(v_x_122_)
{
case 0:
{
lean_object* v___x_152_; uint8_t v___x_153_; 
v___x_152_ = lean_unsigned_to_nat(1024u);
v___x_153_ = lean_nat_dec_le(v___x_152_, v_prec_123_);
if (v___x_153_ == 0)
{
lean_object* v___x_154_; 
v___x_154_ = lean_obj_once(&lp_bmc_instReprCheckStatus_repr___closed__8, &lp_bmc_instReprCheckStatus_repr___closed__8_once, _init_lp_bmc_instReprCheckStatus_repr___closed__8);
v___y_125_ = v___x_154_;
goto v___jp_124_;
}
else
{
lean_object* v___x_155_; 
v___x_155_ = lean_obj_once(&lp_bmc_instReprCheckStatus_repr___closed__9, &lp_bmc_instReprCheckStatus_repr___closed__9_once, _init_lp_bmc_instReprCheckStatus_repr___closed__9);
v___y_125_ = v___x_155_;
goto v___jp_124_;
}
}
case 1:
{
lean_object* v___x_156_; uint8_t v___x_157_; 
v___x_156_ = lean_unsigned_to_nat(1024u);
v___x_157_ = lean_nat_dec_le(v___x_156_, v_prec_123_);
if (v___x_157_ == 0)
{
lean_object* v___x_158_; 
v___x_158_ = lean_obj_once(&lp_bmc_instReprCheckStatus_repr___closed__8, &lp_bmc_instReprCheckStatus_repr___closed__8_once, _init_lp_bmc_instReprCheckStatus_repr___closed__8);
v___y_132_ = v___x_158_;
goto v___jp_131_;
}
else
{
lean_object* v___x_159_; 
v___x_159_ = lean_obj_once(&lp_bmc_instReprCheckStatus_repr___closed__9, &lp_bmc_instReprCheckStatus_repr___closed__9_once, _init_lp_bmc_instReprCheckStatus_repr___closed__9);
v___y_132_ = v___x_159_;
goto v___jp_131_;
}
}
case 2:
{
lean_object* v___x_160_; uint8_t v___x_161_; 
v___x_160_ = lean_unsigned_to_nat(1024u);
v___x_161_ = lean_nat_dec_le(v___x_160_, v_prec_123_);
if (v___x_161_ == 0)
{
lean_object* v___x_162_; 
v___x_162_ = lean_obj_once(&lp_bmc_instReprCheckStatus_repr___closed__8, &lp_bmc_instReprCheckStatus_repr___closed__8_once, _init_lp_bmc_instReprCheckStatus_repr___closed__8);
v___y_139_ = v___x_162_;
goto v___jp_138_;
}
else
{
lean_object* v___x_163_; 
v___x_163_ = lean_obj_once(&lp_bmc_instReprCheckStatus_repr___closed__9, &lp_bmc_instReprCheckStatus_repr___closed__9_once, _init_lp_bmc_instReprCheckStatus_repr___closed__9);
v___y_139_ = v___x_163_;
goto v___jp_138_;
}
}
default: 
{
lean_object* v___x_164_; uint8_t v___x_165_; 
v___x_164_ = lean_unsigned_to_nat(1024u);
v___x_165_ = lean_nat_dec_le(v___x_164_, v_prec_123_);
if (v___x_165_ == 0)
{
lean_object* v___x_166_; 
v___x_166_ = lean_obj_once(&lp_bmc_instReprCheckStatus_repr___closed__8, &lp_bmc_instReprCheckStatus_repr___closed__8_once, _init_lp_bmc_instReprCheckStatus_repr___closed__8);
v___y_146_ = v___x_166_;
goto v___jp_145_;
}
else
{
lean_object* v___x_167_; 
v___x_167_ = lean_obj_once(&lp_bmc_instReprCheckStatus_repr___closed__9, &lp_bmc_instReprCheckStatus_repr___closed__9_once, _init_lp_bmc_instReprCheckStatus_repr___closed__9);
v___y_146_ = v___x_167_;
goto v___jp_145_;
}
}
}
v___jp_124_:
{
lean_object* v___x_126_; lean_object* v___x_127_; uint8_t v___x_128_; lean_object* v___x_129_; lean_object* v___x_130_; 
v___x_126_ = ((lean_object*)(lp_bmc_instReprCheckStatus_repr___closed__1));
lean_inc(v___y_125_);
v___x_127_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_127_, 0, v___y_125_);
lean_ctor_set(v___x_127_, 1, v___x_126_);
v___x_128_ = 0;
v___x_129_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_129_, 0, v___x_127_);
lean_ctor_set_uint8(v___x_129_, sizeof(void*)*1, v___x_128_);
v___x_130_ = l_Repr_addAppParen(v___x_129_, v_prec_123_);
return v___x_130_;
}
v___jp_131_:
{
lean_object* v___x_133_; lean_object* v___x_134_; uint8_t v___x_135_; lean_object* v___x_136_; lean_object* v___x_137_; 
v___x_133_ = ((lean_object*)(lp_bmc_instReprCheckStatus_repr___closed__3));
lean_inc(v___y_132_);
v___x_134_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_134_, 0, v___y_132_);
lean_ctor_set(v___x_134_, 1, v___x_133_);
v___x_135_ = 0;
v___x_136_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_136_, 0, v___x_134_);
lean_ctor_set_uint8(v___x_136_, sizeof(void*)*1, v___x_135_);
v___x_137_ = l_Repr_addAppParen(v___x_136_, v_prec_123_);
return v___x_137_;
}
v___jp_138_:
{
lean_object* v___x_140_; lean_object* v___x_141_; uint8_t v___x_142_; lean_object* v___x_143_; lean_object* v___x_144_; 
v___x_140_ = ((lean_object*)(lp_bmc_instReprCheckStatus_repr___closed__5));
lean_inc(v___y_139_);
v___x_141_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_141_, 0, v___y_139_);
lean_ctor_set(v___x_141_, 1, v___x_140_);
v___x_142_ = 0;
v___x_143_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_143_, 0, v___x_141_);
lean_ctor_set_uint8(v___x_143_, sizeof(void*)*1, v___x_142_);
v___x_144_ = l_Repr_addAppParen(v___x_143_, v_prec_123_);
return v___x_144_;
}
v___jp_145_:
{
lean_object* v___x_147_; lean_object* v___x_148_; uint8_t v___x_149_; lean_object* v___x_150_; lean_object* v___x_151_; 
v___x_147_ = ((lean_object*)(lp_bmc_instReprCheckStatus_repr___closed__7));
lean_inc(v___y_146_);
v___x_148_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_148_, 0, v___y_146_);
lean_ctor_set(v___x_148_, 1, v___x_147_);
v___x_149_ = 0;
v___x_150_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_150_, 0, v___x_148_);
lean_ctor_set_uint8(v___x_150_, sizeof(void*)*1, v___x_149_);
v___x_151_ = l_Repr_addAppParen(v___x_150_, v_prec_123_);
return v___x_151_;
}
}
}
LEAN_EXPORT lean_object* lp_bmc_instReprCheckStatus_repr___boxed(lean_object* v_x_168_, lean_object* v_prec_169_){
_start:
{
uint8_t v_x_233__boxed_170_; lean_object* v_res_171_; 
v_x_233__boxed_170_ = lean_unbox(v_x_168_);
v_res_171_ = lp_bmc_instReprCheckStatus_repr(v_x_233__boxed_170_, v_prec_169_);
lean_dec(v_prec_169_);
return v_res_171_;
}
}
LEAN_EXPORT uint8_t lp_bmc_checkPassed(uint8_t v_x_174_){
_start:
{
if (v_x_174_ == 0)
{
uint8_t v___x_175_; 
v___x_175_ = 1;
return v___x_175_;
}
else
{
uint8_t v___x_176_; 
v___x_176_ = 0;
return v___x_176_;
}
}
}
LEAN_EXPORT lean_object* lp_bmc_checkPassed___boxed(lean_object* v_x_177_){
_start:
{
uint8_t v_x_21__boxed_178_; uint8_t v_res_179_; lean_object* v_r_180_; 
v_x_21__boxed_178_ = lean_unbox(v_x_177_);
v_res_179_ = lp_bmc_checkPassed(v_x_21__boxed_178_);
v_r_180_ = lean_box(v_res_179_);
return v_r_180_;
}
}
LEAN_EXPORT uint8_t lp_bmc_checkDeferred(uint8_t v_x_181_){
_start:
{
if (v_x_181_ == 2)
{
uint8_t v___x_182_; 
v___x_182_ = 1;
return v___x_182_;
}
else
{
uint8_t v___x_183_; 
v___x_183_ = 0;
return v___x_183_;
}
}
}
LEAN_EXPORT lean_object* lp_bmc_checkDeferred___boxed(lean_object* v_x_184_){
_start:
{
uint8_t v_x_21__boxed_185_; uint8_t v_res_186_; lean_object* v_r_187_; 
v_x_21__boxed_185_ = lean_unbox(v_x_184_);
v_res_186_ = lp_bmc_checkDeferred(v_x_21__boxed_185_);
v_r_187_ = lean_box(v_res_186_);
return v_r_187_;
}
}
static lean_object* _init_lp_bmc_instReprBMCReport_repr___redArg___closed__7(void){
_start:
{
lean_object* v___x_201_; lean_object* v___x_202_; 
v___x_201_ = lean_unsigned_to_nat(19u);
v___x_202_ = lean_nat_to_int(v___x_201_);
return v___x_202_;
}
}
static lean_object* _init_lp_bmc_instReprBMCReport_repr___redArg___closed__14(void){
_start:
{
lean_object* v___x_212_; lean_object* v___x_213_; 
v___x_212_ = lean_unsigned_to_nat(15u);
v___x_213_ = lean_nat_to_int(v___x_212_);
return v___x_213_;
}
}
static lean_object* _init_lp_bmc_instReprBMCReport_repr___redArg___closed__17(void){
_start:
{
lean_object* v___x_217_; lean_object* v___x_218_; 
v___x_217_ = lean_unsigned_to_nat(20u);
v___x_218_ = lean_nat_to_int(v___x_217_);
return v___x_218_;
}
}
static lean_object* _init_lp_bmc_instReprBMCReport_repr___redArg___closed__20(void){
_start:
{
lean_object* v___x_222_; lean_object* v___x_223_; 
v___x_222_ = lean_unsigned_to_nat(18u);
v___x_223_ = lean_nat_to_int(v___x_222_);
return v___x_223_;
}
}
static lean_object* _init_lp_bmc_instReprBMCReport_repr___redArg___closed__23(void){
_start:
{
lean_object* v___x_227_; lean_object* v___x_228_; 
v___x_227_ = lean_unsigned_to_nat(17u);
v___x_228_ = lean_nat_to_int(v___x_227_);
return v___x_228_;
}
}
static lean_object* _init_lp_bmc_instReprBMCReport_repr___redArg___closed__28(void){
_start:
{
lean_object* v___x_235_; lean_object* v___x_236_; 
v___x_235_ = lean_unsigned_to_nat(24u);
v___x_236_ = lean_nat_to_int(v___x_235_);
return v___x_236_;
}
}
static lean_object* _init_lp_bmc_instReprBMCReport_repr___redArg___closed__31(void){
_start:
{
lean_object* v___x_240_; lean_object* v___x_241_; 
v___x_240_ = lean_unsigned_to_nat(23u);
v___x_241_ = lean_nat_to_int(v___x_240_);
return v___x_241_;
}
}
static lean_object* _init_lp_bmc_instReprBMCReport_repr___redArg___closed__36(void){
_start:
{
lean_object* v___x_248_; lean_object* v___x_249_; 
v___x_248_ = lean_unsigned_to_nat(21u);
v___x_249_ = lean_nat_to_int(v___x_248_);
return v___x_249_;
}
}
static lean_object* _init_lp_bmc_instReprBMCReport_repr___redArg___closed__39(void){
_start:
{
lean_object* v___x_253_; lean_object* v___x_254_; 
v___x_253_ = lean_unsigned_to_nat(16u);
v___x_254_ = lean_nat_to_int(v___x_253_);
return v___x_254_;
}
}
static lean_object* _init_lp_bmc_instReprBMCReport_repr___redArg___closed__41(void){
_start:
{
lean_object* v___x_256_; lean_object* v___x_257_; 
v___x_256_ = ((lean_object*)(lp_bmc_instReprBMCReport_repr___redArg___closed__0));
v___x_257_ = lean_string_length(v___x_256_);
return v___x_257_;
}
}
static lean_object* _init_lp_bmc_instReprBMCReport_repr___redArg___closed__42(void){
_start:
{
lean_object* v___x_258_; lean_object* v___x_259_; 
v___x_258_ = lean_obj_once(&lp_bmc_instReprBMCReport_repr___redArg___closed__41, &lp_bmc_instReprBMCReport_repr___redArg___closed__41_once, _init_lp_bmc_instReprBMCReport_repr___redArg___closed__41);
v___x_259_ = lean_nat_to_int(v___x_258_);
return v___x_259_;
}
}
LEAN_EXPORT lean_object* lp_bmc_instReprBMCReport_repr___redArg(lean_object* v_x_264_){
_start:
{
uint8_t v_toyAnalysisOnly_265_; uint8_t v_finalTruthClaim_266_; uint8_t v_wdwResidual_267_; uint8_t v_trajectoryFinite_268_; uint8_t v_clockMonotonic_269_; uint8_t v_nodeDetection_270_; uint8_t v_nodeContactFree_271_; uint8_t v_qFiniteAwayFromNodes_272_; uint8_t v_phaseGradientFinite_273_; uint8_t v_classicalLimit_274_; uint8_t v_friedmannResidual_275_; uint8_t v_faithfulness_276_; lean_object* v___x_277_; lean_object* v___x_278_; lean_object* v___x_279_; lean_object* v___x_280_; lean_object* v___x_281_; lean_object* v___x_282_; uint8_t v___x_283_; lean_object* v___x_284_; lean_object* v___x_285_; lean_object* v___x_286_; lean_object* v___x_287_; lean_object* v___x_288_; lean_object* v___x_289_; lean_object* v___x_290_; lean_object* v___x_291_; lean_object* v___x_292_; lean_object* v___x_293_; lean_object* v___x_294_; lean_object* v___x_295_; lean_object* v___x_296_; lean_object* v___x_297_; lean_object* v___x_298_; lean_object* v___x_299_; lean_object* v___x_300_; lean_object* v___x_301_; lean_object* v___x_302_; lean_object* v___x_303_; lean_object* v___x_304_; lean_object* v___x_305_; lean_object* v___x_306_; lean_object* v___x_307_; lean_object* v___x_308_; lean_object* v___x_309_; lean_object* v___x_310_; lean_object* v___x_311_; lean_object* v___x_312_; lean_object* v___x_313_; lean_object* v___x_314_; lean_object* v___x_315_; lean_object* v___x_316_; lean_object* v___x_317_; lean_object* v___x_318_; lean_object* v___x_319_; lean_object* v___x_320_; lean_object* v___x_321_; lean_object* v___x_322_; lean_object* v___x_323_; lean_object* v___x_324_; lean_object* v___x_325_; lean_object* v___x_326_; lean_object* v___x_327_; lean_object* v___x_328_; lean_object* v___x_329_; lean_object* v___x_330_; lean_object* v___x_331_; lean_object* v___x_332_; lean_object* v___x_333_; lean_object* v___x_334_; lean_object* v___x_335_; lean_object* v___x_336_; lean_object* v___x_337_; lean_object* v___x_338_; lean_object* v___x_339_; lean_object* v___x_340_; lean_object* v___x_341_; lean_object* v___x_342_; lean_object* v___x_343_; lean_object* v___x_344_; lean_object* v___x_345_; lean_object* v___x_346_; lean_object* v___x_347_; lean_object* v___x_348_; lean_object* v___x_349_; lean_object* v___x_350_; lean_object* v___x_351_; lean_object* v___x_352_; lean_object* v___x_353_; lean_object* v___x_354_; lean_object* v___x_355_; lean_object* v___x_356_; lean_object* v___x_357_; lean_object* v___x_358_; lean_object* v___x_359_; lean_object* v___x_360_; lean_object* v___x_361_; lean_object* v___x_362_; lean_object* v___x_363_; lean_object* v___x_364_; lean_object* v___x_365_; lean_object* v___x_366_; lean_object* v___x_367_; lean_object* v___x_368_; lean_object* v___x_369_; lean_object* v___x_370_; lean_object* v___x_371_; lean_object* v___x_372_; lean_object* v___x_373_; lean_object* v___x_374_; lean_object* v___x_375_; lean_object* v___x_376_; lean_object* v___x_377_; lean_object* v___x_378_; lean_object* v___x_379_; lean_object* v___x_380_; lean_object* v___x_381_; lean_object* v___x_382_; lean_object* v___x_383_; lean_object* v___x_384_; lean_object* v___x_385_; lean_object* v___x_386_; lean_object* v___x_387_; lean_object* v___x_388_; lean_object* v___x_389_; lean_object* v___x_390_; lean_object* v___x_391_; lean_object* v___x_392_; lean_object* v___x_393_; lean_object* v___x_394_; lean_object* v___x_395_; lean_object* v___x_396_; lean_object* v___x_397_; lean_object* v___x_398_; lean_object* v___x_399_; lean_object* v___x_400_; lean_object* v___x_401_; 
v_toyAnalysisOnly_265_ = lean_ctor_get_uint8(v_x_264_, 0);
v_finalTruthClaim_266_ = lean_ctor_get_uint8(v_x_264_, 1);
v_wdwResidual_267_ = lean_ctor_get_uint8(v_x_264_, 2);
v_trajectoryFinite_268_ = lean_ctor_get_uint8(v_x_264_, 3);
v_clockMonotonic_269_ = lean_ctor_get_uint8(v_x_264_, 4);
v_nodeDetection_270_ = lean_ctor_get_uint8(v_x_264_, 5);
v_nodeContactFree_271_ = lean_ctor_get_uint8(v_x_264_, 6);
v_qFiniteAwayFromNodes_272_ = lean_ctor_get_uint8(v_x_264_, 7);
v_phaseGradientFinite_273_ = lean_ctor_get_uint8(v_x_264_, 8);
v_classicalLimit_274_ = lean_ctor_get_uint8(v_x_264_, 9);
v_friedmannResidual_275_ = lean_ctor_get_uint8(v_x_264_, 10);
v_faithfulness_276_ = lean_ctor_get_uint8(v_x_264_, 11);
v___x_277_ = ((lean_object*)(lp_bmc_instReprBMCReport_repr___redArg___closed__5));
v___x_278_ = ((lean_object*)(lp_bmc_instReprBMCReport_repr___redArg___closed__6));
v___x_279_ = lean_obj_once(&lp_bmc_instReprBMCReport_repr___redArg___closed__7, &lp_bmc_instReprBMCReport_repr___redArg___closed__7_once, _init_lp_bmc_instReprBMCReport_repr___redArg___closed__7);
v___x_280_ = lean_unsigned_to_nat(0u);
v___x_281_ = l_Bool_repr___redArg(v_toyAnalysisOnly_265_);
v___x_282_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_282_, 0, v___x_279_);
lean_ctor_set(v___x_282_, 1, v___x_281_);
v___x_283_ = 0;
v___x_284_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_284_, 0, v___x_282_);
lean_ctor_set_uint8(v___x_284_, sizeof(void*)*1, v___x_283_);
v___x_285_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_285_, 0, v___x_278_);
lean_ctor_set(v___x_285_, 1, v___x_284_);
v___x_286_ = ((lean_object*)(lp_bmc_instReprBMCReport_repr___redArg___closed__9));
v___x_287_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_287_, 0, v___x_285_);
lean_ctor_set(v___x_287_, 1, v___x_286_);
v___x_288_ = lean_box(1);
v___x_289_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_289_, 0, v___x_287_);
lean_ctor_set(v___x_289_, 1, v___x_288_);
v___x_290_ = ((lean_object*)(lp_bmc_instReprBMCReport_repr___redArg___closed__11));
v___x_291_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_291_, 0, v___x_289_);
lean_ctor_set(v___x_291_, 1, v___x_290_);
v___x_292_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_292_, 0, v___x_291_);
lean_ctor_set(v___x_292_, 1, v___x_277_);
v___x_293_ = l_Bool_repr___redArg(v_finalTruthClaim_266_);
v___x_294_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_294_, 0, v___x_279_);
lean_ctor_set(v___x_294_, 1, v___x_293_);
v___x_295_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_295_, 0, v___x_294_);
lean_ctor_set_uint8(v___x_295_, sizeof(void*)*1, v___x_283_);
v___x_296_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_296_, 0, v___x_292_);
lean_ctor_set(v___x_296_, 1, v___x_295_);
v___x_297_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_297_, 0, v___x_296_);
lean_ctor_set(v___x_297_, 1, v___x_286_);
v___x_298_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_298_, 0, v___x_297_);
lean_ctor_set(v___x_298_, 1, v___x_288_);
v___x_299_ = ((lean_object*)(lp_bmc_instReprBMCReport_repr___redArg___closed__13));
v___x_300_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_300_, 0, v___x_298_);
lean_ctor_set(v___x_300_, 1, v___x_299_);
v___x_301_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_301_, 0, v___x_300_);
lean_ctor_set(v___x_301_, 1, v___x_277_);
v___x_302_ = lean_obj_once(&lp_bmc_instReprBMCReport_repr___redArg___closed__14, &lp_bmc_instReprBMCReport_repr___redArg___closed__14_once, _init_lp_bmc_instReprBMCReport_repr___redArg___closed__14);
v___x_303_ = lp_bmc_instReprCheckStatus_repr(v_wdwResidual_267_, v___x_280_);
v___x_304_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_304_, 0, v___x_302_);
lean_ctor_set(v___x_304_, 1, v___x_303_);
v___x_305_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_305_, 0, v___x_304_);
lean_ctor_set_uint8(v___x_305_, sizeof(void*)*1, v___x_283_);
v___x_306_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_306_, 0, v___x_301_);
lean_ctor_set(v___x_306_, 1, v___x_305_);
v___x_307_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_307_, 0, v___x_306_);
lean_ctor_set(v___x_307_, 1, v___x_286_);
v___x_308_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_308_, 0, v___x_307_);
lean_ctor_set(v___x_308_, 1, v___x_288_);
v___x_309_ = ((lean_object*)(lp_bmc_instReprBMCReport_repr___redArg___closed__16));
v___x_310_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_310_, 0, v___x_308_);
lean_ctor_set(v___x_310_, 1, v___x_309_);
v___x_311_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_311_, 0, v___x_310_);
lean_ctor_set(v___x_311_, 1, v___x_277_);
v___x_312_ = lean_obj_once(&lp_bmc_instReprBMCReport_repr___redArg___closed__17, &lp_bmc_instReprBMCReport_repr___redArg___closed__17_once, _init_lp_bmc_instReprBMCReport_repr___redArg___closed__17);
v___x_313_ = lp_bmc_instReprCheckStatus_repr(v_trajectoryFinite_268_, v___x_280_);
v___x_314_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_314_, 0, v___x_312_);
lean_ctor_set(v___x_314_, 1, v___x_313_);
v___x_315_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_315_, 0, v___x_314_);
lean_ctor_set_uint8(v___x_315_, sizeof(void*)*1, v___x_283_);
v___x_316_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_316_, 0, v___x_311_);
lean_ctor_set(v___x_316_, 1, v___x_315_);
v___x_317_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_317_, 0, v___x_316_);
lean_ctor_set(v___x_317_, 1, v___x_286_);
v___x_318_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_318_, 0, v___x_317_);
lean_ctor_set(v___x_318_, 1, v___x_288_);
v___x_319_ = ((lean_object*)(lp_bmc_instReprBMCReport_repr___redArg___closed__19));
v___x_320_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_320_, 0, v___x_318_);
lean_ctor_set(v___x_320_, 1, v___x_319_);
v___x_321_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_321_, 0, v___x_320_);
lean_ctor_set(v___x_321_, 1, v___x_277_);
v___x_322_ = lean_obj_once(&lp_bmc_instReprBMCReport_repr___redArg___closed__20, &lp_bmc_instReprBMCReport_repr___redArg___closed__20_once, _init_lp_bmc_instReprBMCReport_repr___redArg___closed__20);
v___x_323_ = lp_bmc_instReprCheckStatus_repr(v_clockMonotonic_269_, v___x_280_);
v___x_324_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_324_, 0, v___x_322_);
lean_ctor_set(v___x_324_, 1, v___x_323_);
v___x_325_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_325_, 0, v___x_324_);
lean_ctor_set_uint8(v___x_325_, sizeof(void*)*1, v___x_283_);
v___x_326_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_326_, 0, v___x_321_);
lean_ctor_set(v___x_326_, 1, v___x_325_);
v___x_327_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_327_, 0, v___x_326_);
lean_ctor_set(v___x_327_, 1, v___x_286_);
v___x_328_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_328_, 0, v___x_327_);
lean_ctor_set(v___x_328_, 1, v___x_288_);
v___x_329_ = ((lean_object*)(lp_bmc_instReprBMCReport_repr___redArg___closed__22));
v___x_330_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_330_, 0, v___x_328_);
lean_ctor_set(v___x_330_, 1, v___x_329_);
v___x_331_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_331_, 0, v___x_330_);
lean_ctor_set(v___x_331_, 1, v___x_277_);
v___x_332_ = lean_obj_once(&lp_bmc_instReprBMCReport_repr___redArg___closed__23, &lp_bmc_instReprBMCReport_repr___redArg___closed__23_once, _init_lp_bmc_instReprBMCReport_repr___redArg___closed__23);
v___x_333_ = lp_bmc_instReprCheckStatus_repr(v_nodeDetection_270_, v___x_280_);
v___x_334_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_334_, 0, v___x_332_);
lean_ctor_set(v___x_334_, 1, v___x_333_);
v___x_335_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_335_, 0, v___x_334_);
lean_ctor_set_uint8(v___x_335_, sizeof(void*)*1, v___x_283_);
v___x_336_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_336_, 0, v___x_331_);
lean_ctor_set(v___x_336_, 1, v___x_335_);
v___x_337_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_337_, 0, v___x_336_);
lean_ctor_set(v___x_337_, 1, v___x_286_);
v___x_338_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_338_, 0, v___x_337_);
lean_ctor_set(v___x_338_, 1, v___x_288_);
v___x_339_ = ((lean_object*)(lp_bmc_instReprBMCReport_repr___redArg___closed__25));
v___x_340_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_340_, 0, v___x_338_);
lean_ctor_set(v___x_340_, 1, v___x_339_);
v___x_341_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_341_, 0, v___x_340_);
lean_ctor_set(v___x_341_, 1, v___x_277_);
v___x_342_ = lp_bmc_instReprCheckStatus_repr(v_nodeContactFree_271_, v___x_280_);
v___x_343_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_343_, 0, v___x_279_);
lean_ctor_set(v___x_343_, 1, v___x_342_);
v___x_344_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_344_, 0, v___x_343_);
lean_ctor_set_uint8(v___x_344_, sizeof(void*)*1, v___x_283_);
v___x_345_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_345_, 0, v___x_341_);
lean_ctor_set(v___x_345_, 1, v___x_344_);
v___x_346_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_346_, 0, v___x_345_);
lean_ctor_set(v___x_346_, 1, v___x_286_);
v___x_347_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_347_, 0, v___x_346_);
lean_ctor_set(v___x_347_, 1, v___x_288_);
v___x_348_ = ((lean_object*)(lp_bmc_instReprBMCReport_repr___redArg___closed__27));
v___x_349_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_349_, 0, v___x_347_);
lean_ctor_set(v___x_349_, 1, v___x_348_);
v___x_350_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_350_, 0, v___x_349_);
lean_ctor_set(v___x_350_, 1, v___x_277_);
v___x_351_ = lean_obj_once(&lp_bmc_instReprBMCReport_repr___redArg___closed__28, &lp_bmc_instReprBMCReport_repr___redArg___closed__28_once, _init_lp_bmc_instReprBMCReport_repr___redArg___closed__28);
v___x_352_ = lp_bmc_instReprCheckStatus_repr(v_qFiniteAwayFromNodes_272_, v___x_280_);
v___x_353_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_353_, 0, v___x_351_);
lean_ctor_set(v___x_353_, 1, v___x_352_);
v___x_354_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_354_, 0, v___x_353_);
lean_ctor_set_uint8(v___x_354_, sizeof(void*)*1, v___x_283_);
v___x_355_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_355_, 0, v___x_350_);
lean_ctor_set(v___x_355_, 1, v___x_354_);
v___x_356_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_356_, 0, v___x_355_);
lean_ctor_set(v___x_356_, 1, v___x_286_);
v___x_357_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_357_, 0, v___x_356_);
lean_ctor_set(v___x_357_, 1, v___x_288_);
v___x_358_ = ((lean_object*)(lp_bmc_instReprBMCReport_repr___redArg___closed__30));
v___x_359_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_359_, 0, v___x_357_);
lean_ctor_set(v___x_359_, 1, v___x_358_);
v___x_360_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_360_, 0, v___x_359_);
lean_ctor_set(v___x_360_, 1, v___x_277_);
v___x_361_ = lean_obj_once(&lp_bmc_instReprBMCReport_repr___redArg___closed__31, &lp_bmc_instReprBMCReport_repr___redArg___closed__31_once, _init_lp_bmc_instReprBMCReport_repr___redArg___closed__31);
v___x_362_ = lp_bmc_instReprCheckStatus_repr(v_phaseGradientFinite_273_, v___x_280_);
v___x_363_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_363_, 0, v___x_361_);
lean_ctor_set(v___x_363_, 1, v___x_362_);
v___x_364_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_364_, 0, v___x_363_);
lean_ctor_set_uint8(v___x_364_, sizeof(void*)*1, v___x_283_);
v___x_365_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_365_, 0, v___x_360_);
lean_ctor_set(v___x_365_, 1, v___x_364_);
v___x_366_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_366_, 0, v___x_365_);
lean_ctor_set(v___x_366_, 1, v___x_286_);
v___x_367_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_367_, 0, v___x_366_);
lean_ctor_set(v___x_367_, 1, v___x_288_);
v___x_368_ = ((lean_object*)(lp_bmc_instReprBMCReport_repr___redArg___closed__33));
v___x_369_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_369_, 0, v___x_367_);
lean_ctor_set(v___x_369_, 1, v___x_368_);
v___x_370_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_370_, 0, v___x_369_);
lean_ctor_set(v___x_370_, 1, v___x_277_);
v___x_371_ = lp_bmc_instReprCheckStatus_repr(v_classicalLimit_274_, v___x_280_);
v___x_372_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_372_, 0, v___x_322_);
lean_ctor_set(v___x_372_, 1, v___x_371_);
v___x_373_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_373_, 0, v___x_372_);
lean_ctor_set_uint8(v___x_373_, sizeof(void*)*1, v___x_283_);
v___x_374_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_374_, 0, v___x_370_);
lean_ctor_set(v___x_374_, 1, v___x_373_);
v___x_375_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_375_, 0, v___x_374_);
lean_ctor_set(v___x_375_, 1, v___x_286_);
v___x_376_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_376_, 0, v___x_375_);
lean_ctor_set(v___x_376_, 1, v___x_288_);
v___x_377_ = ((lean_object*)(lp_bmc_instReprBMCReport_repr___redArg___closed__35));
v___x_378_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_378_, 0, v___x_376_);
lean_ctor_set(v___x_378_, 1, v___x_377_);
v___x_379_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_379_, 0, v___x_378_);
lean_ctor_set(v___x_379_, 1, v___x_277_);
v___x_380_ = lean_obj_once(&lp_bmc_instReprBMCReport_repr___redArg___closed__36, &lp_bmc_instReprBMCReport_repr___redArg___closed__36_once, _init_lp_bmc_instReprBMCReport_repr___redArg___closed__36);
v___x_381_ = lp_bmc_instReprCheckStatus_repr(v_friedmannResidual_275_, v___x_280_);
v___x_382_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_382_, 0, v___x_380_);
lean_ctor_set(v___x_382_, 1, v___x_381_);
v___x_383_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_383_, 0, v___x_382_);
lean_ctor_set_uint8(v___x_383_, sizeof(void*)*1, v___x_283_);
v___x_384_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_384_, 0, v___x_379_);
lean_ctor_set(v___x_384_, 1, v___x_383_);
v___x_385_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_385_, 0, v___x_384_);
lean_ctor_set(v___x_385_, 1, v___x_286_);
v___x_386_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_386_, 0, v___x_385_);
lean_ctor_set(v___x_386_, 1, v___x_288_);
v___x_387_ = ((lean_object*)(lp_bmc_instReprBMCReport_repr___redArg___closed__38));
v___x_388_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_388_, 0, v___x_386_);
lean_ctor_set(v___x_388_, 1, v___x_387_);
v___x_389_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_389_, 0, v___x_388_);
lean_ctor_set(v___x_389_, 1, v___x_277_);
v___x_390_ = lean_obj_once(&lp_bmc_instReprBMCReport_repr___redArg___closed__39, &lp_bmc_instReprBMCReport_repr___redArg___closed__39_once, _init_lp_bmc_instReprBMCReport_repr___redArg___closed__39);
v___x_391_ = lp_bmc_instReprCheckStatus_repr(v_faithfulness_276_, v___x_280_);
v___x_392_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_392_, 0, v___x_390_);
lean_ctor_set(v___x_392_, 1, v___x_391_);
v___x_393_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_393_, 0, v___x_392_);
lean_ctor_set_uint8(v___x_393_, sizeof(void*)*1, v___x_283_);
v___x_394_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_394_, 0, v___x_389_);
lean_ctor_set(v___x_394_, 1, v___x_393_);
v___x_395_ = lean_obj_once(&lp_bmc_instReprBMCReport_repr___redArg___closed__42, &lp_bmc_instReprBMCReport_repr___redArg___closed__42_once, _init_lp_bmc_instReprBMCReport_repr___redArg___closed__42);
v___x_396_ = ((lean_object*)(lp_bmc_instReprBMCReport_repr___redArg___closed__43));
v___x_397_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_397_, 0, v___x_396_);
lean_ctor_set(v___x_397_, 1, v___x_394_);
v___x_398_ = ((lean_object*)(lp_bmc_instReprBMCReport_repr___redArg___closed__44));
v___x_399_ = lean_alloc_ctor(5, 2, 0);
lean_ctor_set(v___x_399_, 0, v___x_397_);
lean_ctor_set(v___x_399_, 1, v___x_398_);
v___x_400_ = lean_alloc_ctor(4, 2, 0);
lean_ctor_set(v___x_400_, 0, v___x_395_);
lean_ctor_set(v___x_400_, 1, v___x_399_);
v___x_401_ = lean_alloc_ctor(6, 1, 1);
lean_ctor_set(v___x_401_, 0, v___x_400_);
lean_ctor_set_uint8(v___x_401_, sizeof(void*)*1, v___x_283_);
return v___x_401_;
}
}
LEAN_EXPORT lean_object* lp_bmc_instReprBMCReport_repr___redArg___boxed(lean_object* v_x_402_){
_start:
{
lean_object* v_res_403_; 
v_res_403_ = lp_bmc_instReprBMCReport_repr___redArg(v_x_402_);
lean_dec_ref(v_x_402_);
return v_res_403_;
}
}
LEAN_EXPORT lean_object* lp_bmc_instReprBMCReport_repr(lean_object* v_x_404_, lean_object* v_prec_405_){
_start:
{
lean_object* v___x_406_; 
v___x_406_ = lp_bmc_instReprBMCReport_repr___redArg(v_x_404_);
return v___x_406_;
}
}
LEAN_EXPORT lean_object* lp_bmc_instReprBMCReport_repr___boxed(lean_object* v_x_407_, lean_object* v_prec_408_){
_start:
{
lean_object* v_res_409_; 
v_res_409_ = lp_bmc_instReprBMCReport_repr(v_x_407_, v_prec_408_);
lean_dec(v_prec_408_);
lean_dec_ref(v_x_407_);
return v_res_409_;
}
}
lean_object* initialize_Init(uint8_t builtin);
lean_object* initialize_Init(uint8_t builtin);
static bool _G_initialized = false;
LEAN_EXPORT lean_object* initialize_bmc_BMC_ToyReport(uint8_t builtin) {
lean_object * res;
if (_G_initialized) return lean_io_result_mk_ok(lean_box(0));
_G_initialized = true;
res = initialize_Init(builtin);
if (lean_io_result_is_error(res)) return res;
lean_dec_ref(res);
res = initialize_Init(builtin);
if (lean_io_result_is_error(res)) return res;
lean_dec_ref(res);
return lean_io_result_mk_ok(lean_box(0));
}
#ifdef __cplusplus
}
#endif
