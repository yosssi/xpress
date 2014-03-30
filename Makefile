install:
	bower install
	make compile
	make deploy
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
	cat ./bower_components/bootflat/bootstrap/bootstrap.css > ./src/css/all.css
	cat ./bower_components/bootflat/css/font-awesome.css >> ./src/css/all.css
	cat ./bower_components/bootflat/css/bootflat.css >> ./src/css/all.css
	cat ./bower_components/bootflat/css/bootflat-extensions.css >> ./src/css/all.css
	cat ./bower_components/bootflat/css/bootflat-square.css >> ./src/css/all.css
	cat ./src/css/style.css >> ./src/css/all.css
catjs:
	cat ./bower_components/jquery/dist/jquery.js > ./src/js/all.js
	cat ./bower_components/bootflat/js/bootstrap.js >> ./src/js/all.js
	cat ./bower_components/underscore/underscore.js >> ./src/js/all.js
	cat ./bower_components/backbone/backbone.js >> ./src/js/all.js
cat:
	make catcss
	make catjs
uglifycss:
	uglifycss --ugly-comments ./src/css/all.css > ./static/css/all.min.css
uglifyjs:
	uglifyjs ./src/js/all.js -o ./static/js/all.min.js
uglify:
	make uglifycss
	make uglifyjs
compile:
	make stylus
	make cat
	make uglify
deploy:
	cp ./bower_components/bootflat/fonts/* ./static/fonts/
