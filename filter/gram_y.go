// Code generated by goyacc -o gram_y.go gram.y. DO NOT EDIT.

//line gram.y:2
package filter

import __yyfmt__ "fmt"

//line gram.y:2

import (
	"time"
)

//line gram.y:9
type yySymType struct {
	yys   int
	node  Node
	nodes []Node

	item Item

	strings   []string
	float     float64
	duration  time.Duration
	timestamp time.Time
}

const EQ = 57346
const COLON = 57347
const SEMICOLON = 57348
const COMMA = 57349
const COMMENT = 57350
const DURATION = 57351
const EOF = 57352
const ERROR = 57353
const ID = 57354
const LEFT_BRACE = 57355
const LEFT_BRACKET = 57356
const LEFT_PAREN = 57357
const NUMBER = 57358
const RIGHT_BRACE = 57359
const RIGHT_BRACKET = 57360
const RIGHT_PAREN = 57361
const SPACE = 57362
const STRING = 57363
const QUOTED_STRING = 57364
const NAMESPACE = 57365
const DOT = 57366
const operatorsStart = 57367
const ADD = 57368
const DIV = 57369
const GTE = 57370
const GT = 57371
const LT = 57372
const LTE = 57373
const MOD = 57374
const MUL = 57375
const NEQ = 57376
const POW = 57377
const SUB = 57378
const operatorsEnd = 57379
const keywordsStart = 57380
const AS = 57381
const ASC = 57382
const AUTO = 57383
const BY = 57384
const MATCH = 57385
const NOT_MATCH = 57386
const DESC = 57387
const TRUE = 57388
const FALSE = 57389
const FILTER = 57390
const IDENTIFIER = 57391
const IN = 57392
const NOT_IN = 57393
const AND = 57394
const LINK = 57395
const LIMIT = 57396
const SLIMIT = 57397
const OR = 57398
const NIL = 57399
const NULL = 57400
const OFFSET = 57401
const SOFFSET = 57402
const ORDER = 57403
const RE = 57404
const INT = 57405
const FLOAT = 57406
const POINT = 57407
const TIMEZONE = 57408
const WITH = 57409
const keywordsEnd = 57410
const startSymbolsStart = 57411
const START_STMTS = 57412
const START_BINARY_EXPRESSION = 57413
const START_FUNC_EXPRESSION = 57414
const START_WHERE_CONDITION = 57415
const startSymbolsEnd = 57416

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"EQ",
	"COLON",
	"SEMICOLON",
	"COMMA",
	"COMMENT",
	"DURATION",
	"EOF",
	"ERROR",
	"ID",
	"LEFT_BRACE",
	"LEFT_BRACKET",
	"LEFT_PAREN",
	"NUMBER",
	"RIGHT_BRACE",
	"RIGHT_BRACKET",
	"RIGHT_PAREN",
	"SPACE",
	"STRING",
	"QUOTED_STRING",
	"NAMESPACE",
	"DOT",
	"operatorsStart",
	"ADD",
	"DIV",
	"GTE",
	"GT",
	"LT",
	"LTE",
	"MOD",
	"MUL",
	"NEQ",
	"POW",
	"SUB",
	"operatorsEnd",
	"keywordsStart",
	"AS",
	"ASC",
	"AUTO",
	"BY",
	"MATCH",
	"NOT_MATCH",
	"DESC",
	"TRUE",
	"FALSE",
	"FILTER",
	"IDENTIFIER",
	"IN",
	"NOT_IN",
	"AND",
	"LINK",
	"LIMIT",
	"SLIMIT",
	"OR",
	"NIL",
	"NULL",
	"OFFSET",
	"SOFFSET",
	"ORDER",
	"RE",
	"INT",
	"FLOAT",
	"POINT",
	"TIMEZONE",
	"WITH",
	"keywordsEnd",
	"startSymbolsStart",
	"START_STMTS",
	"START_BINARY_EXPRESSION",
	"START_FUNC_EXPRESSION",
	"START_WHERE_CONDITION",
	"startSymbolsEnd",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line gram.y:458

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 11,
	7, 52,
	17, 52,
	-2, 10,
	-1, 12,
	7, 53,
	17, 53,
	-2, 8,
	-1, 20,
	15, 72,
	-2, 12,
	-1, 21,
	15, 73,
	-2, 13,
	-1, 107,
	15, 72,
	-2, 12,
}

const yyPrivate = 57344

const yyLast = 342

var yyAct = [...]int{
	20, 103, 56, 13, 69, 3, 21, 23, 68, 67,
	14, 122, 66, 109, 18, 34, 58, 61, 62, 63,
	64, 65, 34, 101, 59, 60, 29, 123, 12, 11,
	118, 68, 10, 73, 117, 72, 30, 16, 45, 46,
	47, 48, 51, 52, 53, 54, 55, 56, 57, 75,
	76, 77, 78, 79, 80, 81, 82, 83, 84, 85,
	86, 87, 88, 31, 49, 12, 11, 95, 95, 98,
	99, 120, 107, 96, 96, 105, 2, 74, 100, 46,
	108, 94, 97, 119, 53, 54, 71, 56, 70, 92,
	112, 112, 112, 112, 91, 90, 113, 113, 113, 113,
	111, 111, 111, 111, 114, 115, 116, 112, 89, 45,
	46, 7, 124, 113, 4, 53, 54, 111, 56, 57,
	121, 107, 129, 135, 105, 112, 131, 110, 110, 110,
	110, 113, 8, 112, 1, 111, 26, 19, 22, 113,
	24, 124, 124, 111, 110, 29, 134, 132, 15, 32,
	124, 6, 130, 128, 34, 30, 124, 124, 25, 40,
	42, 127, 133, 44, 17, 104, 39, 126, 125, 41,
	110, 5, 102, 43, 9, 28, 33, 0, 0, 37,
	38, 0, 31, 29, 0, 106, 15, 32, 0, 0,
	35, 36, 34, 30, 0, 27, 0, 40, 0, 0,
	0, 0, 29, 0, 39, 15, 32, 41, 0, 0,
	0, 34, 30, 0, 0, 0, 40, 37, 38, 0,
	31, 0, 0, 39, 0, 0, 41, 0, 35, 36,
	0, 58, 0, 27, 0, 0, 37, 38, 0, 31,
	0, 0, 0, 0, 0, 0, 93, 35, 36, 0,
	0, 0, 27, 45, 46, 47, 48, 51, 52, 53,
	54, 55, 56, 57, 29, 0, 0, 0, 32, 0,
	0, 0, 0, 34, 30, 0, 0, 0, 40, 49,
	0, 0, 0, 50, 0, 39, 0, 0, 41, 58,
	0, 0, 0, 0, 0, 0, 0, 0, 37, 38,
	58, 31, 0, 0, 0, 0, 0, 0, 0, 35,
	36, 45, 46, 47, 48, 51, 52, 53, 54, 55,
	56, 57, 45, 46, 47, 48, 51, 52, 53, 54,
	55, 56, 57, 0, 0, 0, 0, 49, 0, 0,
	0, 50,
}

var yyPact = [...]int{
	3, 104, 98, -1000, -1000, 126, -1000, 190, 98, 156,
	-1000, -1000, -1000, 285, -26, 190, -1000, -1000, -12, -15,
	-16, -20, -1000, -1000, -1000, -1000, -1000, 73, 71, -1000,
	-1000, 20, -1000, 17, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 190, 190, 190, 190, 190, 190,
	190, 190, 190, 190, 190, 190, 190, 190, 190, 94,
	81, 80, 75, 227, -1000, -1000, 14, 14, 14, 14,
	1, 171, -6, -1000, -1000, 52, -33, 83, 83, 296,
	12, 83, 83, -33, -33, 83, -33, 52, 83, 252,
	252, 252, 252, -1000, -1000, -16, -20, -1000, -1000, -1000,
	15, 11, 64, -1000, -1000, 285, 252, 7, 8, 150,
	-1000, -1000, -16, -20, 149, 143, 135, -1000, -1000, -1000,
	171, 134, 133, -1000, 252, -1000, -1000, -1000, -1000, -1000,
	-1000, 285, 252, -1000, 105, -1000,
}

var yyPgo = [...]int{
	0, 176, 175, 0, 174, 172, 171, 151, 37, 13,
	6, 21, 3, 1, 14, 165, 20, 32, 164, 10,
	158, 7, 140, 138, 137, 136, 134,
}

var yyR1 = [...]int{
	0, 26, 26, 26, 6, 6, 12, 12, 12, 12,
	12, 12, 19, 19, 10, 10, 1, 1, 21, 22,
	22, 20, 20, 16, 14, 24, 24, 5, 5, 5,
	5, 9, 9, 9, 8, 8, 8, 8, 8, 8,
	25, 13, 13, 13, 15, 15, 7, 7, 4, 4,
	4, 4, 17, 17, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 2, 2, 23, 23, 18, 18, 3, 3,
	3,
}

var yyR2 = [...]int{
	0, 2, 2, 1, 1, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 3, 3, 1, 1, 1, 1,
	1, 1, 1, 3, 4, 3, 3, 3, 2, 1,
	0, 3, 1, 0, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 3, 5, 3, 0, 1, 3,
	2, 0, 1, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 5, 5,
	5, 5, 1, 1, 1, 2, 4, 4, 1, 1,
	4,
}

var yyChk = [...]int{
	-1000, -26, 73, 2, 10, -6, -7, 13, 6, -4,
	-17, -11, -16, -12, -19, 15, -8, -18, -14, -24,
	-3, -10, -23, -21, -22, -20, -25, 62, -2, 12,
	22, 49, 16, -1, 21, 57, 58, 46, 47, 33,
	26, 36, -7, 17, 7, 26, 27, 28, 29, 52,
	56, 30, 31, 32, 33, 34, 35, 36, 4, 50,
	51, 43, 44, -12, -16, -11, 24, 24, 24, 24,
	15, 15, 15, 16, -17, -12, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, 14,
	14, 14, 14, 19, -14, -3, -10, -14, -3, -3,
	-21, 22, -5, -13, -15, -12, 14, -3, -21, -9,
	-8, -19, -3, -10, -9, -9, -9, 19, 19, 19,
	7, -9, 4, 19, 7, 18, 18, 18, 18, -13,
	18, -12, 14, -8, -9, 18,
}

var yyDef = [...]int{
	0, -2, 47, 3, 2, 1, 4, 51, 47, 0,
	48, -2, -2, 0, 36, 0, 6, 7, 9, 11,
	-2, -2, 34, 35, 37, 38, 39, 0, 0, 78,
	79, 0, 74, 0, 18, 19, 20, 21, 22, 40,
	16, 17, 5, 46, 50, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 8, 10, 0, 0, 0, 0,
	0, 30, 0, 75, 49, 54, 55, 56, 57, 58,
	59, 60, 61, 62, 63, 64, 65, 66, 67, 33,
	33, 33, 33, 23, 25, 72, 73, 26, 14, 15,
	0, 0, 0, 29, 41, 42, 33, -2, 0, 0,
	32, 36, 12, 13, 0, 0, 0, 76, 77, 24,
	28, 0, 0, 80, 0, 68, 69, 70, 71, 27,
	43, 44, 33, 31, 0, 45,
}

var yyTok1 = [...]int{
	1,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74,
}

var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:102
		{
			yylex.(*parser).parseResult = yyDollar[2].node
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:107
		{
			yylex.(*parser).unexpected("", "")
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:113
		{
			yyVAL.node = WhereConditions{yyDollar[1].node}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:117
		{
			if yyDollar[3].node != nil {
				arr := yyDollar[1].node.(WhereConditions)
				arr = append(arr, yyDollar[3].node)
				yyVAL.node = arr
			} else {
				yyVAL.node = yyDollar[1].node
			}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:133
		{
			yyVAL.node = &Identifier{Name: yyDollar[1].item.Val}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:137
		{
			yyVAL.node = yyDollar[1].node
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:143
		{
			yyVAL.node = &AttrExpr{
				Obj:  &Identifier{Name: yyDollar[1].item.Val},
				Attr: &Identifier{Name: yyDollar[3].item.Val},
			}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:150
		{
			yyVAL.node = &AttrExpr{
				Obj:  yyDollar[1].node.(*AttrExpr),
				Attr: &Identifier{Name: yyDollar[3].item.Val},
			}
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:163
		{
			yyVAL.node = &StringLiteral{Val: yylex.(*parser).unquoteString(yyDollar[1].item.Val)}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:169
		{
			yyVAL.node = &NilLiteral{}
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:173
		{
			yyVAL.node = &NilLiteral{}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:179
		{
			yyVAL.node = &BoolLiteral{Val: true}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:183
		{
			yyVAL.node = &BoolLiteral{Val: false}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:189
		{
			yyVAL.node = &ParenExpr{Param: yyDollar[2].node}
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:195
		{
			yyVAL.node = yylex.(*parser).newFunc(yyDollar[1].item.Val, yyDollar[3].nodes)
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:201
		{
			yyVAL.node = &CascadeFunctions{Funcs: []*FuncExpr{yyDollar[1].node.(*FuncExpr), yyDollar[3].node.(*FuncExpr)}}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:205
		{
			fc := yyDollar[1].node.(*CascadeFunctions)
			fc.Funcs = append(fc.Funcs, yyDollar[3].node.(*FuncExpr))
			yyVAL.node = fc
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:213
		{
			yyVAL.nodes = append(yyVAL.nodes, yyDollar[3].node)
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:218
		{
			yyVAL.nodes = []Node{yyDollar[1].node}
		}
	case 30:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:222
		{
			yyVAL.nodes = nil
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:228
		{
			nl := yyVAL.node.(NodeList)
			nl = append(nl, yyDollar[3].node)
			yyVAL.node = nl
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:234
		{
			yyVAL.node = NodeList{yyDollar[1].node}
		}
	case 33:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:238
		{
			yyVAL.node = NodeList{}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:252
		{
			yyVAL.node = &Star{}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:260
		{
			yyVAL.node = getFuncArgList(yyDollar[2].node.(NodeList))
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:266
		{
			yyVAL.node = &FuncArg{ArgName: yyDollar[1].item.Val, ArgVal: yyDollar[3].node}
		}
	case 45:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:270
		{
			yyVAL.node = &FuncArg{
				ArgName: yyDollar[1].item.Val,
				ArgVal:  getFuncArgList(yyDollar[4].node.(NodeList)),
			}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:279
		{
			yyVAL.node = yylex.(*parser).newWhereConditions(yyDollar[2].nodes)
		}
	case 47:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:283
		{
			yyVAL.node = nil
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:290
		{
			yyVAL.nodes = []Node{yyDollar[1].node}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:294
		{
			yyVAL.nodes = append(yyVAL.nodes, yyDollar[3].node)
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:299
		{
			yyVAL.nodes = nil
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:306
		{
			yyVAL.node = yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:310
		{
			yyVAL.node = yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:314
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:320
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:326
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:332
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:338
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:344
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:350
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			yyVAL.node = bexpr
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:355
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			yyVAL.node = bexpr
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:360
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:366
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			yyVAL.node = bexpr
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:371
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			yyVAL.node = bexpr
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:376
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 68:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:382
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[4].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 69:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:388
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[4].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 70:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:394
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[4].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 71:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:400
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[4].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:409
		{
			yyVAL.item = yyDollar[1].item
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:413
		{
			yyVAL.item = Item{Val: yyDollar[1].node.(*AttrExpr).String()}
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:420
		{
			yyVAL.node = yylex.(*parser).number(yyDollar[1].item.Val)
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:424
		{
			num := yylex.(*parser).number(yyDollar[2].item.Val)
			switch yyDollar[1].item.Typ {
			case ADD: // pass
			case SUB:
				if num.IsInt {
					num.Int = -num.Int
				} else {
					num.Float = -num.Float
				}
			}
			yyVAL.node = num
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:440
		{
			yyVAL.node = yylex.(*parser).newRegex(yyDollar[3].node.(*StringLiteral).Val)
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:444
		{
			yyVAL.node = yylex.(*parser).newRegex(yylex.(*parser).unquoteString(yyDollar[3].item.Val))
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:451
		{
			yyVAL.item.Val = yylex.(*parser).unquoteString(yyDollar[1].item.Val)
		}
	case 80:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:455
		{
			yyVAL.item.Val = yyDollar[3].node.(*StringLiteral).Val
		}
	}
	goto yystack /* stack new state and value */
}
