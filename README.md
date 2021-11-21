# CVE-2021-41277 PoC

Metabase is an open source data analytics platform. Local File Inclusion issue has been discovered in some versions of metabase. Here is the PoC code in order to determine the target has this vulnerability or not. An adversary could read arbitrary files in metabase server.

## Build
```
go build -o CVE-2021-41277 main.go
```

## Install
```
go get github.com/tahtaciburak/CVE-2021-41277
```

## PoC
```
cat targets.txt | ./CVE-2021-41277
```