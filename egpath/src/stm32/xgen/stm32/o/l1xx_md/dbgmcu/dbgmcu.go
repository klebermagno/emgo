// Peripheral: DBGMCU_Periph  Debug MCU.
// Instances:
//  DBGMCU  mmap.DBGMCU_BASE
// Registers:
//  0x00 32  IDCODE MCU device ID code.
//  0x04 32  CR     Debug MCU configuration register.
//  0x08 32  APB1FZ Debug MCU APB1 freeze register.
//  0x0C 32  APB2FZ Debug MCU APB2 freeze register.
// Import:
//  stm32/o/l1xx_md/mmap
package dbgmcu

const (
	DEV_ID    IDCODE_Bits = 0xFFF << 0   //+ Device Identifier.
	REV_ID    IDCODE_Bits = 0xFFFF << 16 //+ REV_ID[15:0] bits (Revision Identifier).
	REV_ID_0  IDCODE_Bits = 0x01 << 16   //  Bit 0.
	REV_ID_1  IDCODE_Bits = 0x02 << 16   //  Bit 1.
	REV_ID_2  IDCODE_Bits = 0x04 << 16   //  Bit 2.
	REV_ID_3  IDCODE_Bits = 0x08 << 16   //  Bit 3.
	REV_ID_4  IDCODE_Bits = 0x10 << 16   //  Bit 4.
	REV_ID_5  IDCODE_Bits = 0x20 << 16   //  Bit 5.
	REV_ID_6  IDCODE_Bits = 0x40 << 16   //  Bit 6.
	REV_ID_7  IDCODE_Bits = 0x80 << 16   //  Bit 7.
	REV_ID_8  IDCODE_Bits = 0x100 << 16  //  Bit 8.
	REV_ID_9  IDCODE_Bits = 0x200 << 16  //  Bit 9.
	REV_ID_10 IDCODE_Bits = 0x400 << 16  //  Bit 10.
	REV_ID_11 IDCODE_Bits = 0x800 << 16  //  Bit 11.
	REV_ID_12 IDCODE_Bits = 0x1000 << 16 //  Bit 12.
	REV_ID_13 IDCODE_Bits = 0x2000 << 16 //  Bit 13.
	REV_ID_14 IDCODE_Bits = 0x4000 << 16 //  Bit 14.
	REV_ID_15 IDCODE_Bits = 0x8000 << 16 //  Bit 15.
)

const (
	DBG_SLEEP    CR_Bits = 0x01 << 0 //+ Debug Sleep Mode.
	DBG_STOP     CR_Bits = 0x01 << 1 //+ Debug Stop Mode.
	DBG_STANDBY  CR_Bits = 0x01 << 2 //+ Debug Standby mode.
	TRACE_IOEN   CR_Bits = 0x01 << 5 //+ Trace Pin Assignment Control.
	TRACE_MODE   CR_Bits = 0x03 << 6 //+ TRACE_MODE[1:0] bits (Trace Pin Assignment Control).
	TRACE_MODE_0 CR_Bits = 0x01 << 6 //  Bit 0.
	TRACE_MODE_1 CR_Bits = 0x02 << 6 //  Bit 1.
)

const (
	DBG_TIM2_STOP          APB1FZ_Bits = 0x01 << 0  //+ TIM2 counter stopped when core is halted.
	DBG_TIM3_STOP          APB1FZ_Bits = 0x01 << 1  //+ TIM3 counter stopped when core is halted.
	DBG_TIM4_STOP          APB1FZ_Bits = 0x01 << 2  //+ TIM4 counter stopped when core is halted.
	DBG_TIM5_STOP          APB1FZ_Bits = 0x01 << 3  //+ TIM5 counter stopped when core is halted.
	DBG_TIM6_STOP          APB1FZ_Bits = 0x01 << 4  //+ TIM6 counter stopped when core is halted.
	DBG_TIM7_STOP          APB1FZ_Bits = 0x01 << 5  //+ TIM7 counter stopped when core is halted.
	DBG_RTC_STOP           APB1FZ_Bits = 0x01 << 10 //+ RTC Counter stopped when Core is halted.
	DBG_WWDG_STOP          APB1FZ_Bits = 0x01 << 11 //+ Debug Window Watchdog stopped when Core is halted.
	DBG_IWDG_STOP          APB1FZ_Bits = 0x01 << 12 //+ Debug Independent Watchdog stopped when Core is halted.
	DBG_I2C1_SMBUS_TIMEOUT APB1FZ_Bits = 0x01 << 21 //+ SMBUS timeout mode stopped when Core is halted.
	DBG_I2C2_SMBUS_TIMEOUT APB1FZ_Bits = 0x01 << 22 //+ SMBUS timeout mode stopped when Core is halted.
)

const (
	DBG_TIM9_STOP  APB2FZ_Bits = 0x01 << 2 //+ TIM9 counter stopped when core is halted.
	DBG_TIM10_STOP APB2FZ_Bits = 0x01 << 3 //+ TIM10 counter stopped when core is halted.
	DBG_TIM11_STOP APB2FZ_Bits = 0x01 << 4 //+ TIM11 counter stopped when core is halted.
)
