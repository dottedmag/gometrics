#include "textflag.h"

// See https://go.dev/doc/asm
// See https://github.com/teh-cmc/go-internals/blob/master/chapter1_assembly_primer/README.md

// empty package before · is current package
// ·cpuTimer(SB) means that ·cpuTimer is an offset from virtual register SB
// NOSPLIT means "do not insert stack-split preamble", we do not need stack
// $0 is the stack size needed
// 8 is the size of arguments+return values. Can be addressed using FP
TEXT ·cpuTimer(SB),NOSPLIT,$0-8
	ISB $1
	MRS CNTVCT_EL0, R0

	MOVD R0, ret+0(FP)
	RET
