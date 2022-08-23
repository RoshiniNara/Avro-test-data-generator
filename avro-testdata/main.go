package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hamba/avro/ocf"
)

type Container struct {
	Metrics []*Metric `json:"metrics"`
}

type Metric struct {
	Time     *time.Time             `json:"timestamp,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty"`
}

func main() {
	schema := `{
	    "type": "record",
	    "name": "simple",
	    "namespace": "org.hamba.avro",
	    "values": "string",
	    "fields": [
	    	{
			"name": "Metrics", "type": {
				"type": "array",
				"items": {
					"type": "record",
					"name": "metrics",
					"fields": [
						{
							"name": "Metadata",
							"type": {
								"type": "map",
								"values": ["string", "int", "long", "float", "double" ]
							}
						},
						{
							"name": "Data",
							"type": {
								"type": "map",
								"values": ["string", "int", "long", "float", "double" ]
							}
						}
					]
				}
			}
		}
	]
}
`
	by := []byte("{\"metrics\":[{\"timestamp\":\"0001-01-01T00:00:00Z\",\"metadata\":{\"key\":\"value\"}}]}")
	container := Container{}
	json.Unmarshal(by, &container)
	fmt.Println(by)
	//orig := string(by)
	//fmt.Println(orig)
	// Byte array Outputs:
	//	[123 34 109 101 116 114 105 99 115 34 58 91 123 34 116 105 109 101 115 116 97 109 112 34 58 34 50 48 49 57 45 48 51 45 49 53 84 49 49 58 48 56 58 48 50 43 48 49 58 48 48 34 44 34 109 101 116 97 100 97 116 97 34 58 123 34 107 101 121 34 58 34 118 97 108 117 101 34 125 44 34 100 97 116 97 34 58 123 34 115 116 114 105 110 103 34 58 34 116 101 120 116 34 44 34 102 108 111 97 116 34 58 49 46 49 49 44 34 105 110 116 101 103 101 114 34 58 53 125 125 93 125]

	//data, err := avro.Marshal(schema, container)

	f, err := os.Open("file.avro")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	enc, err := ocf.NewEncoder(schema, f)
	if err != nil {
		log.Fatal(err)
	}

	err = enc.Encode(container)
	if err != nil {
		log.Fatal(err)
	}

	if err := enc.Flush(); err != nil {
		log.Fatal(err)
	}

	if err := f.Sync(); err != nil {
		log.Fatal(err)
	}
}

//	out := Container{}
//err = avro.Unmarshal(schema, data, &out)
//if err != nil {
//		log.Fatal(err)
//	}
/*	fmt.Println(out) // Container Output:
	//{[0xc00006e2c0]}
	var b []byte
	b, _ = json.Marshal(&out)
	fmt.Println(b)
	parsed := string(b)
	fmt.Println(parsed)

	//result1 := parsed == orig
	//fmt.Println(result1)
	// Byte array Outputs:
	//	[123 34 109 101 116 114 105 99 115 34 58 91 123 34 116 105 109 101 115 116 97 109 112 34 58 34 50 48 49 57 45 48 51 45 49 53 84 49 49 58 48 56 58 48 50 43 48 49 58 48 48 34 44 34 109 101 116 97 100 97 116 97 34 58 123 34 107 101 121 34 58 34 118 97 108 117 101 34 125 44 34 100 97 116 97 34 58 123 34 115 116 114 105 110 103 34 58 34 116 101 120 116 34 44 34 102 108 111 97 116 34 58 49 46 49 49 44 34 105 110 116 101 103 101 114 34 58 53 125 125 93 125]
}*/
