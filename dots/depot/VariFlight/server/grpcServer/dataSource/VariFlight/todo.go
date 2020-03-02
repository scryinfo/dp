package VariFlight

// todo: fight status push service with websocket.
// Background: flight status varies over time but client do not know exactly what status it is at some point.
// As a matter of fact, two requests of VariFlightCaller may be responded with the same or different pieces of flight dataToCache.

// If the datum responded are the same, the last VariFlightCaller request is just a repeated work, which may not only contribute to VariFlightCaller cost,
// but also suffer from the uncontrollable network delay.
// Thus we should try our best to avoid repeated requests with no different flight dataToCache responded.
// Probably, we can cache flight dataToCache out of VariFlightCaller, and should prefer to inquire effective dataToCache from cache before calling VariFlightCaller.

// todo: flight value onto blockchain
// todo: cacher
// todo: disscuss the disadvantages and neccessity to cache flight value, don't neglect the very nature of flight value, namely dynamic.
// Due to the dynamic nature, flight value may vary overtime while the calling method and its parameters are the same. In other words, the cached
// flight value maybe outdated. If a player were provided with outdated value, what would happen to our predictive game WillCity?
// As a compromise proposal, VariFlightCaller calling intervals can be set, during the validPeriod, we prefer to fetch flight value directly from cache, otherwise
// an wholly new value should be made. This solution comes with benefit that we don't have to value VariFlightCaller too frequently, which can save us some cost
// required by value server provider.
// Besides, when cache value we shouldn't waste computer memory with duplicate records. That means we should cache value that haven't existed yet.
// todo: storer
// todo: storer JSON value directly.
// todo: db Get if cache no.

// todo: value flitering
// todo: think more about limits on flight value. What a piece of value is acceptable if flight value in the past is inquired? ...

// todo: clearly investigation on flight value parametrs, such as page and perpage
// todo: clear the time zone of date in request parameters
// todo: URLDecode of flight data.

// Todo: remove-zero rule for responsed flight number
// Todo: URLDecode of Chinese characters
// Todo: responsed flight local time --> global time
// Todo: multi-languages responsed value

// todo: wrap string function

// todo: zone info of FeiChangZhun corresponding to IANA Time Zone database

// todo: VariFlightCaller.call. How does VariFlightCaller status respond to HTTP response?

// todo: storer.schema. omitempy tag and time conversion with sqlx.

// todo. storer.sqlx.MustExec(schema).
