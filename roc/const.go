// WARNING: This file has automatically been generated on Sun, 03 Nov 2019 20:29:42 CET.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package roc

/*
#cgo CFLAGS: -I"../include"
#cgo LDFLAGS: -lroc
#include "../include/roc/address.h"
#include "../include/roc/config.h"
#include "../include/roc/context.h"
#include "../include/roc/frame.h"
#include "../include/roc/log.h"
#include "../include/roc/platform.h"
#include "../include/roc/receiver.h"
#include "../include/roc/sender.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

// Family as declared in roc/address.h:36
type Family int32

// Family enumeration from roc/address.h:36
const (
	AfInvalid Family = -1
	AfAuto    Family = 0
	AfIpv4    Family = 1
	AfIpv6    Family = 2
)

// PortType as declared in roc/config.h:35
type PortType int32

// PortType enumeration from roc/config.h:35
const (
	PortAudioSource PortType = 1
	PortAudioRepair PortType = 2
)

// Protocol as declared in roc/config.h:54
type Protocol int32

// Protocol enumeration from roc/config.h:54
const (
	ProtoRtp           Protocol = 1
	ProtoRtpRs8mSource Protocol = 2
	ProtoRs8mRepair    Protocol = 3
	ProtoRtpLdpcSource Protocol = 4
	ProtoLdpcRepair    Protocol = 5
)

// FecCode as declared in roc/config.h:81
type FecCode int32

// FecCode enumeration from roc/config.h:81
const (
	FecDisable       FecCode = -1
	FecDefault       FecCode = 0
	FecRs8m          FecCode = 1
	FecLdpcStaircase FecCode = 2
)

// PacketEncoding as declared in roc/config.h:91
type PacketEncoding int32

// PacketEncoding enumeration from roc/config.h:91
const (
	PacketEncodingAvpL16 PacketEncoding = 2
)

// FrameEncoding as declared in roc/config.h:100
type FrameEncoding int32

// FrameEncoding enumeration from roc/config.h:100
const (
	FrameEncodingPcmFloat FrameEncoding = 1
)

// ChannelSet as declared in roc/config.h:108
type ChannelSet int32

// ChannelSet enumeration from roc/config.h:108
const (
	ChannelSetStereo ChannelSet = 2
)

// ResamplerProfile as declared in roc/config.h:128
type ResamplerProfile int32

// ResamplerProfile enumeration from roc/config.h:128
const (
	ResamplerDisable ResamplerProfile = -1
	ResamplerDefault ResamplerProfile = 0
	ResamplerHigh    ResamplerProfile = 1
	ResamplerMedium  ResamplerProfile = 2
	ResamplerLow     ResamplerProfile = 3
)

// LogLevel as declared in roc/log.h:53
type LogLevel int32

// LogLevel enumeration from roc/log.h:53
const (
	LogNone  LogLevel = iota
	LogError LogLevel = 1
	LogInfo  LogLevel = 2
	LogDebug LogLevel = 3
	LogTrace LogLevel = 4
)

const (
	// ddressSize as declared in roc/address.h:40
	addressSize = 64
)
