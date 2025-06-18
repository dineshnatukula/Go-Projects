// Package 'metrics' is a wrapper package around prometheus and expvar metrics.
// It provides an object of type 'Metric', with a set of methods and properties designed to carry out the metrics API - counters and gages,
// both scalar and grouped into arrays.
// 'Metric' will also serve /metrics (prometheus) and /debug/vars (expvar) URIs if appropriate parameters are provided on init.
//
// 'Metric' can be initialized in various ways.
// If parameter Mux, a pointer to http.ServeMux, is provided - it will use it to attach /metrics and /debug/vars to it.
// If parameter 'Port' is provided on init - it will initialize http server on that Port, and use it for /metrics and /debug/vars.
//
// In addition to the definition of the object 'Metric', package provides a global object M of type 'Metric',
// as well as a set of global wrapper functions so it can be easily accessed from client code.
package metrics

// import (
// 	"bytes"
// 	"errors"
// 	"expvar"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	"reflect"
// 	"regexp"
// 	"runtime"
// 	"strconv"
// 	"strings"
// 	"sync"
// 	"time"

// 	L "gitlab.gametechlabs.net/ppbet/backend-services/logger"

// 	"github.com/prometheus/client_golang/prometheus"
// 	"github.com/prometheus/client_golang/prometheus/promhttp"
// )

// // Gauge encapsulates a gauge, with methods Set(int64), SetBool(bool), Inc(), Dec(), Add(int64)
// type Gauge struct {
// 	e *expvar.Int
// 	p prometheus.Gauge
// }

// // Counter encapsulates a counter
// type Counter struct {
// 	e *expvar.Int
// 	p prometheus.Counter
// }

// // GaugeArray encapsulates an array of gauges, with one or more dimensions.
// // Methods are Set(int64, string...), SetBool(bool, string...), Inc(string...), Dec(string...), Add(int64, string...)
// type GaugeArray struct {
// 	e    *expvar.Map
// 	p    *prometheus.GaugeVec
// 	dim  []string
// 	name string
// }

// // CounterArray encapsulates an array of counters
// type CounterArray struct {
// 	e    *expvar.Map
// 	p    *prometheus.CounterVec
// 	dim  []string
// 	name string
// }

// // Metric is the metrics object
// type Metric struct {
// 	Name string
// 	Port string
// 	Mux  *http.ServeMux

// 	initialized   bool
// 	gauges        map[string]*Gauge
// 	counters      map[string]*Counter
// 	gaugeArrays   map[string]*GaugeArray
// 	counterArrays map[string]*CounterArray

// 	gaugesL        sync.RWMutex
// 	countersL      sync.RWMutex
// 	gaugeArraysL   sync.RWMutex
// 	counterArraysL sync.RWMutex

// 	procStat string
// }

// // M is globally defined metric collector variable that may or may not be used.
// // It should be used by packages, and main.go should decide to initialize it or not.
// // If not initialized - no metrics will be produced by packages.
// var M Metric

// // ErrMetricNotDefined is raised when a metric is used before the Register call
// var ErrMetricNotDefined = errors.New("Metric not defined")

// // ErrMetricNotSetup is raised when something is not set up as it should have (bug)
// //var ErrMetricNotSetup = errors.New("Metric not set up properly")

// // ErrMetricAlreadyDefined is raised when a metric is Registered more than once
// var ErrMetricAlreadyDefined = errors.New("Metric already defined")

// // ErrEmptyKey is raised when one of the keys (dimensions in counterArray or gaugeArray) is empty
// var ErrEmptyKey = errors.New("key is empty")

// // Init sets up given metric.
// // If Port is greater than zero - http server will be started at that port.
// // If Mux is provided - "/metrics" and "/debug/vars" will be attached to it
// func (m *Metric) Init(Name, Port string, Mux *http.ServeMux) {
// 	L.L.Info("metric.Init", L.String("Name", Name), L.String("Port", Port))
// 	m.Name = Name
// 	m.Port = Port
// 	m.Mux = Mux
// 	m.procStat = fmt.Sprintf("/proc/%d/stat", os.Getpid())
// 	m.initialized = true

// 	prometheusMustRegister(m.Name, NewBuiltinCPUMetrics(m.Name, m.procStat))
// 	//prometheus.EnableCollectChecks(true)

// 	if m.gauges == nil {
// 		m.gauges = make(map[string]*Gauge)
// 	} else {
// 		for name := range m.gauges {
// 			m.RegisterGauge(name)
// 		}
// 	}

// 	if m.counters == nil {
// 		m.counters = make(map[string]*Counter)
// 	} else {
// 		for name := range m.counters {
// 			m.RegisterCounter(name)
// 		}
// 	}

// 	if m.gaugeArrays == nil {
// 		m.gaugeArrays = make(map[string]*GaugeArray)
// 	} else {
// 		for name := range m.gaugeArrays {
// 			m.RegisterGaugeArray(name)
// 		}
// 	}

// 	if m.counterArrays == nil {
// 		m.counterArrays = make(map[string]*CounterArray)
// 	} else {
// 		for name, v := range m.counterArrays {
// 			m.RegisterCounterArray(name, v.dim...)
// 		}
// 	}

// 	if m.Port != "" {
// 		if m.Mux == nil {
// 			m.Mux = http.NewServeMux()
// 		}

// 		go func() {
// 			L.L.Info("starting metrics mux", L.String("Port", m.Port))
// 			err := http.ListenAndServe(m.Port, m.Mux)
// 			L.L.Error("metrics mux", L.Error(err))
// 		}()
// 	}

// 	if m.Mux != nil {
// 		// add Prometheus handler
// 		m.Mux.Handle("/metrics", promhttp.Handler())
// 		// add expvar handler in non-DefaultServeMux
// 		// (expvar adds itself into http.DefaultServeMux automatically)
// 		if m.Mux != http.DefaultServeMux {
// 			m.Mux.Handle("/debug/vars", expvar.Handler())
// 		}
// 	}
// }

// // InitAll is NewMetrics + a set of RegisterThis and RegisterThat
// func (m *Metric) InitAll(Name, Port string, Mux *http.ServeMux, Gauges, Counters, GaugeArrays, CounterArrays []string) {
// 	L.L.Info("metric.InitAll", L.String("Name", Name), L.String("Port", Port), L.Strings("Gauges", Gauges), L.Strings("Counters", Counters), L.Strings("GaugeArrays", GaugeArrays), L.Strings("CounterArrays", CounterArrays))

// 	m.Init(Name, Port, Mux)
// 	if Gauges != nil {
// 		m.RegisterGauges(Gauges)
// 	}
// 	if Counters != nil {
// 		m.RegisterCounters(Counters)
// 	}
// 	if GaugeArrays != nil {
// 		var Dimensions []string
// 		GaugeArrays, Dimensions = parseNameArrays(GaugeArrays)
// 		m.RegisterGaugeArrays(GaugeArrays, Dimensions...)
// 	}
// 	if CounterArrays != nil {
// 		var Dimensions []string
// 		CounterArrays, Dimensions = parseNameArrays(CounterArrays)
// 		m.RegisterCounterArrays(CounterArrays, Dimensions...)
// 	}
// }

// /*
// 	Gauge
// 	=====
// */

// // RegisterGauge returns a pointer to gauge. Eventually, it may return ErrMetricAlreadyDefined
// func (m *Metric) RegisterGauge(Name string) (gauge *Gauge, err error) {
// 	L.L.Info("metric.RegisterGauge", L.String("Name", Name), L.Bool("initialized", m.initialized))
// 	var f bool

// 	m.gaugesL.Lock()
// 	if m.gauges == nil {
// 		m.gauges = make(map[string]*Gauge)
// 		gauge = &Gauge{}
// 		m.gauges[Name] = gauge
// 	} else {
// 		gauge, f = m.gauges[Name]
// 		if f {
// 			L.L.Warn("metric.RegisterGauge", L.String("ErrMetricAlreadyDefined", Name))
// 			err = ErrMetricAlreadyDefined
// 		} else {
// 			gauge = &Gauge{}
// 			m.gauges[Name] = gauge
// 		}
// 	}
// 	m.gaugesL.Unlock()

// 	if m.initialized {
// 		if gauge.e == nil {
// 			gauge.e = expvarNewInt(Name)
// 			L.L.Debug("metric.RegisterGauge", L.String("ExpVarNewInt", Name))
// 		}
// 		if gauge.p == nil {
// 			subsystem := "gauge"
// 			if strings.Index(Name, "_") > 0 {
// 				subsystem = ""
// 			}
// 			gauge.p = prometheus.NewGauge(prometheus.GaugeOpts{Namespace: m.Name, Subsystem: subsystem, Name: validName(Name), Help: Name})
// 			err = prometheusMustRegister(Name, gauge.p)
// 			if err != nil {
// 				gauge.p = nil
// 				L.L.Error("metric.RegisterGauge", L.String("prometheus.NewGauge", Name), L.Error(err))
// 			} else {
// 				L.L.Debug("metric.RegisterGauge", L.String("prometheus.NewGauge", Name))
// 			}
// 		}
// 	}

// 	return
// }

// func (m *Metric) findGauge(Name string) (*Gauge, error) {
// 	if !m.initialized {
// 		return nil, nil
// 	}

// 	m.gaugesL.RLock()
// 	i, ok := m.gauges[Name]
// 	m.gaugesL.RUnlock()

// 	if !ok {
// 		return nil, ErrMetricNotDefined
// 	}
// 	if i == nil {
// 		L.L.Warn("Gauge not initialized", L.String("Name", Name))
// 		return nil, ErrMetricNotDefined
// 	}
// 	//if i.e == nil || i.p == nil {
// 	//	L.L.Warn("Gauge not set up", L.String("Name", Name))
// 	//	return nil, ErrMetricNotSetup
// 	//}
// 	return i, nil
// }

// // RegisterGauges registers new gauges
// func (m *Metric) RegisterGauges(Names []string) {
// 	for _, Name := range Names {
// 		m.RegisterGauge(Name)
// 	}
// }

// // GaugeSet sets gauge value
// func (m *Metric) GaugeSet(Name string, Value int64) error {
// 	g, err := m.findGauge(Name)
// 	if g == nil {
// 		return err
// 	}

// 	g.Set(Value)
// 	return nil
// }

// // GaugeSetBool sets gauge value to int64 0 or 1
// func (m *Metric) GaugeSetBool(Name string, Value bool) error {
// 	b := int64(0)
// 	if Value {
// 		b = int64(1)
// 	}
// 	return m.GaugeSet(Name, b)
// }

// // GaugeInc increments gauge value
// func (m *Metric) GaugeInc(Name string) error {
// 	g, err := m.findGauge(Name)
// 	if g == nil {
// 		return err
// 	}

// 	g.Inc()
// 	return nil
// }

// // GaugeDec decrements gauge value
// func (m *Metric) GaugeDec(Name string) error {
// 	g, err := m.findGauge(Name)
// 	if g == nil {
// 		return err
// 	}

// 	g.Dec()
// 	return nil
// }

// // GaugeAdd adds a value to gauge
// func (m *Metric) GaugeAdd(Name string, Value int64) error {
// 	g, err := m.findGauge(Name)
// 	if g == nil || Value == 0 {
// 		return err
// 	}

// 	g.Add(Value)
// 	return nil
// }

// // Set sets gauge value
// func (g *Gauge) Set(Value int64) {
// 	if g == nil {
// 		return
// 	}
// 	if g.e != nil {
// 		g.e.Set(Value)
// 	}
// 	if g.p != nil {
// 		g.p.Set(float64(Value))
// 	}
// }

// // SetBool sets gauge value to int64 0 or 1
// func (g *Gauge) SetBool(Value bool) {
// 	if g == nil {
// 		return
// 	}
// 	b := int64(0)
// 	if Value {
// 		b = int64(1)
// 	}

// 	if g.e != nil {
// 		g.e.Set(b)
// 	}
// 	if g.p != nil {
// 		g.p.Set(float64(b))
// 	}
// }

// // Inc increments gauge value
// func (g *Gauge) Inc() {
// 	if g == nil {
// 		return
// 	}
// 	if g.e != nil {
// 		g.e.Add(1)
// 	}
// 	if g.p != nil {
// 		g.p.Inc()
// 	}
// }

// // Dec decrements gauge value
// func (g *Gauge) Dec() {
// 	if g == nil {
// 		return
// 	}
// 	if g.e != nil {
// 		g.e.Add(-1)
// 	}
// 	if g.p != nil {
// 		g.p.Dec()
// 	}
// }

// // Add adds a value to gauge
// func (g *Gauge) Add(Value int64) {
// 	if g == nil {
// 		return
// 	}
// 	if g.e != nil {
// 		g.e.Add(Value)
// 	}
// 	if g.p != nil {
// 		if Value > 0 {
// 			g.p.Add(float64(Value))
// 		} else {
// 			g.p.Sub(float64(-Value))
// 		}
// 	}
// }

// /*
// 	Counter
// 	=======
// */

// // RegisterCounter registers and returns a Counter. Eventually, it may return ErrMetricAlreadyDefined
// func (m *Metric) RegisterCounter(Name string) (counter *Counter, err error) {
// 	L.L.Info("metric.RegisterCounter", L.String("Name", Name), L.Bool("Initialized", m.initialized))
// 	var f bool

// 	m.countersL.Lock()
// 	if m.counters == nil {
// 		m.counters = make(map[string]*Counter)
// 		counter = &Counter{}
// 		m.counters[Name] = counter
// 	} else {
// 		counter, f = m.counters[Name]
// 		if f {
// 			L.L.Warn("metric.RegisterCounter", L.String("ErrMetricAlreadyDefined", Name))
// 			err = ErrMetricAlreadyDefined
// 		} else {
// 			counter = &Counter{}
// 			m.counters[Name] = counter
// 		}
// 	}
// 	m.countersL.Unlock()

// 	if m.initialized {
// 		if counter.e == nil {
// 			counter.e = expvarNewInt(Name)
// 			L.L.Debug("metric.RegisterCounter", L.String("ExpvarNewInt", Name))
// 		}
// 		if counter.p == nil {
// 			subsystem := "counter"
// 			if strings.Index(Name, "_") > 0 {
// 				subsystem = ""
// 			}
// 			counter.p = prometheus.NewCounter(prometheus.CounterOpts{Namespace: m.Name, Subsystem: subsystem, Name: validName(Name), Help: Name})
// 			err = prometheusMustRegister(Name, counter.p)
// 			if err != nil {
// 				counter.p = nil
// 				L.L.Error("metric.RegisterCounter", L.String("prometheus.NewCounter", Name), L.Error(err))
// 			} else {
// 				L.L.Debug("metric.RegisterCounter", L.String("prometheus.NewCounter", Name))
// 			}
// 		}
// 	}

// 	return
// }

// // RegisterCounters registers new counters
// func (m *Metric) RegisterCounters(Names []string) {
// 	for _, Name := range Names {
// 		m.RegisterCounter(Name)
// 	}
// }

// func (m *Metric) findCounter(Name string) (*Counter, error) {
// 	if !m.initialized {
// 		return nil, nil
// 	}

// 	m.countersL.RLock()
// 	i, ok := m.counters[Name]
// 	m.countersL.RUnlock()

// 	if !ok {
// 		return nil, ErrMetricNotDefined
// 	}
// 	if i == nil {
// 		L.L.Warn("Counter not initialized", L.String("Name", Name))
// 		return nil, ErrMetricNotDefined
// 	}
// 	//if i.e == nil || i.p == nil {
// 	//	L.L.Warn("Counter not set up", L.String("Name", Name))
// 	//	return nil, ErrMetricNotSetup
// 	//}
// 	return i, nil
// }

// // CounterInc Incs counter value
// func (m *Metric) CounterInc(Name string) error {
// 	i, err := m.findCounter(Name)
// 	if i == nil {
// 		return err
// 	}

// 	i.Inc()
// 	return nil
// }

// // CounterAdd Adds counter value
// func (m *Metric) CounterAdd(Name string, Value int64) error {
// 	i, err := m.findCounter(Name)
// 	if i == nil {
// 		return err
// 	}

// 	i.Add(Value)
// 	return nil
// }

// // Inc increments a counter
// func (c *Counter) Inc() {
// 	if c == nil {
// 		return
// 	}
// 	if c.e != nil {
// 		c.e.Add(1)
// 	}
// 	if c.p != nil {
// 		c.p.Inc()
// 	}
// }

// // Add adds value to counter
// func (c *Counter) Add(Value int64) {
// 	if c == nil {
// 		return
// 	}
// 	if c.e != nil {
// 		c.e.Add(Value)
// 	}
// 	if c.p != nil {
// 		c.p.Add(float64(Value))
// 	}
// }

// // AddMS adds milliseconds between Since and Now
// func (c *Counter) AddMS(Since time.Time) {
// 	c.Add((time.Now().UnixNano() - Since.UnixNano()) / int64(time.Millisecond))
// }

// /*
// 	Gauge array
// 	===========
// */

// // RegisterGaugeArray returns pointer to new GaugeArray. Eventually, it may return ErrMetricAlreadyDefined
// func (m *Metric) RegisterGaugeArray(Name string, Dimensions ...string) (gaugeArray *GaugeArray, err error) {
// 	L.L.Info("metric.RegisterGaugeArray", L.String("Name", Name), L.Strings("Dimensions", Dimensions), L.Bool("Initialized", m.initialized))
// 	var f bool

// 	if len(Dimensions) == 0 {
// 		Dimensions = []string{"key"}
// 	}
// 	if Dimensions[len(Dimensions)-1] != "service" {
// 		Dimensions = append(Dimensions, "service")
// 	}

// 	m.gaugeArraysL.Lock()
// 	defer m.gaugeArraysL.Unlock()
// 	if m.gaugeArrays == nil {
// 		m.gaugeArrays = make(map[string]*GaugeArray)
// 		gaugeArray = &GaugeArray{dim: Dimensions}
// 		m.gaugeArrays[Name] = gaugeArray
// 	} else {
// 		gaugeArray, f = m.gaugeArrays[Name]
// 		if f {
// 			if !reflect.DeepEqual(gaugeArray.dim, Dimensions) {
// 				L.L.Warn("metric.RegisterGaugeArray", L.String("ErrMetricAlreadyDefined", Name), L.Strings("Old", gaugeArray.dim), L.Strings("New", Dimensions))
// 				err = ErrMetricAlreadyDefined
// 			}
// 		} else {
// 			gaugeArray = &GaugeArray{dim: Dimensions}
// 			m.gaugeArrays[Name] = gaugeArray
// 		}
// 	}

// 	if m.initialized {
// 		gaugeArray.name = m.Name
// 		if gaugeArray.e == nil {
// 			gaugeArray.e = expvarNewMap(Name)
// 			L.L.Debug("metric.RegisterGaugeArray", L.String("ExpVarNewMap", Name))
// 		}
// 		if gaugeArray.p == nil {
// 			subsystem := "gaugearray"
// 			gaugeArray.p = prometheus.NewGaugeVec(prometheus.GaugeOpts{ /*Namespace: m.Name,*/ Subsystem: subsystem, Name: validName(Name), Help: Name}, gaugeArray.dim)
// 			err = prometheusMustRegister(Name, gaugeArray.p)
// 			if err != nil {
// 				gaugeArray.p = nil
// 				L.L.Error("metric.RegisterGaugeArray", L.String("prometheus.NewGaugeVec", Name), L.Error(err))
// 			} else {
// 				L.L.Debug("metric.RegisterGaugeArray", L.String("prometheus.NewGaugeVec", Name))
// 			}
// 		}
// 	}

// 	return
// }

// // RegisterGaugeArrays registers new gaugeArrays
// func (m *Metric) RegisterGaugeArrays(Names []string, Dimensions ...string) {
// 	L.L.Info("metric.RegisterGaugeArrays", L.Strings("Names", Names), L.Strings("Dimensions", Dimensions), L.Bool("Initialized", m.initialized))
// 	for _, Name := range Names {
// 		m.RegisterGaugeArray(Name, Dimensions...)
// 	}
// }

// func (m *Metric) findGaugeArray(Name string) (*GaugeArray, error) {
// 	if !m.initialized {
// 		return nil, nil
// 	}

// 	m.gaugeArraysL.RLock()
// 	i, ok := m.gaugeArrays[Name]
// 	m.gaugeArraysL.RUnlock()

// 	if !ok {
// 		L.L.Warn("GaugeArray not defined", L.String("Name", Name))
// 		return nil, ErrMetricNotDefined
// 	}
// 	if i == nil {
// 		L.L.Warn("GaugeArray not initialized", L.String("Name", Name))
// 		return nil, ErrMetricNotDefined
// 	}
// 	return i, nil
// }

// // GaugeArraySet sets GaugeArray value
// func (m *Metric) GaugeArraySet(Name string, Value int64, Keys ...string) error {
// 	ga, err := m.findGaugeArray(Name)
// 	if ga == nil {
// 		return err
// 	}

// 	return ga.Set(Value, Keys...)
// }

// // GaugeArrayInc Incs GaugeArray value
// func (m *Metric) GaugeArrayInc(Name string, Keys ...string) error {
// 	ga, err := m.findGaugeArray(Name)
// 	if ga == nil {
// 		return err
// 	}

// 	return ga.Inc(Keys...)
// }

// // GaugeArrayDec Decs GaugeArray value
// func (m *Metric) GaugeArrayDec(Name string, Keys ...string) error {
// 	ga, err := m.findGaugeArray(Name)
// 	if ga == nil {
// 		return err
// 	}

// 	return ga.Dec(Keys...)
// }

// // GaugeArrayAdd Adds GaugeArray value
// func (m *Metric) GaugeArrayAdd(Name string, Value int64, Keys ...string) error {
// 	ga, err := m.findGaugeArray(Name)
// 	if ga == nil || Value == 0 {
// 		return err
// 	}

// 	return ga.Add(Value, Keys...)
// }

// // Set sets one member of GaugeArray, determined by Keys
// func (ga *GaugeArray) Set(Value int64, Keys ...string) error {
// 	if ga == nil {
// 		return nil
// 	}
// 	if err := checkKeys(Keys); err != nil {
// 		L.L.Error("GaugeArray.Set.checkKeys", L.Strings("Keys", Keys), L.Error(err))
// 		return err
// 	}
// 	Keys = append(Keys, ga.name)

// 	ld := len(ga.dim)
// 	lk := len(Keys)

// 	if ld == 1 && lk == 1 {
// 		if ga.e != nil {
// 			//ga.e.Set(Key, Value) - no such method...
// 			ga.e.Delete(Keys[0])
// 			ga.e.Add(Keys[0], Value)
// 		}
// 		if ga.p != nil {
// 			ga.p.With(prometheus.Labels{ga.dim[0]: Keys[0]}).Set(float64(Value))
// 		}
// 	} else {
// 		if ld != lk || ld == 0 {
// 			L.L.Panic("GaugeArray.Set.wrongNumberOfDimensions", L.Strings("declared", ga.dim), L.Strings("given", Keys))
// 		}
// 		names := ""
// 		labels := make(prometheus.Labels)
// 		for i, d := range ga.dim {
// 			names += "_" + Keys[i]
// 			labels[d] = Keys[i]
// 		}
// 		if ga.e != nil {
// 			//ga.e.Set(Key, Value) - no such method...
// 			ga.e.Delete(names[1:])
// 			ga.e.Add(names[1:], Value)
// 		}
// 		if ga.p != nil {
// 			ga.p.With(labels).Set(float64(Value))
// 		}
// 	}

// 	return nil
// }

// // SetBool sets value of one member of GaugeArray, determined by Keys to int64 0 or 1
// func (ga *GaugeArray) SetBool(Value bool, Keys ...string) {
// 	if ga == nil {
// 		return
// 	}
// 	b := int64(0)
// 	if Value {
// 		b = int64(1)
// 	}

// 	ga.Set(b, Keys...)
// }

// // Inc increments one member of GaugeArray, determined by Keys
// func (ga *GaugeArray) Inc(Keys ...string) error {
// 	if ga == nil {
// 		return nil
// 	}
// 	if err := checkKeys(Keys); err != nil {
// 		L.L.Error("GaugeArray.Inc.checkKeys", L.Strings("Keys", Keys), L.Error(err))
// 		return err
// 	}
// 	Keys = append(Keys, ga.name)

// 	ld := len(ga.dim)
// 	lk := len(Keys)

// 	if ld == 1 && lk == 1 {
// 		if ga.e != nil {
// 			//ga.e.Set(Key, Value) - no such method...
// 			ga.e.Add(Keys[0], 1)
// 		}
// 		if ga.p != nil {
// 			ga.p.With(prometheus.Labels{ga.dim[0]: Keys[0]}).Add(1)
// 		}
// 	} else {
// 		if ld != lk || ld == 0 {
// 			L.L.Panic("GaugeArray.Dec.wrongNumberOfDimensions", L.Strings("declared", ga.dim), L.Strings("given", Keys))
// 		}
// 		names := ""
// 		labels := make(prometheus.Labels)
// 		for i, d := range ga.dim {
// 			names += "_" + Keys[i]
// 			labels[d] = Keys[i]
// 		}
// 		if ga.e != nil {
// 			//ga.e.Set(Key, Value) - no such method...
// 			ga.e.Add(names[1:], 1)
// 		}
// 		if ga.p != nil {
// 			ga.p.With(labels).Add(1)
// 		}
// 	}

// 	return nil
// }

// // Dec decrements one member of GaugeArray, determined by Keys
// func (ga *GaugeArray) Dec(Keys ...string) error {
// 	if ga == nil {
// 		return nil
// 	}
// 	if err := checkKeys(Keys); err != nil {
// 		L.L.Error("GaugeArray.Dec.checkKeys", L.Strings("Keys", Keys), L.Error(err))
// 		return err
// 	}
// 	Keys = append(Keys, ga.name)

// 	ld := len(ga.dim)
// 	lk := len(Keys)

// 	if ld == 1 && lk == 1 {
// 		if ga.e != nil {
// 			ga.e.Add(Keys[0], -1)
// 		}
// 		if ga.p != nil {
// 			ga.p.With(prometheus.Labels{ga.dim[0]: Keys[0]}).Sub(1)
// 		}
// 	} else {
// 		if ld != lk || ld == 0 {
// 			L.L.Panic("GaugeArray.Dec.wrongNumberOfDimensions", L.Strings("declared", ga.dim), L.Strings("given", Keys))
// 		}
// 		names := ""
// 		labels := make(prometheus.Labels)
// 		for i, d := range ga.dim {
// 			names += "_" + Keys[i]
// 			labels[d] = Keys[i]
// 		}
// 		if ga.e != nil {
// 			ga.e.Add(names[1:], -1)
// 		}
// 		if ga.p != nil {
// 			ga.p.With(labels).Sub(1)
// 		}
// 	}

// 	return nil
// }

// // Add adds a value to one member of GaugeArray, determined by Keys
// func (ga *GaugeArray) Add(Value int64, Keys ...string) error {
// 	if ga == nil {
// 		return nil
// 	}
// 	if err := checkKeys(Keys); err != nil {
// 		L.L.Error("GaugeArray.Add.checkKeys", L.Strings("Keys", Keys), L.Error(err))
// 		return err
// 	}
// 	Keys = append(Keys, ga.name)

// 	ld := len(ga.dim)
// 	lk := len(Keys)

// 	if ld == 1 && lk == 1 {
// 		if ga.e != nil {
// 			ga.e.Add(Keys[0], Value)
// 		}
// 		if ga.p != nil {
// 			if Value > 0 {
// 				ga.p.With(prometheus.Labels{ga.dim[0]: Keys[0]}).Add(float64(Value))
// 			} else {
// 				ga.p.With(prometheus.Labels{ga.dim[0]: Keys[0]}).Sub(float64(-Value))
// 			}
// 		}
// 	} else {
// 		if ld != lk || ld == 0 {
// 			L.L.Panic("GaugeArray.Add.wrongNumberOfDimensions", L.Strings("declared", ga.dim), L.Strings("given", Keys))
// 		}
// 		names := ""
// 		labels := make(prometheus.Labels)
// 		for i, d := range ga.dim {
// 			names += "_" + Keys[i]
// 			labels[d] = Keys[i]
// 		}
// 		if ga.e != nil {
// 			ga.e.Add(names[1:], Value)
// 		}
// 		if ga.p != nil {
// 			if Value > 0 {
// 				ga.p.With(labels).Add(float64(Value))
// 			} else {
// 				ga.p.With(labels).Sub(float64(-Value))
// 			}
// 		}
// 	}

// 	return nil
// }

// /*
// 	Counter array
// 	=============
// */

// // RegisterCounterArray registers and returns new counterArray. Eventually, it may return ErrMetricAlreadyDefined
// func (m *Metric) RegisterCounterArray(Name string, Dimensions ...string) (counterArray *CounterArray, err error) {
// 	L.L.Info("metric.RegisterCounterArray", L.String("Name", Name), L.Strings("Dimensions", Dimensions), L.Bool("initialized", m.initialized))
// 	var f bool

// 	if len(Dimensions) == 0 {
// 		Dimensions = []string{"key"}
// 	}
// 	if Dimensions[len(Dimensions)-1] != "service" {
// 		Dimensions = append(Dimensions, "service")
// 	}

// 	m.counterArraysL.Lock()
// 	if m.counterArrays == nil {
// 		m.counterArrays = make(map[string]*CounterArray)
// 		counterArray = &CounterArray{dim: Dimensions}
// 		m.counterArrays[Name] = counterArray
// 	} else {
// 		counterArray, f = m.counterArrays[Name]
// 		if f {
// 			if !reflect.DeepEqual(counterArray.dim, Dimensions) {
// 				L.L.Warn("metric.RegisterCounterArray", L.String("ErrMetricAlreadyDefined", Name), L.Strings("old", counterArray.dim), L.Strings("new", Dimensions))
// 				err = ErrMetricAlreadyDefined
// 			}
// 		} else {
// 			counterArray = &CounterArray{dim: Dimensions}
// 			m.counterArrays[Name] = counterArray
// 		}
// 	}
// 	m.counterArraysL.Unlock()

// 	if m.initialized {
// 		counterArray.name = m.Name
// 		if counterArray.e == nil {
// 			counterArray.e = expvarNewMap(Name)
// 			L.L.Debug("metric.RegisterCounterArray", L.String("expvarNewMap", Name))
// 		}
// 		if counterArray.p == nil {
// 			subsystem := "counterarray"
// 			counterArray.p = prometheus.NewCounterVec(prometheus.CounterOpts{ /*Namespace: m.Name,*/ Subsystem: subsystem, Name: validName(Name), Help: Name}, counterArray.dim)
// 			err = prometheusMustRegister(Name, counterArray.p)
// 			if err != nil {
// 				counterArray.p = nil
// 				L.L.Error("metric.RegisterCounterArray", L.String("prometheus.NewCounterVec", Name), L.Error(err))
// 			} else {
// 				L.L.Debug("metric.RegisterCounterArray", L.String("prometheus.NewCounterVec", Name))
// 			}
// 		}
// 	}

// 	return
// }

// // RegisterCounterArrays registers new counterArrays
// func (m *Metric) RegisterCounterArrays(Names []string, Dimensions ...string) ([]*CounterArray, error) {
// 	L.L.Info("metric.RegisterCounterArrays", L.Strings("Names", Names), L.Strings("Dimensions", Dimensions), L.Bool("initialized", m.initialized))
// 	ca := []*CounterArray{}
// 	var err error
// 	for _, Name := range Names {
// 		c, e := m.RegisterCounterArray(Name, Dimensions...)
// 		ca = append(ca, c)
// 		if e != nil {
// 			err = e
// 		}
// 	}
// 	return ca, err
// }

// func (m *Metric) findCounterArray(Name string) (*CounterArray, error) {
// 	if !m.initialized {
// 		return nil, nil
// 	}

// 	m.counterArraysL.RLock()
// 	i, ok := m.counterArrays[Name]
// 	m.counterArraysL.RUnlock()

// 	if !ok {
// 		L.L.Warn("CounterArray not defined", L.String("Name", Name))
// 		return nil, ErrMetricNotDefined
// 	}
// 	if i == nil {
// 		L.L.Warn("CounterArray not initialized", L.String("Name", Name))
// 		return nil, ErrMetricNotDefined
// 	}
// 	return i, nil
// }

// // CounterArrayInc Incs counterArray value
// func (m *Metric) CounterArrayInc(Name string, Keys ...string) error {
// 	ca, err := m.findCounterArray(Name)
// 	if ca == nil {
// 		return err
// 	}

// 	return ca.Inc(Keys...)
// }

// // CounterArrayAdd Adds counterArray value
// func (m *Metric) CounterArrayAdd(Name string, Value int64, Keys ...string) error {
// 	ca, err := m.findCounterArray(Name)
// 	if ca == nil || Value <= 0 {
// 		return err
// 	}

// 	return ca.Add(Value, Keys...)
// }

// // Inc increments counterArray value
// func (ca *CounterArray) Inc(Keys ...string) error {
// 	if ca == nil {
// 		return nil
// 	}
// 	if err := checkKeys(Keys); err != nil {
// 		L.L.Error("CounterArray.Inc.checkKeys", L.Strings("Keys", Keys), L.Error(err))
// 		return err
// 	}
// 	Keys = append(Keys, ca.name)

// 	ld := len(ca.dim)
// 	lk := len(Keys)

// 	if ld == 1 && lk == 1 {
// 		if ca.e != nil {
// 			ca.e.Add(Keys[0], 1)
// 		}
// 		if ca.p != nil {
// 			ca.p.With(prometheus.Labels{ca.dim[0]: Keys[0]}).Inc()
// 		}
// 	} else {
// 		if ld != lk || ld == 0 {
// 			L.L.Panic("CounterArray.Inc.wrongNumberOfDimensions", L.Strings("declared", ca.dim), L.Strings("given", Keys))
// 		}
// 		names := ""
// 		labels := make(prometheus.Labels)
// 		for i, d := range ca.dim {
// 			names += "_" + Keys[i]
// 			labels[d] = Keys[i]
// 		}
// 		if ca.e != nil {
// 			ca.e.Add(names[1:], 1)
// 		}
// 		if ca.p != nil {
// 			ca.p.With(labels).Inc()
// 		}
// 	}

// 	return nil
// }

// // Add Adds counterArray value
// func (ca *CounterArray) Add(Value int64, Keys ...string) error {
// 	if ca == nil {
// 		return nil
// 	}
// 	if err := checkKeys(Keys); err != nil {
// 		L.L.Error("CounterArray.Add.checkKeys", L.Strings("Keys", Keys), L.Error(err))
// 		return err
// 	}
// 	Keys = append(Keys, ca.name)

// 	ld := len(ca.dim)
// 	lk := len(Keys)

// 	if ld == 1 && lk == 1 {
// 		if ca.e != nil {
// 			ca.e.Add(Keys[0], Value)
// 		}
// 		if ca.p != nil {
// 			ca.p.With(prometheus.Labels{ca.dim[0]: Keys[0]}).Add(float64(Value))
// 		}
// 	} else {
// 		if ld != lk || ld == 0 {
// 			L.L.Panic("CounterArray.Add.wrongNumberOfDimensions", L.Strings("declared", ca.dim), L.Strings("given", Keys))
// 		}
// 		labels := make(prometheus.Labels)
// 		names := ""
// 		for i, d := range ca.dim {
// 			names += "_" + Keys[i]
// 			labels[d] = Keys[i]
// 		}
// 		if ca.e != nil {
// 			ca.e.Add(names[1:], Value)
// 		}
// 		if ca.p != nil {
// 			ca.p.With(labels).Add(float64(Value))
// 		}
// 	}

// 	return nil
// }

// // AddMS adds milliseconds between Since and Now
// func (ca *CounterArray) AddMS(Since time.Time, Keys ...string) error {
// 	return ca.Add((time.Now().UnixNano()-Since.UnixNano())/int64(time.Millisecond), Keys...)
// }

// /*
// 	Common functions
// 	================
// */

// func parseNameArrays(a []string) ([]string, []string) {
// 	for i, s := range a {
// 		if s == "" {
// 			if len(a[i+1:]) == 0 {
// 				return a[:i], []string{"key"}
// 			}
// 			return a[:i], a[i+1:]
// 		}
// 	}
// 	return a, []string{"key"}
// }

// var validNameRegex = regexp.MustCompile(`[^a-zA-Z0-9:_]`)

// func validName(n string) string {
// 	if n == "" {
// 		return "_"
// 	}
// 	n = validNameRegex.ReplaceAllLiteralString(n, "_")
// 	if n[0] >= '0' && n[0] <= '9' {
// 		n = "_" + n
// 	}
// 	return n
// }

// func checkKeys(keys []string) error {
// 	for _, k := range keys {
// 		if k == "" {
// 			return ErrEmptyKey
// 		}
// 	}
// 	return nil
// }

// /*
// 	Exported package-level functions
// 	================================
// */

// // Init is wrapper for method M.Init of global metrics variable M
// func Init(Name, Port string, Mux *http.ServeMux) {
// 	M.Init(Name, Port, Mux)
// }

// // InitAll is wrapper for method M.InitAll of global metrics variable M
// func InitAll(Name, Port string, Mux *http.ServeMux, Gauges, Counters, GaugeArrays, CounterArrays []string) {
// 	M.InitAll(Name, Port, Mux, Gauges, Counters, GaugeArrays, CounterArrays)
// }

// // RegisterGauge is wrapper for method M.RegisterGauge of global metrics variable M
// func RegisterGauge(Name string) (gauge *Gauge, err error) {
// 	return M.RegisterGauge(Name)
// }

// // RegisterGauges is wrapper for method M.RegisterGauges of global metrics variable M
// func RegisterGauges(Names []string) {
// 	M.RegisterGauges(Names)
// }

// // GaugeSet is wrapper for method M.GaugeSet of global metrics variable M
// func GaugeSet(Name string, Value int64) error {
// 	return M.GaugeSet(Name, Value)
// }

// // GaugeSetBool is wrapper for method M.GaugeSetBool of global metrics variable M
// func GaugeSetBool(Name string, Value bool) error {
// 	return M.GaugeSetBool(Name, Value)
// }

// // GaugeInc is wrapper for method M.GaugeInc of global metrics variable M
// func GaugeInc(Name string) error {
// 	return M.GaugeInc(Name)
// }

// // GaugeDec is wrapper for method M.GaugeDec of global metrics variable M
// func GaugeDec(Name string) error {
// 	return M.GaugeDec(Name)
// }

// // GaugeAdd is wrapper for method M.GaugeAdd of global metrics variable M
// func GaugeAdd(Name string, Value int64) error {
// 	return M.GaugeAdd(Name, Value)
// }

// // RegisterCounter is wrapper for method M.RegisterCounter of global metrics variable M
// func RegisterCounter(Name string) (counter *Counter, err error) {
// 	return M.RegisterCounter(Name)
// }

// // RegisterCounters is wrapper for method M.RegisterCounters of global metrics variable M
// func RegisterCounters(Names []string) {
// 	M.RegisterCounters(Names)
// }

// // CounterInc is wrapper for method M.CounterInc of global metrics variable M
// func CounterInc(Name string) error {
// 	return M.CounterInc(Name)
// }

// // CounterAdd is wrapper for method M.CounterAdd of global metrics variable M
// func CounterAdd(Name string, Value int64) error {
// 	return M.CounterAdd(Name, Value)
// }

// // RegisterGaugeArray is wrapper for method M.RegisterGaugeArray of global metrics variable M
// func RegisterGaugeArray(Name string, Dimensions ...string) (gaugeArray *GaugeArray, err error) {
// 	return M.RegisterGaugeArray(Name, Dimensions...)
// }

// // RegisterGaugeArrays is wrapper for method M.RegisterGaugeArrays of global metrics variable M
// func RegisterGaugeArrays(Names []string, Dimensions ...string) {
// 	M.RegisterGaugeArrays(Names, Dimensions...)
// }

// // GaugeArraySet is wrapper for method M.GaugeArraySet of global metrics variable M
// func GaugeArraySet(Name string, Value int64, Keys ...string) error {
// 	return M.GaugeArraySet(Name, Value, Keys...)
// }

// // GaugeArrayInc is wrapper for method M.GaugeArrayInc of global metrics variable M
// func GaugeArrayInc(Name string, Keys ...string) error {
// 	return M.GaugeArrayInc(Name, Keys...)
// }

// // GaugeArrayDec is wrapper for method M.GaugeArrayDec of global metrics variable M
// func GaugeArrayDec(Name string, Keys ...string) error {
// 	return M.GaugeArrayDec(Name, Keys...)
// }

// // GaugeArrayAdd is wrapper for method M.GaugeArrayAdd of global metrics variable M
// func GaugeArrayAdd(Name string, Value int64, Keys ...string) error {
// 	return M.GaugeArrayAdd(Name, Value, Keys...)
// }

// // RegisterCounterArray is wrapper for method M.RegisterCounterArray of global metrics variable M
// func RegisterCounterArray(Name string, Dimensions ...string) (counterArray *CounterArray, err error) {
// 	return M.RegisterCounterArray(Name, Dimensions...)
// }

// // RegisterCounterArrays is wrapper for method M.RegisterCounterArrays of global metrics variable M
// func RegisterCounterArrays(Names []string, Dimensions ...string) ([]*CounterArray, error) {
// 	return M.RegisterCounterArrays(Names, Dimensions...)
// }

// // CounterArrayInc is wrapper for method M.CounterArrayInc of global metrics variable M
// func CounterArrayInc(Name string, Keys ...string) error {
// 	return M.CounterArrayInc(Name, Keys...)
// }

// // CounterArrayAdd is wrapper for method M.CounterArrayAdd of global metrics variable M
// func CounterArrayAdd(Name string, Value int64, Keys ...string) error {
// 	return M.CounterArrayAdd(Name, Value, Keys...)
// }

// // SinceMS calculates milliseconds between T and Now
// func SinceMS(T time.Time) int64 {
// 	return (time.Now().UnixNano() - T.UnixNano()) / int64(time.Millisecond)
// }

// // Name returns M.Name, a Name property of global metrics variable M
// func Name() string {
// 	return M.Name
// }

// // Name returns M.Port, a Port property of global metrics variable M
// func Port() string {
// 	return M.Port
// }

// // Mux returns M.Port, a pointer to Mux of global metrics variable M
// func Mux() *http.ServeMux {
// 	return M.Mux
// }

// /*
//     Builtin CPU Metrics
// 	===================
// */

// // BuiltinCPUMetrics for all services to use
// type BuiltinCPUMetrics struct {
// 	cpuUsage  *prometheus.GaugeVec
// 	mtx       sync.Mutex
// 	procStat  string
// 	idleTime  uint64
// 	totalTime uint64
// }

// // Describe describes the thing
// func (c *BuiltinCPUMetrics) Describe(ch chan<- *prometheus.Desc) {
// 	c.cpuUsage.Describe(ch)
// }

// // Collect does the collection of metrics
// func (c *BuiltinCPUMetrics) Collect(ch chan<- prometheus.Metric) {

// 	f, err := os.Open(c.procStat)
// 	if err != nil {
// 		L.L.Warn("BuiltinCPUMetrics.Collect.os.Open", L.Error(err))
// 		return
// 	}
// 	defer f.Close()

// 	data, err := ioutil.ReadAll(f)
// 	if err != nil {
// 		L.L.Warn("BuiltinCPUMetrics.Collect.ioutil.ReadAll", L.Error(err))
// 		return
// 	}

// 	var (
// 		x     = bytes.LastIndex(data, []byte(")"))
// 		dummy uint64
// 		dummi int64
// 		state string
// 		UTime float64
// 		STime float64
// 		VSize float64
// 		jiffy = float64(100)
// 	)

// 	if x < 0 {
// 		L.L.Warn("BuiltinCPUMetrics.Collect.bad.format", L.Error(err))
// 		return
// 	}

// 	_, err = fmt.Fscan(
// 		bytes.NewBuffer(data[x+2:]),
// 		&state, // State,
// 		&dummy, // PPID,
// 		&dummy, // PGRP,
// 		&dummy, // Session,
// 		&dummy, // TTY,
// 		&dummi, // TPGID,
// 		&dummy, // Flags,
// 		&dummy, // MinFlt,
// 		&dummy, // CMinFlt,
// 		&dummy, // MajFlt,
// 		&dummy, // CMajFlt,
// 		&UTime, // UTime,
// 		&STime, // STime,
// 		&dummy, // CUTime,
// 		&dummy, // CSTime,
// 		&dummy, // Priority,
// 		&dummy, // Nice,
// 		&dummy, // NumThreads,
// 		&dummy, // itrealvalue,
// 		&dummy, // starttime,
// 		&VSize, // vsize
// 	)
// 	if err != nil {
// 		L.L.Warn("BuiltinCPUMetrics.Collect.Fscan", L.Error(err), L.ByteString("data", data), L.Int("x", x))
// 		return
// 	}

// 	mem := runtime.MemStats{}
// 	runtime.ReadMemStats(&mem)

// 	c.mtx.Lock()
// 	defer c.mtx.Unlock()

// 	c.cpuUsage.WithLabelValues("userTimeSeconds").Set(UTime / jiffy)
// 	c.cpuUsage.WithLabelValues("kernelTimeSeconds").Set(STime / jiffy)
// 	c.cpuUsage.WithLabelValues("totalTimeSeconds").Set((STime + UTime) / jiffy)
// 	c.cpuUsage.WithLabelValues("procCPU").Set(c.GetProcCPU())
// 	c.cpuUsage.WithLabelValues("heapSizeBytes").Set(float64(mem.HeapAlloc))
// 	c.cpuUsage.WithLabelValues("stackSizeBytes").Set(float64(mem.StackSys))
// 	c.cpuUsage.WithLabelValues("totalVmSizeBytes").Set(float64(mem.Sys))

// 	c.cpuUsage.Collect(ch)
// 	c.cpuUsage.Reset()
// }

// // NewBuiltinCPUMetrics creates BuiltinCPUMetrics
// func NewBuiltinCPUMetrics(serviceName string, procStat string) *BuiltinCPUMetrics {
// 	m := BuiltinCPUMetrics{
// 		procStat: procStat,
// 		cpuUsage: prometheus.NewGaugeVec(
// 			prometheus.GaugeOpts{
// 				Subsystem:   "builtin",
// 				Name:        "cpu_usage",
// 				Help:        "CPU usage provided by metrics",
// 				ConstLabels: prometheus.Labels{"service": serviceName},
// 			},
// 			[]string{"item"},
// 		),
// 	}
// 	m.GetProcCPU()
// 	return &m
// }

// // GetProcCPU gets proc cpu usage
// func (c *BuiltinCPUMetrics) GetProcCPU() float64 {
// 	f0, err := os.Open("/proc/stat")
// 	if err != nil {
// 		L.L.Warn("BuiltinCPUMetrics.Collect.os.Open0", L.Error(err))
// 		return 0
// 	}
// 	defer f0.Close()

// 	data0, err := ioutil.ReadAll(f0)
// 	if err != nil {
// 		L.L.Warn("BuiltinCPUMetrics.Collect.ioutil.ReadAll0", L.Error(err))
// 		return 0
// 	}
// 	data1 := strings.Split(string(data0), "\n")

// 	idle, total := uint64(0), uint64(0)

// 	for _, data2 := range data1 {
// 		if strings.Index(data2, "cpu ") == 0 {
// 			fields := strings.Fields(data2)
// 			numFields := len(fields)
// 			for i := 1; i < numFields; i++ {
// 				val, err := strconv.ParseUint(fields[i], 10, 64)
// 				if err != nil {
// 					fmt.Println("Error: ", i, fields[i], err)
// 				}
// 				total += val
// 				if i == 4 {
// 					idle = val
// 				}
// 			}
// 		}
// 	}

// 	if c.totalTime == 0 {
// 		c.totalTime = total
// 		c.idleTime = idle
// 	} else {
// 		total, c.totalTime = total-c.totalTime, total
// 		idle, c.idleTime = idle-c.idleTime, idle
// 	}

// 	procCPU := float64(0)
// 	if total > 0 {
// 		procCPU = float64(total-idle) / float64(total) * 100
// 	}
// 	return procCPU
// }

// // expvar panic control, used to catch panic in case of "Reuse of exported var name",
// // mostly in livedoc for name "storages"

// func expvarNewInt(name string) *expvar.Int {
// 	var e *expvar.Int = nil

// 	defer func() {
// 		if err := recover(); err != nil {
// 			L.L.Error("expvarNewInt", L.String("Name", name), L.String("error", fmt.Sprintf("%+v", err)))
// 			e = nil
// 		}
// 	}()

// 	e = expvar.NewInt(name)

// 	return e
// }

// func expvarNewMap(name string) *expvar.Map {
// 	var m *expvar.Map = nil

// 	defer func() {
// 		if err := recover(); err != nil {
// 			L.L.Error("expvarNewMap", L.String("Name", name), L.String("error", fmt.Sprintf("%+v", err)))
// 			m = nil
// 		}
// 	}()

// 	m = expvar.NewMap(name)

// 	return m
// }

// func prometheusMustRegister(name string, cs ...prometheus.Collector) (err error) {
// 	err = nil
// 	defer func() {
// 		if err := recover(); err != nil {
// 			L.L.Warn("prometheusMustRegister", L.String("Name", name), L.String("error", fmt.Sprintf("%+v", err)))
// 			err = ErrMetricAlreadyDefined
// 		}
// 	}()
// 	prometheus.MustRegister(cs...)
// 	return
// }
