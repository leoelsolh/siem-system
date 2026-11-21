<div align="center">

# **Go SIEM-System**
![Status](https://img.shields.io/badge/status-in%20progress-yellow) ![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white) 

**Security Information & Event Management built on Go**

</div>

### **Why**
SIEM systems are an essential part of system defence when it comes to addressing security threats. Personally, I don't have a reason to use a SIEM in the same way a corporation would, so I thought I'd learn about them by making one instead of just reading the theory. 

This led to me starting this project, to learn more about detection pipelines and how they work, and of coarse, to improve my Go knowledge

### **How**
This system currently uses sample logs, parsing them from a file to a slice of strings which is then passed to `parseLogs` which will turn it from a slice of log lines in string format, to a slice of structs, with one line slotting into one struct. 

`collect → parse → detect → alert`

The slice is then passed back to main, where we loop over the range of logs, printing any log line that is flagged by the `Flag` function. 

```
├── alert.go
├── collect.go
├── detect.go
├── store.go 
├── main.go
├── parse.go
├── README.md
├── logs
│   └── sample.log
└── go.mod
```

### **Usage**

```
git clone https://github.com/leoelsolh/siem-system.git
cd siem-system
```

The sample logs live in `logs/sample.log` so feel free to use that sample, or your own log lines. But again, as of where this project stands today, you won't get the result of a CrowdStrike SIEM... 

Go version (1.25)
```
go run .
```