package sacio

// The length of the header in bytes
const HEADER_LENGTH = 158 * 4

// First significant bit
const (
	MSBFIRST = 0
	LSBFIRST = 1
)

// SAC file format
type (
	F float32 // Floating
	N int32   // Integer
	L bool    // Logical
	I string  // Enumerated
	K string  // Alphanumeric
)

// Header variables
type Header struct {
	Width    int
	Variable string
	DataType string
}

// Enumerated header fields
type Enum struct {
	Index int
	Value string
}

// Timestamp header fields
type Time struct {
	Year int
	Days int
	Hour int
	Min  int
	Sec  int
	Msec int
}

// Contains structed SAC data
type SACData struct {
	// header section
	DELTA    F
	DEPMIN   F
	DEPMAX   F
	SCALE    F
	ODELTA   F
	B        F
	E        F
	O        F
	A        F
	T0       F
	T1       F
	T2       F
	T3       F
	T4       F
	T5       F
	T6       F
	T7       F
	T8       F
	T9       F
	F        F
	RESP0    F
	RESP1    F
	RESP2    F
	RESP3    F
	RESP4    F
	RESP5    F
	RESP6    F
	RESP7    F
	RESP8    F
	RESP9    F
	STLA     F
	STLO     F
	STEL     F
	STDP     F
	EVLA     F
	EVLO     F
	EVEL     F
	EVDP     F
	MAG      F
	USER0    F
	USER1    F
	USER2    F
	USER3    F
	USER4    F
	USER5    F
	USER6    F
	USER7    F
	USER8    F
	USER9    F
	DIST     F
	AZ       F
	BAZ      F
	GCARC    F
	INTERNAL F
	DEPMEN   F
	CMPAZ    F
	CMPINC   F
	XMINIMUM F
	XMAXIMUM F
	YMINIMUM F
	YMAXIMUM F
	UNUSED   F
	NZYEAR   N
	NZJDAY   N
	NZHOUR   N
	NZMIN    N
	NZSEC    N
	NZMSEC   N
	NVHDR    N
	NORID    N
	NEVID    N
	NPTS     N
	NWFID    N
	NXSIZE   N
	NYSIZE   N
	IFTYPE   I
	IDEP     I
	IZTYPE   I
	IINST    I
	ISTREG   I
	IEVREG   I
	IEVTYP   I
	IQUAL    I
	ISYNTH   I
	IMAGTYP  I
	IMAGSRC  I
	LEVEN    L
	LPSPOL   L
	LOVROK   L
	LCALDA   L
	KSTNM    K
	KEVNM    K
	KHOLE    K
	KO       K
	KA       K
	KT0      K
	KT1      K
	KT2      K
	KT3      K
	KT4      K
	KT5      K
	KT6      K
	KT7      K
	KT8      K
	KT9      K
	KF       K
	KUSER0   K
	KUSER1   K
	KUSER2   K
	KCMPNM   K
	KNETWK   K
	KDATRD   K
	KINST    K
	// data section
	Body []F
}

var (
	HEADERS = []Header{
		// 0...4
		{4, "DELTA", "F"},
		{4, "DEPMIN", "F"},
		{4, "DEPMAX", "F"},
		{4, "SCALE", "F"},
		{4, "ODELTA", "F"},
		// 5...9
		{4, "B", "F"},
		{4, "E", "F"},
		{4, "O", "F"},
		{4, "A", "F"},
		{4, "INTERNAL", "F"},
		// 10...14
		{4, "T0", "F"},
		{4, "T1", "F"},
		{4, "T2", "F"},
		{4, "T3", "F"},
		{4, "T4", "F"},
		// 15...19
		{4, "T5", "F"},
		{4, "T6", "F"},
		{4, "T7", "F"},
		{4, "T8", "F"},
		{4, "T9", "F"},
		// 20...24
		{4, "F", "F"},
		{4, "RESP0", "F"},
		{4, "RESP1", "F"},
		{4, "RESP2", "F"},
		{4, "RESP3", "F"},
		// 25...29
		{4, "RESP4", "F"},
		{4, "RESP5", "F"},
		{4, "RESP6", "F"},
		{4, "RESP7", "F"},
		{4, "RESP8", "F"},
		// 30...34
		{4, "RESP9", "F"},
		{4, "STLA", "F"},
		{4, "STLO", "F"},
		{4, "STEL", "F"},
		{4, "STDP", "F"},
		// 35...39
		{4, "EVLA", "F"},
		{4, "EVLO", "F"},
		{4, "EVEL", "F"},
		{4, "EVDP", "F"},
		{4, "MAG", "F"},
		// 40...44
		{4, "USER0", "F"},
		{4, "USER1", "F"},
		{4, "USER2", "F"},
		{4, "USER3", "F"},
		{4, "USER4", "F"},
		// 45...49
		{4, "USER5", "F"},
		{4, "USER6", "F"},
		{4, "USER7", "F"},
		{4, "USER8", "F"},
		{4, "USER9", "F"},
		// 50...54
		{4, "DIST", "F"},
		{4, "AZ", "F"},
		{4, "BAZ", "F"},
		{4, "GCARC", "F"},
		{4, "INTERNAL", "F"},
		// 55...59
		{4, "INTERNAL", "F"},
		{4, "DEPMEN", "F"},
		{4, "CMPAZ", "F"},
		{4, "CMPINC", "F"},
		{4, "XMINIMUM", "F"},
		// 60...64
		{4, "XMAXIMUM", "F"},
		{4, "YMINIMUM", "F"},
		{4, "YMAXIMUM", "F"},
		{4, "UNUSED", "F"},
		{4, "UNUSED", "F"},
		// 65...69
		{4, "UNUSED", "F"},
		{4, "UNUSED", "F"},
		{4, "UNUSED", "F"},
		{4, "UNUSED", "F"},
		{4, "UNUSED", "F"},
		// 70...74
		{4, "NZYEAR", "N"},
		{4, "NZJDAY", "N"},
		{4, "NZHOUR", "N"},
		{4, "NZMIN", "N"},
		{4, "NZSEC", "N"},
		// 75...79
		{4, "NZMSEC", "N"},
		{4, "NVHDR", "N"},
		{4, "NORID", "N"},
		{4, "NEVID", "N"},
		{4, "NPTS", "N"},
		// 80...84
		{4, "INTERNAL", "F"},
		{4, "NWFID", "N"},
		{4, "NXSIZE", "N"},
		{4, "NYSIZE", "N"},
		{4, "UNUSED", "F"},
		// 85...89
		{4, "IFTYPE", "I"},
		{4, "IDEP", "I"},
		{4, "IZTYPE", "I"},
		{4, "UNUSED", "F"},
		{4, "IINST", "I"},
		// 90...94
		{4, "ISTREG", "I"},
		{4, "IEVREG", "I"},
		{4, "IEVTYP", "I"},
		{4, "IQUAL", "I"},
		{4, "ISYNTH", "I"},
		// 95...99
		{4, "IMAGTYP", "I"},
		{4, "IMAGSRC", "I"},
		{4, "UNUSED", "F"},
		{4, "UNUSED", "F"},
		{4, "UNUSED", "F"},
		// 100...104
		{4, "UNUSED", "F"},
		{4, "UNUSED", "F"},
		{4, "UNUSED", "F"},
		{4, "UNUSED", "F"},
		{4, "UNUSED", "F"},
		// 105...109
		{4, "LEVEN", "L"},
		{4, "LPSPOL", "L"},
		{4, "LOVROK", "L"},
		{4, "LCALDA", "L"},
		{4, "UNUSED", "F"},
		// 110...115
		{8, "KSTNM", "K"},
		{16, "KEVNM", "K"},
		// 116...121
		{8, "KHOLE", "K"},
		{8, "KO", "K"},
		{8, "KA", "K"},
		// 122...127
		{8, "KT0", "K"},
		{8, "KT1", "K"},
		{8, "KT2", "K"},
		// 128...133
		{8, "KT3", "K"},
		{8, "KT4", "K"},
		{8, "KT5", "K"},
		// 134...139
		{8, "KT6", "K"},
		{8, "KT7", "K"},
		{8, "KT8", "K"},
		// 140...145
		{8, "KT9", "K"},
		{8, "KF", "K"},
		{8, "KUSER0", "K"},
		// 146...151
		{8, "KUSER1", "K"},
		{8, "KUSER2", "K"},
		{8, "KCMPNM", "K"},
		// 152...157
		{8, "KNETWK", "K"},
		{8, "KDATRD", "K"},
		{8, "KINST", "K"},
	}
	ENUMS = []Enum{
		{1, "itime"},
		{2, "irlim"},
		{3, "iamph"},
		{4, "ixy"},
		{5, "iunkn"},
		{6, "idisp"},
		{7, "ivel"},
		{8, "iacc"},
		{9, "ib"},
		{10, "iday"},
		{11, "io"},
		{12, "ia"},
		{13, "it0"},
		{14, "it1"},
		{15, "it2"},
		{16, "it3"},
		{17, "it4"},
		{18, "it5"},
		{19, "it6"},
		{20, "it7"},
		{21, "it8"},
		{22, "it9"},
		{23, "iradnv"},
		{24, "itannv"},
		{25, "iradev"},
		{26, "itanev"},
		{27, "inorth"},
		{28, "ieast"},
		{29, "ihorza"},
		{30, "idown"},
		{31, "iup"},
		{32, "illlbb"},
		{33, "iwwsn1"},
		{34, "iwwsn2"},
		{35, "ihglp"},
		{36, "isro"},
		{37, "inucl"},
		{38, "ipren"},
		{39, "ipostn"},
		{40, "iquake"},
		{41, "ipreq"},
		{42, "ipostq"},
		{43, "ichem"},
		{44, "iother"},
		{45, "igood"},
		{46, "iglch"},
		{47, "idrop"},
		{48, "ilowsn"},
		{49, "irldta"},
		{50, "ivolts"},
		{52, "imb"},
		{53, "ims"},
		{54, "iml"},
		{55, "imw"},
		{56, "imd"},
		{57, "imx"},
		{58, "ineic"},
		{59, "ipdeq"},
		{60, "ipdew"},
		{61, "ipde"},
		{62, "iisc"},
		{63, "ireb"},
		{64, "iusgs"},
		{65, "ibrk"},
		{66, "icaltech"},
		{67, "illnl"},
		{68, "ievloc"},
		{69, "ijsop"},
		{70, "iuser"},
		{71, "iunknown"},
		{72, "iqb"},
		{73, "iqb1"},
		{74, "iqb2"},
		{75, "iqbx"},
		{76, "iqmt"},
		{77, "ieq"},
		{78, "ieq1"},
		{79, "ieq2"},
		{80, "ime"},
		{81, "iex"},
		{82, "inu"},
		{83, "inc"},
		{84, "io_"},
		{85, "il"},
		{86, "ir"},
		{87, "it"},
		{88, "iu"},
		{89, "ieq3"},
		{90, "ieq0"},
		{91, "iex0"},
		{92, "iqc"},
		{93, "iqb0"},
		{94, "igey"},
		{95, "ilit"},
		{96, "imet"},
		{97, "iodor"},
		{103, "ios"},
	}
)
