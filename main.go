package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"github.com/jcelliott/lumber"
	"path/filepath"
)

const Version="1.0.1"

type (
	Logger interface{
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}
	Driver struct{
		mutex sync.Mutex
		mutexes map[string]*sync.Mutex
		dir string
		log Logger
	}
)

type Options struct{
	Logger
}

func New(dir string, options *Options)(*Driver, error){
dir = filepath.Clean(dir)
opts := Options{}

if options != nill{
	opts = *options
}

if(opts.Logger == nil){
	opts.Logger = lumber.NewConsoleLogger((lumber.INFO))
}
}

func (d *Driver) Write() error{

}

func (d *Driver) Read() error{

}

func (d *Driver) ReadAll(){

}

func (d *Driver) Delete() error{

}

func (d *Driver) getOrCreateMutex() *sync.Mutex{

}

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"
	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error ", err)
	}

	employees := []User{
		{
			"Name": "Eveleen",
			"Agge": 34,
			"Contact": "947-447-5688",
			"Company": "Jabbersphere",
			Address{"San Francisco","Central","USA","420029"}
		  }, {
			"Name": "Winslow",
			"Agge": 66,
			"Contact": "420-607-1115",
			"Company": "Gabvine",
			Address{"San Francisco","Central","USA","420029"}
		  }, {
			"Name": "Nellie",
			"Agge": 47,
			"Contact": "689-843-4027",
			"Company": "Wikizz",
			Address{"San Francisco","Central","USA","420029"}
		  }, {
			"Name": "Neils",
			"Agge": 81,
			"Contact": "364-555-9387",
			"Company": "Twitterlist",
			Address{"Texas","Colorado","USA","420111"}
		  }, {
			"Name": "Doro",
			"Agge": 25,
			"Contact": "411-155-7706",
			"Company": "Ainyx",
			Address{"New York","New York City","USA","420035"}
		  }, {
			"Name": "Starr",
			"Agge": 73,
			"Contact": "818-377-3109",
			"Company": "Vimbo",
			Address{"Texas","Colorado","USA","420111"}
		  }, {
			"Name": "Kaitlin",
			"Agge": 83,
			"Contact": "609-985-9276",
			"Company": "Gevee",
			Address{"Texas","Colorado","USA","420111"}
		  }, {
			"Name": "Melissa",
			"Agge": 77,
			"Contact": "742-655-9010",
			"Company": "JumpXS",
			Address{"New York","New York City","USA","420035"}
		  },
	}

	for _, value := range employees{
		db.Write("users", value.Name, User{
			Name: value.Name,
			Age: value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error ", err)
	}
	fmt.Println(records)

	allusers := []Users{}

	for _, f := range records{
		employeeFound := User{}
		if err := json.Unmarshal([]byte(f), &employeeFound);err != nil{
			fmt.Println("Error ",err)
		}
		allusers = append(allusers, employeeFound)
	}
	
	fmt.Println((allusers))
}
