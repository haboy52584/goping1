package main

import (
    "testing"
    "time"
)

// Test for GetSystemInfo
func TestGetSystemInfo(t *testing.T) {
    cmd := NewCmd()
    info, err := cmd.SysInfo()

    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if info.Host == "" {
        t.Error("Expected hostname to be non-empty")
    }

    if info.IP == "" {
        t.Error("Expected IP address to be non-empty")
    }
}

// Test for Ping
func TestPing(t *testing.T) {
    cmd := NewCmd()
    host := "google.com" // Or any reliable host
    result, err := cmd.Ping(host)

    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if !result.Success {
        t.Error("Expected ping to be successful")
    }

    if result.Time > time.Second {
        t.Error("Expected ping time to be less than 1 second")
    }
}

