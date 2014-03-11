install:
	bower install
run:
	go run main.go xpress
stop:
	ps aux | grep xpress | grep exe/main | grep -v grep | awk '{print $$2}' | xargs kill -9
rerun:
	make stop
	make run
stylus:
	stylus ./src/css/style.styl
catcss:
	cat ./bower_components/bootflat/bootstrap/bootstrap.css > ./dist/css/all.css
	cat ./bower_components/bootflat/css/font-awesome.css >> ./dist/css/all.css
	cat ./bower_components/bootflat/css/bootflat.css >> ./dist/css/all.css
	cat ./bower_components/bootflat/css/bootflat-extensions.css >> ./dist/css/all.css
	cat ./bower_components/bootflat/css/bootflat-square.css >> ./dist/css/all.css
	cat ./src/css/style.css >> ./dist/css/all.css
catjs:
	cat ./bower_components/jquery/dist/jquery.js > ./dist/js/all.js
	cat ./bower_components/bootflat/js/bootstrap.js >> ./dist/js/all.js
	cat ./bower_components/underscore/underscore.js >> ./dist/js/all.js
	cat ./bower_components/backbone/backbone.js >> ./dist/js/all.js
cat:
	make catcss
	make catjs
uglifycss:
	uglifycss --ugly-comments ./dist/css/all.css > ./public/css/all.min.css
uglifyjs:
	uglifyjs ./dist/js/all.js -o ./public/js/all.min.js
uglify:
	make uglifycss
	make uglifyjs
compile:
	make stylus
	make cat
	make uglify
