// Lean compiler output
// Module: BMC
// Imports: public import Init public meta import Init public import BMC.ToyReport public import BMC.Promotion public import BMC.Robustness public import BMC.ClockFragility public import BMC.ClockReadiness public import BMC.FriedmannSpec
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
lean_object* initialize_Init(uint8_t builtin);
lean_object* initialize_Init(uint8_t builtin);
lean_object* initialize_bmc_BMC_ToyReport(uint8_t builtin);
lean_object* initialize_bmc_BMC_Promotion(uint8_t builtin);
lean_object* initialize_bmc_BMC_Robustness(uint8_t builtin);
lean_object* initialize_bmc_BMC_ClockFragility(uint8_t builtin);
lean_object* initialize_bmc_BMC_ClockReadiness(uint8_t builtin);
lean_object* initialize_bmc_BMC_FriedmannSpec(uint8_t builtin);
static bool _G_initialized = false;
LEAN_EXPORT lean_object* initialize_bmc_BMC(uint8_t builtin) {
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
res = initialize_bmc_BMC_Promotion(builtin);
if (lean_io_result_is_error(res)) return res;
lean_dec_ref(res);
res = initialize_bmc_BMC_Robustness(builtin);
if (lean_io_result_is_error(res)) return res;
lean_dec_ref(res);
res = initialize_bmc_BMC_ClockFragility(builtin);
if (lean_io_result_is_error(res)) return res;
lean_dec_ref(res);
res = initialize_bmc_BMC_ClockReadiness(builtin);
if (lean_io_result_is_error(res)) return res;
lean_dec_ref(res);
res = initialize_bmc_BMC_FriedmannSpec(builtin);
if (lean_io_result_is_error(res)) return res;
lean_dec_ref(res);
return lean_io_result_mk_ok(lean_box(0));
}
#ifdef __cplusplus
}
#endif
