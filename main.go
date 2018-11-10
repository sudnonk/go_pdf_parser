package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"regexp"
	"strings"
)

type PDF struct {
	Version string
	Objs    []Object
}

type Text struct {
	Font Font   //フォント
	Text string //内容
	X    int    //X座標
	Y    int    //Y座標
	Tc   int    //文字間隔
	Tw   int    //単語間隔
	Tz   int    //文字の幅
	TL   int    //行間隔
	Tf   int    //フォントサイズ
	Tr   int    //レンダリングモード
	Ts   int    //テキストライズの方向と距離(|Ts|が距離で正負が方向)
}

type Font struct {
	Name     string
	Encoding string
	BaseFont string
}

var debug debugT

type debugT bool

func (d debugT) Println(msg ...interface{}) {
	if d {
		log.Println(msg...)
	}
}

func main() {
	app := cli.NewApp()

	app.Name = "PDF parser"
	app.Usage = "This app do just parse pdf."
	app.Version = "1.0.0"
	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "Echo version",
	}

	app.Action = func(ctx *cli.Context) error {
		fname := ctx.String("filename")
		verbose := ctx.Bool("verbose")
		debug = debugT(verbose)

		parsePDF(fname)

		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose, V",
			Usage: "If set, echo debug logs.",
		},
		cli.StringFlag{
			Name:  "filename, f",
			Usage: "Find from the list in `PATH`",
		},
	}

	app.Run(os.Args)
}

// obj > directory = stream
// trailer > directory
// top level = version, eof, obj,xref,trailer,startxref
var versionRegexp = regexp.MustCompile(`^%PDF-(.+)$`)
var eofRegexp = regexp.MustCompile(`^%%EOF$`)
var commentRegexp = regexp.MustCompile(`^%.+$`)
var objStartRegexp = regexp.MustCompile(`^(\d+) (\d+) obj$`)
var objEndRegexp = regexp.MustCompile(`^endobj$`)
var directoryStartRegexp = regexp.MustCompile(`^<<.*$`)
var directoryEndRegexp = regexp.MustCompile(`^.*>>$`)
var streamStartRegexp = regexp.MustCompile(`^stream$`)
var streamEndRegexp = regexp.MustCompile(`^endstream$`)
var xrefStartRegexp = regexp.MustCompile(`^xref$`)
var xrefSubRegexp = regexp.MustCompile(`^(\d+) (\d+)$`)
var xrefRegexp = regexp.MustCompile(`^(\d{10}) (\d{5}) (.)$`)
var trailerRegexp = regexp.MustCompile(`^trailer$`)
var startxrefRegexp = regexp.MustCompile(`^startxref$`)

//Parse file designated by `fname`, and return list of Packs.
func parsePDF(fname string) (PDF, error) {
	file, err := os.Open(fname)
	if err != nil {
		return PDF{}, err
	}
	defer file.Close()

	var pdf PDF

	isInObj, isInxref, isIntrailer := false, false, false

	coo := make(chan Object)
	clo := make(chan string)
	cxx := make(chan Xref)
	clx := make(chan string)
	ctt := make(chan Trailer)
	clt := make(chan string)

	s := bufio.NewScanner(file)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())

		//PDFのバージョンが記載されてる行
		if versionRegexp.MatchString(l) {
			re := versionRegexp.FindStringSubmatch(l)
			if len(re) != 2 {
				return pdf, fmt.Errorf("Failed to parse Version. %v", l)
			}

			pdf.Version = re[1]
			continue
		}
		//終端行
		if eofRegexp.MatchString(l) {
			break
		}

		//コメント行
		if commentRegexp.MatchString(l) {
			continue
		}

		//endobj
		if objEndRegexp.MatchString(l) {
			isInObj = false
			clo <- l
			pdf.Objs = append(pdf.Objs, <-cooo)
			continue
		}

		if isInObj {
			clo <- l
			continue
		}

		// 6 0 obj
		if objStartRegexp.MatchString(l) {
			isInObj = true
			go parseObj(coo, clo)
			clo <- l
			continue
		}

		if xrefSubRegexp.MatchString(l) {
			//todo: implement
			continue
		}

		if isInxref {
			clx <- l
			continue
		}

		if xrefStartRegexp.MatchString(l) {
			isInxref = true
			go parseXref(cxx, clx)
			clx <- l
			continue
		}

		if isIntrailer {
			clt <- l
			continue
		}

		if trailerRegexp.MatchString(l) {
			isInxref = false
			isIntrailer = true
			go parseTrailer(ctt, clt)
			clt <- l
			continue
		}

		if startxrefRegexp.MatchString(l) {
			isIntrailer = false
			//todo: implement
			continue
		}

		log.Println("Unknown line.")
	}
}

func parseObj(co chan<- Object, cl <-chan string) {

}

func parseXref(cx chan<- Xref, cl <-chan string) {

}

func parseTrailer(ct chan<- Trailer, cl <-chan string) {

}
