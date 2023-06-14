#include "textflag.h"

// See https://go.dev/doc/asm
// See https://github.com/teh-cmc/go-internals/blob/master/chapter1_assembly_primer/README.md

// "" is current package
// ""·rtdsc(SB) means that ""·rtdsc is a offset from virtual register SB
// NOSPLIT means "do not insert stack-split preamble", we do not need stack
// $0 is the stack size needed
// 8 is the size of arguments+return values. Can be addressed using FP
TEXT ·cpuTimer(SB),NOSPLIT,$0-8
	// RTDSC
	BYTE $0x0F; BYTE $0x31

	SHLQ $32, DX
	ORQ AX, DX
	MOVQ DX, ret+0(FP)
	RET
