package main

import (
    "encoding/json"
    "net"
    "os"
    "os/exec"
    "time"
    "io/ioutil"
    "log"
    "net/http"
)

// Structs for response data
type CmdReq struct {
    Type string `json:"type"`
    Data string `json:"data"`
}

type CmdRes struct {
    OK  bool        `json:"ok"`
    Res interface{} `json:"res"`
    Err string      `json:"err,omitempty"`
}

type PingRes struct {
    Success bool          `json:"success"`
    Time    time.Duration `json:"time"`
}

type SysInfo struct {
    Host string `json:"host"`
    IP   string `json:"ip"`
}

// Core functionality interface
type Cmdr interface {
    Ping(host string) (PingRes, error)
    SysInfo() (SysInfo, error)
}

type cmd struct{}

func NewCmd() Cmdr {
    return &cmd{}
}

// Shortened ping function
func (c *cmd) Ping(host string) (PingRes, error) {
    start := time.Now()
    err := exec.Command("ping", "-c", "1", host).Run()
    return PingRes{Success: err == nil, Time: time.Since(start)}, err
}

// Fetch system information
func (c *cmd) SysInfo() (SysInfo, error) {
    host, err := os.Hostname()
    if err != nil {
        return SysInfo{}, err
    }
    ip, err := getIP()
    if err != nil {
        return SysInfo{}, err
    }
    return SysInfo{Host: host, IP: ip}, nil
}

// Helper function to get the first non-loopback IP address
func getIP() (string, error) {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return "", err
    }
    for _, addr := range addrs {
        if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() && ip.IP.To4() != nil {
            return ip.IP.String(), nil
        }
    }
    return "", nil
}

// Handle HTTP requests
func main() {
    http.HandleFunc("/exec", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
            return
        }

        var req CmdReq
        body, _ := ioutil.ReadAll(r.Body)
        json.Unmarshal(body, &req)

        cmd := NewCmd()
        var res CmdRes

        switch req.Type {
        case "ping":
            ping, err := cmd.Ping(req.Data)
            res = CmdRes{OK: err == nil, Res: ping, Err: errMsg(err)}
        case "sysinfo":
            info, err := cmd.SysInfo()
            res = CmdRes{OK: err == nil, Res: info, Err: errMsg(err)}
        default:
            res = CmdRes{OK: false, Err: "Unknown command"}
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(res)
    })

    log.Fatal(http.ListenAndServe(":8081", nil))
}

// Helper function to return error messages
func errMsg(err error) string {
    if err != nil {
        return err.Error()
    }
    return ""
}

