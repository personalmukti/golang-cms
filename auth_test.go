package main

import "testing"

func TestRegisterAndLogin(t *testing.T) {
	users = make(map[string]*User) // reset store
	token, err := Login("user", "pass")
	if err == nil || token != "" {
		t.Fatalf("expected login failure for unknown user")
	}
	if err := Register("user", "pass", "user@example.com"); err != nil {
		t.Fatalf("register failed: %v", err)
	}
	if err := Register("user", "pass", "user@example.com"); err == nil {
		t.Fatalf("expected duplicate register error")
	}
	token, err = Login("user", "pass")
	if err != nil || token == "" {
		t.Fatalf("login failed: %v", err)
	}
}

func TestRolePermission(t *testing.T) {
	users = make(map[string]*User)
	roles = make(map[string]*Role)
	if err := Register("admin", "secret", "admin@example.com"); err != nil {
		t.Fatalf("register failed: %v", err)
	}
	AddRole("admin", []Permission{"read", "write"})
	if err := AssignRole("admin", "admin"); err != nil {
		t.Fatalf("assign role failed: %v", err)
	}
	if !HasPermission("admin", "write") {
		t.Fatal("expected to have write permission")
	}
	if HasPermission("admin", "delete") {
		t.Fatal("unexpected delete permission")
	}
}
