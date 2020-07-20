package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func gfwlist2Rules(outfile string) error {
	url := "https://raw.githubusercontent.com/gfwlist/gfwlist/master/gfwlist.txt"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	encoded, err := ioutil.ReadAll(resp.Body)
	encodedString := string(encoded)
	decoded, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return err
	}
	var output bytes.Buffer
	dict := make(map[string]bool)

	reader := bufio.NewReader(bytes.NewReader(decoded))
	processor := makeProcessor()
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		line = strings.TrimSpace(line)
		result := processor(line)
		if result == "" {
			continue
		}
		if _, ok := dict[result]; !ok {
			_, err = output.WriteString(result + "\n")
			if err != nil {
				return err
			}
		}
	}
	return ioutil.WriteFile(outfile, output.Bytes(), 0777)
}

func makeProcessor() func(string) string {
	topLevelDomainStrForWebURLExpand := "(?:com|net|org|gov|mil|edu|biz|info|pro|name|coop|travel|xxx|idv|aero|museum|mobi|asia|tel|int|post|jobs|cat|ac|ad|ae|af|ag|ai|al|am|an|ao|aq|ar|as|at|au|aw|az|ba|bb|bd|be|bf|bg|bh|bi|bj|bm|bn|bo|br|bs|bt|bv|bw|by|bz|ca|cc|cd|cf|cg|ch|ci|ck|cl|cm|cn|co|cr|cu|cv|cx|cy|cz|de|dj|dk|dm|do|dz|ec|ee|eg|eh|er|es|et|eu|fi|fj|fk|fm|fo|fr|ga|gd|ge|gf|gg|gh|gi|gl|gm|gn|gp|gq|gr|gs|gt|gu|gw|gy|hk|hm|hn|hr|ht|hu|id|ie|il|im|in|io|iq|ir|is|it|je|jm|jo|jp|ke|kg|kh|ki|km|kn|kp|kr|kw|ky|kz|la|lb|lc|li|lk|lr|ls|ma|mc|md|me|mg|mh|mk|ml|mm|mn|mo|mp|mq|mr|ms|mt|mu|mv|mw|mx|my|mz|na|nc|ne|nf|ng|ni|nl|no|np|nr|nu|nz|om|pa|pe|pf|pg|ph|pk|pl|pm|pn|pr|ps|pt|pw|py|qa|re|ro|ru|rw|sa|sb|sc|sd|se|sg|sh|si|sj|sk|sm|sn|so|sr|st|sv|sy|sz|tc|td|tf|tg|th|tj|tk|tl|tm|tn|to|tp|tr|tt|tv|tw|tz|ua|ug|uk|um|us|uy|uz|va|vc|ve|vg|vi|vn|vu|wf|ws|ye|yt|yu|yr|za|zm|zw|accountant|club|coach|college|company|construction|consulting|contractors|cooking|corp|credit|creditcard|dance|dealer|democrat|dental|dentist|design|diamonds|direct|doctor|drive|eco|education|energy|engineer|engineering|equipment|events|exchange|expert|express|faith|farm|farmers|fashion|finance|financial|fish|fit|fitness|flights|florist|flowers|food|football|forsale|furniture|game|games|garden|gmbh|golf|health|healthcare|hockey|holdings|holiday|home|hospital|hotel|hotels|house|inc|industries|insurance|insure|investments|islam|jewelry|justforu|kid|kids|law|lawyer|legal|lighting|limited|live|llc|llp|loft|ltd|ltda|managment|marketing|media|medical|men|money|mortgage|moto|motorcycles|music|mutualfunds|ngo|partners|party|pharmacy|photo|photography|photos|physio|pizza|plumbing|press|prod|productions|radio|rehab|rent|repair|report|republican|restaurant|room|rugby|safe|sale|sarl|save|school|secure|security|services|shoes|show|soccer|spa|sport|sports|spot|srl|storage|studio|tattoo|taxi|team|tech|technology|thai|tips|tour|tours|toys|trade|trading|travelers|university|vacations|ventures|versicherung|versicherung|vet|wedding|wine|winners|work|works|yachts|zone|archi|architect|casa|contruction|estate|haus|house|immo|immobilien|lighting|loft|mls|realty|academy|arab|bible|care|catholic|charity|christmas|church|college|community|contact|degree|education|faith|foundation|gay|halal|hiv|indiands|institute|irish|islam|kiwi|latino|mba|meet|memorial|ngo|phd|prof|school|schule|science|singles|social|swiss|thai|trust|university|uno|auction|best|bid|boutique|center|cheap|compare|coupon|coupons|deal|deals|diamonds|discount|fashion|forsale|free|gift|gold|gratis|hot|jewelry|kaufen|luxe|luxury|market|moda|pay|promo|qpon|review|reviews|rocks|sale|shoes|shop|shopping|store|tienda|top|toys|watch|zero|bar|bio|cafe|catering|coffee|cooking|diet|eat|food|kitchen|menu|organic|pizza|pub|rest|restaurant|vodka|wine|abudhabi|africa|alsace|amsterdam|barcelona|bayern|berlin|boats|booking|boston|brussels|budapest|caravan|casa|catalonia|city|club|cologne|corsica|country|cruise|cruises|deal|deals|doha|dubai|durban|earth|flights|fly|fun|gent|guide|hamburg|helsinki|holiday|hotel|hoteles|hotels|ist|istanbul|joburg|koeln|land|london|madrid|map|melbourne|miami|moscow|nagoya|nrw|nyc|osaka|paris|party|persiangulf|place|quebec|reise|reisen|rio|roma|room|ruhr|saarland|stockholm|swiss|sydney|taipei|tickets|tirol|tokyo|tour|tours|town|travelers|vacations|vegas|wales|wien|world|yokohama|zuerich|art|auto|autos|baby|band|baseball|beats|beauty|beknown|bike|book|boutique|broadway|car|cars|club|coach|contact|cool|cricket|dad|dance|date|dating|design|dog|events|family|fan|fans|fashion|film|final|fishing|football|fun|furniture|futbol|gallery|game|games|garden|gay|golf|guru|hair|hiphop|hockey|home|horse|icu|joy|kid|kids|life|lifestyle|like|living|lol|makeup|meet|men|moda|moi|mom|movie|movistar|music|party|pet|pets|photo|photography|photos|pics|pictures|play|poker|rodeo|rugby|run|salon|singles|ski|skin|smile|soccer|social|song|soy|sport|sports|star|style|surf|tatoo|tennis|theater|theatre|tunes|vip|wed|wedding|win|winners|yoga|you|analytics|antivirus|app|blog|call|camera|channel|chat|click|cloud|computer|contact|data|dev|digital|direct|docs|domains|dot|download|email|foo|forum|graphics|guide|help|home|host|hosting|idn|link|lol|mail|mobile|network|online|open|page|phone|pin|search|site|software|webcam|airforce|army|black|blue|box|buzz|casa|cool|day|discover|donuts|exposed|fast|finish|fire|fyi|global|green|help|here|how|international|ira|jetzt|jot|like|live|kim|navy|new|news|next|ninja|now|one|ooo|pink|plus|red|solar|tips|today|weather|wow|wtf|xyz|abogado|adult|anquan|aquitaine|attorney|audible|autoinsurance|banque|bargains|bcn|beer|bet|bingo|blackfriday|bom|boo|bot|broker|builders|business|bzh|cab|cal|cam|camp|cancerresearch|capetown|carinsurance|casino|ceo|cfp|circle|claims|cleaning|clothing|codes|condos|connectors|courses|cpa|cymru|dds|delivery|desi|directory|diy|dvr|ecom|enterprises|esq|eus|fail|feedback|financialaid|frontdoor|fund|gal|gifts|gives|giving|glass|gop|got|gripe|grocery|group|guitars|hangout|homegoods|homes|homesense|hotels|ing|ink|juegos|kinder|kosher|kyoto|lat|lease|lgbt|liason|loan|loans|locker|lotto|love|maison|markets|matrix|meme|mov|okinawa|ong|onl|origins|parts|patch|pid|ping|porn|progressive|properties|property|protection|racing|read|realestate|realtor|recipes|rentals|sex|sexy|shopyourway|shouji|silk|solutions|stroke|study|sucks|supplies|supply|tax|tires|total|training|translations|travelersinsurcance|ventures|viajes|villas|vin|vivo|voyage|vuelos|wang|watches|测试|集团|在线|公益|公司|移动|我爱你|商标|商城|中文网|中信|中国|中國|測試|网络|香港|台湾|台灣|机构|组织机构|世界|网址|游戏|新加坡|政务|परीक्षा|한국|ভারত|موقع|বাংলা|москва|испытание|қаз|онлайн|сайт|срб|테스트|орг|삼성|சிங்கப்பூர்|дети|мкд|טעסט|భారత్|ලංකා|ભારત|भारत|آزمایشی|பரிட்சை|संगठन|укр|δοκιμή|إختبار|мон|الجزائر|عمان|ایران|امارات|بازار|پاکستان|الاردن|بھارت|المغرب|السعودية|سودان|مليسيا|شبكة|გე|ไทย|سورية|рф|تونس|みんな|ਭਾਰਤ|مصر|قطر|இலங்கை|இந்தியா|فلسطين|テスト)\\b"
	domainReg := regexp.MustCompile("([a-zA-Z0-9][-a-zA-Z0-9]*\\.)+" + topLevelDomainStrForWebURLExpand)
	return func(line string) string {
		prefix, content := splitLine(line)
		if content == "" {
			return ""
		}
		switch prefix {
		case "!", "/", "@@":
			return ""
		default:
			// prefix: "", "||", "|"
			return domainReg.FindString(content)
		}
	}
}

func splitLine(line string) (prefix string, content string) {
	if line == "" {
		return "", ""
	}
	if len(line) > 1 {
		prefix = line[0:2]
		content = strings.TrimSpace(line[2:])
		switch prefix {
		case "@@", "||":
			return
		}
	}
	prefix = line[0:1]
	content = strings.TrimSpace(line[1:])
	switch prefix {
	case "!", "/", "|":
		return
	default:
		return "", line
	}
}
