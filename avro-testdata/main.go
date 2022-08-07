package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hamba/avro"
)

type Container struct {
	Metrics []*Metric `json:"metrics"`
}

type Metric struct {
	Time     string                 `json:"timestamp,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty"`
}

func main() {
	schema, err := avro.Parse(`{
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
							"name": "Time",
							"type": "string"
							
						},
						{
							"name": "Metadata",
							"type": {
								"type": "map",
								"values": "string"
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
`)
	if err != nil {
		log.Fatal(err)
	}

	by := []byte("{\"metrics\":[{\"timestamp\":\"2019-03-15T11:08:02+01:00\",\"metadata\":{\"key\":\"value\"},\"data\":{\"string\":\"text\",\"float\":1.11,\"integer\":5}}]}")
	container := Container{}
	json.Unmarshal(by, &container)
	fmt.Println(by)
	// Byte array Outputs:
	//	[123 34 109 101 116 114 105 99 115 34 58 91 123 34 116 105 109 101 115 116 97 109 112 34 58 34 50 48 49 57 45 48 51 45 49 53 84 49 49 58 48 56 58 48 50 43 48 49 58 48 48 34 44 34 109 101 116 97 100 97 116 97 34 58 123 34 107 101 121 34 58 34 118 97 108 117 101 34 125 44 34 100 97 116 97 34 58 123 34 115 116 114 105 110 103 34 58 34 116 101 120 116 34 44 34 102 108 111 97 116 34 58 49 46 49 49 44 34 105 110 116 101 103 101 114 34 58 53 125 125 93 125]
	in := container

	data, err := avro.Marshal(schema, in)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
	// Avro Outputs:
	//	[1 174 1 50 50 48 49 57 45 48 51 45 49 53 84 49 49 58 48 56 58 48 50 43 48 49 58 48 48 1 20 6 107 101 121 10 118 97 108 117 101 0 5 90 12 115 116 114 105 110 103 0 8 116 101 120 116 10 102 108 111 97 116 8 195 245 40 92 143 194 241 63 14 105 110 116 101 103 101 114 8 0 0 0 0 0 0 20 64 0 0]

	out := Container{}
	err = avro.Unmarshal(schema, data, &out)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
	// Container Output:
	//{[0xc00006e2c0]}
	var b []byte
	b, _ = json.Marshal(&out)
	fmt.Println(b)
	// Byte array Outputs:
	//	[123 34 109 101 116 114 105 99 115 34 58 91 123 34 116 105 109 101 115 116 97 109 112 34 58 34 50 48 49 57 45 48 51 45 49 53 84 49 49 58 48 56 58 48 50 43 48 49 58 48 48 34 44 34 109 101 116 97 100 97 116 97 34 58 123 34 107 101 121 34 58 34 118 97 108 117 101 34 125 44 34 100 97 116 97 34 58 123 34 115 116 114 105 110 103 34 58 34 116 101 120 116 34 44 34 102 108 111 97 116 34 58 49 46 49 49 44 34 105 110 116 101 103 101 114 34 58 53 125 125 93 125]
}
