type TimeMap struct {
    Store map[string][]ValueObject
}

type ValueObject struct {
    Timestamp int
    Value string
}


func Constructor() TimeMap {
    return TimeMap{Store: make(map[string][]ValueObject, 0)}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    if _, ok := this.Store[key]; !ok {
        this.Store[key] = make([]ValueObject, 0)
    }

    this.Store[key] = append(this.Store[key], ValueObject{Timestamp: timestamp, Value: value})
}


func (this *TimeMap) Get(key string, timestamp int) string {
    
    value := ""

    if values, ok := this.Store[key]; ok {
        low, hi := 0, len(values)-1
        for low <= hi {
            mid := low + (hi - low) / 2
            if values[mid].Timestamp == timestamp {
                return values[mid].Value
            }

            if values[mid].Timestamp < timestamp {
                low = mid + 1
                value = values[mid].Value
            } else {
                hi = mid - 1
            }
        }
    }

    
    return value
}
