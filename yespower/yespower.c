#include "yespower.h"

void yespower_hash(const char* input, uint32_t inputLen, uint32_t N, uint32_t r, char* pers, uint32_t persLen, char* output) {
    const yespower_params_t params = {
        .version = YESPOWER_1_0,
        .N = N,
        .r = r,
        .pers = (const uint8_t*)pers,
        .perslen = persLen
    };
    yespower_tls((const uint8_t*)input, inputLen, &params, (yespower_binary_t*)output);
}