package helpers

import (
	"testing"
)

func TestConnect(t *testing.T) {
	if Db != nil {
		t.Errorf("variable Db not empty")
	}
	if Client != nil {
		t.Errorf("variable Client not empty")
	}
	if DbName != "" {
		t.Errorf("variable DbName not empty")
	}
	if Url != "" {
		t.Errorf("variable Url not empty")
	}

	err := Connect("mongodb://127.0.0.1:27017", "app")
	if err != nil {
		t.Errorf("error connect mongodb %v", err)
	}

	if Client == nil {
		t.Errorf("variable Client not empty")
	}
	if Db == nil {
		t.Errorf("variable Db empty")
	}
	if DbName == "" {
		t.Errorf("variable DbName empty")
	}
	if Url == "" {
		t.Errorf("variable Url empty")
	}
}

func TestConnectError(t *testing.T) {
	if Db == nil {
		t.Errorf("variable Db not empty")
	}
	if Client == nil {
		t.Errorf("variable Client not empty")
	}
	if DbName == "" {
		t.Errorf("variable DbName not empty")
	}
	if Url == "" {
		t.Errorf("variable Url not empty")
	}

	err := Connect("", "")
	if err == nil {
		t.Errorf("error connect mongodb %v", err)
	}

	if Db == nil {
		t.Errorf("variable Db not empty")
	}
	if Client == nil {
		t.Errorf("variable Client not empty")
	}
	if DbName != "" {
		t.Errorf("variable DbName not empty")
	}
	if Url != "" {
		t.Errorf("variable Url not empty")
	}
}

func TestGetCollection(t *testing.T) {
	err := Connect("mongodb://127.0.0.1:27017", "app")
	if err != nil {
		t.Errorf("error connect mongodb %v", err)
	}

	collection := GetCollection("accounts")
	if collection == nil {
		t.Errorf("error open collection mongodb %v", collection)
	}
}
