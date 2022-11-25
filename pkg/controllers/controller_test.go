package controllers

import (
	"encoding/json"
	"testing"
)

func TestMazeTraverer1(t *testing.T) {

	s := `{"forward": "exit"}`
	data := []byte(s)
	var Path map[string]interface{}
	json.Unmarshal(data, &Path)
	currentPath := make([]string, 0)
	pathFound := false
	ans := mazeTraverer(Path, currentPath, &pathFound)
	if len(ans) > 0 && ans[0] == "forward" {
		t.Logf("Success")
	} else {
		t.Error("Failed")
	}

}

func TestMazeTraverer2(t *testing.T) {

	s := `{"forward": "tiger", "left": "ogre", "right": "demon"}`
	data := []byte(s)
	var Path map[string]interface{}
	json.Unmarshal(data, &Path)
	currentPath := make([]string, 0)
	pathFound := false
	ans := mazeTraverer(Path, currentPath, &pathFound)
	if len(ans) == 0 {
		t.Logf("Success")
	} else {
		t.Error("Failed")
	}

}

func TestMazeTraverer3(t *testing.T) {

	s := `{"forward": "tiger", "left": {"forward": {"upstairs": "exit"}, "left": "dragon"}, "right": {"forward": "dead end"}}`
	data := []byte(s)
	var Path map[string]interface{}
	json.Unmarshal(data, &Path)
	currentPath := make([]string, 0)
	pathFound := false
	ans := mazeTraverer(Path, currentPath, &pathFound)
	if len(ans) == 3 && ans[0] == "left" && ans[1] == "forward" && ans[2] == "upstairs" {
		t.Logf("Success")
	} else {
		t.Error("Failed")
	}

}

func TestMazeTraverer4(t *testing.T) {

	s := `{"forward": "tiger", "left": {"forward": {"upstairs": "exit"}, "left": "exit"}, "right": {"forward": "dead end"}}`
	data := []byte(s)
	var Path map[string]interface{}
	json.Unmarshal(data, &Path)
	currentPath := make([]string, 0)
	pathFound := false
	ans := mazeTraverer(Path, currentPath, &pathFound)
	if len(ans) == 3 && ans[0] == "left" && ans[1] == "forward" && ans[2] == "upstairs" {
		t.Logf("Success")
	} else if len(ans) == 2 && ans[0] == "left" && ans[1] == "left" {
		t.Logf("Success")
	} else {
		t.Error("Failed")
	}

}
