// Peripheral: FLASH_Periph  FLASH Registers.
// Instances:
//  FLASH  mmap.FLASH_R_BASE
// Registers:
//  0x00 32  ACR     Access control register.
//  0x04 32  KEYR    Key register.
//  0x08 32  OPTKEYR Option key register.
//  0x0C 32  SR      Status register.
//  0x10 32  CR      Control register.
//  0x14 32  OPTCR   Option control register.
//  0x18 32  OPTCR1  Option control register 1.
// Import:
//  stm32/o/f411xe/mmap
package flash

const (
	LATENCY       ACR_Bits = 0x0F << 0 //+
	LATENCY_0WS   ACR_Bits = 0x00 << 0
	LATENCY_1WS   ACR_Bits = 0x01 << 0
	LATENCY_2WS   ACR_Bits = 0x02 << 0
	LATENCY_3WS   ACR_Bits = 0x03 << 0
	LATENCY_4WS   ACR_Bits = 0x04 << 0
	LATENCY_5WS   ACR_Bits = 0x05 << 0
	LATENCY_6WS   ACR_Bits = 0x06 << 0
	LATENCY_7WS   ACR_Bits = 0x07 << 0
	LATENCY_8WS   ACR_Bits = 0x08 << 0
	LATENCY_9WS   ACR_Bits = 0x09 << 0
	LATENCY_10WS  ACR_Bits = 0x0A << 0
	LATENCY_11WS  ACR_Bits = 0x0B << 0
	LATENCY_12WS  ACR_Bits = 0x0C << 0
	LATENCY_13WS  ACR_Bits = 0x0D << 0
	LATENCY_14WS  ACR_Bits = 0x0E << 0
	LATENCY_15WS  ACR_Bits = 0x0F << 0
	PRFTEN        ACR_Bits = 0x01 << 8  //+
	ICEN          ACR_Bits = 0x01 << 9  //+
	DCEN          ACR_Bits = 0x01 << 10 //+
	ICRST         ACR_Bits = 0x01 << 11 //+
	DCRST         ACR_Bits = 0x01 << 12 //+
	BYTE0_ADDRESS ACR_Bits = 0x10008F << 10
	BYTE2_ADDRESS ACR_Bits = 0x40023C03 << 0
)

const (
	EOP    SR_Bits = 0x01 << 0  //+
	SOP    SR_Bits = 0x01 << 1  //+
	WRPERR SR_Bits = 0x01 << 4  //+
	PGAERR SR_Bits = 0x01 << 5  //+
	PGPERR SR_Bits = 0x01 << 6  //+
	PGSERR SR_Bits = 0x01 << 7  //+
	BSY    SR_Bits = 0x01 << 16 //+
)

const (
	PG      CR_Bits = 0x01 << 0 //+
	SER     CR_Bits = 0x01 << 1 //+
	MER     CR_Bits = 0x01 << 2 //+
	SNB     CR_Bits = 0x1F << 3 //+
	SNB_0   CR_Bits = 0x01 << 3
	SNB_1   CR_Bits = 0x02 << 3
	SNB_2   CR_Bits = 0x04 << 3
	SNB_3   CR_Bits = 0x08 << 3
	SNB_4   CR_Bits = 0x08 << 3
	PSIZE   CR_Bits = 0x03 << 8 //+
	PSIZE_0 CR_Bits = 0x01 << 8
	PSIZE_1 CR_Bits = 0x02 << 8
	MER2    CR_Bits = 0x01 << 15 //+
	STRT    CR_Bits = 0x01 << 16 //+
	EOPIE   CR_Bits = 0x01 << 24 //+
	LOCK    CR_Bits = 0x01 << 31 //+
)

const (
	OPTLOCK    OPTCR_Bits = 0x01 << 0 //+
	OPTSTRT    OPTCR_Bits = 0x01 << 1 //+
	BOR_LEV_0  OPTCR_Bits = 0x01 << 2 //+
	BOR_LEV_1  OPTCR_Bits = 0x01 << 3 //+
	BOR_LEV    OPTCR_Bits = 0x03 << 2
	BFB2       OPTCR_Bits = 0x01 << 4 //+
	WDG_SW     OPTCR_Bits = 0x01 << 5 //+
	nRST_STOP  OPTCR_Bits = 0x01 << 6 //+
	nRST_STDBY OPTCR_Bits = 0x01 << 7 //+
	RDP        OPTCR_Bits = 0xFF << 8 //+
	RDP_0      OPTCR_Bits = 0x01 << 8
	RDP_1      OPTCR_Bits = 0x02 << 8
	RDP_2      OPTCR_Bits = 0x04 << 8
	RDP_3      OPTCR_Bits = 0x08 << 8
	RDP_4      OPTCR_Bits = 0x10 << 8
	RDP_5      OPTCR_Bits = 0x20 << 8
	RDP_6      OPTCR_Bits = 0x40 << 8
	RDP_7      OPTCR_Bits = 0x80 << 8
	nWRP       OPTCR_Bits = 0xFFF << 16 //+
	nWRP_0     OPTCR_Bits = 0x01 << 16
	nWRP_1     OPTCR_Bits = 0x02 << 16
	nWRP_2     OPTCR_Bits = 0x04 << 16
	nWRP_3     OPTCR_Bits = 0x08 << 16
	nWRP_4     OPTCR_Bits = 0x10 << 16
	nWRP_5     OPTCR_Bits = 0x20 << 16
	nWRP_6     OPTCR_Bits = 0x40 << 16
	nWRP_7     OPTCR_Bits = 0x80 << 16
	nWRP_8     OPTCR_Bits = 0x100 << 16
	nWRP_9     OPTCR_Bits = 0x200 << 16
	nWRP_10    OPTCR_Bits = 0x400 << 16
	nWRP_11    OPTCR_Bits = 0x800 << 16
	DB1M       OPTCR_Bits = 0x01 << 30 //+
	SPRMOD     OPTCR_Bits = 0x01 << 31 //+
)

const (
	nWRP    OPTCR1_Bits = 0xFFF << 16 //+
	nWRP_0  OPTCR1_Bits = 0x01 << 16
	nWRP_1  OPTCR1_Bits = 0x02 << 16
	nWRP_2  OPTCR1_Bits = 0x04 << 16
	nWRP_3  OPTCR1_Bits = 0x08 << 16
	nWRP_4  OPTCR1_Bits = 0x10 << 16
	nWRP_5  OPTCR1_Bits = 0x20 << 16
	nWRP_6  OPTCR1_Bits = 0x40 << 16
	nWRP_7  OPTCR1_Bits = 0x80 << 16
	nWRP_8  OPTCR1_Bits = 0x100 << 16
	nWRP_9  OPTCR1_Bits = 0x200 << 16
	nWRP_10 OPTCR1_Bits = 0x400 << 16
	nWRP_11 OPTCR1_Bits = 0x800 << 16
)
