package food

import "time"

// create custom protein data type
type protein struct {
    id int
    name string
    amount string // A string is used here because my gf and I often refer to item amounts in units of "1 bag" or "half a box". This will be easier to represent as a string
    purchase_date time.Time
}


