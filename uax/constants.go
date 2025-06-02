package uax

type PlatformType int

const (
	PlatformType_Unknown = iota
	PlatformType_Desktop
	PlatformType_SmartMobile // Smartphone/Tablet
	PlatformType_DumbMobile  // Feature Phone
	PlatformType_Console
	PlatformType_Bot
	PlatformType_Headless
	PlatformType_Application
	PlatformType_Tools
	PlatformType_Appliance
	PlatformType_Other
	PlatformType_Schizophrenic
)

type Quriks int

const (
	Quriks_DegradedWebKit = (1 << iota) // see README section "Quriks"
)

type BrowserEngine string

const (
	BrowserEngine_None     = "None"
	BrowserEngine_Other    = "Other"
	BrowserEngine_KHTML    = "KHTML"
	BrowserEngine_WebKit   = "WebKit"
	BrowserEngine_Blink    = "Blink"
	BrowserEngine_Presto   = "Presto"  // Old Opera Engine
	BrowserEngine_Ladybird = "LibWeb"  // Ladybird
	BrowserEngine_Gecko    = "Gecko"   // Firefox
	BrowserEngine_Goanna   = "Goanna"  // New Moon
	BrowserEngine_Tasman   = "Tasman"  // IE5 for Mac
	BrowserEngine_Trident  = "Trident" // IE4+
	BrowserEngine_Spartan  = "Spartan" // EdgeHTML
	BrowserEngine_Netscape = "Mozilla" // NetScape 4 and under
	BrowserEngine_Mosaic   = "Mosaic"  // NCSA Mosaic
	BrowserEngine_NetSurf  = "NetSurf"
	BrowserEngine_Servo    = "Servo"
)

type PlatformOS string

const (
	PlatformOS_None          = "None"
	PlatformOS_Other         = "Other"
	PlatformOS_Windows       = "Windows"
	PlatformOS_WindowsPhone  = "Windows Phone"
	PlatformOS_WindowsMobile = "Windows Mobile"
	PlatformOS_WindowsCE     = "Windows CE"
	PlatformOS_Android       = "Android"
	PlatformOS_macOS         = "macOS"
	PlatformOS_iOS           = "iOS"
	PlatformOS_Linux         = "Linux"
	PlatformOS_FreeBSD       = "FreeBSD"
	PlatformOS_ChromeOS      = "ChromeOS"
	PlatformOS_CrOS          = "CrOS" // shortname for ChromeOS
	PlatformOS_BlackBerry    = "BlackBerry"
	PlatformOS_Symbian       = "SymbOS"
	PlatformOS_HotJava       = "HotJava"
	PlatformOS_RiscOS        = "RISC OS"
)

const (
	Browser_Netscape         = "Netscape"
	Browser_Firebird         = "Firebird"
	Browser_Firefox          = "Firefox"
	Browser_TenFourFox       = "TenFourFox"
	Browser_KMeleon          = "K-Meleon"
	Browser_PaleMoon         = "PaleMoon"
	Browser_Basilisk         = "Basilisk"
	Browser_Mypal            = "Mypal"
	Browser_SucklessSurf     = "Surf"
	Browser_QtWeb            = "QtWeb"
	Browser_iCab             = "iCab"
	Browser_InternetExplorer = "Internet Explorer"
	Browser_LegacyEdge       = "Legacy Edge"
	Browser_AOL              = "AOL"
	Browser_Chrome           = "Chrome"
	Browser_HeadlessChrome   = "HeadlessChrome"
	Browser_MSEdge           = "Microsoft Edge"
	Browser_Vivaldi          = "Vivaldi"
	Browser_Brave            = "Brave"
	Browser_Opera            = "Opera"
	Browser_OperaMini        = "Opera Mini"
	Browser_Safari           = "Safari"
	Browser_NintendoBrowser  = "NintendoBrowser"
	Browser_GenericWebKit    = "Generic WebKit"
	Browser_SamsungBrowser   = "Samsung Browser"
	Browser_HuaweiBrowser    = "Huawei Browser"
	Browser_MiUiBrowser      = "MiUi Browser"
	Browser_PlanetWeb        = "PlanetWeb"
	Browser_DreamKey         = "DreamKey"
	Browser_DreamPassport    = "DreamPassport"
	Browser_NCSAMosaic       = "NCSA Mosaic"
	Browser_NetSurf          = "NetSurf"
	Browser_Konqueror        = "Konqueror"
	Browser_HotJava          = "HotJava"
	Browser_NetFront         = "NetFront"
	Browser_Surveyon         = "Surveyon"

	CrawlerBot_Google    = "Googlebot"
	CrawlerBot_Twitter   = "Twitterbot"
	CrawlerBot_Facebook  = "FacebookBot"
	CrawlerBot_Apple     = "Applebot"
	CrawlerBot_Bing      = "Bingbot"
	CrawlerBot_Yandex    = "YandexBot"
	CrawlerBot_Yahoo     = "YahooBot"
	CrawlerBot_Semrush   = "SemrushBot"
	CralwerBot_Discord   = "Discordbot"
	CrawlerBot_Bytedance = "BytedanceBot"
	CrawlerBot_MSN       = "MSNbot"
	CrawlerBot_BUbiNGbot = "BUbiNG"

	Console_NintendoDS          = "Nintendo DS"
	Console_NintendoDSi         = "Nintendo DSi"
	Console_Nintendo3DS         = "Nintendo 3DS"
	Console_NintendoNew3DS      = "New Nintendo 3DS"
	Console_NintendoSwitch      = "Nintendo Switch"
	Console_NintendoWii         = "Nintendo Wii"
	Console_NintendoWiiU        = "Nintendo Wii U"
	Console_SegaDreamcast       = "Sega Dreamcast"
	Console_XboxOriginal        = "Original Xbox"
	Console_Xbox360             = "Xbox 360"
	Console_XboxOne             = "Xbox One"
	Console_XboxOne_S           = "Xbox One S"
	Console_XboxOne_X           = "Xbox One X"
	Console_XboxSeries_X        = "Xbox Series X" // idk if theres one for the S?
	Console_PlayStation2        = "PlayStation 2"
	Console_PlayStation3        = "PlayStation 3"
	Console_PlayStation4        = "PlayStation 4"
	Console_PlayStation5        = "PlayStation 5"
	Console_PlayStationVita     = "PlayStation Vita"
	Console_PlayStationPortable = "PlayStation Portable"

	Appliance_WebTV  = "WebTV"
	Appliance_MSNTV2 = "MSNTV2"

	Application_Facebook  = "Facebook App"
	Application_Instagram = "Instagram App"
	Application_TikTok    = "TikTok App"

	Tools_Curl    = "curl"
	Tools_Wget    = "Wget"
	Tools_Seafile = "Seafile"
)
