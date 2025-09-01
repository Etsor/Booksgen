# Booksgen
Books metadata generator with support for JSON, CSV, and XML output formats, book cover generation, and both CLI and HTTP API interfaces.

> [!WARNING]
> Project is in a very early stage of development. Especially the API and cover generation
>
> So far cover generation is only available in standalone app  

## Output examples
```json
{
  "books": [
    {
      "isbn": "5858958900109",
      "title": "Breeze Forest",
      "author": "Joshua Brown",
      "genre": "Biography",
      "publisher": "Tor Books",
      "year": 1536,
      "pages": 615
    },
    {
      "isbn": "0210219808245",
      "title": "Flame Village",
      "author": "Penelope Carter",
      "genre": "Mythology",
      "publisher": "Hachette Livre",
      "year": 448,
      "pages": 1392
    }
  ]
}
```

```xml
<?xml version="1.0" encoding="UTF-8"?>
<books>
    <book>
        <isbn>5858958900109</isbn>
        <title>Breeze Forest</title>
        <author>Joshua Brown</author>
        <genre>Biography</genre>
        <publisher>Tor Books</publisher>
        <year>1536</year>
        <pages>615</pages>
    </book>
    <book>
        <isbn>0210219808245</isbn>
        <title>Flame Village</title>
        <author>Penelope Carter</author>
        <genre>Mythology</genre>
        <publisher>Hachette Livre</publisher>
        <year>448</year>
        <pages>1392</pages>
    </book>
</books>
```

```csv
isbn,title,author,genre,publisher,year,pages
8780226264412,Mind Dream,Matthew Green,Biography,Pearson,956,555
5542422879817,Echo Thief,Jack Anderson,Thriller,Grand Central Publishing,1248,1980
9423048128989,Roots Petal,Jonathan Green,Urban Fantasy,Little, Brown and Company,1630,1699
9699057371624,Echo Rose,Lucas Perez,Children's,Random House Children's Books,431,1201
1292087887917,Circle Rift,Avery Smith,Children's,Little, Brown and Company,1762,1601
4599389468857,Realm Storm,James Anderson,Fairy Tale,Scholastic,157,64
2029166779416,Shield Mage,Amelia King,Steampunk,Orbit Books,728,2044
3209489817574,Truth Legend,Ethan King,Fantasy,Knopf Doubleday,1591,1661
3819262826174,Roots Twilight,Penelope Thompson,Detective,Hachette Livre,1223,1995
9070815733095,Fate Island,Amelia Wright,Mythology,Europa Editions,801,969
```


<p>
<img src="res/github/example1.png" width="200">
<img src="res/github/example2.png" width="200">
<img src="res/github/example3.png" width="200">
</p>

## Installation

Requires Go 1.22 or higher

```sh
git clone https://github.com/Etsor/Booksgen
cd Booksgen
./build.sh
```

## Usage
```sh
-st  --standalone        use standalone app
--api                    start HTTP server (default port: 8080)
-h   --help              show help
```

## Usage standalone
```sh
-j    --json             generate JSON file
-x    --xml              generate XML file
-c    --csv              generate CSV file
-a    --amount           specify amount of books (default: 1)
-o    --output           specify output directory (default: ./output/)

-cov  --cover            generate book cover
-cova --covamount        specify amount of covers (default: 1)
-covo --covoutput        specify covers output directory (default: ./covoutput/)
-covw --covwidth         specify covers width (default: 400)
-covh --covheight        specify covers height (default: 650)
```
### Examples
```sh
./build/booksgen -st -j -x -c -a 1000000 -o ./
```
Generates 3 files (.json, .xml, .csv) in ./ directory with 1000000 books in them

```sh
./build/booksgen -st -cov -cova 100 -covo ./covers/
```
Generates 100 book covers in ./covers/ directory

## Usage HTTP server:
```sh
-p   --port              specify port (default: 8080)
-i   --ip                specify IP (default: localhost)
--log                    log IP addresses of incoming requests in SQLite database
                         (id, ip, endpoint, amount, timestamp)
-db  --dbpath            specify database file path (default: ./requests.db)
```

### Examples
```sh
./build/booksgen --api -p 6969
```
Starts HTTP server on 127.0.0.1:6969

```sh
./build/booksgen --api --ip 0.0.0.0 --log
```
Starts HTTP server on 0.0.0.0:8080 and stores logs into SQLite database

#### Example request
```sh
                                 ↓----[json|xml|csv]
curl http://127.0.0.1:6969/books/xml?amount=1000
```

Returns XML with 1000 books


## Dependencies

- [gg](https://github.com/fogleman/gg) - Graphics library for covers generation
- [go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite driver