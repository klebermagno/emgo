// Peripheral: CAN_Periph  Controller Area Network.
// Instances:
//  CAN1  mmap.CAN1_BASE
//  CAN2  mmap.CAN2_BASE
// Registers:
//  0x00  32  MCR   Master control register.
//  0x04  32  MSR   Master status register.
//  0x08  32  TSR   Transmit status register.
//  0x0C  32  RF0R  Receive FIFO 0 register.
//  0x10  32  RF1R  Receive FIFO 1 register.
//  0x14  32  IER   Interrupt enable register.
//  0x18  32  ESR   Error status register.
//  0x1C  32  BTR   Bit timing register.
//  0x1B0 32  FMR   Filter master register.
//  0x1B4 32  FM1R  Filter mode register.
//  0x1BC 32  FS1R  Filter scale register.
//  0x1C4 32  FFA1R Filter FIFO assignment register.
//  0x1CC 32  FA1R  Filter activation register.
// Import:
//  stm32/o/f40_41xxx/mmap
package can

const (
	INRQ  MCR_Bits = 0x01 << 0  //+ Initialization Request.
	SLEEP MCR_Bits = 0x01 << 1  //+ Sleep Mode Request.
	TXFP  MCR_Bits = 0x01 << 2  //+ Transmit FIFO Priority.
	RFLM  MCR_Bits = 0x01 << 3  //+ Receive FIFO Locked Mode.
	NART  MCR_Bits = 0x01 << 4  //+ No Automatic Retransmission.
	AWUM  MCR_Bits = 0x01 << 5  //+ Automatic Wakeup Mode.
	ABOM  MCR_Bits = 0x01 << 6  //+ Automatic Bus-Off Management.
	TTCM  MCR_Bits = 0x01 << 7  //+ Time Triggered Communication Mode.
	RESET MCR_Bits = 0x01 << 15 //+ bxCAN software master reset.
)

const (
	INAK  MSR_Bits = 0x01 << 0  //+ Initialization Acknowledge.
	SLAK  MSR_Bits = 0x01 << 1  //+ Sleep Acknowledge.
	ERRI  MSR_Bits = 0x01 << 2  //+ Error Interrupt.
	WKUI  MSR_Bits = 0x01 << 3  //+ Wakeup Interrupt.
	SLAKI MSR_Bits = 0x01 << 4  //+ Sleep Acknowledge Interrupt.
	TXM   MSR_Bits = 0x01 << 8  //+ Transmit Mode.
	RXM   MSR_Bits = 0x01 << 9  //+ Receive Mode.
	SAMP  MSR_Bits = 0x01 << 10 //+ Last Sample Point.
	RX    MSR_Bits = 0x01 << 11 //+ CAN Rx Signal.
)

const (
	RQCP0 TSR_Bits = 0x01 << 0  //+ Request Completed Mailbox0.
	TXOK0 TSR_Bits = 0x01 << 1  //+ Transmission OK of Mailbox0.
	ALST0 TSR_Bits = 0x01 << 2  //+ Arbitration Lost for Mailbox0.
	TERR0 TSR_Bits = 0x01 << 3  //+ Transmission Error of Mailbox0.
	ABRQ0 TSR_Bits = 0x01 << 7  //+ Abort Request for Mailbox0.
	RQCP1 TSR_Bits = 0x01 << 8  //+ Request Completed Mailbox1.
	TXOK1 TSR_Bits = 0x01 << 9  //+ Transmission OK of Mailbox1.
	ALST1 TSR_Bits = 0x01 << 10 //+ Arbitration Lost for Mailbox1.
	TERR1 TSR_Bits = 0x01 << 11 //+ Transmission Error of Mailbox1.
	ABRQ1 TSR_Bits = 0x01 << 15 //+ Abort Request for Mailbox 1.
	RQCP2 TSR_Bits = 0x01 << 16 //+ Request Completed Mailbox2.
	TXOK2 TSR_Bits = 0x01 << 17 //+ Transmission OK of Mailbox 2.
	ALST2 TSR_Bits = 0x01 << 18 //+ Arbitration Lost for mailbox 2.
	TERR2 TSR_Bits = 0x01 << 19 //+ Transmission Error of Mailbox 2.
	ABRQ2 TSR_Bits = 0x01 << 23 //+ Abort Request for Mailbox 2.
	CODE  TSR_Bits = 0x03 << 24 //+ Mailbox Code.
	TME   TSR_Bits = 0x07 << 26 //+ TME[2:0] bits.
	TME0  TSR_Bits = 0x01 << 26 //  Transmit Mailbox 0 Empty.
	TME1  TSR_Bits = 0x02 << 26 //  Transmit Mailbox 1 Empty.
	TME2  TSR_Bits = 0x04 << 26 //  Transmit Mailbox 2 Empty.
	LOW   TSR_Bits = 0x07 << 29 //+ LOW[2:0] bits.
	LOW0  TSR_Bits = 0x01 << 29 //  Lowest Priority Flag for Mailbox 0.
	LOW1  TSR_Bits = 0x02 << 29 //  Lowest Priority Flag for Mailbox 1.
	LOW2  TSR_Bits = 0x04 << 29 //  Lowest Priority Flag for Mailbox 2.
)

const (
	FMP0  RF0R_Bits = 0x03 << 0 //+ FIFO 0 Message Pending.
	FULL0 RF0R_Bits = 0x01 << 3 //+ FIFO 0 Full.
	FOVR0 RF0R_Bits = 0x01 << 4 //+ FIFO 0 Overrun.
	RFOM0 RF0R_Bits = 0x01 << 5 //+ Release FIFO 0 Output Mailbox.
)

const (
	FMP1  RF1R_Bits = 0x03 << 0 //+ FIFO 1 Message Pending.
	FULL1 RF1R_Bits = 0x01 << 3 //+ FIFO 1 Full.
	FOVR1 RF1R_Bits = 0x01 << 4 //+ FIFO 1 Overrun.
	RFOM1 RF1R_Bits = 0x01 << 5 //+ Release FIFO 1 Output Mailbox.
)

const (
	TMEIE  IER_Bits = 0x01 << 0  //+ Transmit Mailbox Empty Interrupt Enable.
	FMPIE0 IER_Bits = 0x01 << 1  //+ FIFO Message Pending Interrupt Enable.
	FFIE0  IER_Bits = 0x01 << 2  //+ FIFO Full Interrupt Enable.
	FOVIE0 IER_Bits = 0x01 << 3  //+ FIFO Overrun Interrupt Enable.
	FMPIE1 IER_Bits = 0x01 << 4  //+ FIFO Message Pending Interrupt Enable.
	FFIE1  IER_Bits = 0x01 << 5  //+ FIFO Full Interrupt Enable.
	FOVIE1 IER_Bits = 0x01 << 6  //+ FIFO Overrun Interrupt Enable.
	EWGIE  IER_Bits = 0x01 << 8  //+ Error Warning Interrupt Enable.
	EPVIE  IER_Bits = 0x01 << 9  //+ Error Passive Interrupt Enable.
	BOFIE  IER_Bits = 0x01 << 10 //+ Bus-Off Interrupt Enable.
	LECIE  IER_Bits = 0x01 << 11 //+ Last Error Code Interrupt Enable.
	ERRIE  IER_Bits = 0x01 << 15 //+ Error Interrupt Enable.
	WKUIE  IER_Bits = 0x01 << 16 //+ Wakeup Interrupt Enable.
	SLKIE  IER_Bits = 0x01 << 17 //+ Sleep Interrupt Enable.
)

const (
	EWGF  ESR_Bits = 0x01 << 0  //+ Error Warning Flag.
	EPVF  ESR_Bits = 0x01 << 1  //+ Error Passive Flag.
	BOFF  ESR_Bits = 0x01 << 2  //+ Bus-Off Flag.
	LEC   ESR_Bits = 0x07 << 4  //+ LEC[2:0] bits (Last Error Code).
	LEC_0 ESR_Bits = 0x01 << 4  //  Bit 0.
	LEC_1 ESR_Bits = 0x02 << 4  //  Bit 1.
	LEC_2 ESR_Bits = 0x04 << 4  //  Bit 2.
	TEC   ESR_Bits = 0xFF << 16 //+ Least significant byte of the 9-bit Transmit Error Counter.
	REC   ESR_Bits = 0xFF << 24 //+ Receive Error Counter.
)

const (
	BRP  BTR_Bits = 0x3FF << 0 //+ Baud Rate Prescaler.
	TS1  BTR_Bits = 0x0F << 16 //+ Time Segment 1.
	TS2  BTR_Bits = 0x07 << 20 //+ Time Segment 2.
	SJW  BTR_Bits = 0x03 << 24 //+ Resynchronization Jump Width.
	LBKM BTR_Bits = 0x01 << 30 //+ Loop Back Mode (Debug).
	SILM BTR_Bits = 0x01 << 31 //+ Silent Mode.
)

const (
	FINIT FMR_Bits = 0x01 << 0 //+ Filter Init Mode.
)

const (
	FBM   FM1R_Bits = 0x3FFF << 0 //+ Filter Mode.
	FBM0  FM1R_Bits = 0x01 << 0   //  Filter Init Mode bit 0.
	FBM1  FM1R_Bits = 0x02 << 0   //  Filter Init Mode bit 1.
	FBM2  FM1R_Bits = 0x04 << 0   //  Filter Init Mode bit 2.
	FBM3  FM1R_Bits = 0x08 << 0   //  Filter Init Mode bit 3.
	FBM4  FM1R_Bits = 0x10 << 0   //  Filter Init Mode bit 4.
	FBM5  FM1R_Bits = 0x20 << 0   //  Filter Init Mode bit 5.
	FBM6  FM1R_Bits = 0x40 << 0   //  Filter Init Mode bit 6.
	FBM7  FM1R_Bits = 0x80 << 0   //  Filter Init Mode bit 7.
	FBM8  FM1R_Bits = 0x100 << 0  //  Filter Init Mode bit 8.
	FBM9  FM1R_Bits = 0x200 << 0  //  Filter Init Mode bit 9.
	FBM10 FM1R_Bits = 0x400 << 0  //  Filter Init Mode bit 10.
	FBM11 FM1R_Bits = 0x800 << 0  //  Filter Init Mode bit 11.
	FBM12 FM1R_Bits = 0x1000 << 0 //  Filter Init Mode bit 12.
	FBM13 FM1R_Bits = 0x2000 << 0 //  Filter Init Mode bit 13.
)

const (
	FSC   FS1R_Bits = 0x3FFF << 0 //+ Filter Scale Configuration.
	FSC0  FS1R_Bits = 0x01 << 0   //  Filter Scale Configuration bit 0.
	FSC1  FS1R_Bits = 0x02 << 0   //  Filter Scale Configuration bit 1.
	FSC2  FS1R_Bits = 0x04 << 0   //  Filter Scale Configuration bit 2.
	FSC3  FS1R_Bits = 0x08 << 0   //  Filter Scale Configuration bit 3.
	FSC4  FS1R_Bits = 0x10 << 0   //  Filter Scale Configuration bit 4.
	FSC5  FS1R_Bits = 0x20 << 0   //  Filter Scale Configuration bit 5.
	FSC6  FS1R_Bits = 0x40 << 0   //  Filter Scale Configuration bit 6.
	FSC7  FS1R_Bits = 0x80 << 0   //  Filter Scale Configuration bit 7.
	FSC8  FS1R_Bits = 0x100 << 0  //  Filter Scale Configuration bit 8.
	FSC9  FS1R_Bits = 0x200 << 0  //  Filter Scale Configuration bit 9.
	FSC10 FS1R_Bits = 0x400 << 0  //  Filter Scale Configuration bit 10.
	FSC11 FS1R_Bits = 0x800 << 0  //  Filter Scale Configuration bit 11.
	FSC12 FS1R_Bits = 0x1000 << 0 //  Filter Scale Configuration bit 12.
	FSC13 FS1R_Bits = 0x2000 << 0 //  Filter Scale Configuration bit 13.
)

const (
	FFA   FFA1R_Bits = 0x3FFF << 0 //+ Filter FIFO Assignment.
	FFA0  FFA1R_Bits = 0x01 << 0   //  Filter FIFO Assignment for Filter 0.
	FFA1  FFA1R_Bits = 0x02 << 0   //  Filter FIFO Assignment for Filter 1.
	FFA2  FFA1R_Bits = 0x04 << 0   //  Filter FIFO Assignment for Filter 2.
	FFA3  FFA1R_Bits = 0x08 << 0   //  Filter FIFO Assignment for Filter 3.
	FFA4  FFA1R_Bits = 0x10 << 0   //  Filter FIFO Assignment for Filter 4.
	FFA5  FFA1R_Bits = 0x20 << 0   //  Filter FIFO Assignment for Filter 5.
	FFA6  FFA1R_Bits = 0x40 << 0   //  Filter FIFO Assignment for Filter 6.
	FFA7  FFA1R_Bits = 0x80 << 0   //  Filter FIFO Assignment for Filter 7.
	FFA8  FFA1R_Bits = 0x100 << 0  //  Filter FIFO Assignment for Filter 8.
	FFA9  FFA1R_Bits = 0x200 << 0  //  Filter FIFO Assignment for Filter 9.
	FFA10 FFA1R_Bits = 0x400 << 0  //  Filter FIFO Assignment for Filter 10.
	FFA11 FFA1R_Bits = 0x800 << 0  //  Filter FIFO Assignment for Filter 11.
	FFA12 FFA1R_Bits = 0x1000 << 0 //  Filter FIFO Assignment for Filter 12.
	FFA13 FFA1R_Bits = 0x2000 << 0 //  Filter FIFO Assignment for Filter 13.
)

const (
	FACT   FA1R_Bits = 0x3FFF << 0 //+ Filter Active.
	FACT0  FA1R_Bits = 0x01 << 0   //  Filter 0 Active.
	FACT1  FA1R_Bits = 0x02 << 0   //  Filter 1 Active.
	FACT2  FA1R_Bits = 0x04 << 0   //  Filter 2 Active.
	FACT3  FA1R_Bits = 0x08 << 0   //  Filter 3 Active.
	FACT4  FA1R_Bits = 0x10 << 0   //  Filter 4 Active.
	FACT5  FA1R_Bits = 0x20 << 0   //  Filter 5 Active.
	FACT6  FA1R_Bits = 0x40 << 0   //  Filter 6 Active.
	FACT7  FA1R_Bits = 0x80 << 0   //  Filter 7 Active.
	FACT8  FA1R_Bits = 0x100 << 0  //  Filter 8 Active.
	FACT9  FA1R_Bits = 0x200 << 0  //  Filter 9 Active.
	FACT10 FA1R_Bits = 0x400 << 0  //  Filter 10 Active.
	FACT11 FA1R_Bits = 0x800 << 0  //  Filter 11 Active.
	FACT12 FA1R_Bits = 0x1000 << 0 //  Filter 12 Active.
	FACT13 FA1R_Bits = 0x2000 << 0 //  Filter 13 Active.
)
