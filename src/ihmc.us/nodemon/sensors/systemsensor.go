package sensors

import (
	log "github.com/Sirupsen/logrus"
	"ihmc.us/nodemon/measure"
	"github.com/cloudfoundry/gosigar"
	"time"
	"ihmc.us/nodemon/sensors/system"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/dustin/go-humanize"
	"ihmc.us/nodemon/subjects/memory"
	"ihmc.us/nodemon/util"
	"ihmc.us/nodemon/subjects/cpu"
	"strconv"
)

// fetch hardware statistics from golang sigar library
type SystemSensor struct {
	cpu      chan<- *measure.Measure
	memory   chan<- *measure.Measure
	logDebug bool
	tag      string
}

func NewSystemSensor(cpu, memory chan<- *measure.Measure, logDebug bool) *SystemSensor {

	ss := new(SystemSensor)
	ss.cpu = cpu
	ss.memory = memory
	ss.logDebug = logDebug
	ss.tag = "SystemSensor"
	return ss
}

// start goroutines of SystemSensor
func (ss *SystemSensor) Start() {
	go ss.CPUPoll()
	go ss.MemoryPoll()
}

func (ss *SystemSensor) CPUPoll() {

	sigCpu := sigar.Cpu{}

	for {
		sigCpu.Get()
		m := &measure.Measure{
			Subject:   measure.Subject_cpu,
			Strings:   make(map[string]string),
			Integers:  make(map[string]int64),
			Doubles:   make(map[string]float64),
			Timestamp: &timestamp.Timestamp{Seconds: time.Now().Unix()},
		}

		//TODO fill sensor IP
		m.GetIntegers()[cpu.Int_total.String()] = int64(sigCpu.Total())
		m.GetIntegers()[cpu.Int_user.String()] = int64(sigCpu.User)
		m.GetIntegers()[cpu.Int_nice.String()] = int64(sigCpu.Nice)
		m.GetIntegers()[cpu.Int_sys.String()] = int64(sigCpu.Sys)
		m.GetIntegers()[cpu.Int_idle.String()] = int64(sigCpu.Idle)
		m.GetIntegers()[cpu.Int_wait.String()] = int64(sigCpu.Wait)
		m.GetIntegers()[cpu.Int_irq.String()] = int64(sigCpu.Irq)
		m.GetIntegers()[cpu.Int_softirq.String()] = int64(sigCpu.SoftIrq)
		m.GetIntegers()[cpu.Int_stolen.String()] = int64(sigCpu.Stolen)

		if ss.logDebug {
			log.Debug(ss.TAG() + "CPUPoll produced measure. " +
				" total: " + strconv.Itoa(int(m.GetIntegers()[cpu.Int_total.String()])) +
				" user: " + strconv.Itoa(int(m.GetIntegers()[cpu.Int_user.String()])) +
				" nice: " + strconv.Itoa(int(m.GetIntegers()[cpu.Int_nice.String()])) +
				" sys: " + strconv.Itoa(int(m.GetIntegers()[cpu.Int_sys.String()])))
		}

		//put in channel
		ss.cpu <- m

		time.Sleep(system.DefaultCPUPollingSeconds * time.Second)
	}
}

func (ss *SystemSensor) MemoryPoll() {

	mem := sigar.Mem{}

	for {
		mem.Get()
		m := &measure.Measure{
			Subject:   measure.Subject_memory,
			Strings:   make(map[string]string),
			Integers:  make(map[string]int64),
			Doubles:   make(map[string]float64),
			Timestamp: &timestamp.Timestamp{Seconds: time.Now().Unix()},
		}

		//TODO fill sensor IP
		m.GetIntegers()[memory.Int_total.String()] = int64(mem.Total)
		m.GetIntegers()[memory.Int_free.String()] = int64(mem.Free)
		m.GetIntegers()[memory.Int_used.String()] = int64(mem.Used)
		m.GetIntegers()[memory.Int_actual_free.String()] = int64(mem.ActualFree)
		m.GetIntegers()[memory.Int_actual_used.String()] = int64(mem.ActualUsed)

		if ss.logDebug {
			log.Debug(ss.TAG() + "MemoryPoll produced measure. " +
				" total: " + humanize.Bytes(uint64(m.GetIntegers()[memory.Int_total.String()])) +
				" free: " + humanize.Bytes(uint64(m.GetIntegers()[memory.Int_free.String()])) +
				" used: " + humanize.Bytes(uint64(m.GetIntegers()[memory.Int_used.String()])) +
				" actual_free: " + humanize.Bytes(uint64(m.GetIntegers()[memory.Int_actual_free.String()])) +
				" actual_used: " + humanize.Bytes(uint64(m.GetIntegers()[memory.Int_actual_used.String()])))
		}

		//put in channel
		ss.memory <- m

		time.Sleep(system.DefaultMemoryPollingSeconds * time.Second)
	}
}

func (ss *SystemSensor) Stop() {
	//ns.wg.Done()
}

func (ss *SystemSensor) TAG() string {
	return ss.tag + " [" + util.GetGIDString() + "] "
}
