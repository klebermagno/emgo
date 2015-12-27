package rtc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f411xe/mmap"
)

type RTC_Periph struct {
	TR      TR
	DR      DR
	CR      CR
	ISR     ISR
	PRER    PRER
	WUTR    WUTR
	CALIBR  CALIBR
	ALRMR   [2]ALRMR
	WPR     WPR
	SSR     SSR
	SHIFTR  SHIFTR
	TSTR    TSTR
	TSDR    TSDR
	TSSSR   TSSSR
	CALR    CALR
	TAFCR   TAFCR
	ALRMSSR [2]ALRMSSR
	_       uint32
	BKPR    [20]BKPR
}

var RTC = (*RTC_Periph)(unsafe.Pointer(uintptr(mmap.RTC_BASE)))

type TR_Bits uint32

func (b TR_Bits) Field(mask TR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TR_Bits) J(v int) TR_Bits {
	return TR_Bits(bits.Make32(v, uint32(mask)))
}

type TR struct{ mmio.U32 }

func (r *TR) Bits(mask TR_Bits) TR_Bits { return TR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TR) StoreBits(mask, b TR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TR) SetBits(mask TR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *TR) ClearBits(mask TR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *TR) Load() TR_Bits             { return TR_Bits(r.U32.Load()) }
func (r *TR) Store(b TR_Bits)           { r.U32.Store(uint32(b)) }

type TR_Mask struct{ mmio.UM32 }

func (rm TR_Mask) Load() TR_Bits   { return TR_Bits(rm.UM32.Load()) }
func (rm TR_Mask) Store(b TR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) PM() TR_Mask {
	return TR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(PM)}}
}

func (p *RTC_Periph) HT() TR_Mask {
	return TR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(HT)}}
}

func (p *RTC_Periph) HU() TR_Mask {
	return TR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(HU)}}
}

func (p *RTC_Periph) MNT() TR_Mask {
	return TR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MNT)}}
}

func (p *RTC_Periph) MNU() TR_Mask {
	return TR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MNU)}}
}

func (p *RTC_Periph) ST() TR_Mask {
	return TR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(ST)}}
}

func (p *RTC_Periph) SU() TR_Mask {
	return TR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(SU)}}
}

type DR_Bits uint32

func (b DR_Bits) Field(mask DR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DR_Bits) J(v int) DR_Bits {
	return DR_Bits(bits.Make32(v, uint32(mask)))
}

type DR struct{ mmio.U32 }

func (r *DR) Bits(mask DR_Bits) DR_Bits { return DR_Bits(r.U32.Bits(uint32(mask))) }
func (r *DR) StoreBits(mask, b DR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DR) SetBits(mask DR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DR) ClearBits(mask DR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DR) Load() DR_Bits             { return DR_Bits(r.U32.Load()) }
func (r *DR) Store(b DR_Bits)           { r.U32.Store(uint32(b)) }

type DR_Mask struct{ mmio.UM32 }

func (rm DR_Mask) Load() DR_Bits   { return DR_Bits(rm.UM32.Load()) }
func (rm DR_Mask) Store(b DR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) YT() DR_Mask {
	return DR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(YT)}}
}

func (p *RTC_Periph) YU() DR_Mask {
	return DR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(YU)}}
}

func (p *RTC_Periph) WDU() DR_Mask {
	return DR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(WDU)}}
}

func (p *RTC_Periph) MT() DR_Mask {
	return DR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(MT)}}
}

func (p *RTC_Periph) MU() DR_Mask {
	return DR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(MU)}}
}

func (p *RTC_Periph) DT() DR_Mask {
	return DR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(DT)}}
}

func (p *RTC_Periph) DU() DR_Mask {
	return DR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(DU)}}
}

type CR_Bits uint32

func (b CR_Bits) Field(mask CR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR_Bits) J(v int) CR_Bits {
	return CR_Bits(bits.Make32(v, uint32(mask)))
}

type CR struct{ mmio.U32 }

func (r *CR) Bits(mask CR_Bits) CR_Bits { return CR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR) StoreBits(mask, b CR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR) SetBits(mask CR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CR) ClearBits(mask CR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CR) Load() CR_Bits             { return CR_Bits(r.U32.Load()) }
func (r *CR) Store(b CR_Bits)           { r.U32.Store(uint32(b)) }

type CR_Mask struct{ mmio.UM32 }

func (rm CR_Mask) Load() CR_Bits   { return CR_Bits(rm.UM32.Load()) }
func (rm CR_Mask) Store(b CR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) COE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(COE)}}
}

func (p *RTC_Periph) OSEL() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(OSEL)}}
}

func (p *RTC_Periph) POL() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(POL)}}
}

func (p *RTC_Periph) COSEL() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(COSEL)}}
}

func (p *RTC_Periph) BCK() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(BCK)}}
}

func (p *RTC_Periph) SUB1H() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(SUB1H)}}
}

func (p *RTC_Periph) ADD1H() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(ADD1H)}}
}

func (p *RTC_Periph) TSIE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(TSIE)}}
}

func (p *RTC_Periph) WUTIE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(WUTIE)}}
}

func (p *RTC_Periph) ALRBIE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(ALRBIE)}}
}

func (p *RTC_Periph) ALRAIE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(ALRAIE)}}
}

func (p *RTC_Periph) TSE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(TSE)}}
}

func (p *RTC_Periph) WUTE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(WUTE)}}
}

func (p *RTC_Periph) ALRBE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(ALRBE)}}
}

func (p *RTC_Periph) ALRAE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(ALRAE)}}
}

func (p *RTC_Periph) DCE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(DCE)}}
}

func (p *RTC_Periph) FMT() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(FMT)}}
}

func (p *RTC_Periph) BYPSHAD() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(BYPSHAD)}}
}

func (p *RTC_Periph) REFCKON() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(REFCKON)}}
}

func (p *RTC_Periph) TSEDGE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(TSEDGE)}}
}

func (p *RTC_Periph) WUCKSEL() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(WUCKSEL)}}
}

type ISR_Bits uint32

func (b ISR_Bits) Field(mask ISR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ISR_Bits) J(v int) ISR_Bits {
	return ISR_Bits(bits.Make32(v, uint32(mask)))
}

type ISR struct{ mmio.U32 }

func (r *ISR) Bits(mask ISR_Bits) ISR_Bits { return ISR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ISR) StoreBits(mask, b ISR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ISR) SetBits(mask ISR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ISR) ClearBits(mask ISR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ISR) Load() ISR_Bits              { return ISR_Bits(r.U32.Load()) }
func (r *ISR) Store(b ISR_Bits)            { r.U32.Store(uint32(b)) }

type ISR_Mask struct{ mmio.UM32 }

func (rm ISR_Mask) Load() ISR_Bits   { return ISR_Bits(rm.UM32.Load()) }
func (rm ISR_Mask) Store(b ISR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) RECALPF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(RECALPF)}}
}

func (p *RTC_Periph) TAMP1F() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(TAMP1F)}}
}

func (p *RTC_Periph) TSOVF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(TSOVF)}}
}

func (p *RTC_Periph) TSF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(TSF)}}
}

func (p *RTC_Periph) WUTF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(WUTF)}}
}

func (p *RTC_Periph) ALRBF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(ALRBF)}}
}

func (p *RTC_Periph) ALRAF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(ALRAF)}}
}

func (p *RTC_Periph) INIT() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(INIT)}}
}

func (p *RTC_Periph) INITF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(INITF)}}
}

func (p *RTC_Periph) RSF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(RSF)}}
}

func (p *RTC_Periph) INITS() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(INITS)}}
}

func (p *RTC_Periph) SHPF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(SHPF)}}
}

func (p *RTC_Periph) WUTWF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(WUTWF)}}
}

func (p *RTC_Periph) ALRBWF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(ALRBWF)}}
}

func (p *RTC_Periph) ALRAWF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(ALRAWF)}}
}

type PRER_Bits uint32

func (b PRER_Bits) Field(mask PRER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PRER_Bits) J(v int) PRER_Bits {
	return PRER_Bits(bits.Make32(v, uint32(mask)))
}

type PRER struct{ mmio.U32 }

func (r *PRER) Bits(mask PRER_Bits) PRER_Bits { return PRER_Bits(r.U32.Bits(uint32(mask))) }
func (r *PRER) StoreBits(mask, b PRER_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PRER) SetBits(mask PRER_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *PRER) ClearBits(mask PRER_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *PRER) Load() PRER_Bits               { return PRER_Bits(r.U32.Load()) }
func (r *PRER) Store(b PRER_Bits)             { r.U32.Store(uint32(b)) }

type PRER_Mask struct{ mmio.UM32 }

func (rm PRER_Mask) Load() PRER_Bits   { return PRER_Bits(rm.UM32.Load()) }
func (rm PRER_Mask) Store(b PRER_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) PREDIV_A() PRER_Mask {
	return PRER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(PREDIV_A)}}
}

func (p *RTC_Periph) PREDIV_S() PRER_Mask {
	return PRER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(PREDIV_S)}}
}

type WUTR_Bits uint32

func (b WUTR_Bits) Field(mask WUTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WUTR_Bits) J(v int) WUTR_Bits {
	return WUTR_Bits(bits.Make32(v, uint32(mask)))
}

type WUTR struct{ mmio.U32 }

func (r *WUTR) Bits(mask WUTR_Bits) WUTR_Bits { return WUTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *WUTR) StoreBits(mask, b WUTR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WUTR) SetBits(mask WUTR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *WUTR) ClearBits(mask WUTR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *WUTR) Load() WUTR_Bits               { return WUTR_Bits(r.U32.Load()) }
func (r *WUTR) Store(b WUTR_Bits)             { r.U32.Store(uint32(b)) }

type WUTR_Mask struct{ mmio.UM32 }

func (rm WUTR_Mask) Load() WUTR_Bits   { return WUTR_Bits(rm.UM32.Load()) }
func (rm WUTR_Mask) Store(b WUTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) WUT() WUTR_Mask {
	return WUTR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(WUT)}}
}

type CALIBR_Bits uint32

func (b CALIBR_Bits) Field(mask CALIBR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CALIBR_Bits) J(v int) CALIBR_Bits {
	return CALIBR_Bits(bits.Make32(v, uint32(mask)))
}

type CALIBR struct{ mmio.U32 }

func (r *CALIBR) Bits(mask CALIBR_Bits) CALIBR_Bits { return CALIBR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CALIBR) StoreBits(mask, b CALIBR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CALIBR) SetBits(mask CALIBR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *CALIBR) ClearBits(mask CALIBR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *CALIBR) Load() CALIBR_Bits                 { return CALIBR_Bits(r.U32.Load()) }
func (r *CALIBR) Store(b CALIBR_Bits)               { r.U32.Store(uint32(b)) }

type CALIBR_Mask struct{ mmio.UM32 }

func (rm CALIBR_Mask) Load() CALIBR_Bits   { return CALIBR_Bits(rm.UM32.Load()) }
func (rm CALIBR_Mask) Store(b CALIBR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) DCS() CALIBR_Mask {
	return CALIBR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(DCS)}}
}

func (p *RTC_Periph) DC() CALIBR_Mask {
	return CALIBR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(DC)}}
}

type ALRMR_Bits uint32

func (b ALRMR_Bits) Field(mask ALRMR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ALRMR_Bits) J(v int) ALRMR_Bits {
	return ALRMR_Bits(bits.Make32(v, uint32(mask)))
}

type ALRMR struct{ mmio.U32 }

func (r *ALRMR) Bits(mask ALRMR_Bits) ALRMR_Bits { return ALRMR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ALRMR) StoreBits(mask, b ALRMR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ALRMR) SetBits(mask ALRMR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *ALRMR) ClearBits(mask ALRMR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *ALRMR) Load() ALRMR_Bits                { return ALRMR_Bits(r.U32.Load()) }
func (r *ALRMR) Store(b ALRMR_Bits)              { r.U32.Store(uint32(b)) }

type ALRMR_Mask struct{ mmio.UM32 }

func (rm ALRMR_Mask) Load() ALRMR_Bits   { return ALRMR_Bits(rm.UM32.Load()) }
func (rm ALRMR_Mask) Store(b ALRMR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) MSK4(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&(*[2]mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28))[n], uint32(MSK4)}}
}

func (p *RTC_Periph) WDSEL(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&(*[2]mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28))[n], uint32(WDSEL)}}
}

func (p *RTC_Periph) DT(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&(*[2]mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28))[n], uint32(DT)}}
}

func (p *RTC_Periph) DU(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&(*[2]mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28))[n], uint32(DU)}}
}

func (p *RTC_Periph) MSK3(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&(*[2]mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28))[n], uint32(MSK3)}}
}

func (p *RTC_Periph) PM(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&(*[2]mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28))[n], uint32(PM)}}
}

func (p *RTC_Periph) HT(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&(*[2]mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28))[n], uint32(HT)}}
}

func (p *RTC_Periph) HU(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&(*[2]mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28))[n], uint32(HU)}}
}

func (p *RTC_Periph) MSK2(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&(*[2]mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28))[n], uint32(MSK2)}}
}

func (p *RTC_Periph) MNT(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&(*[2]mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28))[n], uint32(MNT)}}
}

func (p *RTC_Periph) MNU(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&(*[2]mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28))[n], uint32(MNU)}}
}

func (p *RTC_Periph) MSK1(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&(*[2]mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28))[n], uint32(MSK1)}}
}

func (p *RTC_Periph) ST(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&(*[2]mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28))[n], uint32(ST)}}
}

func (p *RTC_Periph) SU(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&(*[2]mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28))[n], uint32(SU)}}
}

type WPR_Bits uint32

func (b WPR_Bits) Field(mask WPR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WPR_Bits) J(v int) WPR_Bits {
	return WPR_Bits(bits.Make32(v, uint32(mask)))
}

type WPR struct{ mmio.U32 }

func (r *WPR) Bits(mask WPR_Bits) WPR_Bits { return WPR_Bits(r.U32.Bits(uint32(mask))) }
func (r *WPR) StoreBits(mask, b WPR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WPR) SetBits(mask WPR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *WPR) ClearBits(mask WPR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *WPR) Load() WPR_Bits              { return WPR_Bits(r.U32.Load()) }
func (r *WPR) Store(b WPR_Bits)            { r.U32.Store(uint32(b)) }

type WPR_Mask struct{ mmio.UM32 }

func (rm WPR_Mask) Load() WPR_Bits   { return WPR_Bits(rm.UM32.Load()) }
func (rm WPR_Mask) Store(b WPR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) KEY() WPR_Mask {
	return WPR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 36)), uint32(KEY)}}
}

type SSR_Bits uint32

func (b SSR_Bits) Field(mask SSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SSR_Bits) J(v int) SSR_Bits {
	return SSR_Bits(bits.Make32(v, uint32(mask)))
}

type SSR struct{ mmio.U32 }

func (r *SSR) Bits(mask SSR_Bits) SSR_Bits { return SSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *SSR) StoreBits(mask, b SSR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SSR) SetBits(mask SSR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *SSR) ClearBits(mask SSR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *SSR) Load() SSR_Bits              { return SSR_Bits(r.U32.Load()) }
func (r *SSR) Store(b SSR_Bits)            { r.U32.Store(uint32(b)) }

type SSR_Mask struct{ mmio.UM32 }

func (rm SSR_Mask) Load() SSR_Bits   { return SSR_Bits(rm.UM32.Load()) }
func (rm SSR_Mask) Store(b SSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) SS() SSR_Mask {
	return SSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 40)), uint32(SS)}}
}

type SHIFTR_Bits uint32

func (b SHIFTR_Bits) Field(mask SHIFTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SHIFTR_Bits) J(v int) SHIFTR_Bits {
	return SHIFTR_Bits(bits.Make32(v, uint32(mask)))
}

type SHIFTR struct{ mmio.U32 }

func (r *SHIFTR) Bits(mask SHIFTR_Bits) SHIFTR_Bits { return SHIFTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *SHIFTR) StoreBits(mask, b SHIFTR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SHIFTR) SetBits(mask SHIFTR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *SHIFTR) ClearBits(mask SHIFTR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *SHIFTR) Load() SHIFTR_Bits                 { return SHIFTR_Bits(r.U32.Load()) }
func (r *SHIFTR) Store(b SHIFTR_Bits)               { r.U32.Store(uint32(b)) }

type SHIFTR_Mask struct{ mmio.UM32 }

func (rm SHIFTR_Mask) Load() SHIFTR_Bits   { return SHIFTR_Bits(rm.UM32.Load()) }
func (rm SHIFTR_Mask) Store(b SHIFTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) SUBFS() SHIFTR_Mask {
	return SHIFTR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 44)), uint32(SUBFS)}}
}

func (p *RTC_Periph) ADD1S() SHIFTR_Mask {
	return SHIFTR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 44)), uint32(ADD1S)}}
}

type TSTR_Bits uint32

func (b TSTR_Bits) Field(mask TSTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TSTR_Bits) J(v int) TSTR_Bits {
	return TSTR_Bits(bits.Make32(v, uint32(mask)))
}

type TSTR struct{ mmio.U32 }

func (r *TSTR) Bits(mask TSTR_Bits) TSTR_Bits { return TSTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TSTR) StoreBits(mask, b TSTR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TSTR) SetBits(mask TSTR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *TSTR) ClearBits(mask TSTR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *TSTR) Load() TSTR_Bits               { return TSTR_Bits(r.U32.Load()) }
func (r *TSTR) Store(b TSTR_Bits)             { r.U32.Store(uint32(b)) }

type TSTR_Mask struct{ mmio.UM32 }

func (rm TSTR_Mask) Load() TSTR_Bits   { return TSTR_Bits(rm.UM32.Load()) }
func (rm TSTR_Mask) Store(b TSTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) PM() TSTR_Mask {
	return TSTR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 48)), uint32(PM)}}
}

func (p *RTC_Periph) HT() TSTR_Mask {
	return TSTR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 48)), uint32(HT)}}
}

func (p *RTC_Periph) HU() TSTR_Mask {
	return TSTR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 48)), uint32(HU)}}
}

func (p *RTC_Periph) MNT() TSTR_Mask {
	return TSTR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 48)), uint32(MNT)}}
}

func (p *RTC_Periph) MNU() TSTR_Mask {
	return TSTR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 48)), uint32(MNU)}}
}

func (p *RTC_Periph) ST() TSTR_Mask {
	return TSTR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 48)), uint32(ST)}}
}

func (p *RTC_Periph) SU() TSTR_Mask {
	return TSTR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 48)), uint32(SU)}}
}

type TSDR_Bits uint32

func (b TSDR_Bits) Field(mask TSDR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TSDR_Bits) J(v int) TSDR_Bits {
	return TSDR_Bits(bits.Make32(v, uint32(mask)))
}

type TSDR struct{ mmio.U32 }

func (r *TSDR) Bits(mask TSDR_Bits) TSDR_Bits { return TSDR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TSDR) StoreBits(mask, b TSDR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TSDR) SetBits(mask TSDR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *TSDR) ClearBits(mask TSDR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *TSDR) Load() TSDR_Bits               { return TSDR_Bits(r.U32.Load()) }
func (r *TSDR) Store(b TSDR_Bits)             { r.U32.Store(uint32(b)) }

type TSDR_Mask struct{ mmio.UM32 }

func (rm TSDR_Mask) Load() TSDR_Bits   { return TSDR_Bits(rm.UM32.Load()) }
func (rm TSDR_Mask) Store(b TSDR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) WDU() TSDR_Mask {
	return TSDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(WDU)}}
}

func (p *RTC_Periph) MT() TSDR_Mask {
	return TSDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(MT)}}
}

func (p *RTC_Periph) MU() TSDR_Mask {
	return TSDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(MU)}}
}

func (p *RTC_Periph) DT() TSDR_Mask {
	return TSDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(DT)}}
}

func (p *RTC_Periph) DU() TSDR_Mask {
	return TSDR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(DU)}}
}

type TSSSR_Bits uint32

func (b TSSSR_Bits) Field(mask TSSSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TSSSR_Bits) J(v int) TSSSR_Bits {
	return TSSSR_Bits(bits.Make32(v, uint32(mask)))
}

type TSSSR struct{ mmio.U32 }

func (r *TSSSR) Bits(mask TSSSR_Bits) TSSSR_Bits { return TSSSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TSSSR) StoreBits(mask, b TSSSR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TSSSR) SetBits(mask TSSSR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *TSSSR) ClearBits(mask TSSSR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *TSSSR) Load() TSSSR_Bits                { return TSSSR_Bits(r.U32.Load()) }
func (r *TSSSR) Store(b TSSSR_Bits)              { r.U32.Store(uint32(b)) }

type TSSSR_Mask struct{ mmio.UM32 }

func (rm TSSSR_Mask) Load() TSSSR_Bits   { return TSSSR_Bits(rm.UM32.Load()) }
func (rm TSSSR_Mask) Store(b TSSSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) SS() TSSSR_Mask {
	return TSSSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 56)), uint32(SS)}}
}

type CALR_Bits uint32

func (b CALR_Bits) Field(mask CALR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CALR_Bits) J(v int) CALR_Bits {
	return CALR_Bits(bits.Make32(v, uint32(mask)))
}

type CALR struct{ mmio.U32 }

func (r *CALR) Bits(mask CALR_Bits) CALR_Bits { return CALR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CALR) StoreBits(mask, b CALR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CALR) SetBits(mask CALR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CALR) ClearBits(mask CALR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CALR) Load() CALR_Bits               { return CALR_Bits(r.U32.Load()) }
func (r *CALR) Store(b CALR_Bits)             { r.U32.Store(uint32(b)) }

type CALR_Mask struct{ mmio.UM32 }

func (rm CALR_Mask) Load() CALR_Bits   { return CALR_Bits(rm.UM32.Load()) }
func (rm CALR_Mask) Store(b CALR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) CALP() CALR_Mask {
	return CALR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 60)), uint32(CALP)}}
}

func (p *RTC_Periph) CALW8() CALR_Mask {
	return CALR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 60)), uint32(CALW8)}}
}

func (p *RTC_Periph) CALW16() CALR_Mask {
	return CALR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 60)), uint32(CALW16)}}
}

func (p *RTC_Periph) CALM() CALR_Mask {
	return CALR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 60)), uint32(CALM)}}
}

type TAFCR_Bits uint32

func (b TAFCR_Bits) Field(mask TAFCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TAFCR_Bits) J(v int) TAFCR_Bits {
	return TAFCR_Bits(bits.Make32(v, uint32(mask)))
}

type TAFCR struct{ mmio.U32 }

func (r *TAFCR) Bits(mask TAFCR_Bits) TAFCR_Bits { return TAFCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TAFCR) StoreBits(mask, b TAFCR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TAFCR) SetBits(mask TAFCR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *TAFCR) ClearBits(mask TAFCR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *TAFCR) Load() TAFCR_Bits                { return TAFCR_Bits(r.U32.Load()) }
func (r *TAFCR) Store(b TAFCR_Bits)              { r.U32.Store(uint32(b)) }

type TAFCR_Mask struct{ mmio.UM32 }

func (rm TAFCR_Mask) Load() TAFCR_Bits   { return TAFCR_Bits(rm.UM32.Load()) }
func (rm TAFCR_Mask) Store(b TAFCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) ALARMOUTTYPE() TAFCR_Mask {
	return TAFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 64)), uint32(ALARMOUTTYPE)}}
}

func (p *RTC_Periph) TSINSEL() TAFCR_Mask {
	return TAFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 64)), uint32(TSINSEL)}}
}

func (p *RTC_Periph) TAMPINSEL() TAFCR_Mask {
	return TAFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 64)), uint32(TAMPINSEL)}}
}

func (p *RTC_Periph) TAMPPUDIS() TAFCR_Mask {
	return TAFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 64)), uint32(TAMPPUDIS)}}
}

func (p *RTC_Periph) TAMPPRCH() TAFCR_Mask {
	return TAFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 64)), uint32(TAMPPRCH)}}
}

func (p *RTC_Periph) TAMPFLT() TAFCR_Mask {
	return TAFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 64)), uint32(TAMPFLT)}}
}

func (p *RTC_Periph) TAMPFREQ() TAFCR_Mask {
	return TAFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 64)), uint32(TAMPFREQ)}}
}

func (p *RTC_Periph) TAMPTS() TAFCR_Mask {
	return TAFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 64)), uint32(TAMPTS)}}
}

func (p *RTC_Periph) TAMPIE() TAFCR_Mask {
	return TAFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 64)), uint32(TAMPIE)}}
}

func (p *RTC_Periph) TAMP1TRG() TAFCR_Mask {
	return TAFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 64)), uint32(TAMP1TRG)}}
}

func (p *RTC_Periph) TAMP1E() TAFCR_Mask {
	return TAFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 64)), uint32(TAMP1E)}}
}

type ALRMSSR_Bits uint32

func (b ALRMSSR_Bits) Field(mask ALRMSSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ALRMSSR_Bits) J(v int) ALRMSSR_Bits {
	return ALRMSSR_Bits(bits.Make32(v, uint32(mask)))
}

type ALRMSSR struct{ mmio.U32 }

func (r *ALRMSSR) Bits(mask ALRMSSR_Bits) ALRMSSR_Bits { return ALRMSSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ALRMSSR) StoreBits(mask, b ALRMSSR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ALRMSSR) SetBits(mask ALRMSSR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *ALRMSSR) ClearBits(mask ALRMSSR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *ALRMSSR) Load() ALRMSSR_Bits                  { return ALRMSSR_Bits(r.U32.Load()) }
func (r *ALRMSSR) Store(b ALRMSSR_Bits)                { r.U32.Store(uint32(b)) }

type ALRMSSR_Mask struct{ mmio.UM32 }

func (rm ALRMSSR_Mask) Load() ALRMSSR_Bits   { return ALRMSSR_Bits(rm.UM32.Load()) }
func (rm ALRMSSR_Mask) Store(b ALRMSSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) MASKSS(n int) ALRMSSR_Mask {
	return ALRMSSR_Mask{mmio.UM32{&(*[2]mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 68))[n], uint32(MASKSS)}}
}

func (p *RTC_Periph) SS(n int) ALRMSSR_Mask {
	return ALRMSSR_Mask{mmio.UM32{&(*[2]mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 68))[n], uint32(SS)}}
}

type BKPR_Bits uint32

func (b BKPR_Bits) Field(mask BKPR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BKPR_Bits) J(v int) BKPR_Bits {
	return BKPR_Bits(bits.Make32(v, uint32(mask)))
}

type BKPR struct{ mmio.U32 }

func (r *BKPR) Bits(mask BKPR_Bits) BKPR_Bits { return BKPR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BKPR) StoreBits(mask, b BKPR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BKPR) SetBits(mask BKPR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *BKPR) ClearBits(mask BKPR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *BKPR) Load() BKPR_Bits               { return BKPR_Bits(r.U32.Load()) }
func (r *BKPR) Store(b BKPR_Bits)             { r.U32.Store(uint32(b)) }

type BKPR_Mask struct{ mmio.UM32 }

func (rm BKPR_Mask) Load() BKPR_Bits   { return BKPR_Bits(rm.UM32.Load()) }
func (rm BKPR_Mask) Store(b BKPR_Bits) { rm.UM32.Store(uint32(b)) }
