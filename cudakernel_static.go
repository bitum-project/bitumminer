// Copyright (c) 2016 The Bitum developers.

//+build linux,cuda darwin,cuda

package main

/*
#include "bitum.h"
*/
import "C"
import (
	"unsafe"

	"github.com/barnex/cuda5/cu"
)

func cudaPrecomputeTable(input *[192]byte) {
	if input == nil {
		panic("input is nil")
	}
	C.bitum_cpu_setBlock_52((*C.uint32_t)(unsafe.Pointer(input)))
}

func cudaInvokeKernel(gridx, blockx, threads uint32, startNonce uint32, nonceResults cu.DevicePtr, targetHigh uint32) {
	C.bitum_hash_nonce(C.uint32_t(gridx), C.uint32_t(blockx), C.uint32_t(threads),
		C.uint32_t(startNonce), (*C.uint32_t)(unsafe.Pointer(nonceResults)), C.uint32_t(targetHigh))
}
