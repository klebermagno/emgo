// Peripheral: GPIO_Periph  General Purpose I/O.
// Instances:
//  GPIOA  mmap.GPIOA_BASE
//  GPIOB  mmap.GPIOB_BASE
//  GPIOC  mmap.GPIOC_BASE
//  GPIOD  mmap.GPIOD_BASE
//  GPIOE  mmap.GPIOE_BASE
//  GPIOF  mmap.GPIOF_BASE
//  GPIOG  mmap.GPIOG_BASE
// Registers:
//  0x00 32  CRL
//  0x04 32  CRH
//  0x08 32  IDR
//  0x0C 32  ODR
//  0x10 32  BSRR
//  0x14 32  BRR
//  0x18 32  LCKR
// Import:
//  stm32/o/f10x_hd/mmap
package gpio

const (
	MODE    CRL_Bits = 0x33333333 << 0 //+ Port x mode bits.
	MODE0   CRL_Bits = 0x03 << 0       //  MODE0[1:0] bits (Port x mode bits, pin 0).
	MODE0_0 CRL_Bits = 0x01 << 0       //  Bit 0.
	MODE0_1 CRL_Bits = 0x02 << 0       //  Bit 1.
	MODE1   CRL_Bits = 0x30 << 0       //  MODE1[1:0] bits (Port x mode bits, pin 1).
	MODE1_0 CRL_Bits = 0x10 << 0       //  Bit 0.
	MODE1_1 CRL_Bits = 0x20 << 0       //  Bit 1.
	MODE2   CRL_Bits = 0x300 << 0      //  MODE2[1:0] bits (Port x mode bits, pin 2).
	MODE2_0 CRL_Bits = 0x100 << 0      //  Bit 0.
	MODE2_1 CRL_Bits = 0x200 << 0      //  Bit 1.
	MODE3   CRL_Bits = 0x3000 << 0     //  MODE3[1:0] bits (Port x mode bits, pin 3).
	MODE3_0 CRL_Bits = 0x1000 << 0     //  Bit 0.
	MODE3_1 CRL_Bits = 0x2000 << 0     //  Bit 1.
	MODE4   CRL_Bits = 0x30000 << 0    //  MODE4[1:0] bits (Port x mode bits, pin 4).
	MODE4_0 CRL_Bits = 0x10000 << 0    //  Bit 0.
	MODE4_1 CRL_Bits = 0x20000 << 0    //  Bit 1.
	MODE5   CRL_Bits = 0x300000 << 0   //  MODE5[1:0] bits (Port x mode bits, pin 5).
	MODE5_0 CRL_Bits = 0x100000 << 0   //  Bit 0.
	MODE5_1 CRL_Bits = 0x200000 << 0   //  Bit 1.
	MODE6   CRL_Bits = 0x3000000 << 0  //  MODE6[1:0] bits (Port x mode bits, pin 6).
	MODE6_0 CRL_Bits = 0x1000000 << 0  //  Bit 0.
	MODE6_1 CRL_Bits = 0x2000000 << 0  //  Bit 1.
	MODE7   CRL_Bits = 0x30000000 << 0 //  MODE7[1:0] bits (Port x mode bits, pin 7).
	MODE7_0 CRL_Bits = 0x10000000 << 0 //  Bit 0.
	MODE7_1 CRL_Bits = 0x20000000 << 0 //  Bit 1.
	CNF     CRL_Bits = 0x33333333 << 2 //+ Port x configuration bits.
	CNF0    CRL_Bits = 0x03 << 2       //  CNF0[1:0] bits (Port x configuration bits, pin 0).
	CNF0_0  CRL_Bits = 0x01 << 2       //  Bit 0.
	CNF0_1  CRL_Bits = 0x02 << 2       //  Bit 1.
	CNF1    CRL_Bits = 0x30 << 2       //  CNF1[1:0] bits (Port x configuration bits, pin 1).
	CNF1_0  CRL_Bits = 0x10 << 2       //  Bit 0.
	CNF1_1  CRL_Bits = 0x20 << 2       //  Bit 1.
	CNF2    CRL_Bits = 0x300 << 2      //  CNF2[1:0] bits (Port x configuration bits, pin 2).
	CNF2_0  CRL_Bits = 0x100 << 2      //  Bit 0.
	CNF2_1  CRL_Bits = 0x200 << 2      //  Bit 1.
	CNF3    CRL_Bits = 0x3000 << 2     //  CNF3[1:0] bits (Port x configuration bits, pin 3).
	CNF3_0  CRL_Bits = 0x1000 << 2     //  Bit 0.
	CNF3_1  CRL_Bits = 0x2000 << 2     //  Bit 1.
	CNF4    CRL_Bits = 0x30000 << 2    //  CNF4[1:0] bits (Port x configuration bits, pin 4).
	CNF4_0  CRL_Bits = 0x10000 << 2    //  Bit 0.
	CNF4_1  CRL_Bits = 0x20000 << 2    //  Bit 1.
	CNF5    CRL_Bits = 0x300000 << 2   //  CNF5[1:0] bits (Port x configuration bits, pin 5).
	CNF5_0  CRL_Bits = 0x100000 << 2   //  Bit 0.
	CNF5_1  CRL_Bits = 0x200000 << 2   //  Bit 1.
	CNF6    CRL_Bits = 0x3000000 << 2  //  CNF6[1:0] bits (Port x configuration bits, pin 6).
	CNF6_0  CRL_Bits = 0x1000000 << 2  //  Bit 0.
	CNF6_1  CRL_Bits = 0x2000000 << 2  //  Bit 1.
	CNF7    CRL_Bits = 0x30000000 << 2 //  CNF7[1:0] bits (Port x configuration bits, pin 7).
	CNF7_0  CRL_Bits = 0x10000000 << 2 //  Bit 0.
	CNF7_1  CRL_Bits = 0x20000000 << 2 //  Bit 1.
)

const (
	MODE     CRH_Bits = 0x33333333 << 0 //+ Port x mode bits.
	MODE8    CRH_Bits = 0x03 << 0       //  MODE8[1:0] bits (Port x mode bits, pin 8).
	MODE8_0  CRH_Bits = 0x01 << 0       //  Bit 0.
	MODE8_1  CRH_Bits = 0x02 << 0       //  Bit 1.
	MODE9    CRH_Bits = 0x30 << 0       //  MODE9[1:0] bits (Port x mode bits, pin 9).
	MODE9_0  CRH_Bits = 0x10 << 0       //  Bit 0.
	MODE9_1  CRH_Bits = 0x20 << 0       //  Bit 1.
	MODE10   CRH_Bits = 0x300 << 0      //  MODE10[1:0] bits (Port x mode bits, pin 10).
	MODE10_0 CRH_Bits = 0x100 << 0      //  Bit 0.
	MODE10_1 CRH_Bits = 0x200 << 0      //  Bit 1.
	MODE11   CRH_Bits = 0x3000 << 0     //  MODE11[1:0] bits (Port x mode bits, pin 11).
	MODE11_0 CRH_Bits = 0x1000 << 0     //  Bit 0.
	MODE11_1 CRH_Bits = 0x2000 << 0     //  Bit 1.
	MODE12   CRH_Bits = 0x30000 << 0    //  MODE12[1:0] bits (Port x mode bits, pin 12).
	MODE12_0 CRH_Bits = 0x10000 << 0    //  Bit 0.
	MODE12_1 CRH_Bits = 0x20000 << 0    //  Bit 1.
	MODE13   CRH_Bits = 0x300000 << 0   //  MODE13[1:0] bits (Port x mode bits, pin 13).
	MODE13_0 CRH_Bits = 0x100000 << 0   //  Bit 0.
	MODE13_1 CRH_Bits = 0x200000 << 0   //  Bit 1.
	MODE14   CRH_Bits = 0x3000000 << 0  //  MODE14[1:0] bits (Port x mode bits, pin 14).
	MODE14_0 CRH_Bits = 0x1000000 << 0  //  Bit 0.
	MODE14_1 CRH_Bits = 0x2000000 << 0  //  Bit 1.
	MODE15   CRH_Bits = 0x30000000 << 0 //  MODE15[1:0] bits (Port x mode bits, pin 15).
	MODE15_0 CRH_Bits = 0x10000000 << 0 //  Bit 0.
	MODE15_1 CRH_Bits = 0x20000000 << 0 //  Bit 1.
	CNF      CRH_Bits = 0x33333333 << 2 //+ Port x configuration bits.
	CNF8     CRH_Bits = 0x03 << 2       //  CNF8[1:0] bits (Port x configuration bits, pin 8).
	CNF8_0   CRH_Bits = 0x01 << 2       //  Bit 0.
	CNF8_1   CRH_Bits = 0x02 << 2       //  Bit 1.
	CNF9     CRH_Bits = 0x30 << 2       //  CNF9[1:0] bits (Port x configuration bits, pin 9).
	CNF9_0   CRH_Bits = 0x10 << 2       //  Bit 0.
	CNF9_1   CRH_Bits = 0x20 << 2       //  Bit 1.
	CNF10    CRH_Bits = 0x300 << 2      //  CNF10[1:0] bits (Port x configuration bits, pin 10).
	CNF10_0  CRH_Bits = 0x100 << 2      //  Bit 0.
	CNF10_1  CRH_Bits = 0x200 << 2      //  Bit 1.
	CNF11    CRH_Bits = 0x3000 << 2     //  CNF11[1:0] bits (Port x configuration bits, pin 11).
	CNF11_0  CRH_Bits = 0x1000 << 2     //  Bit 0.
	CNF11_1  CRH_Bits = 0x2000 << 2     //  Bit 1.
	CNF12    CRH_Bits = 0x30000 << 2    //  CNF12[1:0] bits (Port x configuration bits, pin 12).
	CNF12_0  CRH_Bits = 0x10000 << 2    //  Bit 0.
	CNF12_1  CRH_Bits = 0x20000 << 2    //  Bit 1.
	CNF13    CRH_Bits = 0x300000 << 2   //  CNF13[1:0] bits (Port x configuration bits, pin 13).
	CNF13_0  CRH_Bits = 0x100000 << 2   //  Bit 0.
	CNF13_1  CRH_Bits = 0x200000 << 2   //  Bit 1.
	CNF14    CRH_Bits = 0x3000000 << 2  //  CNF14[1:0] bits (Port x configuration bits, pin 14).
	CNF14_0  CRH_Bits = 0x1000000 << 2  //  Bit 0.
	CNF14_1  CRH_Bits = 0x2000000 << 2  //  Bit 1.
	CNF15    CRH_Bits = 0x30000000 << 2 //  CNF15[1:0] bits (Port x configuration bits, pin 15).
	CNF15_0  CRH_Bits = 0x10000000 << 2 //  Bit 0.
	CNF15_1  CRH_Bits = 0x20000000 << 2 //  Bit 1.
)

const (
	IDR0  IDR_Bits = 0x01 << 0  //+ Port input data, bit 0.
	IDR1  IDR_Bits = 0x01 << 1  //+ Port input data, bit 1.
	IDR2  IDR_Bits = 0x01 << 2  //+ Port input data, bit 2.
	IDR3  IDR_Bits = 0x01 << 3  //+ Port input data, bit 3.
	IDR4  IDR_Bits = 0x01 << 4  //+ Port input data, bit 4.
	IDR5  IDR_Bits = 0x01 << 5  //+ Port input data, bit 5.
	IDR6  IDR_Bits = 0x01 << 6  //+ Port input data, bit 6.
	IDR7  IDR_Bits = 0x01 << 7  //+ Port input data, bit 7.
	IDR8  IDR_Bits = 0x01 << 8  //+ Port input data, bit 8.
	IDR9  IDR_Bits = 0x01 << 9  //+ Port input data, bit 9.
	IDR10 IDR_Bits = 0x01 << 10 //+ Port input data, bit 10.
	IDR11 IDR_Bits = 0x01 << 11 //+ Port input data, bit 11.
	IDR12 IDR_Bits = 0x01 << 12 //+ Port input data, bit 12.
	IDR13 IDR_Bits = 0x01 << 13 //+ Port input data, bit 13.
	IDR14 IDR_Bits = 0x01 << 14 //+ Port input data, bit 14.
	IDR15 IDR_Bits = 0x01 << 15 //+ Port input data, bit 15.
)

const (
	ODR0  ODR_Bits = 0x01 << 0  //+ Port output data, bit 0.
	ODR1  ODR_Bits = 0x01 << 1  //+ Port output data, bit 1.
	ODR2  ODR_Bits = 0x01 << 2  //+ Port output data, bit 2.
	ODR3  ODR_Bits = 0x01 << 3  //+ Port output data, bit 3.
	ODR4  ODR_Bits = 0x01 << 4  //+ Port output data, bit 4.
	ODR5  ODR_Bits = 0x01 << 5  //+ Port output data, bit 5.
	ODR6  ODR_Bits = 0x01 << 6  //+ Port output data, bit 6.
	ODR7  ODR_Bits = 0x01 << 7  //+ Port output data, bit 7.
	ODR8  ODR_Bits = 0x01 << 8  //+ Port output data, bit 8.
	ODR9  ODR_Bits = 0x01 << 9  //+ Port output data, bit 9.
	ODR10 ODR_Bits = 0x01 << 10 //+ Port output data, bit 10.
	ODR11 ODR_Bits = 0x01 << 11 //+ Port output data, bit 11.
	ODR12 ODR_Bits = 0x01 << 12 //+ Port output data, bit 12.
	ODR13 ODR_Bits = 0x01 << 13 //+ Port output data, bit 13.
	ODR14 ODR_Bits = 0x01 << 14 //+ Port output data, bit 14.
	ODR15 ODR_Bits = 0x01 << 15 //+ Port output data, bit 15.
)

const (
	BS0  BSRR_Bits = 0x01 << 0  //+ Port x Set bit 0.
	BS1  BSRR_Bits = 0x01 << 1  //+ Port x Set bit 1.
	BS2  BSRR_Bits = 0x01 << 2  //+ Port x Set bit 2.
	BS3  BSRR_Bits = 0x01 << 3  //+ Port x Set bit 3.
	BS4  BSRR_Bits = 0x01 << 4  //+ Port x Set bit 4.
	BS5  BSRR_Bits = 0x01 << 5  //+ Port x Set bit 5.
	BS6  BSRR_Bits = 0x01 << 6  //+ Port x Set bit 6.
	BS7  BSRR_Bits = 0x01 << 7  //+ Port x Set bit 7.
	BS8  BSRR_Bits = 0x01 << 8  //+ Port x Set bit 8.
	BS9  BSRR_Bits = 0x01 << 9  //+ Port x Set bit 9.
	BS10 BSRR_Bits = 0x01 << 10 //+ Port x Set bit 10.
	BS11 BSRR_Bits = 0x01 << 11 //+ Port x Set bit 11.
	BS12 BSRR_Bits = 0x01 << 12 //+ Port x Set bit 12.
	BS13 BSRR_Bits = 0x01 << 13 //+ Port x Set bit 13.
	BS14 BSRR_Bits = 0x01 << 14 //+ Port x Set bit 14.
	BS15 BSRR_Bits = 0x01 << 15 //+ Port x Set bit 15.
	BR0  BSRR_Bits = 0x01 << 16 //+ Port x Reset bit 0.
	BR1  BSRR_Bits = 0x01 << 17 //+ Port x Reset bit 1.
	BR2  BSRR_Bits = 0x01 << 18 //+ Port x Reset bit 2.
	BR3  BSRR_Bits = 0x01 << 19 //+ Port x Reset bit 3.
	BR4  BSRR_Bits = 0x01 << 20 //+ Port x Reset bit 4.
	BR5  BSRR_Bits = 0x01 << 21 //+ Port x Reset bit 5.
	BR6  BSRR_Bits = 0x01 << 22 //+ Port x Reset bit 6.
	BR7  BSRR_Bits = 0x01 << 23 //+ Port x Reset bit 7.
	BR8  BSRR_Bits = 0x01 << 24 //+ Port x Reset bit 8.
	BR9  BSRR_Bits = 0x01 << 25 //+ Port x Reset bit 9.
	BR10 BSRR_Bits = 0x01 << 26 //+ Port x Reset bit 10.
	BR11 BSRR_Bits = 0x01 << 27 //+ Port x Reset bit 11.
	BR12 BSRR_Bits = 0x01 << 28 //+ Port x Reset bit 12.
	BR13 BSRR_Bits = 0x01 << 29 //+ Port x Reset bit 13.
	BR14 BSRR_Bits = 0x01 << 30 //+ Port x Reset bit 14.
	BR15 BSRR_Bits = 0x01 << 31 //+ Port x Reset bit 15.
)

const (
	BR0  BRR_Bits = 0x01 << 0  //+ Port x Reset bit 0.
	BR1  BRR_Bits = 0x01 << 1  //+ Port x Reset bit 1.
	BR2  BRR_Bits = 0x01 << 2  //+ Port x Reset bit 2.
	BR3  BRR_Bits = 0x01 << 3  //+ Port x Reset bit 3.
	BR4  BRR_Bits = 0x01 << 4  //+ Port x Reset bit 4.
	BR5  BRR_Bits = 0x01 << 5  //+ Port x Reset bit 5.
	BR6  BRR_Bits = 0x01 << 6  //+ Port x Reset bit 6.
	BR7  BRR_Bits = 0x01 << 7  //+ Port x Reset bit 7.
	BR8  BRR_Bits = 0x01 << 8  //+ Port x Reset bit 8.
	BR9  BRR_Bits = 0x01 << 9  //+ Port x Reset bit 9.
	BR10 BRR_Bits = 0x01 << 10 //+ Port x Reset bit 10.
	BR11 BRR_Bits = 0x01 << 11 //+ Port x Reset bit 11.
	BR12 BRR_Bits = 0x01 << 12 //+ Port x Reset bit 12.
	BR13 BRR_Bits = 0x01 << 13 //+ Port x Reset bit 13.
	BR14 BRR_Bits = 0x01 << 14 //+ Port x Reset bit 14.
	BR15 BRR_Bits = 0x01 << 15 //+ Port x Reset bit 15.
)

const (
	LCK0  LCKR_Bits = 0x01 << 0  //+ Port x Lock bit 0.
	LCK1  LCKR_Bits = 0x01 << 1  //+ Port x Lock bit 1.
	LCK2  LCKR_Bits = 0x01 << 2  //+ Port x Lock bit 2.
	LCK3  LCKR_Bits = 0x01 << 3  //+ Port x Lock bit 3.
	LCK4  LCKR_Bits = 0x01 << 4  //+ Port x Lock bit 4.
	LCK5  LCKR_Bits = 0x01 << 5  //+ Port x Lock bit 5.
	LCK6  LCKR_Bits = 0x01 << 6  //+ Port x Lock bit 6.
	LCK7  LCKR_Bits = 0x01 << 7  //+ Port x Lock bit 7.
	LCK8  LCKR_Bits = 0x01 << 8  //+ Port x Lock bit 8.
	LCK9  LCKR_Bits = 0x01 << 9  //+ Port x Lock bit 9.
	LCK10 LCKR_Bits = 0x01 << 10 //+ Port x Lock bit 10.
	LCK11 LCKR_Bits = 0x01 << 11 //+ Port x Lock bit 11.
	LCK12 LCKR_Bits = 0x01 << 12 //+ Port x Lock bit 12.
	LCK13 LCKR_Bits = 0x01 << 13 //+ Port x Lock bit 13.
	LCK14 LCKR_Bits = 0x01 << 14 //+ Port x Lock bit 14.
	LCK15 LCKR_Bits = 0x01 << 15 //+ Port x Lock bit 15.
	LCKK  LCKR_Bits = 0x01 << 16 //+ Lock key.
)
