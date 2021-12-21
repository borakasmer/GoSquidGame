package shared

var Config = configuration{
	FILEURL:  `/Users/borakasmer/go/src/SquidGameGo/SquidPlayer.csv`,
	MONGOURL: "mongodb://localhost/squid",
	SPERATOR: ';',
	ISHEADER: true,
	ROWCOUNT: 5,
}

type configuration struct {
	FILEURL  string
	MONGOURL string
	ISHEADER bool
	SPERATOR rune
	ROWCOUNT int
}
