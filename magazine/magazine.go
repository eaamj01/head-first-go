package magazine

// Address Common struct for addresses
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

// Subscriber Struct that holds subscriber info and has reference HomeAddress to Address struct
type Subscriber struct {
	Name        string
	Rate        float64
	Active      bool
	HomeAddress Address
}

// Employee struct holding employee data and has a direct reference to Address struct
type Employee struct {
	Name   string
	Salary float64
	Address
}
