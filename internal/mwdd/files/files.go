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
		mtime: time.Unix(1616096260, 0),
		size:  1167,
	},
	"mediawiki.yml": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\x84\x93_o\xda0\x14\xc5\xdf\xf3)\xae\xbaI\x95&\xd9\x06\xa6m\xc8\x12\x0f\xace4\x1a\x94\xaa@\xbb\xb7ȍ/\x89Eb{\xb6I\xa8\xa6~\xf7)\xb4\r\u007f:u\xe2ɾ\xe7\xfe\x8e\xef\xe1\xa6B\xe7\x95\xd1\x1c\xce{\xb4w\x1eE\x1e]\xa5R\xf4<\x8a\x00J\x94J\xd4j\xadx\x04\x00\xa0J\x91!\ai\xd25:\xe20S>\xb8G\xda\bvJj\\\xc6$V\xcc\a\x87!͉\xcd\xed\xb7\x1eYْ\b+\xd2\x1c{d+\xf1a\x93\xf1\x0e\xfdJ\xbb\xc4ww\xd8\xca\x14\x9b\xb2ql\x0e@\x80\xb2Z\xa8@V\xc6\x11\x15\xa8\xcf\xf9\xe9ٙ\xbdt\xffDV\xd6R\x1e\xd4\xce>\xfe\x99\x8e.\xe3\xe1}\xfc3N\xeef\x93\xe5t4O.f\x97\xa3'\xce*\xe1X]\xd7,\x0fe\xc1j\x9e6\x8f\x93g\xefv\xc6\xd3\xe1x4\xe7\xa45$\xbb8\xfc\x1b\x1a{\xbeg\xcf1q\x89\x05f\"\xfc\x0f?\x99\x8d\x8f\xe0\x85\xc9^х\xc9\x0e\xc6<ᡮ\x943\xbaD\x1d^\x02\x84\x0f\xb0\xf4(\xe1\xe1\x11*\xe1\x94\xd9x(\x85\xd2\x01\xb5\xd0)\x82O\x9d\xb2\xc1C0\xb0RZ´!߫\xb5\xa2m\xfb\xb0\xf0\x06\x1c\xfe\xde(\x87\x12V\xc6A;\xa2\xd2\x12\xb7\xd4\xe6\x16\b\xe4!X\xcf\x19\xb3\xb9xp*\x15\xc1\xb8\x93]Xt\xbf|\xee\xf7{/`\x02\xd3\xfb$\xbe\x9e/\x86\x93Ir3\\\\\rN\x92kuw\xf1\xedb9\x9c$W\xb3\xf9b\xf0\x89\xb6\xc3\xd3\xe6/\xa6\x85IE\x91\x1b\x1f\xde\xe8of\xb7\x8bA\xbf\xd3ﴕ_\x97\xa3\xef\xcbqr1\xbb\xfe\x11\x8f\a\x87\xc1\x1fU8yڵ4T-J\xe4\xf0\xae\xa7D\x8bZ\xfa\xc4h\xde:\xe9L\xe9-\xb1\xcel\x1f\x9f5\xda\xef\x8b\xdd\x0em~\xdd\xcek)\xf1(\\\x9a\xef\x15\xffp\xd1\x18j\xe3\xd6\a\x18i}\x14\x1d|.\xa7\xbbx|\u05ec\x10\xff\x1b\x00\x00\xff\xff$e\xe0a\xe0\x03\x00\x00",
		hash:  "44b00e209e0d40cea69db7c3d00dc09ba1e6816f2cc9cee0f72ba56a21b3f187",
		mime:  "",
		mtime: time.Unix(1616096225, 0),
		size:  992,
	},
	"mediawiki/MwddSettings.php": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xc4WmO#9\x12\xfeL~Em\x18M'RHgn\xa5\xbd\x05.7\atf'7\xbc\xe4HX> \x149\xed\xean\x87n\xbb\xc7v\xde4\u2fdf\xca\xee\x84\x04\xe6\x06\xe6\xbe,R\xa4\xc4\xf5T\x95\xeb\xedq\xf1\x8f\x8feV\xd6j\xfb\xaf\xfc\xd5\xf6\xe1\xe26\x8a`8:\xb9\x1e\xbd\x8e\xae\xed\xc3@+\x8b\xb1\x05\x962!\x8d\x85\x05N\x00\xa5ի\x9aH\xa0\x01\xbfpL\x84Dހ\xe0\xa2\x17\xf5On\xfb_\xfa\x014\xa1\t\xdfj{\xb8\x14\xf6\xb8\xf6\xf8\x13\xd7B;+\xdft\xad\xdb\f%\xcc\fr\x98\v\x06g\xe7\xfd\x16\xfd\x02\x9b!pL\xd8,\xb7\x10\x9d\x82H@*\xb8\xb8\x1d\xd3w\x03\xa6\xc4X$\x02\xb9\xbf\xfb\xe0\xf3`<<\x19\xf4\xa1\xdb\xedB\x10\xe7\"\x80\xf7\xefw\"\"\xc5u4\x00\x00^\xb4\x91\xb4 \xa8\x9c\x05\xd0\xf4\x81BO\xd8\f\xf5\xe62\xde7\xca9̙n\x81Ґ\xa2u\x92\xe8\x14\x12\xad\n\xf7]\xe3\xd7\x19\x1a믵\xf1_w\xca\xf5-\xff︊\x1fPG\x13\xe8z\xcb\xc7\xfex\x91\x0eQ\xcfQC\x17\xeaa\xb8A\xb5\v\xe4\x82-ăh\x17\v\xce۹\x8aY\x9e)c\x8f~\xefԏk\x8f\x80\xb9A\x914\x80i\xcdV\xe3\a\\\x8dq)\x8c5\r\b\x86\xbd\xeb?{\xd7\xe3˓\x8b^Ђwc\xff\xfb\xe5]>+c\aL[\x03]\xc0e\x99+N\xf9io\xe9\xdc\xedغ\xa7L=\x8f幭\xbb\xce\xfd\xcb\xc8nqr\xed\xf3ttđ\x9aҋ\x1a\xcdu(\xeb*\t\xba\u008dd\x93\x1c\xc1*\xe0\x18\v\x8e\xb0\xc8D\x9c\xc1\x05\xa5\xe4V<\b*\x80U\xaeP\rW\b*\x92\xd2\xebZ4\xdbO5\xbd\x92\xf9\xca\x01\xeb\x8cϙ\x8c\x91\xd7\xc1\xa0\x9e\x8b\x18\r\xf5\x98\xcdp\x051\x930A0\x88\xb2\x05Lr\x12,\x10\x98F\x90ʂ\x90`\xd1XS{G\xa5\x18\xae\xb5\xbbpW\xdb\v\xf8\xe4@c\x99\x8b\x98\x05\xd0\xfd'\xb5\bUi\xb2\x92\xac\xc0ƶ\xb4\t\xbfP\xafn\xe3_\xb4\xec\xe0\xf3\xe0\xe6\xb2?\x1a\x8fz\xc3Q\x00\xcdVm/\xd0ȅ\xf9\x9ei/\xa8\xacV\xa8\xb7\x18L5+3a\xf1\xc0Xf\r\xff\x9e\xe9\xe7\x90\xca\xc9\v\xcd\xd7\xdd\xdd\x1f\xbf\x9dA\"fل\x19|]\xe1\xdd\"\x8dN\xe9\xaaO\xfd\x17M\x8ek\xfeܸΪ\xcasW\xdb\xdb\v(4\x17f\x9dO\x0e\nf,\xeaz\x8b\x04|BF\x9chc\xc6\tf\x06\xb5;\x0e\xb4R6pg%3f\xa1\xb4OX`\x95\xd2\xfeܮJo\xa2^\xac\xcc\xd7\xdc[Nr\x96\xfa\xaaE\xa7W\xe3\xa8\xf7\xe9\xe4\xe6|\xe4$\xb9b\xde\xc4N3\xddm\xf7\xc5=|\x84\x0e\x1c\xc1\x87Vm\xef\xde\xe5P$\x8d\x1f\xc1\x1dco\a\u007fw\xef\xc3\u007f\x16{\xa5\xf1\x97\x05\xbf\x0f's%8\x98\x85\xb0q&dJS\xac\x91qESj\x95\x02d:_A#Q\x1apɊ2G\xe03M\xc8Yə\xc5v\x99\x95M\xf2Q\xb0%\xe4,u^~\xeb\xecf\xd6\xe5\xcd\xcd\u007f\x18\x82\xbb\x178\a\xc6Β\x04\x1a\v5\xcb9HDN\xee\xe3\x8c\xc9\x14\x81<\x9a\xaf\xb9\xb0\xf8\xb1\xe9\xfb\xa8Ԙ\x88%\xd1r\xfd؟\x8c\x88\x95\xaeJ+\x94\xa4\xf6\xaa\xf7.\xff\xe8_\xf6\xba})Utڂ*P8\xfb|r=썺\x13!\x99^\xd5\u007fb\x02\xaei\x8a_G\x8b\xa4\xf1:\x17\xac{\xe2j2\xc5؞\xb18\xa3\xb6\xf1\xb2Mw\xc493\xbeR\x81\xf3}\xcaҫ!e\xc9W\xb7\xea&\a\xb8\xab\f\x1f\xfd\xf6\xeb\xdf\x0f\x03\xb8\xdf\xe4\xf8\xcd\xd1\xfdQ\xd1\a\xbc\x87\xa1#\x90\xff'\xd27R\xd3:z\xef\xe8\xe9\x85}\x86\xab\xff܊C;\b\xcba\xc0\xd27\x90\x14=FB\xe3X\xc9\x18a<\x8e\xfa\xd7\xe31\xb4!\b/h\x90\xbd)\xb2D-\x1d\xb8\x0e\xdb:4w\x01\xc1\\\xa5\xb6\xf0GG\xaeb\x0e\xdd[Z\x94F(y\x81Ɛ\xca'\x91\xef\xe8\xfd\xc0\xe7\x01\xcb\x053h*\xdfo\xa7h\x9cӢg\x85L\xdfЧ\xfb\xfb\x1b00Α\x83\xa22XQ\xa0\v7S\vZ\x1b\xa8\xb04OV\xcf\xd0Evb̬\xc0\x81VK\x81\xe6\xc6`\xe4\x175ZgU\xac\xf2\x81\xf2+K\xc2r\x83\xee\xf20\xa4\xb62닑\xedX\x8b\xd2\x0e\x98͞\x06\x98@գQ'2=\xd8О\x97_\xa0e\x97t\x97\x92\xc5\x0e4Њf\xa7~\faH\x91\x10]l\x0e[@;52\x0e*\xd9\xd9]\xb7<\x11\xff\x8c\xae\xa2+\xd0H\t\x80\x84\xcdE\xac$44\x16j\x8e\x1cJ\xa5mE\x83\xa9j\xd6\xc2\xf0\xdd\"\xfdT\x81\xbaP\xff\xb6\x13\xc9c\xd8\xf6\x17\x0e+;m\x11+O/\xf0'\xd3B\xcd\fp\xa11\xb6J\vti\xb8)\x89\x14\xa3\xeap\xe5m\xf6\a\x8f\xa1(\xa8c\xc2\xca\u07b7M&\x1e}*FE\xf9L酭\xc7\xd0\x16\xa5G;ry\x1d\x1f\x13\xac\xee\x9fi/\\\xd7\xe7y\x9c?\xba\x1dEKm\xf8\x1e\"\x9c\xcc\xd2Z%<W\xe9\xce\r\xc29\xd3a\xae\xd2p\xb3G\xd7\x1d\x89\x93ιJiV\b\xf6\x1d\xed\x90\x13\xa6\x9d\xab\x94\xbc\t)\xc6\x06m\x03\x82\xa5?\x9f3=\xe6\u00949[\x8d\v\xb6\x1cs,m\x16\xb4\xe0\xe0\x03휯\xc2\xe3L\xe4\\\xa3|\xbb\x06g\x96m\xd05\xd4Z\xe9\xb1ƪs\x1a/\xac\xacU\x1d\xd0\x04-p\xf2j\xdcz\xcb\x18\xdd\xfb\x15\xa1e\"ߙ:\x92\x0f\xffs\xdesz\xdb\x02\x97\xb3hV\x94ï9<\u05c8N\x9d\xc2)\x8b\x1f\xac\xf6s\xb3\xa58\xc7\\\x95\x05J{˴t<\xb0%\xee\xb9-\xff\xdfl\xce|\xedGh\xecFNe\xee\x15L\xe4\xb4Ļ\x19\xa2%ި\x02\x81N!f6\xceP\u007f|2\xe4\xd0;\xf6\v\xd4)\xcaxu\xa6\xa4e\xb1us\xbf\xee\x86\u007f=\x1b\xfeA\xb5\xe0\fQr\xffT\xfcO\xa8wwcP\u007f\xc7%\x13\xf9\xc9\xccf(\xad\x88\x19ez'\xa2KeER\tL\v\xecLKb\xc5$\x01f\xe8?\x0e\xaed`\xc1\xa0\xe4.J\xef\x8dt\xc8ۈ\xe5\x0fO\xb4\xb7\x11\xddR&r\xe1r\xb7E\x89\xee5غ\xae\x1b\xb8]\x9a\xcds\xb58S\xe5\xea;\xb2\x1b\x83}i,\x93\xf6L\x15\x85\xdfw\xb6\x8c\u007f\xc1U\xc5/\xa9f\x1c\xbf\xa0\x9b\xb9\xce\xf4\xb0cX'\x99\x1avؙ&\x1d6M\x18\xab\xb8\x17c\x8d\xb6\xc2M\u007f\xef\x1c\xfe:=\xec\xd0\xe7o\xd3D\x1fN?t\xfc'\xf9@\x92\x84~\x91\x85i\xe70\x99~\x98&\x0fRƱ,\x96\xb2XV\f0\xf8<\x80s\xe5\x13\xb9a\xd98\xc3\xf8\x01\xfc\xe32\xc8ʳ\\@\x17\x82pf\x88\tb\x96\x87\x13!ß|\xf1z\x97ѫ\xd8\xff\x06\x00\x00\xff\xff\x9a[Vx\xc1\x11\x00\x00",
		hash:  "a9d129c02b539e3f317947cfa034bc59a49554032e26329ec9df40517f9ecd79",
		mime:  "",
		mtime: time.Unix(1616091482, 0),
		size:  4545,
	},
	"mediawiki/MwddSpecialPage-aliases.php": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xb2\xb1/\xc8(\xe0\xe2R).HM\xceL\xcc\tHLOu\xcc\xc9L,N-V\xb0U\x88\x8e\xb5\xc6&\x13\xad\x9e\x9a\xa7\x1e\v\x92\xe7\xe2T\xf7-OIQW\xb0\xb5S\x88V\x80\xb2cu\xb8b\xad\xb9\x00\x01\x00\x00\xff\xff\x97Dl2Y\x00\x00\x00",
		hash:  "4ab3af2702babee43b165550753b199d3d25397464bd083293ed0e259fd9cf9a",
		mime:  "",
		mtime: time.Unix(1616091482, 0),
		size:  89,
	},
	"mediawiki/MwddSpecialPage.php": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xb4R]\xab\x1a1\x10}N~\xc5 \x82\xbb\x82H_W\xd1\">Xpi\xa1\x05\x1f\xda\"\xb3\xc94\x9b\x1a\x93%\x1fU\xb0\xfe\xf7\x92\xbam-\xf7r\xb9\x0f\xf7\xbe%sΜ9s\x98\xf9\xb2k;΅\xc1\x10\xa0>I\xf9\xb1#\xa1\xd1\x00\x9d#Y\x19\xa0\xff\u007f@Ep\xe1\x9cu\xa91Z\xc0\xb7dE\xd4\xce\xc2~/\x9c\r\xd1'\x11\x8b\x12.\x9c\xb1\x0e=\xd9XU\xf7\b\x8c\xb2\xf6\b\xca\x19gW\xce\xd9t<\xe6\f\xc6\xf06\x10ݏ\xa8*:\x93H\x912zct\xe8\xf1\b!zm\xd5O\x9b\x8c\x81aHM&g|\xfa\xd0Q\xafP\xfc\xe5\xc1\xff\xbe\x1e\xc1g\x9c1e\\\x83\x06\x86\xc7\x1c\x02\xf9\x1fZP\x98q\xce\xd80\xb6:L\x16\x8a\xe2\xfb\x14\xbb\x14\x8br\xb2@)7\x9f\xeam\x01\x83]\xabE\v\xa1o\x00\xf4\x04>Y\xab\xadZ\x0en\xc2O\xf7ϧ\x8d_<\x87\xf9=8\xbb'+\x9c\xcc\xd6\xef]B\xf9\xa2\x83\x06\x1bw\x02\xe9(\xc0z\x05\x06\x15\x18\xe7\x0e\xaf\xb9̗\x9a\xa4Ɲ>\xe8\u007f\xaf?\xbbU\x95\xa2\xf8Ά\x88VP\x96P\x14\u05eb\xadC\xb9B\x93k\xbe/\xd6xޢ*\xca\xcfo\xbe\xde\xf2\xf8}h\xd7_\x01\x00\x00\xff\xffǍ\xda4\xe0\x02\x00\x00",
		hash:  "630fed6deedeabaa0cdf43c2d14451ed66159473c73fc23cf3f4cb161a45d44c",
		mime:  "",
		mtime: time.Unix(1616091482, 0),
		size:  736,
	},
	"nginx/client_max_body_size.conf": {
		data:  "client_max_body_size 1024m;",
		hash:  "87e1334f7ba74c1f53d7fb2b1b035b14cb6105e94dce8c5d75303595cdc91831",
		mime:  "",
		mtime: time.Unix(1616091482, 0),
		size:  0,
	},
	"nginx/timeouts.conf": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff*(ʯ\xa8\x8cO\xce\xcf\xcbKM.\x89/\xc9\xccM\xcd/-Q042(\xb6\xe6\x82\xc8\x15\xa7\xe6\xa5\xc0%\x14\x14P\xe4\x8aR\x131\xe5\x00\x01\x00\x00\xff\xffp\x17(\x9eS\x00\x00\x00",
		hash:  "431b33796adb902aa6dfaa1d01bc7cf0365511d7c89b923a5066f529c5bc2958",
		mime:  "",
		mtime: time.Unix(1616091482, 0),
		size:  83,
	},
	"wait-for-it.sh": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xa4W[S\xdb\xc8\x12~ׯ\xe8\b\x1d\x8e͉\"\x9b:\xe7l\xad\x89Ȳ\xc0\x82\xabH`\xb1\xfd\xb0EQ\x94,\xb5\xd0\x14\xf2\x8c23\xe2\x12\x87\xfd\xed[s\xd1\xcd8\x89\x93\xe8\xc5\xd2\\z\xbe\xee\xfe\xe6\xeb\xf6֫\xa0\x14<\x98\x13\x1a \xbd\x87y$2g\v\x00f\x02AfD\x80\x889)$H\x06\x12\x85\x04\x92B\x04\xb7\xe4\x1e)L\x0f/ cB\x06\x05\xe3\x12\"\x8e\x10\xddG$\x8f\xe69:N\xbcHh\xb4\xc0\xd0\xeb\xcd#\x81\xea\x15\xbcA\xdfq0\xce\x18r\xde\xeb\xc3Rٺ\xba\x02\xef\xcf\xd9\xf8x\n>E\x18\xc2\xf5\xf5\x1e\xc8\f)\xa8u\xe0z\xbf\xb90\xdc\xdf\xde݃\x94\xc0\xb3\xe3\x94\"\xba\xc5^\xdfY:\x00\x00q$\xe1\xed[\x98M\x0eN\x8ea\u007f{י\xa9鑞\xf3,\x00\x8dp\xa4\x11^\xf9\xe2\x1a\xae|\t\x92,\x90\x95R}\xf8\x10\xb3\xc5\"\xa2\tD\xfcV\\\xeb\xad~\x06\xa7\xe7\x93)|\x06\xdfW\xbbC\xfde\x9eS&$0\x0e\xe3\v(i\x82\\\a\xc5\xec*\xe0\xe2\xfc\xd2\xecRǅ\xfa\xcb<*R\x1a\xc2ʞ\xaf=\a\xb9DN#I\xee1\u007fz\rO\xac\x04Q`L\xd2'\x15 \xed\x16(\xdc&\xf8\xa2\xf1Ӡ\x11\x1a\x88\x90\x9c\xc4r\xc5\xf29͟\x00\x1f1.%\x82(\xe7U\x04H\xaa-\xeb4\x8b2\x8e\x11\x13a\x8c}\xd4\xc6>\x96\x04Wm\xc1\x11\xa3\xff\x96\xc0JY\x94\n\xcf\x13\b\x19\xc9R\xc0\x02\x85J\x865 a:~\u007f|>3\xe1\xb1\xf1\x0f\xed\xd87C15\xeb\x81P\x10\x183\x9a\x88\xd7\xf0\t9\x83\x94q\xa0\xacʧ9ʇ\xc3\xf3\xf7\xef\x0f>\x1c\xc1\xc1\xe5ɤc\xe7\xd8\xfa\\9\xfc@d\xa6\xf3\x0eQ*U^*\xefSB\x89\xc8P8\x9aZ\xda.>\x12\tC\xe7\xd9q\x1e\"\"oR\xc6k\x1aZ\x1aW\x1e\xfa\xb7\x12\x065\x91k\xe7,\xf1\xc1\xad\x989\x02e\x89\xd0[톧H6\xf24i\f\xae\xca-`ic\xdc\xfa\xef\x1aL\xb9\xc0\x1f\xb5\xaf\f\xd7G\x18s)\xd1?BF\\\xdeH\x11z\xbd$\x92\b\xff\xf9\x97\xe8뉇\x8c\xe4\b\xe6z%\xac>غ?\x9e\xfc>\x9b\xfc\x05>~l]\xe3Nji\f\xfe'\x03\x044\x90\xce,GQ\xe62\xf4\xde5\x1e\xb5\xddSOO\x8b\xc2>\x04\t\xde\a2.\x02m+ж\xfa\xb0\xaf\x87i\x99簻\xbf=\xfc\x86q\xebl\xcb\x01\xb3F;0X\xef\x00\xd2d]\\\xbe\x92\x81vԉh\xf4\xd1\x12\xce\xeb\xf5\x8cM\xf0\xeb\xb0\xf7\xfb\xdd$WϜct\xb7\x0e\xbe\xc8\x11\v\x18ڴP\x132\x8e\xb2\xe4\xb4r\xaaMۛ\a\x1e\x15\x056\xf4݂1\x05Ƶ01\x10e\xa1\x15e2>\x19\u007f\x98BRr\xc5 ˓\x11dR\x16\xa3 ()y|#d\x14\xdf\xe1c\x9cE\xf4\x16\xdf\xc4l\x11D\xc1\xff~\xf9\xff\xaf\xbb\xed[a\xc5}=+*\x86{\x8a:\x8a\xe2\u007f\x9c\x1d\x9c4d\xf7\x06\xb5\xea\xf8~\x9c\x91<\xa9\x14ِ\xc8\n\xad\to#+\xf5\xfe헷d\x93\x13\u007f\xf8$\x9b\x94\x8b\xf1Q\xe8\xbdү\x92G\x05\xb8w$\xcf\xc1W\xe1\xf4\xbd\x8b\xf1\x91\v\xe3\x0f\x86\xfb*'\xa0\x86\xf4\xd7\xe5\xf1dv6\xadHj\xe3g\x06uu\xdcLTjՈ\xe3\x92sL,\xd5*-XU\x92Uq\xe8HA\xc5\"\x03B\xb1h\v\n\xceb\x14B\xe9f\xb9@*\x85ctA\x81ݪ\xc5ϱ\x02\x11G\x02\xc1\xf5\x86.\x90\x06\xf5\xceh\a\x9a\xab\xa3¬\xa3\xdb\xf3\x96\xc3 \x18\x05\xf0\xdcL*`\xa1\xb7\xac\xd6\\\r\xae\x9f\xeb9\x05\xb7=7l͉\x8c\xa4\x12\x1a\x11\xd8۫_m\x82\x9b3\x0eO\xc7gG\xe1p\xa3\xbd\xadJ\xd8\x18\xd0\x1c\xdf\xd0@\xbb.7\x16&\xd3\xcb\xf1\xe1\xa6&\xb2\x95\xf0\xb8ޮ\xbb*f\x9a\xb5a\b\xae۴TZA\xf6:ҡO\xd9]\x1f%M\xff\x9d\x17g-\x87[;\u1cfb\x11Ң\xdfM\xd6:\xa4\xfaJ\xfd\x1cR͞\x9d\x17g}\x0f\xd2V.\xec\x05Y\v\xb6\xba<?\x87\xb7\x92\x8f\x9du\x87~\x0fj\xbf\xdf]\xd6P\xfal\x1c\xaa\xe6\xd9Y_=\xbayƼ\x95'\xdda\xaf[\xd8\x02[\xabΌ\xdeQ\xf6@k-\x18\x817t\xbfn\nE\x14;\xbaL9&\xa6\xae\xe6\xaakC\xfa\xf93\xb8F\x87Vb\xect\x0e>\xe6\x9c\xf1\x91n\x8a)b\xa2\nW\xc1\xd9=I\x10\xa2\x95\xd6\xd8\xfeqycp\x19L)q\x9c*\xe2\xdeҾ\x8d\xfc\xdd\xff\x0e\x9e\x1d{\x15\xbd\xa5y\x19\xf9\x83g\xc7(\x84\xb7Կz\xc4\\yo\xa9\u007f\xf5\x88\xb3\x05q\x86\xf1\x9d.\xa2\x88\xba\x9b\xae\xfaV\x01)g\v\x98\x97\xe2i\xce\x1e\xdf}\xcfR\v\xee\xe6\xe2`z\x1az=\x8eQ^D2\x03\xaf\xf7\x90\x918\xab6\xf6\xfbN\x97\xa2z\x03\x84\u007f\x83kM\xb9/\xab\x87i\xd9Z\xaa\xd3.\x8a\xa1\xebK\xd7\xe9\x14O\xbb~\xf0\x85\xf5\xae\x8e\xab\x85\xa1C\xb5\xae\x17\xae\xfa\x905%O\xf7\xd8U\xb1\xa9Oެ\xbd^moꉕ#\xda\xfet\xa0\xbc\\\x9b\x92\x8eCgcx\xf5\x92\x92\xeb\xca\xf4\xf66x\x86=_j|\xd6\xd4m\xfbWm\xc1\x12|\r\x1c\xd3R\xe8Ƌ\xb5\xff\xa9\xd9\xda\xdb\\\xb1N\xc4Z\x85[\xedѐ\x9b0v\x96\xa6\xe4\x9f\x00\x00\x00\xff\xff\xaf\xabL\xf5\xf8\x0f\x00\x00",
		hash:  "6a2b7c49ead02dcaf820d25c8df99043aca0c0937a38c0c1a75725e0faa42326",
		mime:  "text/x-sh; charset=utf-8",
		mtime: time.Unix(1616091482, 0),
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
