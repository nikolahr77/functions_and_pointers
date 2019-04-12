package main

import "testing"



func TestPerson_SetName(t *testing.T) {
	v := Person{"Nikola", 21}
	if v.Name!="Nikola" || v.Age!=21{
		t.Errorf("There is a problem with the struct, want name: %s age : %d, got name: %s age : %d",v.Name,v.Age,"Nikola",21)
	}
	v.SetName("Pesho")
	if v.Name!="Nikola"{
		t.Errorf("The name has change, but it should't, want %s, got %s","Nikola",v.Name)
	}
}

func TestPerson_SetAge(t *testing.T) {
	v := Person{"Nikola", 21}
	if v.Name!="Nikola" || v.Age!=21{
		t.Errorf("There is a problem with the struct, want name: %s age : %d, got name: %s age : %d",v.Name,v.Age,"Nikola",21)
	}
	v.SetAge(25)
	if v.Age!=25{
		t.Errorf("The age was not set, want %d, got %d",25,v.Age)
	}
}