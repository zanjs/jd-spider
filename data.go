package main

// Attribute is
type Attribute struct {
	Name string `json:"name"`
	Sort int8   `json:"sort"`
}

//ResData is
type ResData struct {
	URL       string      `json:"url"`
	Title     string      `json:"title"`
	Thumb     string      `json:"thumb"`
	Domain    string      `json:"domain"`
	Attribute []Attribute `json:"attribute"`
}
