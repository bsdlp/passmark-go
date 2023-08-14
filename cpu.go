package passmark

/*
	{
		"data": [
		    {
				"id": "4017",
				"name": "AArch64 rev 2 (aarch64)",
				"price": "NA",
				"cpumark": "2,274",
				"thread": "901",
				"value": "NA",
				"threadValue": "NA",
				"tdp": "NA",
				"powerPerf": "NA",
				"date": "Feb 2021",
				"socket": "Unknown",
				"cat": "Unknown",
				"speed": "2189",
				"turbo": "NA",
				"cpuCount": 1,
				"cores": "8",
				"logicals": "1",
				"secondaryCores": "0",
				"secondaryLogicals": "0",
				"rank": 2425,
				"samples": "51",
				"href": "AArch64+rev+2+%28aarch64%29&amp;id=4017",
				"output": true
		    }
		]
	}
*/
type CPU struct {
	ID                string      `json:"id"`
	Name              string      `json:"name"`
	Price             string      `json:"price"`
	Cpumark           string      `json:"cpumark"`
	Thread            string      `json:"thread"`
	Value             string      `json:"value"`
	ThreadValue       string      `json:"threadValue"`
	TDP               string      `json:"tdp"`
	PowerPerf         string      `json:"powerPerf"`
	Date              string      `json:"date"`
	Socket            string      `json:"socket"`
	Cat               string      `json:"cat"`
	Speed             string      `json:"speed"`
	Turbo             string      `json:"turbo"`
	CPUCount          interface{} `json:"cpuCount"`
	Cores             string      `json:"cores"`
	Logicals          string      `json:"logicals"`
	SecondaryCores    interface{} `json:"secondaryCores"`
	SecondaryLogicals interface{} `json:"secondaryLogicals"`
	Rank              int         `json:"rank"`
	Samples           interface{} `json:"samples"`
	Href              string      `json:"href"`
	Output            bool        `json:"output"`
}

func (cpu *CPU) Link() string {
	//  "href": "AArch64+rev+2+%28aarch64%29&amp;id=4017",
	// https://www.cpubenchmark.net/cpu_lookup.php?cpu=AArch64+rev+2+%28aarch64%29&id=4017
	return ""
}
