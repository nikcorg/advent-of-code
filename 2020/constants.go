package main

// exit codes from sysexits.h
const (
	exOk          = 0  /* successful termination */
	exUsage       = 64 /* command line usage error */
	exDataerr     = 65 /* data format error */
	exNoinput     = 66 /* cannot open input */
	exNouser      = 67 /* addressee unknown */
	exNohost      = 68 /* host name unknown */
	exUnavailable = 69 /* service unavailable */
	exSoftware    = 70 /* internal software error */
	exOserr       = 71 /* system error (e.g., can't fork) */
	exOsfile      = 72 /* critical OS file missing */
	exCantcreat   = 73 /* can't create (user) output file */
	exIoerr       = 74 /* input/output error */
	exTempfail    = 75 /* temp failure; user is invited to retry */
	exProtocol    = 76 /* remote error in protocol */
	exNoperm      = 77 /* permission denied */
	exConfig      = 78 /* configuration error */
)
