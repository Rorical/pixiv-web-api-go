package apis

type PixivResponseUrls struct {
	Mini     string `json:"mini"`
	Thumb    string `json:"thumb"`
	Small    string `json:"small"`
	Regular  string `json:"regular"`
	Original string `json:"original"`
}
type PixivResponseTranslation struct {
	En string `json:"en"`
}
type PixivResponseTags struct {
	Tag         string                     `json:"tag"`
	Locked      bool                       `json:"locked"`
	Deletable   bool                       `json:"deletable"`
	Userid      string                     `json:"userId"`
	Translation []PixivResponseTranslation `json:"translation"`
	Username    string                     `json:"userName"`
}
type PixivResponseTags struct {
	Authorid string              `json:"authorId"`
	Islocked bool                `json:"isLocked"`
	Tags     []PixivResponseTags `json:"tags"`
	Writable bool                `json:"writable"`
}
type PixivResponseTitlecaptiontranslation struct {
	Worktitle   string `json:"workTitle"`
	Workcaption string `json:"workCaption"`
}
type PixivResponse89125708 struct {
	Id                      string                                 `json:"id"`
	Title                   string                                 `json:"title"`
	Illusttype              int                                    `json:"illustType"`
	Xrestrict               int                                    `json:"xRestrict"`
	Restrict                int                                    `json:"restrict"`
	Sl                      int                                    `json:"sl"`
	Url                     string                                 `json:"url"`
	Description             string                                 `json:"description"`
	Tags                    []string                               `json:"tags"`
	Userid                  string                                 `json:"userId"`
	Username                string                                 `json:"userName"`
	Width                   int                                    `json:"width"`
	Height                  int                                    `json:"height"`
	Pagecount               int                                    `json:"pageCount"`
	Isbookmarkable          bool                                   `json:"isBookmarkable"`
	Bookmarkdata            interface{}                            `json:"bookmarkData"`
	Alt                     string                                 `json:"alt"`
	Titlecaptiontranslation []PixivResponseTitlecaptiontranslation `json:"titleCaptionTranslation"`
	Createdate              string                                 `json:"createDate"`
	Updatedate              string                                 `json:"updateDate"`
	Isunlisted              bool                                   `json:"isUnlisted"`
	Ismasked                bool                                   `json:"isMasked"`
}
type PixivResponseTitlecaptiontranslation struct {
	Worktitle   string `json:"workTitle"`
	Workcaption string `json:"workCaption"`
}
type PixivResponse88372007 struct {
	Id                      string                                 `json:"id"`
	Title                   string                                 `json:"title"`
	Illusttype              int                                    `json:"illustType"`
	Xrestrict               int                                    `json:"xRestrict"`
	Restrict                int                                    `json:"restrict"`
	Sl                      int                                    `json:"sl"`
	Url                     string                                 `json:"url"`
	Description             string                                 `json:"description"`
	Tags                    []string                               `json:"tags"`
	Userid                  string                                 `json:"userId"`
	Username                string                                 `json:"userName"`
	Width                   int                                    `json:"width"`
	Height                  int                                    `json:"height"`
	Pagecount               int                                    `json:"pageCount"`
	Isbookmarkable          bool                                   `json:"isBookmarkable"`
	Bookmarkdata            interface{}                            `json:"bookmarkData"`
	Alt                     string                                 `json:"alt"`
	Titlecaptiontranslation []PixivResponseTitlecaptiontranslation `json:"titleCaptionTranslation"`
	Createdate              string                                 `json:"createDate"`
	Updatedate              string                                 `json:"updateDate"`
	Isunlisted              bool                                   `json:"isUnlisted"`
	Ismasked                bool                                   `json:"isMasked"`
	Profileimageurl         string                                 `json:"profileImageUrl"`
}
type PixivResponseUserillusts struct {
	_89125708 []PixivResponse89125708 `json:"89125708"`
	_88372007 []PixivResponse88372007 `json:"88372007"`
	_87344915 interface{}             `json:"87344915"`
	_86552103 interface{}             `json:"86552103"`
	_85578826 interface{}             `json:"85578826"`
	_85492210 interface{}             `json:"85492210"`
	_84665888 interface{}             `json:"84665888"`
	_84665842 interface{}             `json:"84665842"`
	_81490736 interface{}             `json:"81490736"`
	_80502339 interface{}             `json:"80502339"`
	_79909255 interface{}             `json:"79909255"`
	_79793887 interface{}             `json:"79793887"`
	_79749855 interface{}             `json:"79749855"`
	_79180683 interface{}             `json:"79180683"`
	_78622141 interface{}             `json:"78622141"`
	_78009456 interface{}             `json:"78009456"`
	_77970951 interface{}             `json:"77970951"`
	_77921513 interface{}             `json:"77921513"`
	_77879039 interface{}             `json:"77879039"`
	_77727497 interface{}             `json:"77727497"`
	_77696800 interface{}             `json:"77696800"`
	_77569511 interface{}             `json:"77569511"`
	_77426268 interface{}             `json:"77426268"`
	_77231798 interface{}             `json:"77231798"`
	_77216942 interface{}             `json:"77216942"`
	_76651620 interface{}             `json:"76651620"`
	_76645181 interface{}             `json:"76645181"`
	_76532272 interface{}             `json:"76532272"`
	_76348201 interface{}             `json:"76348201"`
	_76087238 interface{}             `json:"76087238"`
	_76032619 interface{}             `json:"76032619"`
	_75733719 interface{}             `json:"75733719"`
	_74921460 interface{}             `json:"74921460"`
	_74553719 interface{}             `json:"74553719"`
	_74387013 interface{}             `json:"74387013"`
	_73726667 interface{}             `json:"73726667"`
	_73633485 interface{}             `json:"73633485"`
	_73487503 interface{}             `json:"73487503"`
	_73187064 interface{}             `json:"73187064"`
	_73122457 interface{}             `json:"73122457"`
	_72645957 interface{}             `json:"72645957"`
	_72601604 interface{}             `json:"72601604"`
	_72522686 interface{}             `json:"72522686"`
	_72481387 interface{}             `json:"72481387"`
	_72455259 interface{}             `json:"72455259"`
	_72269579 interface{}             `json:"72269579"`
	_72210581 interface{}             `json:"72210581"`
	_72194193 interface{}             `json:"72194193"`
	_71989455 interface{}             `json:"71989455"`
	_71833055 interface{}             `json:"71833055"`
	_71782684 interface{}             `json:"71782684"`
	_71622613 interface{}             `json:"71622613"`
	_71531631 interface{}             `json:"71531631"`
	_71476107 interface{}             `json:"71476107"`
	_71434083 interface{}             `json:"71434083"`
	_71306324 interface{}             `json:"71306324"`
	_71209219 interface{}             `json:"71209219"`
	_71157453 interface{}             `json:"71157453"`
	_70847677 interface{}             `json:"70847677"`
	_70826824 interface{}             `json:"70826824"`
	_70644304 interface{}             `json:"70644304"`
	_70469561 interface{}             `json:"70469561"`
	_70440995 interface{}             `json:"70440995"`
	_70426699 interface{}             `json:"70426699"`
	_70315635 interface{}             `json:"70315635"`
	_70208481 interface{}             `json:"70208481"`
	_70171193 interface{}             `json:"70171193"`
	_70088526 interface{}             `json:"70088526"`
	_69666468 interface{}             `json:"69666468"`
	_69224856 interface{}             `json:"69224856"`
	_69088998 interface{}             `json:"69088998"`
	_68745883 interface{}             `json:"68745883"`
	_67857235 interface{}             `json:"67857235"`
	_66956431 interface{}             `json:"66956431"`
	_66912942 interface{}             `json:"66912942"`
	_66450775 interface{}             `json:"66450775"`
	_66352334 interface{}             `json:"66352334"`
	_66352295 interface{}             `json:"66352295"`
	_66207209 interface{}             `json:"66207209"`
	_65892644 interface{}             `json:"65892644"`
	_65859375 interface{}             `json:"65859375"`
	_65778669 interface{}             `json:"65778669"`
	_65701396 interface{}             `json:"65701396"`
	_65627395 interface{}             `json:"65627395"`
	_65413746 interface{}             `json:"65413746"`
	_65142202 interface{}             `json:"65142202"`
	_64938903 interface{}             `json:"64938903"`
	_64936452 interface{}             `json:"64936452"`
	_64152613 interface{}             `json:"64152613"`
	_63865077 interface{}             `json:"63865077"`
	_63845095 interface{}             `json:"63845095"`
	_63842713 interface{}             `json:"63842713"`
	_63842669 interface{}             `json:"63842669"`
	_62898137 interface{}             `json:"62898137"`
	_62558465 interface{}             `json:"62558465"`
	_62269065 interface{}             `json:"62269065"`
	_62121527 interface{}             `json:"62121527"`
	_61891438 interface{}             `json:"61891438"`
	_61779068 interface{}             `json:"61779068"`
	_61722233 interface{}             `json:"61722233"`
	_59959150 interface{}             `json:"59959150"`
	_54311056 interface{}             `json:"54311056"`
	_54146699 interface{}             `json:"54146699"`
	_53376065 interface{}             `json:"53376065"`
	_53271712 interface{}             `json:"53271712"`
	_53142453 interface{}             `json:"53142453"`
	_46572596 interface{}             `json:"46572596"`
	_39750054 interface{}             `json:"39750054"`
	_35671572 interface{}             `json:"35671572"`
	_35081087 interface{}             `json:"35081087"`
	_31347435 interface{}             `json:"31347435"`
	_23322524 interface{}             `json:"23322524"`
	_18882071 interface{}             `json:"18882071"`
	_17646967 interface{}             `json:"17646967"`
	_17513394 interface{}             `json:"17513394"`
	_6430115  interface{}             `json:"6430115"`
	_4948507  interface{}             `json:"4948507"`
	_3672447  interface{}             `json:"3672447"`
	_3671967  interface{}             `json:"3671967"`
	_3291654  interface{}             `json:"3291654"`
	_3204992  interface{}             `json:"3204992"`
	_2393248  interface{}             `json:"2393248"`
	_2321247  interface{}             `json:"2321247"`
}
type PixivResponseResponsive struct {
	Url string `json:"url"`
}
type PixivResponseRectangle struct {
	Url string `json:"url"`
}
type PixivResponse500x500 struct {
	Url string `json:"url"`
}
type PixivResponseHeader struct {
	Url string `json:"url"`
}
type PixivResponseFooter struct {
	Url string `json:"url"`
}
type PixivResponseExpandedfooter struct {
	Url string `json:"url"`
}
type PixivResponseLogo struct {
	Url string `json:"url"`
}
type PixivResponseRelatedworks struct {
	Url string `json:"url"`
}
type PixivResponseZoneconfig struct {
	Responsive     []PixivResponseResponsive     `json:"responsive"`
	Rectangle      []PixivResponseRectangle      `json:"rectangle"`
	_500x500       []PixivResponse500x500        `json:"500x500"`
	Header         []PixivResponseHeader         `json:"header"`
	Footer         []PixivResponseFooter         `json:"footer"`
	Expandedfooter []PixivResponseExpandedfooter `json:"expandedFooter"`
	Logo           []PixivResponseLogo           `json:"logo"`
	Relatedworks   []PixivResponseRelatedworks   `json:"relatedworks"`
}
type PixivResponseAlternatelanguages struct {
	Ja string `json:"ja"`
	En string `json:"en"`
}
type PixivResponseOgp struct {
	Description string `json:"description"`
	Image       string `json:"image"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}
type PixivResponseTwitter struct {
	Description string `json:"description"`
	Image       string `json:"image"`
	Title       string `json:"title"`
	Card        string `json:"card"`
}
type PixivResponseMeta struct {
	Title              string                            `json:"title"`
	Description        string                            `json:"description"`
	Canonical          string                            `json:"canonical"`
	Alternatelanguages []PixivResponseAlternatelanguages `json:"alternateLanguages"`
	Descriptionheader  string                            `json:"descriptionHeader"`
	Ogp                []PixivResponseOgp                `json:"ogp"`
	Twitter            []PixivResponseTwitter            `json:"twitter"`
}
type PixivResponseExtradata struct {
	Meta []PixivResponseMeta `json:"meta"`
}
type PixivResponseTitlecaptiontranslation struct {
	Worktitle   string `json:"workTitle"`
	Workcaption string `json:"workCaption"`
}
type PixivResponseBody struct {
	Illustid                string                                 `json:"illustId"`
	Illusttitle             string                                 `json:"illustTitle"`
	Illustcomment           string                                 `json:"illustComment"`
	Id                      string                                 `json:"id"`
	Title                   string                                 `json:"title"`
	Description             string                                 `json:"description"`
	Illusttype              int                                    `json:"illustType"`
	Createdate              string                                 `json:"createDate"`
	Uploaddate              string                                 `json:"uploadDate"`
	Restrict                int                                    `json:"restrict"`
	Xrestrict               int                                    `json:"xRestrict"`
	Sl                      int                                    `json:"sl"`
	Urls                    []PixivResponseUrls                    `json:"urls"`
	Tags                    []PixivResponseTags                    `json:"tags"`
	Alt                     string                                 `json:"alt"`
	Storabletags            []string                               `json:"storableTags"`
	Userid                  string                                 `json:"userId"`
	Username                string                                 `json:"userName"`
	Useraccount             string                                 `json:"userAccount"`
	Userillusts             []PixivResponseUserillusts             `json:"userIllusts"`
	Likedata                bool                                   `json:"likeData"`
	Width                   int                                    `json:"width"`
	Height                  int                                    `json:"height"`
	Pagecount               int                                    `json:"pageCount"`
	Bookmarkcount           int                                    `json:"bookmarkCount"`
	Likecount               int                                    `json:"likeCount"`
	Commentcount            int                                    `json:"commentCount"`
	Responsecount           int                                    `json:"responseCount"`
	Viewcount               int                                    `json:"viewCount"`
	Ishowto                 bool                                   `json:"isHowto"`
	Isoriginal              bool                                   `json:"isOriginal"`
	Imageresponseoutdata    []interface{}                          `json:"imageResponseOutData"`
	Imageresponsedata       []interface{}                          `json:"imageResponseData"`
	Imageresponsecount      int                                    `json:"imageResponseCount"`
	Polldata                interface{}                            `json:"pollData"`
	Seriesnavdata           interface{}                            `json:"seriesNavData"`
	Descriptionboothid      interface{}                            `json:"descriptionBoothId"`
	Descriptionyoutubeid    interface{}                            `json:"descriptionYoutubeId"`
	Comicpromotion          interface{}                            `json:"comicPromotion"`
	Fanboxpromotion         interface{}                            `json:"fanboxPromotion"`
	Contestbanners          []interface{}                          `json:"contestBanners"`
	Isbookmarkable          bool                                   `json:"isBookmarkable"`
	Bookmarkdata            interface{}                            `json:"bookmarkData"`
	Contestdata             interface{}                            `json:"contestData"`
	Zoneconfig              []PixivResponseZoneconfig              `json:"zoneConfig"`
	Extradata               []PixivResponseExtradata               `json:"extraData"`
	Titlecaptiontranslation []PixivResponseTitlecaptiontranslation `json:"titleCaptionTranslation"`
	Isunlisted              bool                                   `json:"isUnlisted"`
	Request                 interface{}                            `json:"request"`
}
type PixivResponse struct {
	Error   bool                `json:"error"`
	Message string              `json:"message"`
	Body    []PixivResponseBody `json:"body"`
}
