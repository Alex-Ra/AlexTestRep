package main

import (
    "fmt"
    "strconv"
    "os"
    "io"
    "io/ioutil"
    "time"
    "encoding/json"
    //"flag"
)

type MyObj1 struct {
	//Rtu_info json.RawMessage
	//Readings json.RawMessage
	Rtu_info RtuType
}

type Jsonobject struct {
	Rtu_info RtuType
	Readings []ReadingsType
}

type RtuType struct {
	ID int
	Type string
}
 
type ReadingsType struct {
	Sensor int
	Time int
	Value  float64
}


func main() {
	//var species = flag.String("species", "gopher", "the species we are studying")
	/*wordPtr := flag.String("word", "foo", "a string")
	flag.Parse()
	fmt.Println("Print flag: ", *wordPtr)	

    fmt.Println("The time is", time.Now())*/
    
    
/*for _, arg := range os.Args {
    if arg == "-help" {
      fmt.Printf ("I need somebody\n")
    }else if arg == "-version" {
      fmt.Printf ("Version Zero\n")
    } else {
      fmt.Printf("arg %s: \n", arg)
    }
  }*/
  var jsFile, outFile  string
  for i :=1; i< len(os.Args); i++ {
  	if os.Args[i] == "-f" {
  	  jsFile = os.Args[i+1]
  	  i++
      fmt.Printf ("jsFile Name %s %s %d\n", jsFile, "index: ", i)
    }else if os.Args[i] == "-o" {
      outFile = os.Args[i+1]
      i++
      fmt.Printf ("outFile Name %s %s %d\n", outFile, "index: ", i)
    }else {
      fmt.Printf ("Wrong argument '%s %s %d\n", os.Args[i], "' in the command line. Index: ", i)
    }       
  }
  
  //fmt.Printf ("jsFile Name outside %s \n", jsFile)
   // fmt.Printf ("outFile Name outside %s \n", outFile)
    
    
// Wrong way to reead? (not working)
    fmt.Println("And if you try to open a file:")
    fmt.Println(os.Open("test/data.json"))
// Right way to read files
    file, e := ioutil.ReadFile(jsFile)
    if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file)) 

	
	var myJsObj Jsonobject
	//reading json file into Jsonobject struct
	json.Unmarshal(file, &myJsObj)
	fmt.Printf("Rtu_info\nID: %d\n", myJsObj.Rtu_info.ID) 
	fmt.Printf("Type: %s\n", myJsObj.Rtu_info.Type)
	fmt.Printf("\nReadings\nSensor    Time             Value\n")	
   
    //var ofile string = "test/b.txt"
    /*f, err := os.OpenFile(outFile, 1, 0666)
        if err != nil {
                fmt.Println(err)
                return
        }*/
        f, err := os.Create(outFile)

        n, err :=io.WriteString(f, "Rtu_info\nID: " + strconv.Itoa(myJsObj.Rtu_info.ID) +
        "\nType: " + myJsObj.Rtu_info.Type)
		io.WriteString(f,"\n\nReadings\nSensor    Time             Value")
        if err != nil {
                fmt.Println(n, err)
                return
        }   
    
    
// Looping Readings	
	for _, r:=range myJsObj.Readings{
	//Converting float to string to show wright number of decimals
		var strValue string = strconv.FormatFloat(r.Value, 'f', -1, 64)
		var mTime = time.Unix(0, int64(r.Time)*int64(time.Second))
		//formating date-time
		//const layout = "Mon, 02 Jan 2006 15:04:05 MST"
		const layout = "2006-01-02 15:04:05"		
		fmt.Printf("%d %s %s\n", r.Sensor, mTime.Format(layout), strValue)
		io.WriteString(f, "\n" + strconv.Itoa(r.Sensor) + " " + mTime.Format(layout) + " " +strValue)
	}
	// Closing output file
	f.Close()
	

	/*var x map[string]interface{}
	
	fmt.Printf("Results: %+v\n", x)
	json.Unmarshal(file, &x)
	fmt.Printf("Results: %+v\n", x) 
	
	for key, value := range x {
    	
    	if key == "readings"{ 
    		fmt.Println("Readings", "\n") 		
    	}else if key == "rtu_info"{
			fmt.Println("Rtu_info", "\n")
			var rtu = map[string]string{}
			rtu = value
			fmt.Println("Rtu", rtu, "\n")
			//var r map[string]
    	}
    	fmt.Println("Key:", key, "Value:", value, "\n")
	}*/	
		
}


