all:app.js background.js

app.js:go/app.go
	gopherjs build $^ -o $@

background.js:go/background.go
	gopherjs build $^ -o $@