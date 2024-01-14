package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
)

type Receiver interface {
	Recv(src *ReceiverKey, high bool) []Pulse
	IsZeroed() bool
}

const (
	FlipFlopByte = '%'
	ConjunctByte = '&'
)

var (
	ButtonKey    ReceiverKey = [2]byte{'>', '>'}
	BroadCastKey ReceiverKey = [2]byte{'.', '.'}
)

type Logger struct {
	id                                               int
	stateChangePushes                                []int
	rcvCount, hiSndcount, loSndcount, broadcastCount int
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
	//fmt.Printf("Recv:[%04d] %s : %s \n", l.rcvCount, l.PulseRepr(srcKey, dstKey, high), stateStr)

}

func (l *Logger) LogSend(srcKey, dstKey *ReceiverKey, high bool) {
	if high {
		l.hiSndcount += 1
	} else {
		l.loSndcount += 1
	}

	//totalSndCount := l.hiSndcount + l.loSndcount
	//fmt.Printf("Send:[%04d] %s \n", totalSndCount, l.PulseRepr(srcKey, dstKey, high))
}

func (l *Logger) LogInputStateChange(
	srcKey *ReceiverKey,
	on bool,
) {
	var state string

	if on {
		state = "On"
	} else {
		state = "Off"
	}

	if on {
		l.stateChangePushes = append(l.stateChangePushes, l.broadcastCount)
	}

	fmt.Printf(
		"(%06d) %s High %s:[%04d] \n",
		l.broadcastCount,
		srcKey[:2],
		state,
		l.rcvCount,
	)
}

type ReceiverKey [2]byte
type Network map[ReceiverKey]Receiver

func (n *Network) IsZeroed() bool {
	for _, recvr := range *n {
		if !recvr.IsZeroed() {
			return false
		}
	}
	return true
}

func (n *Network) Get(k string) Receiver {
	var rk ReceiverKey = [2]byte([]byte(k)[:2])
	return (*n)[rk]
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
	network  *Network
	logger   *Logger
	initKeys []ReceiverKey
	srcsMap  map[ReceiverKey][]ReceiverKey
}

func (b *Broadcaster) Broadcast(high bool) {
	q := []Pulse{}
	b.logger.broadcastCount += 1

	b.logger.LogSend(&ButtonKey, &BroadCastKey, high)
	b.logger.LogReceive(&ButtonKey, &BroadCastKey, high, "button pushed")

	for _, dstKey := range b.initKeys {
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

// Flip-flop modules (prefix %) are either on or off; they are initially off. If
// a flip-flop module receives a high pulse, it is ignored and nothing happens.
// However, if a flip-flop module receives a low pulse, it flips between on and
// off. If it was off, it turns on and sends a high pulse. If it was on, it
// turns off and sends a low pulse.
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

func (ff *FlipFlop) IsZeroed() bool {
	return !ff.on
}

// Conjunction modules (prefix &) remember the type of the most recent pulse
// received from each of their connected input modules; they initially default
// to remembering a low pulse for each input. When a pulse is received, the
// conjunction module first updates its memory for that input. Then, if it
// remembers high pulses for all inputs, it sends a low pulse; otherwise, it
// sends a high pulse.
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

	if *cj.module.key == [2]byte{'m', 'f'} {
		if high && !cj.inputStates[*src] {
			cj.module.logger.LogInputStateChange(src, true)
		}
		if !high && cj.inputStates[*src] {
			cj.module.logger.LogInputStateChange(src, false)
		}
	}

	cj.inputStates[*src] = high

	for _, state := range cj.inputStates {
		if !state {
			return cj.module.Send(true)
		}
	}

	return cj.module.Send(false)
}

func (cj *Conjunction) IsZeroed() bool {
	for _, on := range cj.inputStates {
		if on {
			return false
		}
	}
	return true
}

var moduleExpr = regexp.MustCompile(`([%&])(\w{2})`)

func Initdsts(dstsBytes []byte) (dsts []ReceiverKey) {
	dstSplit := bytes.Split(dstsBytes, []byte(", "))
	if len(dstSplit) == 0 {
		return
	}
	for _, destStr := range dstSplit {
		if len(destStr) != 2 {
			continue
		}
		dsts = append(dsts, [2]byte(destStr))
	}
	return
}

func InitBroadcaster(filename string) *Broadcaster {
	input, _ := os.ReadFile(filename)
	network := make(Network)
	logger := Logger{id: 0}

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
				&logger,
				dsts,
				srcsMap,
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
			&logger,
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

	return &broadcaster
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func PartOne() {
	broadcaster := InitBroadcaster("20.txt")

	for i := 0; i < 1000; i++ {
		broadcaster.Broadcast(false)
	}
	println("Part One:", broadcaster.HiLoProduct())
}

func PartTwo() {

	broadcaster := InitBroadcaster("20.txt")

	for i := 0; i < 5000; i++ {
		broadcaster.Broadcast(false)
	}

	pushLCM := broadcaster.logger.stateChangePushes[0]

	for _, push := range broadcaster.logger.stateChangePushes[1:4] {
		pushLCM = lcm(pushLCM, push)
	}

	println("Part Two:", pushLCM)
}

func main() {
	PartOne()
	PartTwo()
}
