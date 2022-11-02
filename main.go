package main

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
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

driver := Driver{
	dir: dir,
	mutexes: make(map[string]*sync.Mutex),
	log: opts.Logger,
}

if _,err := os.Stat(dir); err == nil{
	opts.Logger.Debug("Using '%s' (database alredy exists)\n",dir)
	return &driver, nil
}

opts.Logger.Debug("Creating the database at '%s'...\n",dir)
return &driver, os.MkdirAll(dir,0755)

}

func (d *Driver) Write(collection, resource string,v interface{}) error{

if collection == ""{
	return fmt.Errorf("Missing collection - no place to save record!")
}

if resource == ""{
	return fmt.Errorf("Missing resource - unable to save record (no name)!")
}

mutex:= g.getOrCreateMutex(collection)
mutex.Lock()
defer mutex.Unlock()

dir := filepath.Join(d.dir,collection)
fnlPath := filepath.Join(dir, resource+".json")
tmpPath := fnlPath + ".tmp"

if err := os.MkdirAll(dir,0755);err != nil{
	return err
}

b, err := json.MarshalIndent(v, "", "\t")

if err != nil{
	return err
}

b = append(b,byte('\n'))

if err := ioutil.WriteFile(tmpPath, b, 0644);err != nil{
	return err
}

}

func (d *Driver) Read() error{

}

func (d *Driver) ReadAll(){

}

func (d *Driver) Delete() error{

}

func (d *Driver) getOrCreateMutex() *sync.Mutex{

}

func stat(path string)(fi os.FileInfo, err error){
	if fi, err = os.Stat(path);os.IsNotExist(err){
		fi, err = os.Stat(path + ".json")
	}
	return
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
