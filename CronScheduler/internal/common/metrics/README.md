# Metrics package

The purpose of this package is to make the process of generating metrics from Go code easier and more unified. Two different use cases are supported by the API, and they can be used in parallel.

## Definitions

The following metrics elements exist:

- `Gauge`: scalar int64 metric that can be set to arbitrary values during program/service execution. Supported methods/operations are: Set, SetBool, Inc, Dec, Add.
- `Counter`: scalar int64 metric that starts from zero and can only be incremented, either by one or by a given positive value. Supported methods/operations are: Inc, Add.
- `GaugeArray`: an array of Gauges, indexed by a key (string value).
- `CounterArray`: an array of Counters, indexed by a key (string value).

Array (GaugeArray or CounterArray) is a set of gauges bearing the same and having multiple values identified by keys - string identifiers. An example would be the CounterArray which counts the number of bets per brand,which would look like the following:
```javascript
bets_per_brand{key="BrandA"} 200
bets_per_brand{key="BrandB"} 130
bets_per_brand{key="default"} 30
```


## API

### Functions

##### API functions are divided into three groups:

* _Initialization functions_: used to start metrics mechanism, with a set of attributes provided during initialization.
* _Registration functions_: used for registering metric elements. Each metric element is identified by its name (string), and can not be used unless it is registered.
* _Update functions_: used to set or update values of particular metrics.

##### The above functions operate according to following rules:

1. _Registration functions_ can be called at any time during the life of program/service, before and/or after _initialization functions_. Prior to initialization, _registration functions_ have no impact on metrics, but metric configuration that they are carrying is recorded, and used during execution of _initialization functions_. If _initialization functions_ are never called - no metrics are produced, but also no errors/exceptions from _registration functions_. This ensures that metrics can be registered in packages, and their existence will depend whether _initialization functions_ are called from main.go or not.
1. _Registration function_ for single metric object can be called more than once for the same metric object. The first call will set up and return new metric object, subsequent calls will return the same metric object and ErrMetricAlreadyDefined. This can be used to detect if metric names overlap between packages or between a package and main.go, and also to use the metric object even if it is a possible duplicate.
1. _Update functions_ will return error if called prior to (or without) corresponding _registration functions_.
1. _Update functions_ will have no effect, nor will return errors, if called for registered metrics, but prior to calling _initialization functions_.

### Global metrics variable `M`

In metrics package, one global metric variable, `M` of type Metric is defined. Although other local metrics variables can be defined and used, even more than one at a time, this global variable should be used for all metrics management.

### Initialization functions

#### `Init()`
```go
func (*Metric) Init(
        Name string,          // Global metric name, should match the program/service name
        Port string,          // Port to which to expose the metrics URLs, e.g. :8880
        Mux *http.ServeMux    // http multiplexer to use to expose the metrics
     )
```

Function **Init()** is used to establish and activate metrics system.

If **Port** is given, **Init()** will _ListenAndServe_ on that port, exposing metrics URLs either via **Mux**, if provided, or via newly created multiplexer. The multiplexer that is used for _ListenAndServe_ will be available to the rest of the world in **M.Mux**. If **Port** is empty string - no _ListenAndServe_ will be started.

If **Mux** is provided (not _nil_) - metrics URLs will be added to it. If **Mux** is not provided, but **Port** is provided - new multiplexer will be created, as stated above, to serve metrics URLs. If neither **Mux** nor **Port** are provided - there will be no URLs attached to metrics in the **Init()** function.

So, combining **Port** and **Mux**, metrics can be set up to set up a new multiplexer, use external multiplexer, or work without one. If a multiplexer is used (internal or external), it is exposed in **M.Mux**, so other parts of the program/service can attach their URLs to it.

#### `InitAll()`
```go
func (*Metric) InitAll(
        Name string,               // Global metric name, should match the program/service name
        Port string,               // Port to which to expose the metrics URLs, e.g. :8880
        Mux *http.ServeMux,        // http multiplexer to use to expose the metrics
        Gauges        []string,    // equivalent to parameter 'Names' in RegisterGauges()
        Counters      []string,    // equivalent to parameter 'Names' in RegisterCounters()
        GaugeArrays   []string,    // equivalent to parameter 'Names' in RegisterGaugeArrays()
        CounterArrays []string     // equivalent to parameter 'Names' in RegisterCounterArrays()
     )
```

Function **InitAll()** is equivalent to calling the following functions separately: **Init()**, **RegisterGauges()**, **RegisterCounters()**, **RegisterGaugeArrays()**, and **RegisterCounterArrays()**, i.e. to initialize metrics system and register all metrics in a single call.

First three parameters, **Name**, **Port** and **Mux** have the same meaning as in **Init()**. Subsequent parameters are slices containing the names of corresponding metrics elements.

Example:
```go
metrics.M.InitAll("betupdates", ":8880", nil, nil, nil, nil, []string{"operations", "durationsMS", "", "brand", "operation", "status"})
```

### Registration functions

* `func (*Metric) RegisterGauge(Name string) (gauge *Gauge, err error)` is used to register a single gauge, identified by **Name**. It returns a pointer to a gauge as the first parameter, and an error or _nil_ as the second. Error, if present, is **ErrMetricAlreadyDefined**, which means that the same function was already called with the same **Name**. However, even if error is reported, the first parameter still contains a pointer to a gauge, so multiple packages can call the same function and use the same metric.

* `func (*Metric) RegisterGauges(Names []string)` is used to register a number of gauges. It is equivalent to calling **RegisterGauge()** without keeping the returned values - gauge pointers and errors.

* `func (*Metric) RegisterCounter(Name string) (counter *Counter, err error)` is used to register a single counter, identified by **Name**. It returns a pointer to a counter as the first parameter, and an error or _nil_ as the second. Error, if present, is **ErrMetricAlreadyDefined**, which means that the same function was already called with the same **Name**. However, even if error is reported, the first parameter still contains a pointer to a counter, so multiple packages can call the same function and use the same metric.

* `func (*Metric) RegisterCounters(Names []string)` is used to register a number of counters. It is equivalent to calling **RegisterCounter()** without keeping the returned values - counter pointers and errors.

* `func (*Metric) RegisterGaugeArray(Name string, Dimensions ...string) (gaugeArray *GaugeArray, err error)` registers gauge array and returns a pointer to it, and error or _nil_. Error, if present, is **ErrMetricAlreadyDefined**, which means that the same function was already called with the same **Name**. However, even if error is reported, the first parameter still contains a pointer to a gauge array, so multiple packages can call the same function and use the same metric.

The second parameter - **Dimensions** is optional slice of array dimensions, if omitted - one dimension named "key" is assumed (**[]string{"key"}**). An example of counter array with multiple dimensions would be a counter of bets per brand and status (`Dimensions = []string{"brand","status"}`):
```javascript
bet_status{brand="BrandA", status="OK"} 20
bet_status{brand="BrandA", status="Failed"} 5
bet_status{brand="BrandB", status="OK"} 120
bet_status{brand="BrandA", status="Failed"} 11
bet_status{brand="default",status="OK"} 300
```

* `func (*Metric) RegisterGaugeArrays(Names []string, Dimensions ...string)` is used to register a number of gauge arrays, all with the same set of dimensions. It is equivalent to calling **RegisterGaugeArray()** without keeping the returned values - gauge array pointers and errors. If different dimensions are needed for different gauges - the method **RegisterGaugeArray()** should be used.

* `func (*Metric) RegisterCounterArray(Name string, Dimensions ...string) (counterArray *CounterArray, err error)` registers counter array and returns a pointer to it, and error or _nil_. Error, if present, is **ErrMetricAlreadyDefined**, which means that the same function was already called with the same **Name**. However, even if error is reported, the first parameter still contains a pointer to a counter array, so multiple packages can call the same function and use the same metric.

* `func (*Metric) RegisterCounterArrays(Names []string, Dimensions ...string)` is used to register a number of counter arrays, all with the same set of dimensions. It is equivalent to calling **RegisterCounterArray()** without keeping the returned values - counter array pointers and errors. If different dimensions are needed for different counters - the method **RegisterCounterArray()** should be used.

* `func (*Metric) InitAll(Name, Port string, Mux *http.ServeMux, Gauges, Counters, GaugeArrays, CounterArrays []string)` **one more time**: since GaugeArrays and CounterArrays may have custom dimensions (apart from the default dimension **key**), and functions **RegisterGaugeArray()** and **RegisterCounterArray()** have a parameter **Dimension**, an implicit way to specify custom dimensions via **InitAll()** has been introduced as follows: parameters **GaugeArrays** and **CounterArrays**, which normally contain the list of array names, can be used to also hold the dimensions, provided that these two groups are separated by one empty string. For example:
```go
metrics.M.InitAll("livedoc", ":8880", nil, nil, nil, nil, []string{"operations", "durationsMS", "", "brand", "operation", "status"})
```
is used to register counter arrays "operations" and "durationMS", which both have dimensions "brand", "operation" and "status".

### Update functions

For each registered metric object (gauge, counter, gauge array, counter array), update operations can be performed "by name" or "by reference".

#### By name

"By name" updates call functions/methods of the metric object (preferably global **metric.M**), providing the name of the metric that should be updated. They may return **ErrMetricNotDefined** if update is called prior/without appropriate **Register...** call.

* `func (*Metric) GaugeSet(Name string, Value int64) error` sets the value of a gauge identified by **Name**.
* `func (*Metric) GaugeSetBool(Name string, Value bool) error` sets the value of a gauge identified by **Name** based on boolean input: 0 for false, 1 for true.
* `func (*Metric) GaugeInc(Name string) error` increments current gauge value by 1.
* `func (*Metric) GaugeDec(Name string) error` decrements current gauge value by 1.
* `func (*Metric) GaugeAdd(Name string, Value int64) error` adds the given value to a current gauge value. **Value** can be negative.

* `func (*Metric) CounterInc(Name string) error` increments current counter value by 1.
* `func (*Metric) CounterAdd(Name string, Value int64) error` increments current counter value by given **Value**.

The following are update methods for arrays. They will panic if the number of **Keys** provided does not match the number of dimensions given on registration. Please remember that absence of **Dimensions** on registration means default dimension **key**, in which case one **Keys** value must be provided to the following methods.

* `func (*Metric) GaugeArraySet(Name string, Value int64, Keys ...string) error` sets the value of a gauge array identified by **Name**, for dimensions given in **Keys...**.
* `func (*Metric) GaugeArrayInc(Name string, Keys ...string) error` increments gauge array identified by **Name** by 1, for dimensions given in **Keys...**.
* `func (*Metric) GaugeArrayDec(Name string, Keys ...string) error` decrements gauge array identified by **Name** by 1, for dimensions given in **Keys...**.
* `func (*Metric) GaugeArrayAdd(Name string, Value int64, Keys ...string) error` adds the given value to a gauge array identified by **Name**, for dimensions given in **Keys...**.

* `func (*Metric) CounterArrayInc(Name string, Keys ...string) error` increments counter array identified by **Name** by 1, for dimensions given in **Keys...**.
* `func (*Metric) CounterArrayAdd(Name string, Value int64, Keys ...string) error` increments counter array identified by **Name** by **Value**, for dimensions given in **Keys...**.


#### By reference

If metrics are registered using single registration functions, the pointers to registered objects were given back, and they can be used for easier and faster access to metric's update methods. As opposed to "By name" methods, these methods will not return **ErrMetricNotDefined**, since pointers were obtained in **Register...** calls.

* `func (*Gauge) Set(Value int64)`  sets the value of a gauge.
* `func (*Gauge) SetBool(Value bool)`  sets the value of a gauge based on boolean input: 0 for false, 1 for true.
* `func (*Gauge) Inc()` increments current gauge value by 1.
* `func (*Gauge) Dec()` decrements current gauge value by 1.
* `func (*Gauge) Add(Value int64)` adds the given value to a current gauge value. **Value** can be negative.

* `func (*Counter) Inc()` increments current counter value by 1.
* `func (*Counter) Add(Value int64)` increments current counter value by given **Value**.

The following are update methods for arrays. They will panic if the number of **Keys** provided does not match the number of dimensions given on registration. Please remember that absence of **Dimensions** on registration means default dimension **key**, in which case one **Keys** value must be provided to the following methods.

* `func (*GaugeArray) Set(Value int64, Keys ...string) error` sets the value of a gauge array, for dimensions given in **Keys...**.
* `func (*GaugeArray) Inc(Keys ...string) error` increments gauge array by 1, for dimensions given in **Keys...**.
* `func (*GaugeArray) Dec(Keys ...string) error` decrements gauge array by 1, for dimensions given in **Keys...**.
* `func (*GaugeArray) Add(Value int64, Keys ...string) error` adds the given value to a gauge array, for dimensions given in **Keys...**.

* `func (*CounterArray) Inc(Keys ...string) error` increments counter array by 1, for dimensions given in **Keys...**.
* `func (*CounterArray) Add(Value int64, Keys ...string) error` increments counter array by **Value**, for dimensions given in **Keys...**.

## Current implementation

Currently, the metrics package feeds the data provided via API to *Prometheus* and *expvar* metrics systems, with good prospects of opting out *expvar* in the future. Consequently, URLs that are provided to **Mux** for exposing the metrics are **/metrics** for *Prometheus* and **/debug/vars** for *expvar*. *Prometheus* data is gathered by external Prometheus services, and displayed in Grafana later on.

The gymnastics regarding **Mux** coming in to metrics, and going out, are used to tie the logging URL **/log** from the logging system into the same multiplexer. Obviously, other management and instrumentation URLs can be added in the same way.

## Example

Sample usage of metrics package can be found in `e3dummyd` service, together with the logging system which coexists with metrics on the same **Mux**.
