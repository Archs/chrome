all:popup_tests.js content_scripts_tests.js

popup_tests.js:go/popup_tests.go
	gopherjs build $^ -o $@

content_scripts_tests.js:go/content_scripts_tests.go
	gopherjs build $^ -o $@

clean:
	rm -rf *.js *.js.map