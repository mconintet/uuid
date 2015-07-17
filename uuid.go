package uuid

// Details see [rfc4122](https://tools.ietf.org/html/rfc4122)

// 0                   1                   2                   3
// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                          time_low                             |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |       time_mid                |         time_hi_and_version   |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |clk_seq_hi_res |  clk_seq_low  |         node (0-1)            |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                         node (2-5)                            |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

// Field                  Data Type     Octet  Note
//                                      #
//
// time_low               unsigned 32   0-3    The low field of the
//                        bit integer          timestamp
//
// time_mid               unsigned 16   4-5    The middle field of the
//                        bit integer          timestamp
//
// time_hi_and_version    unsigned 16   6-7    The high field of the
//                        bit integer          timestamp multiplexed
//                                             with the version number
//
// clock_seq_hi_and_rese  unsigned 8    8      The high field of the
// rved                   bit integer          clock sequence
//                                             multiplexed with the
//                                             variant
//
// clock_seq_low          unsigned 8    9      The low field of the
//                        bit integer          clock sequence
//
// node                   unsigned 48   10-15  The spatially unique
//                        bit integer          node identifier

import (
	"encoding/hex"
	"fmt"
	"time"
)

const (
	timeOffset = uint64(12219292800 * 1e9)
	variant    = 0x80
	ver1       = 0x10
)

type Uuid struct {
	TimeLow   []byte
	TimeMid   []byte
	TimeHiVer []byte
	ClockSeq  []byte
	Node      []byte
}

func (u *Uuid) String() string {
	return fmt.Sprintf("%s-%s-%s-%s-%s",
		hex.EncodeToString(u.TimeLow),
		hex.EncodeToString(u.TimeMid),
		hex.EncodeToString(u.TimeHiVer),
		hex.EncodeToString(u.ClockSeq),
		hex.EncodeToString(u.Node),
	)
}

func GetTime() int64 {
	return int64((uint64(time.Now().UnixNano()) + timeOffset) / 100)
}
