package main

const HELPMSG string = "" +
	"Usage:\n" +
	"  -st  --standalone        use \033[1mstandalone\033[0m app\n" +
	"  --api                    start \033[1mHTTP server\033[0m (default port: 8080)\n" +
	"  -h   --help              show help\n" +
	"________________________________________________________________________________\n"

const STANDHELPMSG string = "" +
	"Usage standalone:\n" +
	"  -j    --json              generate \033[1mJSON\033[0m file\n" +
	"  -x    --xml               generate \033[1mXML\033[0m file\n" +
	"  -c    --csv               generate \033[1mCSV\033[0m file\n" +
	"  -a    --amount            specify \033[1mamount\033[0m of books (default: 1)\n" +
	"  -o    --output            specify \033[1moutput directory\033[0m (default: ./output/)\n" +
	"\n" +
	"  -cov  --cover             generate book \033[1mcover\033[0m\n" +
	"  -cova --covamount         specify \033[1mamount\033[0m of covers (default: 1)\n" +
	"  -covo --covoutput         specify covers \033[1moutput directory\033[0m (default: ./covoutput/)\n" +
	"  -covw --covwidth          specify covers \033[1mwidth\033[0m (default: 400)\n" +
	"  -covh --covheight         specify covers \033[1mheight\033[0m (default: 650)\n\n" +

	"Example:\n" +
	"  ./booksgen -st -j -x -c -a 10000 -o ./\n" +
	"  Generates 3 \033[1mfiles (.json, .xml, .csv)\033[0m in \033[1m./\033[0m directory with \033[1m10000\033[0m books in them\n\n" +

	"  ./booksgen -st -cov -cova 10 -covo ./covers/\n" +
	"   Generates \033[1m10\033[0m book \033[1mcovers\033[0m in \033[1m./covers/\033[0m directory\n" +
	"________________________________________________________________________________\n"

const APIHELPMSG string = "" +
	"Usage HTTP server:\n" +
	"  -p   --port               specify \033[1mport\033[0m (default: 8080)\n" +
	"  -i   --ip                 specify \033[1mIP\033[0m (default: localhost)\n" +
	"  --log                     \033[1mlog\033[0m IP addresses of incoming requests in \033[1mSQLite database\033[0m\n" +
	"                            (id, ip, endpoint, amount, timestamp)\n" +
	"  -db  --dbpath             specify \033[1mdatabase file path\033[0m (default: ./requests.db)\n\n" +

	"Example:\n" +
	"  ./booksgen --api -p 6969\n" +
	"  Starts \033[1mHTTP server\033[0m on 127.0.0.1:\033[1m6969\033[0m\n\n" +

	"  ./booksgen --api --ip 0.0.0.0 --log \n" +
	"  Starts \033[1mHTTP server\033[0m on \033[1m0.0.0.0\033[0m:8080 and \033[1mstores logs into SQLite database\033[0m\n" +
	"  (server will be available on whole local network)\n" +
	"  \033[41m\033[30m!DO IT ONLY IF YOU KNOW WHAT YOU ARE DOING!\033[0m\n\n" +

	"Example request:\n" +
	"                                  ↓----[json|xml|csv]\n" +
	"  GET http://127.0.0.1:6969/books/xml?amount=1000\n" +
	"  Returns \033[1mXML\033[0m with \033[1m1000\033[0m books\n" +
	"________________________________________________________________________________\n"
