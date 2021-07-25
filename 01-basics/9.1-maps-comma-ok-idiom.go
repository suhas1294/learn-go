m := map[string]int{
	"James":           32,
	"Miss Moneypenny": 27,
}
fmt.Println(m)

fmt.Println(m["does_not_exist"]) // 0

v, ok := m["does_not_exist"] // v=0, ok = false

if v, ok := m["does_not_exist"]; ok {
	fmt.Println("record exist");
}

// adding new element to map
m["new key"] = 23;

// ranging

for k, v := range m{
	fmt.Println(k, v);
}

// deleting an entry
delete(m, "new key");