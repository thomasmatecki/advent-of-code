package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
)

type Receiver interface {
	Recv(src *ReceiverKey, high bool) []Pulse
}

const (
	FlipFlopByte = '%'
	ConjunctByte = '&'
)

var (
	ButtonKey       ReceiverKey = [2]byte{'>', '>'}
	BroadCastKey    ReceiverKey = [2]byte{'.', '.'}
	RxKey           ReceiverKey = [2]byte{'r', 'x'}
	MfKey           ReceiverKey = [2]byte{'m', 'f'}
	MfKeyInputFlips map[ReceiverKey]int
)

type Logger struct {
	rcvCount, hiSndcount, loSndcount int
	rXDone                           bool
}

func init() {
	MfKeyInputFlips = make(map[ReceiverKey]int)
}

func (l *Logger) PulseRepr(srcKey, dstKey *ReceiverKey, high bool) string {

	var hiLoStr string
	if high {
		hiLoStr = "hi"
	} else {
		hiLoStr = "lo"
	}

	if srcKey != nil {
		return fmt.Sprintf("%s -%s-> %s", srcKey[:2], hiLoStr, dstKey[:2])
	} else {
		return fmt.Sprintf(".. -%s-> %s", hiLoStr, dstKey[:2])
	}
}

func (l *Logger) LogReceive(srcKey, dstKey *ReceiverKey, high bool, stateStr string) {
	l.rcvCount += 1
	fmt.Printf("Recv:[%04d] %s : %s \n", l.rcvCount, l.PulseRepr(srcKey, dstKey, high), stateStr)
}

func (l *Logger) LogSend(srcKey, dstKey *ReceiverKey, high bool) {
	if high {
		l.hiSndcount += 1
	} else {
		l.loSndcount += 1
	}
	if *dstKey == RxKey {
		if !high {
			l.rXDone = true
		}
	}
	//fmt.Printf("Send:[%04d] %s \n", l.hiSndcount+l.loSndcount, l.PulseRepr(srcKey, dstKey, high))
}

type ReceiverKey [2]byte
type Network map[ReceiverKey]Receiver

func (n *Network) Get(k string) Receiver {
	var cnKey ReceiverKey = [2]byte{k[0], k[1]}
	r := (*n)[cnKey]
	return r
}

type Pulse struct {
	src  *ReceiverKey
	dst  *Receiver
	high bool
}

func (p *Pulse) Push() []Pulse {
	if *p.dst == nil {
		return []Pulse{}
	}
	return (*p.dst).Recv(p.src, p.high)
}

type Module struct {
	key     *ReceiverKey
	network *Network
	logger  *Logger
	dsts    []ReceiverKey
}

func (m *Module) Send(high bool) []Pulse {
	result := make([]Pulse, 0)

	for _, dstKey := range m.dsts {
		dst := (*m.network)[dstKey]

		m.logger.LogSend(m.key, &dstKey, high)

		result = append(result, Pulse{
			m.key,
			&dst,
			high,
		})
	}
	return result
}

type Broadcaster struct {
	network *Network
	logger  *Logger
	dstKeys []ReceiverKey
}

func (b *Broadcaster) Broadcast(high bool) {
	q := []Pulse{}

	b.logger.LogSend(&ButtonKey, &BroadCastKey, high)
	b.logger.LogReceive(&ButtonKey, &BroadCastKey, high, "button pushed")

	for _, dstKey := range b.dstKeys {
		dst := (*b.network)[dstKey]
		b.logger.LogSend(&BroadCastKey, &dstKey, high)

		q = append(q, Pulse{
			src:  nil,
			high: high,
			dst:  &dst,
		})
	}

	for len(q) > 0 {
		p := q[0]
		nextPulses := p.Push()
		q = append(q[1:], nextPulses...)
	}
}

func (b *Broadcaster) HiLoProduct() int {
	return b.logger.hiSndcount * b.logger.loSndcount
}

func (b *Broadcaster) RxDone() bool {
	return b.logger.rXDone
}

/*
// Flip-flop modules (prefix %) are either on or off; they are initially off. If
// a flip-flop module receives a high pulse, it is ignored and nothing happens.
// However, if a flip-flop module receives a low pulse, it flips between on and
// off. If it was off, it turns on and sends a high pulse. If it was on, it
// turns off and sends a low pulse.
*/
type FlipFlop struct {
	module Module
	on     bool
}

func (ff *FlipFlop) Recv(src *ReceiverKey, high bool) []Pulse {
	var stateStr string

	if ff.on {
		stateStr = "is on"
	} else {
		stateStr = "is off"
	}

	ff.module.logger.LogReceive(src, ff.module.key, high, stateStr)

	if high {
		return []Pulse{}
	} else {
		ff.on = !ff.on
		return ff.module.Send(ff.on)
	}
}

/*
// Conjunction modules (prefix &) remember the type of the most recent pulse
// received from each of their connected input modules; they initially default
// to remembering a low pulse for each input. When a pulse is received, the
// conjunction module first updates its memory for that input. Then, if it
// remembers high pulses for all inputs, it sends a low pulse; otherwise, it
// sends a high pulse.
*/
type Conjunction struct {
	module      Module
	inputStates map[ReceiverKey]bool
}

func (cj *Conjunction) Recv(src *ReceiverKey, high bool) []Pulse {

	cj.module.logger.LogReceive(src, cj.module.key, high, "---")

	_, ok := cj.inputStates[*src]

	if !ok {
		panic("no!")

	}
	//var JfKey ReceiverKey = [2]byte{'j', 'f'}
	//if *cj.module.key == MfKey && *src == JfKey {
	//	if cj.inputStates[*src] != high {
	//		if !high {
	//			fmt.Printf("%s -> mf input high [%08d - %08d] (%d)\n",
	//				src[:2],
	//				MfKeyInputFlips[*src],
	//				cj.module.logger.rcvCount-1,
	//				cj.module.logger.rcvCount-1-MfKeyInputFlips[*src],
	//			)
	//		}
	//	}
	//	MfKeyInputFlips[*src] = cj.module.logger.rcvCount
	//}

	cj.inputStates[*src] = high

	for _, state := range cj.inputStates {
		if !state {
			return cj.module.Send(true)
		}
	}

	if *cj.module.key == MfKey {
		println("found it!")
	}

	return cj.module.Send(false)
}

var moduleExpr = regexp.MustCompile(`([%&])(\w{2})`)

func Initdsts(dstsBytes []byte) (dsts []ReceiverKey) {
	for _, destStr := range bytes.Split(dstsBytes, []byte(", ")) {
		dsts = append(dsts, [2]byte(destStr))
	}
	return
}

func InitBroadcaster(filename string) *Broadcaster {
	input, _ := os.ReadFile(filename)
	network := make(Network)
	logger := new(Logger)

	conjunctions := make([]*Conjunction, 0)
	srcsMap := make(map[ReceiverKey][]ReceiverKey)
	var broadcaster Broadcaster

	for _, line := range bytes.Split(input, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		splitLine := bytes.Split(line, []byte(" -> "))
		moduleBytes, dstsBytes := splitLine[0], splitLine[1]
		dsts := Initdsts(dstsBytes)

		if string(moduleBytes) == "broadcaster" {

			broadcaster = Broadcaster{
				&network,
				logger,
				dsts,
			}
			continue
		}

		moduleMatch := moduleExpr.FindSubmatch(moduleBytes)
		var key ReceiverKey = [2]byte(moduleMatch[2])

		for _, dst := range dsts {
			if srcsMap[dst] == nil {
				srcsMap[dst] = []ReceiverKey{key}
			} else {
				srcsMap[dst] = append(srcsMap[dst], key)
			}
		}

		module := Module{
			&key,
			&network,
			logger,
			dsts,
		}

		if moduleMatch[1][0] == FlipFlopByte {
			network[key] = &FlipFlop{
				module,
				false,
			}
		} else {
			conjunctions = append(conjunctions, &Conjunction{
				module,
				map[ReceiverKey]bool{},
			})
		}
	}

	for _, conjunction := range conjunctions {
		for _, inputKey := range srcsMap[*conjunction.module.key] {
			conjunction.inputStates[inputKey] = false
		}
		network[*conjunction.module.key] = conjunction
	}

	network.Get("cn")

	return &broadcaster
}

func PartOne() {
	broadcaster := InitBroadcaster("20.txt")

	for i := 0; i < 1000; i++ {
		broadcaster.Broadcast(false)
	}
	println("Part One:", broadcaster.HiLoProduct())
}

func PartTwo() {

	//broadcaster := InitBroadcaster("20.txt")
	count := 0
	//	for !broadcaster.RxDone() {
	//		count++
	//		broadcaster.Broadcast(false)
	//	}
	println("Part Two:", count)
}

func main() {
	PartOne()
	PartTwo()
}
