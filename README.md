## About

Uuid in golang, only supports Version 1 (MAC address & date-time) now

## Usage

	if uuid, err := NewV1(nil); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(uuid.String())
	}
	
	str := uuid.String() // "b2bbeaa8-2c38-11e5-b341-14109fd7e511"
