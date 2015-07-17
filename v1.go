package uuid

import (
	"crypto/rand"
	"errors"
	"math/big"
	"net"
	"sync"
)

var v1mutex = &sync.Mutex{}

func NewV1(node net.HardwareAddr) (uuid *Uuid, err error) {
	var (
		inter *net.Interface
		t     int64
		clk   *big.Int
		clkB  []byte
	)

	v1mutex.Lock()
	defer v1mutex.Unlock()

	uuid = new(Uuid)

	if node == nil {
		if inter, err = GetFirstNetInterface(); err == nil {
			uuid.Node = inter.HardwareAddr[0:6]
		} else if err == errNoInterface {
			if uuid.Node, err = GetRandomNode(); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		uuid.Node = node[0:6]
	}

	if clk, err = rand.Prime(rand.Reader, 14); err != nil {
		return nil, err
	}

	t = GetTime()

	uuid.TimeLow = []byte{
		byte(t >> 24),
		byte(t >> 16),
		byte(t >> 8),
		byte(t),
	}

	uuid.TimeMid = []byte{
		byte(t >> 40),
		byte(t >> 32),
	}

	uuid.TimeHiVer = []byte{
		byte(t>>56) | ver1,
		byte(t >> 48),
	}

	clkB = clk.Bytes()
	uuid.ClockSeq = []byte{
		clkB[0] | variant,
		clkB[1],
	}

	return uuid, err
}

var errNoInterface = errors.New("no interface")

func GetFirstNetInterface() (interPtr *net.Interface, err error) {
	var (
		inters []net.Interface
		inter  net.Interface
	)

	if inters, err = net.Interfaces(); err != nil {
		return nil, err
	}

	if len(inters) == 0 {
		return nil, errNoInterface
	}

	for _, inter = range inters {
		if len(inter.HardwareAddr) > 0 {
			break
		}
	}

	if len(inter.HardwareAddr) > 0 {
		return &inter, nil
	}

	return nil, errNoInterface
}

func GetRandomNode() (node []byte, err error) {
	var bi *big.Int

	if bi, err = rand.Prime(rand.Reader, 48); err != nil {
		return nil, err
	}

	return bi.Bytes(), nil
}
