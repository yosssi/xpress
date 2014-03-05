install:
	bower install
	sudo npm install
run:
	go run main.go xpress &
stop:
	ps aux | grep xpress | grep exe/main | grep -v grep | awk '{print $$2}' | xargs kill -9
rerun:
	make stop
	make run
stylus:
	stylus ./src/css/style.styl
catcss:
	cat ./bower_components/bootstrap/dist/css/bootstrap.css ./src/css/style.css > ./dist/css/all.css
catjs:
	cat ./bower_components/bootstrap/dist/js/bootstrap.js ./src/js/app.js > ./dist/js/all.js
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
