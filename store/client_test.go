package store

import (
	"encoding/hex"
	"testing"
)

func TestClient_GetObject(t *testing.T) {
	client, err := NewClient("../testrepo")
	if err !=nil{
		t.Fatal(err)
	}
	hashString := "cf48af45ad4fb7199c6e3c6b4de23f329c4027cf"
	hash,err := hex.DecodeString(hashString)
	if err!=nil{
		t.Fatal(err)
	}
	obj, err:=client.GetObject(hash)
	if err != nil{
		t.Fatal(err)
	}
	t.Log(string(obj.Data))
}