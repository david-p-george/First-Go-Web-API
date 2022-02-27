package main

import (
	"encoding/json"
	"log"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
)

type Verse struct {
	Pk          int
	Translation string
	Book        int
	Chapter     int
	Verse       int
	Text        string
	Comment     string
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World 👌")
	})

	app.Get("/api/:book/:chapter", func(c *fiber.Ctx) error {
		verses := readAndReturn()

		book := returnBookID(c.Params("book"))

		chapterNo, chapterNoErr := strconv.Atoi(c.Params("chapter"))

		if chapterNoErr != nil {
			log.Fatal("Error converting chapter number to string", chapterNoErr)
		}

		return c.JSON(returnArrayOfVerses(verses, book, chapterNo))
	})

	app.Listen(":3002")
}

func returnArrayOfVerses(verses []Verse, book int, chapterNo int) []Verse {
	var filteredArr []Verse

	for _, obj := range verses {
		if obj.Book == book && obj.Chapter == chapterNo {
			filteredArr = append(filteredArr, obj)
		}
	}

	return filteredArr
}

func readAndReturn() []Verse {
	content, err := ioutil.ReadFile("./esv.json")

	if err != nil {
		log.Fatal("Error while opening file", err)
	}

	var verses []Verse
	umErr := json.Unmarshal(content, &verses)

	if umErr != nil {
		log.Fatal("Error while unmarshaling json", umErr)
	}

	return verses
}

func returnBookID(bookAbbr string) int {
	switch bookAbbr {
	case "Genesis", "genesis", "gen", "GEN":
		return 1

	case "Exodus", "exodus", "exo", "EXO":
		return 2

	case "Leviticus", "leviticus", "lev", "LEV":
		return 3

	case "Numbers", "numbers", "num", "NUM":
		return 4

	case "Deuteronomy", "deuteronomy", "deu", "DEU":
		return 5

	case "Joshua", "joshua", "jos", "JOS":
		return 6

	case "Judges", "judges", "jud", "JUD":
		return 7

	case "Ruth", "ruth", "rut", "RUT", "RUTH":
		return 8

	case "1Samuel", "1samuel", "1sa", "1SAM", "1SA":
		return 9

	case "2Samuel", "2samuel", "2sa", "2SAM", "2SA":
		return 10

	case "1Kings", "1kings", "1ki", "1KI":
		return 11

	case "2Kings", "2kings", "2ki", "2KI":
		return 12

	case "1Chronicles", "1chronicles", "1ch", "1CH":
		return 13

	case "2Chronicles", "2chronicles", "2ch", "2CH":
		return 14

	case "Ezra", "ezra", "ezr", "EZR":
		return 15

	case "Nehemiah", "nehemiah", "neh", "NEH":
		return 16

	case "Esther", "esther", "est", "EST":
		return 17

	case "Job", "job", "JOB":
		return 18

	case "Psalms", "psalms", "psa", "PSA":
		return 19

	case "Proverbs", "proverbs", "pro", "PRO":
		return 20

	case "Ecclesiastes", "ecclesiastes", "ecc", "ECC":
		return 21

	case "SongofSolomon", "songofsongs", "songofsolomon", "sos", "sng", "SOS", "SNG":
		return 22

	case "Isaiah", "isaiah", "isa", "ISA":
		return 23

	case "Jeremiah", "jeremiah", "jer", "JER":
		return 24

	case "Lamentations", "lamentations", "lam", "LAM":
		return 25

	case "Ezekiel", "ezekiel", "eze", "EZE":
		return 26

	case "Daniel", "daniel", "dan", "DAN":
		return 27

	case "Hosea", "hosea", "hos", "HOS":
		return 28

	case "Joel", "joel", "joe", "JOEL", "JOE":
		return 29

	case "Amos", "amos", "amo", "AMOS", "AMO":
		return 30

	case "Obadiah", "obadiah", "oba", "OBA", "OBAD":
		return 31

	case "Jonah", "jonah", "jon", "JON":
		return 32

	case "Micah", "micah", "mic", "MIC":
		return 33

	case "Nahum", "nahum", "nah", "NAH":
		return 34

	case "Habbakkuk", "habbakkuk", "hab", "HAB":
		return 35

	case "Zephaniah", "zephaniah", "zep", "ZEP":
		return 36

	case "Haggai", "haggai", "hag", "HAG":
		return 37

	case "Zechariah", "zechariah", "zec", "ZEC":
		return 38

	case "Malachi", "malachi", "mal", "MAL":
		return 39

	case "Matthew", "matthew", "mat", "MAT":
		return 40

	case "Mark", "mark", "mar", "MAR":
		return 41

	case "Luke", "luke", "luk", "LUK":
		return 42

	case "John", "john", "joh", "JOH":
		return 43

	case "ActsofApostles", "acts", "act", "ACTS", "ACT":
		return 44

	case "Romans", "romans", "rom", "ROM":
		return 45

	case "1Corinthians", "1corinthians", "1cor", "1co", "1COR", "1CO":
		return 46

	case "2 Corinthians", "2corinthians", "2cor", "2co", "2COR", "2CO":
		return 47

	case "Galatians", "galatians", "gal", "GAL":
		return 48

	case "Ephesians", "ephesians", "eph", "EPH":
		return 49

	case "Philippians", "philippians", "phi", "PHI":
		return 50

	case "Colossians", "colossians", "col", "COL":
		return 51

	case "1 Thessalonians", "1thessalonians", "1the", "1th", "1THE", "1TH":
		return 52

	case "2 Thessalonians", "2thessalonians", "2the", "2th", "2THE", "2TH":
		return 53

	case "1 Timothy", "1timothy", "1tim", "1ti", "1TIM", "1TI":
		return 54

	case "2 Timothy", "2timothy", "2tim", "2ti", "2TIM", "2TI":
		return 55

	case "Titus", "titus", "tit", "TIT":
		return 56

	case "Philemon", "philemon", "phil", "PHIL":
		return 57

	case "Hebrews", "hebrews", "heb", "HEB":
		return 58

	case "James", "james", "jam", "JAM":
		return 59

	case "1 Peter", "1peter", "1pet", "1pe", "1PET", "1PE":
		return 60

	case "2 Peter", "2peter", "2pet", "2pe", "2PET", "2PE":
		return 61

	case "1 John", "1john", "1jo", "1joh", "1JOHN", "1JO", "1JOH":
		return 62

	case "2 John", "2john", "2jo", "2joh", "2JOHN", "2JO", "2JOH":
		return 63

	case "3 John", "3john", "3jo", "3joh", "3JOHN", "3JO", "3JOH":
		return 64

	case "Jude", "jude", "JUDE":
		return 65

	case "Revelation", "revelation", "rev", "REV":
		return 66
	}

	return 0
}
