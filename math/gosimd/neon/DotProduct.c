#include <arm_neon.h>
#include <stdbool.h>

float DotProduct(float *left, float *right, int len, float result) {
    int i;
    float32x4_t sum_vec = vdupq_n_f32(0.0f);
    for (i = 0; i <= len - 4; i += 4) {
        float32x4_t left_vec = vld1q_f32(&left[i]);
        float32x4_t right_vec = vld1q_f32(&right[i]);
        sum_vec = vmlaq_f32(sum_vec, left_vec, right_vec);
    }
    float sum = vaddvq_f32(sum_vec);
    for (; i < len; i++) {
        sum += left[i] * right[i];
    }
    return sum + result;
}

bool Supported() {
    return true;
}