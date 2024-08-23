  go build -o without_pgo -gcflags="-m" main.go

# command-line-arguments
./main.go:44:6: can inline extractNames
./main.go:75:14: inlining call to fmt.Println
./main.go:76:14: inlining call to fmt.Println
./main.go:77:14: inlining call to fmt.Println
./main.go:80:30: inlining call to extractNames
./main.go:81:14: inlining call to fmt.Println
./main.go:55:13: inlining call to log.Println
./main.go:56:31: inlining call to http.ListenAndServe
./main.go:44:19: leaking param content: data
./main.go:45:15: make([]string, 0) escapes to heap
./main.go:59:16: leaking param: w
./main.go:59:39: r does not escape
./main.go:66:6: moved to heap: bios
./main.go:75:14: ... argument does not escape
./main.go:75:15: "Name:" escapes to heap
./main.go:75:40: bio.PersonalInfo.Name escapes to heap
./main.go:76:14: ... argument does not escape
./main.go:76:15: "University:" escapes to heap
./main.go:76:54: bio.Education.University.SchoolName escapes to heap
./main.go:77:14: ... argument does not escape
./main.go:77:15: "Current Job:" escapes to heap
./main.go:77:54: bio.WorkExperience.Job2.Role escapes to heap
./main.go:80:30: make([]string, 0) escapes to heap
./main.go:81:14: ... argument does not escape
./main.go:81:15: "Courses:" escapes to heap
./main.go:81:27: courseNames escapes to heap
./main.go:85:17: ([]byte)("Data processed successfully") escapes to heap
../../sdk/go1.22.6/src/log/log.go:405:24: leaking param: b to result ~r0 level=0
./main.go:55:13: ... argument does not escape
./main.go:55:14: "Starting server on :8080" escapes to heap
./main.go:55:13: func literal does not escape
./main.go:56:11: ... argument does not escape
./main.go:56:31: &http.Server{...} escapes to heap
$ 


go build -o with_pgo -pgo=auto -gcflags="-m" main.go

# command-line-arguments
./main.go:109:9: PGO devirtualizing interface call w.Write to http.(*response).Write
./main.go:51:6: can inline extractNames
./main.go:67:6: can inline processor
./main.go:89:13: inlining call to fmt.Println
./main.go:92:29: inlining call to extractNames
./main.go:109:9: inlining call to http.(*response).Write
./main.go:62:13: inlining call to log.Println
./main.go:63:31: inlining call to http.ListenAndServe
./main.go:51:19: leaking param content: data
./main.go:52:15: make([]string, 0, len(data)) escapes to heap
./main.go:67:16: leaking param: w
./main.go:67:39: r does not escape
./main.go:76:6: moved to heap: bioWrapper
./main.go:87:13: ... argument does not escape
./main.go:87:14: "Name:" escapes to heap
./main.go:87:39: bio.PersonalInfo.Name escapes to heap
./main.go:88:13: ... argument does not escape
./main.go:88:14: "University:" escapes to heap
./main.go:88:53: bio.Education.University.SchoolName escapes to heap
./main.go:89:13: ... argument does not escape
./main.go:89:14: "Current Job:" escapes to heap
./main.go:89:53: bio.WorkExperience.Job2.Role escapes to heap
./main.go:92:29: make([]string, 0, len(data)) escapes to heap
./main.go:93:13: ... argument does not escape
./main.go:93:14: "Courses:" escapes to heap
./main.go:93:26: courseNames escapes to heap
./main.go:103:42: bioWrapper escapes to heap
../../sdk/go1.22.6/src/log/log.go:405:24: leaking param: b to result ~r0 level=0
./main.go:62:13: ... argument does not escape
./main.go:62:14: "Starting server on :8080" escapes to heap
./main.go:62:13: func literal does not escape
./main.go:63:11: ... argument does not escape
./main.go:63:31: &http.Server{...} escapes to heap
$ 