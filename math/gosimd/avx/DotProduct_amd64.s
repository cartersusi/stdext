// +build amd64

// func DotProduct(left, right []float32, result float32) float32
TEXT ·DotProduct(SB), 4, $0
    // go slice lengths
    MOVQ    leftLen+8(FP), AX
    MOVQ    rightLen+32(FP), BX
    
    // min length
    CMPQ    AX, BX
    CMOVQLT AX, BX
    
    // pointers to slice data
    MOVQ    leftData+0(FP), SI
    MOVQ    rightData+24(FP), DX
    
    // init accum registers
    VXORPS  Y2, Y2, Y2  // Y2 holds partial sums for vec ops
    XORPS   X3, X3      // X3 holds final sum
    
    // loop index
    MOVQ    $0, CX
vectorLoop:
    MOVQ    BX, DI
    SUBQ    CX, DI
    CMPQ    DI, $8
    JL      singleLoop
    
    // proc 8 float32
    VMOVUPS (SI)(CX*4), Y0
    VMOVUPS (DX)(CX*4), Y1
    VMULPS  Y0, Y1, Y0  // vec mul
    VADDPS  Y0, Y2, Y2  // accumm
    
    ADDQ    $8, CX
    JMP     vectorLoop

singleLoop:
    CMPQ    CX, BX
    JGE     reduction
    
    // proc one float32 value
    MOVSS   (SI)(CX*4), X0
    MOVSS   (DX)(CX*4), X1
    MULSS   X1, X0
    ADDSS   X0, X3
    
    INCQ    CX
    JMP     singleLoop

reduction:
    // reduce Y2 vector register to scalar
    VEXTRACTF128 $1, Y2, X1
    VEXTRACTF128 $0, Y2, X0
    ADDPS   X1, X0
    HADDPS  X0, X0
    HADDPS  X0, X0
    
    // add vec sum to the scalar sum
    ADDSS   X0, X3
    
    // add input res val
    MOVSS   result+48(FP), X0
    ADDSS   X0, X3
    
    // final result storage
    MOVSS   X3, ret+56(FP)
    RET
