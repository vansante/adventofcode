package assignment

import (
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
)

type Day16 struct{}

const (
	d16PacketTypeLiteral = 4
)

type d16Packet struct {
	ver     int
	typ     int
	lenTyp  bool
	val     int
	packets []*d16Packet
}

func (p *d16Packet) numSubPackets(data string) (int, string) {
	if p.typ == d16PacketTypeLiteral {
		panic("wrong type of packet")
	}
	if !p.lenTyp {
		panic("wrong type mode")
	}
	num, err := strconv.ParseInt(data[:11], 2, 64)
	CheckErr(err)
	return int(num), data[11:]
}

func (p *d16Packet) subPacketBits(data string) (int, string) {
	if p.typ == d16PacketTypeLiteral {
		panic("wrong type of packet")
	}
	if p.lenTyp {
		panic("wrong type mode")
	}
	num, err := strconv.ParseInt(data[:15], 2, 64)
	CheckErr(err)
	return int(num), data[15:]
}

func (p *d16Packet) sumVersions() int64 {
	sum := int64(p.ver)
	for _, sub := range p.packets {
		sum += sub.sumVersions()
	}
	return sum
}

func (p *d16Packet) calculate() int64 {
	if p.typ >= 5 {
		if len(p.packets) != 2 {
			panic("invalid packet count")
		}
	}

	result := int64(0)
	switch p.typ {
	case 0:
		for _, sub := range p.packets {
			result += sub.calculate()
		}
		return result
	case 1:
		if len(p.packets) == 1 {
			return p.packets[0].calculate()
		}
		result = 1
		for _, sub := range p.packets {
			result *= sub.calculate()
		}
		return result
	case 2:
		result = int64(math.MaxInt)
		for _, sub := range p.packets {
			res := sub.calculate()
			if res < result {
				result = res
			}
		}
		return result
	case 3:
		result = int64(math.MinInt)
		for _, sub := range p.packets {
			res := sub.calculate()
			if res > result {
				result = res
			}
		}
		return result
	case 4:
		return int64(p.val)
	case 5:
		if p.packets[0].calculate() > p.packets[1].calculate() {
			return 1
		}
		return 0
	case 6:
		if p.packets[0].calculate() < p.packets[1].calculate() {
			return 1
		}
		return 0
	case 7:
		if p.packets[0].calculate() == p.packets[1].calculate() {
			return 1
		}
		return 0
	}
	panic("invalid packet type")
}

func (d *Day16) setSubPackets(p *d16Packet, data string) string {
	if p.lenTyp {
		var totalPackets int
		totalPackets, data = p.numSubPackets(data)
		for i := 0; i < totalPackets; i++ {
			var pack *d16Packet
			pack, data = d.getPacket(data)
			p.packets = append(p.packets, pack)
		}
		return data
	}

	var bits int
	bits, data = p.subPacketBits(data)
	dataLen := len(data)
	for dataLen-len(data) < bits {
		var pack *d16Packet
		pack, data = d.getPacket(data)
		p.packets = append(p.packets, pack)
	}
	return data
}

func (d *Day16) getPacket(data string) (*d16Packet, string) {
	v, err := strconv.ParseInt(data[:3], 2, 8)
	CheckErr(err)
	data = data[3:]
	tp, err := strconv.ParseInt(data[:3], 2, 8)
	CheckErr(err)
	data = data[3:]

	p := &d16Packet{
		ver: int(v),
		typ: int(tp),
	}

	if tp != d16PacketTypeLiteral {
		p.lenTyp = data[:1] == "1"
		data = data[1:]

		data = d.setSubPackets(p, data)
		return p, data
	}

	nwStr := ""
	for {
		hasNext := data[:1] == "1"
		data = data[1:]
		nwStr = nwStr + data[:4]
		data = data[4:]
		if !hasNext {
			break
		}
	}
	num, err := strconv.ParseInt(nwStr, 2, 64)
	CheckErr(err)

	p.val = int(num)
	return p, data
}

func (d *Day16) getInput(input string) string {
	bytes, err := hex.DecodeString(input)
	CheckErr(err)

	str := ""
	for _, b := range bytes {
		str += fmt.Sprintf("%08b", b)
	}
	return str
}

func (d *Day16) SolveI(input string) int64 {
	dt := d.getInput(input)
	p, _ := d.getPacket(dt)

	return p.sumVersions()
}

func (d *Day16) SolveII(input string) int64 {
	dt := d.getInput(input)
	p, _ := d.getPacket(dt)

	return p.calculate()
}
