go install -v golang.org/x/tools/cmd/present@latest
go get golang.org/x/tools/cmd/present

present  -http 127.0.0.1:3999
present  -http 127.0.0.1:3999 -notes
present  -http 127.0.0.1:3999 -notes -use_playground false



https://commonmark.org/help/tutorial/