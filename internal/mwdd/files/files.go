package files

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type staticFilesFile struct {
	data  string
	mime  string
	mtime time.Time
	// size is the size before compression. If 0, it means the data is uncompressed
	size int
	// hash is a sha256 hash of the file contents. Used for the Etag, and useful for caching
	hash string
}

var staticFiles = map[string]*staticFilesFile{
	"base.yml": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\x8c\x93\xdfk\xdb0\x10\xc7\xdf\xfdW\x1c\xed\xa0/\x8d݆\xc2\xc0PJ\xa1\x83\r\xb6\xa5\xb4\xd9^\x8d\"]k-\xd2I\xe8d\xbb\xdd\xd8\xff>d;\xb1[\xd2Q\xfc \xa3\xfb\xdc\xcf\xef\xa9\xc5\xc0\xdaQ\t'\xcb|y\x92e\x8c\xa1\xd5\x12\xb9\xcc2\x00\xe5\xb9\xcc\x00\x00\xb4\x15\x8fX\x82\u0087\x80:\n.\x14\xf1\xc2\a\xf7\xf4\xbcH\x0e\x18z\xaau\xa6\xb18\xba\x00,\xa0hE(BC\x85rr\x8b!g'\xb7\xe5\xa1\xcbޣv\x1cIؔ\xc7sn;\xa5r\xe3\xa40\xe9\xbe\a\bc\xe7\xc2v\x9f`_^_\xa2o/*\xa1T@\xe6\x12\xce\xcf\xf2\xf4\x9d\x9f\xa56\xe8Q\xd3\xd3P\xee\xc0\x1f\xc3zu\xb3*!\xa07B\"t:\xd6\xf0\xab\xd3Fa(f\xf4)8\x92\b\x8dW\"\xa2\x9aO\x82\xb5i\x05uV\xe1\x9c/\x8d\x888\x16\x8b\xd4\xea\xe0\xc8\"\xc5i ?\xbfܭ\u007f\\\u007f\xad>\xaf\xeeח\xbdˡ>\x13\x99\x88\xef\xd7\xdf>\xdd_\xe6\x16\x95\x16\x9d\xde\xeaW\xec\x80\x1eC\xaa[\x8a\xa0 \r\x0f\x02\xb23MԎN!ւ\xb6\f\xd1\xc1\xcd\xed\xfd\x14z\xbd\xbe\xadnWw\xeb\xcb\x0f\u007f\xd2\xf1w\x17HS\xc4@\u0080wa\xa8$\xfd\xcc\xf4<\x1a\x1d\xca\xf1<\xea-\n=\x92\xe2\xca\xd1D*ϯ4}\xb3[E\xb3\f\x93n\xa3\xa9b\x14A\xd6\x13\U0004ef58ҿs#\xa3\xf5/.\x82\xdb;䃾\x854\x1a)VV<U\x1b\xa7\x9e+ֿ1\x97\x8e\x1e\xca\x02\xa3\xdcA\x8e\x1er\xf5\x1f\xf6@\xe0\xa8-\xba&\xf2[\xc1^ڃ˲y\xbb\xd3\v\xf5\xc2\xee\xdaL\xa4~\x9c\x9eư\xee\x10k\xcdI\x86\x8dؘg\xe8\x04\xc5~568:4Al\f^\x81\v`]@\b\x82\x94\xb3W \xb8\a\x8c\x96\x91A\n\x82Zx\x8f\x04\x9a \xd6\b\xc3\xe0v\"̒&\xa5\xa0ad8\xff\xb8\xece=+\x96\x17{`\x01\xdcl\b\xe3\xfe\xb1&\xeb\xbf\x00\x00\x00\xff\xffCH\xcfn\x8f\x04\x00\x00",
		hash:  "d0e6a14c396eaa6260bb93cc8786cb22d40507a21b127babf7b062bdbe3dfda4",
		mime:  "",
		mtime: time.Unix(1619249758, 0),
		size:  1167,
	},
	"mediawiki.yml": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\x8c\x93oo\xda>\x10ǟ\xe7U\x9c\xf8\xfd\xa4J\x93\x9c\x00\xd36d\x89\aYa4\x1a\x94\xaa\x84v\xcf\"7>\x12\x8b\xc4\xf6l\x93\x80\xa6\xbe\xf7)І?\x9d\xba\x89G\xbe\xfb\xde縻o*4V(I\xe1\xaa\xef\xf7\xaf<Ϣ\xa9D\x8a\x96z\x1e@\x89\\\xb0Z\xac\x05\xf5\x00\x00D\xc92\xa4\xc0U\xbaFC\ff\xc2:\xb3\xf3\x1b\xc1^\xe9+\x93\x05\x1c\xab\xc0:\x83.͉\xce\xf5\x97>Y\xe9\x920\xcd\xd2\x1c\xfbd\xcb\xf1i\x93Ѯ\xff\xd9\xef\x11\xdb\xdbcQ:\xb3\xd3JHG\xa1\x13\x945\xe7\xc11\xe4ۼ\xb3WU\xaaؔ\xcd\xffj\x1e@\xc0\x0fj&\x1cY)CD\xa3\xa2\x97o\xa3\x8e\xd2\xe3 {\xfeI\xae\xf3\xff\xaf\xd9x\x14\x85\x8f\xd1\xf7(y\x98O\x97\xb3\xf1\"\xb9\x9e\x8f\xc6\xcf4\xa8\x98\t\xea\xba\x0erW\x16AM\xd3f\x04\xdey\xb7r\x14\xc6!%m;\u0099coHA\x13\xa5\x1c\v̘\xfb\x1b1\x9a\x85\x93\xf1┹?\x83}K=ă\xc3y\xfe\x15?\x9dO\xce\xe0\x85\xca^х\xcaN\x16w\xc1CY\t\xa3d\x89ҽ\x9c\x04\xfe\x83\xa5E\x0eO;\xa8\x98\x11jc\xa1dB:\x94L\xa6\b65B;\vN\xc1JH\x0e\xb3\x86\xfc(\xd6\xc2o\xcb\xc3\xc2*0\xf8s#\frX)\x03\xed\x88Br\xdc\xfa:\xd7@ wN[\x1a\x04:gOF\xa4\xcc)s\xe1\xc1\xb8\xf7\xe9\xe3`\xd0\u007f\x01\x13\x98=&\xd1\xed\"\x0e\xa7\xd3\xe4.\x8co\x86\x17\x9bku\x0f\xd1}\xbc\f\xa7\xc9\xcd|\x11\x0f?\xf8\xed\xf0~c\x1a\xbfP)+re\xdd\x1b\xfd\xdd\xfc>\x1e\x0e\xba\x83n\x9b\xf91\x1a\u007f]N\x92\xeb\xf9\xed\xb7h2<]\xfcY\x86\x92\xe7}IC\x95\xacD\n\xef\xf6\xe4\xa8Qr\x9b(I\xdbN2\x13rK\xb4Q\xdb\xddA#\xed1\xd9\xeb\xfaͯ\xd7}M%\x16\x99I\xf3\xa3\xe2\x0f]$\xbaZ\x99\xf5\t\x86k\xeby'\x1f\u0e7f\xcf#\a\x17\x9e\xc7\x1aS\xd1\xdf\x01\x00\x00\xff\xff\xa6\x85\xc1\xb2j\x04\x00\x00",
		hash:  "fb52785be77a00cfc1746b15f9fade0ef4b99619c9937ec6a8192d16a9337a5d",
		mime:  "",
		mtime: time.Unix(1619265173, 0),
		size:  1130,
	},
	"mediawiki/MwddSettings.php": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xccX[O\xe3H\x16~&\xbf\xe2L@\xed \x858\xbd#\xcd\x0e\xb0\x99^ \xe9\xe9\xecpɒ0< \x14U\xecc\xa7B\xb9\xca]U\xceE-\xfe\xfb\xeaT9!\t\xf4\x84^\xed\xc3\"E\xc2>߹ߪ\xfc\x8fO\xf98\xafT\xf6w\xfcU\xf6\xe1\xea\xbe݆\xfe\xe0\xecv\xb0\x1b]ه\x9eV\x16#\v,e\\\x1a\v3\x1c\x01J\xab\x17\x15\x9e@\r~\x8a1\xe1\x12\xe3\x1a\x04W\x9dv\xf7\xec\xbe\xfbG7\x80C8\x84o\x95=\x9cs{Zy\xfe\x01\xb3\xd0\x16\xf9\xbb̺\x1f\xa3\x84\xc2`\fS\xce\xe0\xe2\xb2[\xa7'\xb0c\x84\x18\x13V\b\v\xeds\xe0\tH\x05W\xf7C\xfa߀\xc91\xe2\tǘ\xf41.-J&#\x04\x13i\x9e[\x033n\xc7pt4\xe3O\x1crfH\xfa\x8c\v\x01\x06\xad\x17\xe2}\xee}\xe9\r\xfbg\xbd.\xb4Z-\b\"\xc1\x03\xf8\xf0a#\x12\x84]F\x01\x00\xc0\x93V\x94:\x04\xa5\x91\x01\x1c\xfa\x00\xc1Ua,\x8c\x10\xd8HM\x11\xeeqt\x8b_\v4\xf6\xe4$FJ@\x1f\xf5\x14u\xa3\xb2\x0f\xd9,\x8e\xc9Y\x03\fr\xad\xe6\v0\x8e\xe6͗j\x15\x80\\ik\x1a\x95\x83YzfL\x91aO\xab9Gsg\xb0\xed\x11\x94[\x15)\xd1# \xb4 a\xc2\xe0)Y\xd3\xe1v\x8cz\x15R\x1fA\x94S\x982]\a\xa5!E\xeb(\xedsH\xb4\xca\xdc\xff\xda[샴\x8aF\xd51Wעq\x10\xab\xe8\tu{\x04-/\xf9Կ\x9e\xa5\xdeIhA5\fW\xa8F\x861g\x94\x94\x06\xb9\xde\x10*bb\xac\x8c=\xf9\xb5Y=\xad<\x03\n\x83<\xa9\x01Ӛ-\x86O\xb8\x18\xe2\x9c\x1bkj\x10\xf4;\xb7\u007fvn\x87\xd7gW\x9d\xa0\x0e\aC\xff\xfcږ/\xca\xd8\x1e\xf3Q\xc0y.TL\xd9j\xac\xf1<l\xc8z\xa4\xbcm\xfb\xb2-\xeb\xa1\xf9\xf8ڳ\xefe\xb6v\xb8teY3\x9cL\xb8\x93l$\x10,e5\xe21\xc2ḷ1\\QH\xee\xa9N\xdb\xe7D\xa4D\xd5\\\"(IJ/sq\xd8x\xa9\xb0\x1b)\x16\x0eXe\xf1\x94\xea>\xae\xba\xca\xe1\x11\x1a\xea\x14;\xc6\x05DLR\x15\x1aDY\a&c\"\xcc\x10\x98F\x90\xca\x02\x97`\xd1XS9\xa0T\xf4\x97\xdc-x\xa8\xec\x05\xd9\xc2|\x15G\x1as\xc1#\x16@\xeb7\xaa\x12J\xd4h!Y\x86\xb5-\xc0!\xfcD\xfd\xb3\xc5\xf5\xaa\x93z_zw\xd7\xdd\xc1p\xd0\xe9\x0f\x028\xacW\xf6\x02\x8d17o)\xf0\x84Rp\x89z\x8f\xc0T\xb3|\xcc-\x1e\x19ˬ\x89\xdf\x12\xbd\r)\x95\xbc\xe2ܭ\xee\xf1\xf4\xfd\x03\xb1\xcd,\x1b1\x83\xbb\x19\xc2\x10>\xf3\xb4\xd0\b\xaa\xb0kY+\f\x97)\x98\xaf\x82[t\x9d\xeb2M\xc3p\xac\n\x11S\xb2]\x06\x1a\x8dO$cpӾ\x01\x9c\xa2\xb4\x05\x13b\x01\x89\xd2\x04\x97\xa9\x01\xc1\x9f\x10rel\xaa\xd1x\x19n:J\xc4\x18\x18dJ#,+\v\xcc\xc2X\xcc@%\x10\x8d1z\xe22mT\xa8G\x13.p՞\a\xdd\x1e4 \bcfY\x18@c\xad\x9b\x1a\x104\xbcѫm\xb2\"\x0e\x169B\v\x82\x92\xbe\xd67\xaf1η\xc0u\xc0\xc1,m\x9fS2_Z\xb5=:\xad8\xb3\xb6\xf8Z/\xd2K\xdd\xc4\xeb\xc7lY\xed{\xf4\xdb\v\xe2Q\x9bk\x8c\xac\xd2\vW5\xeb.\x05\xf5\x12BJ=u\xa9\xc6S\xec\"\xdfzO\xea=-\x11,\xf55\xde>\xbf\x19\xb6;\x9f\xcf\xee.\a\x9e$\x14\xf3%\xfa\x91\x9e\x1f땽G\xe7\xe0w\x1c\xf1!\xd8\xe1\a\x95\xba\x93Yu\xf0\xea.\xd3\v\x83ڽ\x0f\xb4R\xb6\xf4\x94v\xe6Lio\\`\x95\xd2\xc1\xff\xc6эi\xf3\xb052\x1e\xe1\x134\xe1d#\x1a{<\xa9\xed`r\xf1\xd8\b\xc8\xc3c\x19\x92W\xe1X\xb2\xfd_\x84e\x1fΦ\x8a\xc7`f\xdcFԚ\xb4\x004\xb2Xр\xb7J\x012-\x16P\xa3\xde\xc59\xcbr\x81\x10\x17\x9a\x90E\x1e3\x8b\x8d|\x9c\x1f:-\x19\x9b\x83`\xa9\xd3\xf3K\xf3\xcd\xf2:\xad\xec=W*{a\xe8\xe7\x048-\xc6\x16I\x02\xb5\x99\x1b!n\x00X\x05ј\xc9\x14\xdd\xc8\xf0\xcd\xf3鰬\xb8\\c\xc2\xe7\xb4\u05eb\xa7\xe5\xab\x01\xed\xb5\x9b\xdcr%\xa9\x12\xab\x9d\xeb\u07fbםVWJ\xd5>\xafC\xe90\\|9\xbb\xedw\x06\xad\x11\x97L/\xaa?t\x9a\xbc\xa5\x05\xb0\x1b͓\xda\xee5\xb2잛\xd1\x04#{\xc1\xa21Ք\xa7-\xcb&\x88\x043>g\x81\xd3}\xceқ>E\xca\xe5;(\xcb\xcc\x01\x1eJ\xc1'\xbf\xfc\xfc\xf7\xe3\x00^\x9a\xf8\xdd\xde\xfd^n\x1e\xf8\x00}\xb7{\xfe\x1bO߹Ֆ\xde{E/\xa7\xb4-\u070f\xa5\xa7O\xa7q&\xa0\xc7\xd2w\xec7:\xd0p\x8dCE\a\xf6\xe1\xb0ݽ\x1d\x0eݤ\xbd\xa2.\xf7\xa2H\x12\xd5vpJ\xc3~\xed\xa5y\b\b\xe62\xb5\x86?9q\x19s\xe8\xceܢ4\\\xc9+4\x86X>s\xb1\xc1\xf7\x17:\x8f\x98\xe0̠)u\xbf\u007f\xbb㔮<\x96\xd6\xebn&ri\xacft\xbc\xa4\xe4Q\xd7X]\xb8\xf3\xfa>\xf4\xa9\x12\xccR\x16Aݝ\xa6\xc7\xec\xd8\x1d\xa7gU\x1f\x13Nw\x1e\xb7\x02\xab4\x1e\x8fV\x93\xc7ӯвk\x12\x9e\xb3ȁzZQ\xc1WO!\f\xc9V\xea\xf3\xd5\xcb:Е\x10YL{~\xfd굦iu\xa8p\xa7\x83\x18\x126呒PӘ\xa9)\xc6\xee\x96R\x0e\xb1T\x1dV\xc2\xf0`\x96~.A-\xa8~\xdb\xf0\xe59lx\x83\xc3RN\x83G\xaa\xeaC\xf0'\xd3\\\x15\x06\xe2r!st\x81\xb8\xcbi\xa0\xad\xb6\xb4\x97\xd9\xed=\x87<\xa34\x87\xa5\xbco\xabH<\xfbP\f\xb2|\x8b镬\xe7\xd0f\xb9G\xbb\x89\xb0\x1b\x1f\x11\x8c\xec]\x11\x97\x19\xda\xf6\xf3\xaf\xac#o\xa9v>@\x1bGEZ)\x89\x97*ݰ \x9c2\x1d\n\x95\x86\xab\v\x94\xb7\xd51]\xaa\x94*\x9cpo\xb0\x871a\x1aB\xa5\xa4\x8eK>4hk\x10\xcc\xfd\xfb)\xd3Ø\x9b\\\xb0\xc50c\xf3a\x8c\xb9\x1d\au8\xfaH\xb7\x8d\x9d\xf0h\xccE\xacQ\xbe\x9f\xc3\x1f\xa7Jt\x05\xb5Vz\xa8\xb1,\x9d\xda+)KV\a4A\x1d\x1c\xbdl\xa0\xce<B\xb7w\xdah\x19\x17/}T\xd2\xfb\xff\xbe\xec8\xbeu\x82\x8bY\xbb\xc8\xf2\xfeW\x01\xdb\x1c\xeds\xc7p\u03a2'\xab}\xe3\xac1NQ\xa8<Ci\uf656\xee$\xbdF\xee\xb8\xfbݿؔ\xf9\xe4\x0f\xd0؍\xc6\xeed\x8c\v\xba\xbe\xb9&\xa2\xeb\x9bQ\x19\x02\xbd\x85\x88\xd9h\x8c\xfaӋ \x87ސ\x9f\xa1NQF\x8b\v%-\x8b\xack\xfce9\xfcs\xab\xfb{\xe5A\xa5\x8f2\xf6\x03\xfe\xbbP\xaf\xeeΠ~C%\xe3⬰c\x94\x96G\x8c\"\xbd\xe1ѵ\xb2<)\t\xa6\x0e\xb6\xd0\x12i\x82$\xc0\f\xddZb%\x03\v\x06e\xec\xbc\xf4ڈ\x87\xb4\r\x98xz\xf9R\xb1\"\xddS$\x04w\xb1[}\xc5\xd8\a7\xc3\xd7\xccu\x1d\xb7\x11\xff3!\xd4\xecB\xe5\x8b7hw\x06\xbb\xd2X&\xed\x85\xca2\u007fNY\x13\xfe\a.\xca\x01\x93j\x16\xe3\x1f蚮99n\x1a\xd6L&\x86\x1d7'I\x93M\x12\xc6\xcaዑF[\xe2&\xbf6\x8f\u007f\x9e\x1c7\xe9\xf7\xb7I\xa2\x8f'\x1f\x9b\xfe\x97|$JBO$a\xd2<N&\x1f'ɓ\x94Q$\xb3\xb9\xcc\xe6\xe5\b\xe8}\xe9\xc1\xa5\xf2\x81\\\x8dYw݂\f]>\xc7\xf9\x85\xe0t\r\n\vC\xa3 b\"\x1cq\x19\xfe\xe0\x9e\xea\\\xb7\xdf\xf1Qni\x01MyVX\x05\x14P\x9a\xec*\x01\xe5>'\x91\xb1\xa2\\x\xb4+\xddm\xd0\xf8\x0fHl5\xb7\x17@7\xde\x18\x94\x04\u007fІ\xa3\xdf\x1c\x12\xe8\xa1\xd1h\xfc'\x00\x00\xff\xff\x1fW\xe0\x92\xec\x14\x00\x00",
		hash:  "1d4ddbe8252c2ecd646f1f98ffbb73b92da53022637841b494f3b74f4012c4d1",
		mime:  "",
		mtime: time.Unix(1619265412, 0),
		size:  5356,
	},
	"mediawiki/MwddSpecialPage-aliases.php": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xb2\xb1/\xc8(\xe0\xe2R).HM\xceL\xcc\tHLOu\xcc\xc9L,N-V\xb0U\x88\x8e\xb5\xc6&\x13\xad\x9e\x9a\xa7\x1e\v\x92\xe7\xe2T\xf7-OIQW\xb0\xb5S\x88V\x80\xb2cu\xb8b\xad\xb9\x00\x01\x00\x00\xff\xff\x97Dl2Y\x00\x00\x00",
		hash:  "4ab3af2702babee43b165550753b199d3d25397464bd083293ed0e259fd9cf9a",
		mime:  "",
		mtime: time.Unix(1619264411, 0),
		size:  89,
	},
	"mediawiki/MwddSpecialPage.php": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xb4Rߋ\x131\x10~N\xfe\x8a\xa1\x14o\xb7x,\xben\x8f\x9e\x94>\x9c\xd0R\xc1\x83{P)\xb3ɘ\x8d\x97&K2\xf1*\xe7\xfd\xef\x92vՊ\">\xe8\xdb\xee|\xdf|?\x86\\]\x0f\xfd \xa5r\x98\x12l\x1e\xb4~3\x90\xb2\xe8\x80\x0eL^'\x18\xff_\xa3!x\x94R\f\xb9sV\xc1\x87\xec\x15\xdb\xe0a\xb7S\xc1'\x8eYqUã\x14b\xc0H\x9e\xdb\xf6\x1c\x81\x8b\xa2}\x01\xf5\\\x8a')E3\x9bI\x013x\x99\x88\xce-ږ\x0e\xa42SAO\x8c\x01#\xee!q\xb4\xde|\xf1\xd99\x98\xa6\xdc\x15r\xc1\x9b_\x13\x8d\n\xd5w\x1e\xfc\x9c\xeb7\xf8\\\na\\\xe8\xd0\xc1t_\x8e@\xf1\x93U\x94\xe6R\n\xd14p\xbb]m!d\x1e2\x83\xee\x80?\x0f\xf4\x1ct\xe7qO\xf0\f\x02\xf7\x14\xcb<\x11\xb3\xf5&AO\x91\xa4\x10S\xeem\xba\\\x18\xe2\xedq\xb7\xaa/\x17\xa8\xf5\xcd\xedf]\xc1䮷\xaa\x874Z\x01F\x82\x98\xbd\xb7\xde\\ON\x91\xfe\xbc\u007f\xd5tq\xf17̏)\xf8\x1dy\x15t)}\xde\x0f\xea\u007fj4\xb9\t\x0f\xa0\x03%X-\xc1\xa1\x01\x17\xc2\xfd\xff,\xf3nC\xda❽\xb7?\xbe\xbeuk[C\xfc\xca'F\xaf\xa8H\x18\xe2\xd5r\x1dP/ѕY\x1c\x87\x1b<\xac\xd1T\xf5\xdb\x17\xefO\xf78>ѧ\xaf\x01\x00\x00\xff\xff\x03\x01\x88\xd7\x1a\x03\x00\x00",
		hash:  "ec610b9429e18c498c19f00d1f4e8a3516f61c115fbdcfe9a6d68710d3e13706",
		mime:  "",
		mtime: time.Unix(1619265173, 0),
		size:  794,
	},
	"mediawiki/entrypoint.sh": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xffDαj\xf40\x10\x04\xe0\xdeO\xb1?\xfe\xbb`/\\\n\x87T)Ҧ\v\xa4\x96\xad\xb1\xb4\x9c%\x8b\xd5\xfaty\xfbpǑt\xbb0\xc3|\xfd?\x9e%\xf3\xecj캞>\xe0\xc5}\xc9Y\x88\xa9\xb56xg\x8e2\xe0+\xd9N3\xc8\xcd\x1bngS1P\x84\u07bfE\xe1\f\xe4E+\xc1\x16Zw%/\xeb\nE6\xaab\xa8\xdd\x12\xd3\xeei\x9a&\xe2\x8bSn\xadq\xb4\xb4qcI.\xa0\xb2ߗ3\xf4\xa6\xf8\x8cȄ+\x96\xc3@\x16A\x8aplN\t\xd9\xf4\xbb쒭\xeb)\x9a\x95\xfa\xca\x1c\xa0*669K\xba\xf1\xc7]\x03+\x97\xed\b\x92+\a1\xd9PY\xb1!\a\xf6\xb8\f\x8f\xc1'V\xac\x95#\x9c\xaf\x9c\\5\xe8\x03\xb1\xde\x1b\xd5\x14\xb6ġ\xc42\x9d\x86\xb5\xa4\xc1\x15\xb7D\x9c\x86\xab\xc7|\x04~\xff\r\x8f\x86T6g\xe8\x9f_:\xfec\x8e5\xd2\xff\xb7\xee'\x00\x00\xff\xff\xf4\x97\xe2,f\x01\x00\x00",
		hash:  "99e1f477f10531d53e933fe3b99bd1fa04fe3f9b2fe8f91a86381d750f8243ef",
		mime:  "text/x-sh; charset=utf-8",
		mtime: time.Unix(1619265173, 0),
		size:  358,
	},
	"mysql.yml": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xffL\x90Ak\xdc@\f\x85\xef\xfe\x15b{H\v\xddq\x92\xe3@\x0f\x81\xb4\xa5\xd0\xc5m6\x10z\n\x8aGY\x8b\xcch\x92\x91\xd6n\xfa\xeb\x8b\xc7!k|\x19\xeb{zz\xd2HE9\x8b\x87\xb3Kwy\xd64Je\xe4\x9e\xd47\x00\xe9U_\xe2\xfc\x00\xf8\x00\xb7\xddu\a\x18c\x9e\x80\x13\x1e\b\xfa\xa3ZN\xfc\x0f\x8d\xb3|\x06͉l`9@\xe4'\x82ݟ\xfd\xef\x9f\xf7?vW߿\xc2GT\x88Y\x0e\x80\nl\x90\x88L\xc1\x06\x82B/G.\x14\x80Ũ<bO\x9f\xd6ӎJp\xb7\xfb\xb6\xccS\x98\x06\x12@\t\xc0\x8fs\xb7\x12\xd0_Vs\xceզ\xaa\xf2\x90\xb00\x86\a\x1f\xd1H\xad\x12\x92\x91K\x96Db\xcb6\x00۷\x807]w{\xff\xebj\xbf\xbf\xebn\xae\xbfXΥ\n\x86\xac&\x98f\xbb\xf9\x06.M!\xb8\x98{\x8c3\xa9\x92 z2\xbb8w\xf3wq^+B6\xe5\xf2\xb4\xe2\xe1Y\xeb{\xcc\xf1\x98h\x05\xaa\xfd6\xa0\xa1oG,m䇶\xd6\xde\x15n\xf9o\x13\xb2\xf8v\x0e\xb2].\xff\xb6\x9a\x95\xd7\xe7\xccb\x1e6kڞ\x88\xd3aS\xc5}N\t%x\xd8T˰i\x9aU\xa0U\x94\xff\x01\x00\x00\xff\xff\xeetJ\xa9\x15\x02\x00\x00",
		hash:  "5fbeeb4b3528f7f2bdd670876de160e695dab778f50f844ab405075a1601efb6",
		mime:  "",
		mtime: time.Unix(1619265412, 0),
		size:  533,
	},
	"mysql/main/custom.cnf": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xffT\xca=\xae\x02!\x10\a\xf0\x9eSp\x01B\xf6\xf5TO\v+\x13cg\f\x99]ƅ\b\xcc\xca\xcc\x12\xbd\xbdq\xad\xec\xfe\x1f\xbfKy\xf1#\x87\xabbl\x1d\x9bI\xc1\r*\xd3l\xc6T\xdd\xf6}Ҷp\x86\x8ef]\x02\b\xb2\x1bT\x00\x81\x90\x9a\xb3\x1d\x9a\xcdi\xb4\x1b\xb7*\xd5Ja\xf4\xb7\xbcr\xf4\x99f\x0f\xe2\xa5=\xfdD\xa5$\xd1N\xff\xfd\x92\x82\x12)h\xa7\x8f~w8\xed\xffϊ\xefi1\x91X\xcc\x04S\xc4o\xafP\xd04d\xca\x1d\xdf\x01\x00\x00\xff\xff@mH7\xb6\x00\x00\x00",
		hash:  "2326c1d81574d6d6525c116b5d3a7c71e6d6d330eaaee8d4b3bd2e7bb65a8485",
		mime:  "",
		mtime: time.Unix(1619265412, 0),
		size:  182,
	},
	"mysql/main/entrypoint.sh": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\x8c\x91\xb1j$1\x10Ds}E\x1dsppܮ.X\xd8\xf4\xfe\xe0\x02\x1b\x1ck\xa5\x9e\x9d\xc6R\xf7\xb8\xa5\xf1x\xfe\xdehv\xc1\x0e\x1c8*\x90\xaa\xeaI\xdd\xc3\x0f\u007fa\xf1\x97P'\xe7\x06\xfc\x17<\xb1$]+x\xc4J(\xbaHC\x9b\b#gBb\xa3\xd8\xf2\x06nX9g\x90$,3\xa6\xf0\xcar\xc5\xf9|\xc6LV\xb8VV\xa9\b\x92P\xb6\xfa\x92\xb1\xaa\xfcj0\ni/\x8b*#_݀\a\x05Km\xfd\xfc\x8eR\x044*\xb3Z\xb0\xed\x0eT\xdb\xf6\xae\xa8\xf3\x86Ѵ\xf4\x12\xa3?]\x04q*\x9a:\xbe)\xfe\x9eN'g\x05\a\x1b\xe1\xa9E\xbf\xe3}\xe7\x1d\x93\xff}\x8c2\xba8×5\xa5C\\j\xd3\xe2oү\xbe\x88|2\xde\xc2\x1d\xb6c\xbeav\x03\x1e\xfb\v\xe9\x8d\xe2\xd2h\xff\xba\xd1u\xc9\xc1\xees\xf1(\xc18\xa4\vH\x9am\xb3\xb24\xe7\x97j>k\fy_N\xd2\xf8Lv\xf80\x1c넟\xff\xdc{\x00\x00\x00\xff\xff\x90\x80&\xed\xbc\x01\x00\x00",
		hash:  "99e324d8429452d5682f735dc203d282a5d67d39b40cdbfa80d553001969866e",
		mime:  "text/x-sh; charset=utf-8",
		mtime: time.Unix(1619265412, 0),
		size:  444,
	},
	"nginx/client_max_body_size.conf": {
		data:  "client_max_body_size 1024m;",
		hash:  "87e1334f7ba74c1f53d7fb2b1b035b14cb6105e94dce8c5d75303595cdc91831",
		mime:  "",
		mtime: time.Unix(1619249758, 0),
		size:  0,
	},
	"nginx/timeouts.conf": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff*(ʯ\xa8\x8cO\xce\xcf\xcbKM.\x89/\xc9\xccM\xcd/-Q042(\xb6\xe6\x82\xc8\x15\xa7\xe6\xa5\xc0%\x14\x14P\xe4\x8aR\x131\xe5\x00\x01\x00\x00\xff\xffp\x17(\x9eS\x00\x00\x00",
		hash:  "431b33796adb902aa6dfaa1d01bc7cf0365511d7c89b923a5066f529c5bc2958",
		mime:  "",
		mtime: time.Unix(1619249758, 0),
		size:  83,
	},
	"wait-for-it.sh": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xa4W[S\xdb\xc8\x12~ׯ\xe8\b\x1d\x8e͉\"\x9b:\xe7l\xad\x89Ȳ\xc0\x82\xabH`\xb1\xfd\xb0EQ\x94,\xb5\xd0\x14\xf2\x8c23\xe2\x12\x87\xfd\xed[s\xd1\xcd8\x89\x93\xe8\xc5\xd2\\z\xbe\xee\xfe\xe6\xeb\xf6֫\xa0\x14<\x98\x13\x1a \xbd\x87y$2g\v\x00f\x02AfD\x80\x889)$H\x06\x12\x85\x04\x92B\x04\xb7\xe4\x1e)L\x0f/ cB\x06\x05\xe3\x12\"\x8e\x10\xddG$\x8f\xe69:N\xbcHh\xb4\xc0\xd0\xeb\xcd#\x81\xea\x15\xbcA\xdfq0\xce\x18r\xde\xeb\xc3Rٺ\xba\x02\xef\xcf\xd9\xf8x\n>E\x18\xc2\xf5\xf5\x1e\xc8\f)\xa8u\xe0z\xbf\xb90\xdc\xdf\xde݃\x94\xc0\xb3\xe3\x94\"\xba\xc5^\xdfY:\x00\x00q$\xe1\xed[\x98M\x0eN\x8ea\u007f{י\xa9鑞\xf3,\x00\x8dp\xa4\x11^\xf9\xe2\x1a\xae|\t\x92,\x90\x95R}\xf8\x10\xb3\xc5\"\xa2\tD\xfcV\\\xeb\xad~\x06\xa7\xe7\x93)|\x06\xdfW\xbbC\xfde\x9eS&$0\x0e\xe3\v(i\x82\\\a\xc5\xec*\xe0\xe2\xfc\xd2\xecRǅ\xfa\xcb<*R\x1a\xc2ʞ\xaf=\a\xb9DN#I\xee1\u007fz\rO\xac\x04Q`L\xd2'\x15 \xed\x16(\xdc&\xf8\xa2\xf1Ӡ\x11\x1a\x88\x90\x9c\xc4r\xc5\xf29͟\x00\x1f1.%\x82(\xe7U\x04H\xaa-\xeb4\x8b2\x8e\x11\x13a\x8c}\xd4\xc6>\x96\x04Wm\xc1\x11\xa3\xff\x96\xc0JY\x94\n\xcf\x13\b\x19\xc9R\xc0\x02\x85J\x865 a:~\u007f|>3\xe1\xb1\xf1\x0f\xed\xd87C15\xeb\x81P\x10\x183\x9a\x88\xd7\xf0\t9\x83\x94q\xa0\xacʧ9ʇ\xc3\xf3\xf7\xef\x0f>\x1c\xc1\xc1\xe5ɤc\xe7\xd8\xfa\\9\xfc@d\xa6\xf3\x0eQ*U^*\xefSB\x89\xc8P8\x9aZ\xda.>\x12\tC\xe7\xd9q\x1e\"\"oR\xc6k\x1aZ\x1aW\x1e\xfa\xb7\x12\x065\x91k\xe7,\xf1\xc1\xad\x989\x02e\x89\xd0[톧H6\xf24i\f\xae\xca-`ic\xdc\xfa\xef\x1aL\xb9\xc0\x1f\xb5\xaf\f\xd7G\x18s)\xd1?BF\\\xdeH\x11z\xbd$\x92\b\xff\xf9\x97\xe8뉇\x8c\xe4\b\xe6z%\xac>غ?\x9e\xfc>\x9b\xfc\x05>~l]\xe3Nji\f\xfe'\x03\x044\x90\xce,GQ\xe62\xf4\xde5\x1e\xb5\xddSOO\x8b\xc2>\x04\t\xde\a2.\x02m+ж\xfa\xb0\xaf\x87i\x99簻\xbf=\xfc\x86q\xebl\xcb\x01\xb3F;0X\xef\x00\xd2d]\\\xbe\x92\x81vԉh\xf4\xd1\x12\xce\xeb\xf5\x8cM\xf0\xeb\xb0\xf7\xfb\xdd$WϜct\xb7\x0e\xbe\xc8\x11\v\x18ڴP\x132\x8e\xb2\xe4\xb4r\xaaMۛ\a\x1e\x15\x056\xf4݂1\x05Ƶ01\x10e\xa1\x15e2>\x19\u007f\x98BRr\xc5 ˓\x11dR\x16\xa3 ()y|#d\x14\xdf\xe1c\x9cE\xf4\x16\xdf\xc4l\x11D\xc1\xff~\xf9\xff\xaf\xbb\xed[a\xc5}=+*\x86{\x8a:\x8a\xe2\u007f\x9c\x1d\x9c4d\xf7\x06\xb5\xea\xf8~\x9c\x91<\xa9\x14ِ\xc8\n\xad\to#+\xf5\xfe헷d\x93\x13\u007f\xf8$\x9b\x94\x8b\xf1Q\xe8\xbdү\x92G\x05\xb8w$\xcf\xc1W\xe1\xf4\xbd\x8b\xf1\x91\v\xe3\x0f\x86\xfb*'\xa0\x86\xf4\xd7\xe5\xf1dv6\xadHj\xe3g\x06uu\xdcLTjՈ\xe3\x92sL,\xd5*-XU\x92Uq\xe8HA\xc5\"\x03B\xb1h\v\n\xceb\x14B\xe9f\xb9@*\x85ctA\x81ݪ\xc5ϱ\x02\x11G\x02\xc1\xf5\x86.\x90\x06\xf5\xceh\a\x9a\xab\xa3¬\xa3\xdb\xf3\x96\xc3 \x18\x05\xf0\xdcL*`\xa1\xb7\xac\xd6\\\r\xae\x9f\xeb9\x05\xb7=7l͉\x8c\xa4\x12\x1a\x11\xd8۫_m\x82\x9b3\x0eO\xc7gG\xe1p\xa3\xbd\xadJ\xd8\x18\xd0\x1c\xdf\xd0@\xbb.7\x16&\xd3\xcb\xf1\xe1\xa6&\xb2\x95\xf0\xb8ޮ\xbb*f\x9a\xb5a\b\xae۴TZA\xf6:ҡO\xd9]\x1f%M\xff\x9d\x17g-\x87[;\u1cfb\x11Ң\xdfM\xd6:\xa4\xfaJ\xfd\x1cR͞\x9d\x17g}\x0f\xd2V.\xec\x05Y\v\xb6\xba<?\x87\xb7\x92\x8f\x9du\x87~\x0fj\xbf\xdf]\xd6P\xfal\x1c\xaa\xe6\xd9Y_=\xbayƼ\x95'\xdda\xaf[\xd8\x02[\xabΌ\xdeQ\xf6@k-\x18\x817t\xbfn\nE\x14;\xbaL9&\xa6\xae\xe6\xaakC\xfa\xf93\xb8F\x87Vb\xect\x0e>\xe6\x9c\xf1\x91n\x8a)b\xa2\nW\xc1\xd9=I\x10\xa2\x95\xd6\xd8\xfeqycp\x19L)q\x9c*\xe2\xdeҾ\x8d\xfc\xdd\xff\x0e\x9e\x1d{\x15\xbd\xa5y\x19\xf9\x83g\xc7(\x84\xb7Կz\xc4\\yo\xa9\u007f\xf5\x88\xb3\x05q\x86\xf1\x9d.\xa2\x88\xba\x9b\xae\xfaV\x01)g\v\x98\x97\xe2i\xce\x1e\xdf}\xcfR\v\xee\xe6\xe2`z\x1az=\x8eQ^D2\x03\xaf\xf7\x90\x918\xab6\xf6\xfbN\x97\xa2z\x03\x84\u007f\x83kM\xb9/\xab\x87i\xd9Z\xaa\xd3.\x8a\xa1\xebK\xd7\xe9\x14O\xbb~\xf0\x85\xf5\xae\x8e\xab\x85\xa1C\xb5\xae\x17\xae\xfa\x905%O\xf7\xd8U\xb1\xa9Oެ\xbd^moꉕ#\xda\xfet\xa0\xbc\\\x9b\x92\x8eCgcx\xf5\x92\x92\xeb\xca\xf4\xf66x\x86=_j|\xd6\xd4m\xfbWm\xc1\x12|\r\x1c\xd3R\xe8Ƌ\xb5\xff\xa9\xd9\xda\xdb\\\xb1N\xc4Z\x85[\xedѐ\x9b0v\x96\xa6\xe4\x9f\x00\x00\x00\xff\xff\xaf\xabL\xf5\xf8\x0f\x00\x00",
		hash:  "6a2b7c49ead02dcaf820d25c8df99043aca0c0937a38c0c1a75725e0faa42326",
		mime:  "text/x-sh; charset=utf-8",
		mtime: time.Unix(1619249758, 0),
		size:  4088,
	},
}

// NotFound is called when no asset is found.
// It defaults to http.NotFound but can be overwritten
var NotFound = http.NotFound

// ServeHTTP serves a request, attempting to reply with an embedded file.
func ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	path := strings.TrimPrefix(req.URL.Path, "/")
	f, ok := staticFiles[path]
	if !ok {
		if path != "" && !strings.HasSuffix(path, "/") {
			NotFound(rw, req)
			return
		}
		f, ok = staticFiles[path+"index.html"]
		if !ok {
			NotFound(rw, req)
			return
		}
	}
	header := rw.Header()
	if f.hash != "" {
		if hash := req.Header.Get("If-None-Match"); hash == f.hash {
			rw.WriteHeader(http.StatusNotModified)
			return
		}
		header.Set("ETag", f.hash)
	}
	if !f.mtime.IsZero() {
		if t, err := time.Parse(http.TimeFormat, req.Header.Get("If-Modified-Since")); err == nil && f.mtime.Before(t.Add(1*time.Second)) {
			rw.WriteHeader(http.StatusNotModified)
			return
		}
		header.Set("Last-Modified", f.mtime.UTC().Format(http.TimeFormat))
	}
	header.Set("Content-Type", f.mime)

	// Check if the asset is compressed in the binary
	if f.size == 0 {
		header.Set("Content-Length", strconv.Itoa(len(f.data)))
		io.WriteString(rw, f.data)
	} else {
		if header.Get("Content-Encoding") == "" && strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
			header.Set("Content-Encoding", "gzip")
			header.Set("Content-Length", strconv.Itoa(len(f.data)))
			io.WriteString(rw, f.data)
		} else {
			header.Set("Content-Length", strconv.Itoa(f.size))
			reader, _ := gzip.NewReader(strings.NewReader(f.data))
			io.Copy(rw, reader)
			reader.Close()
		}
	}
}

// Server is simply ServeHTTP but wrapped in http.HandlerFunc so it can be passed into net/http functions directly.
var Server http.Handler = http.HandlerFunc(ServeHTTP)

// Open allows you to read an embedded file directly. It will return a decompressing Reader if the file is embedded in compressed format.
// You should close the Reader after you're done with it.
func Open(name string) (io.ReadCloser, error) {
	f, ok := staticFiles[name]
	if !ok {
		return nil, fmt.Errorf("Asset %s not found", name)
	}

	if f.size == 0 {
		return ioutil.NopCloser(strings.NewReader(f.data)), nil
	}
	return gzip.NewReader(strings.NewReader(f.data))
}

// ModTime returns the modification time of the original file.
// Useful for caching purposes
// Returns zero time if the file is not in the bundle
func ModTime(file string) (t time.Time) {
	if f, ok := staticFiles[file]; ok {
		t = f.mtime
	}
	return
}

// Hash returns the hex-encoded SHA256 hash of the original file
// Used for the Etag, and useful for caching
// Returns an empty string if the file is not in the bundle
func Hash(file string) (s string) {
	if f, ok := staticFiles[file]; ok {
		s = f.hash
	}
	return
}
