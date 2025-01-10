package model

type Root[T any] struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Ttl     int64  `json:"ttl"`
	Data    T      `json:"data"`
}
type LoginData struct {
	IsLogin            bool                `json:"isLogin"`
	EmailVerified      int64               `json:"email_verified"`
	Face               string              `json:"face"`
	FaceNft            int64               `json:"face_nft"`
	FaceNftType        int64               `json:"face_nft_type"`
	LevelInfo          LoginLevelInfo      `json:"level_info"`
	Mid                int64               `json:"mid"`
	MobileVerified     int64               `json:"mobile_verified"`
	Money              float64             `json:"money"`
	Moral              int64               `json:"moral"`
	Official           LoginOfficial       `json:"official"`
	OfficialVerify     LoginOfficialVerify `json:"officialVerify"`
	Pendant            LoginPendant        `json:"pendant"`
	Scores             int64               `json:"scores"`
	Uname              string              `json:"uname"`
	VipDueDate         int64               `json:"vipDueDate"`
	VipStatus          int64               `json:"vipStatus"`
	VipType            int64               `json:"vipType"`
	VipPayType         int64               `json:"vip_pay_type"`
	VipThemeType       int64               `json:"vip_theme_type"`
	VipLabel           LoginVipLabel       `json:"vip_label"`
	VipAvatarSubscript int64               `json:"vip_avatar_subscript"`
	VipNicknameColor   string              `json:"vip_nickname_color"`
	Vip                LoginVip            `json:"vip"`
	Wallet             LoginWallet         `json:"wallet"`
	HasShop            bool                `json:"has_shop"`
	ShopUrl            string              `json:"shop_url"`
	AnswerStatus       int64               `json:"answer_status"`
	IsSeniorMember     int64               `json:"is_senior_member"`
	WbiImg             LoginWbiImg         `json:"wbi_img"`
	IsJury             bool                `json:"is_jury"`
}
type LoginLevelInfo struct {
	CurrentLevel int64  `json:"current_level"`
	CurrentMin   int64  `json:"current_min"`
	CurrentExp   int64  `json:"current_exp"`
	NextExp      string `json:"next_exp"`
}
type LoginOfficial struct {
	Role  int64  `json:"role"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Type  int64  `json:"type"`
}
type LoginOfficialVerify struct {
	Type int64  `json:"type"`
	Desc string `json:"desc"`
}
type LoginPendant struct {
	Pid               int64  `json:"pid"`
	Name              string `json:"name"`
	Image             string `json:"image"`
	Expire            int64  `json:"expire"`
	ImageEnhance      string `json:"image_enhance"`
	ImageEnhanceFrame string `json:"image_enhance_frame"`
	NPid              int64  `json:"n_pid"`
}
type LoginVipLabel struct {
	Path                  string `json:"path"`
	Text                  string `json:"text"`
	LabelTheme            string `json:"label_theme"`
	TextColor             string `json:"text_color"`
	BgStyle               int64  `json:"bg_style"`
	BgColor               string `json:"bg_color"`
	BorderColor           string `json:"border_color"`
	UseImgLabel           bool   `json:"use_img_label"`
	ImgLabelUriHans       string `json:"img_label_uri_hans"`
	ImgLabelUriHant       string `json:"img_label_uri_hant"`
	ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"`
	ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"`
}
type LoginVip struct {
	Type               int64           `json:"type"`
	Status             int64           `json:"status"`
	DueDate            int64           `json:"due_date"`
	VipPayType         int64           `json:"vip_pay_type"`
	ThemeType          int64           `json:"theme_type"`
	Label              LoginLabel      `json:"label"`
	AvatarSubscript    int64           `json:"avatar_subscript"`
	NicknameColor      string          `json:"nickname_color"`
	Role               int64           `json:"role"`
	AvatarSubscriptUrl string          `json:"avatar_subscript_url"`
	TvVipStatus        int64           `json:"tv_vip_status"`
	TvVipPayType       int64           `json:"tv_vip_pay_type"`
	TvDueDate          int64           `json:"tv_due_date"`
	AvatarIcon         LoginAvatarIcon `json:"avatar_icon"`
}
type LoginLabel struct {
	Path                  string `json:"path"`
	Text                  string `json:"text"`
	LabelTheme            string `json:"label_theme"`
	TextColor             string `json:"text_color"`
	BgStyle               int64  `json:"bg_style"`
	BgColor               string `json:"bg_color"`
	BorderColor           string `json:"border_color"`
	UseImgLabel           bool   `json:"use_img_label"`
	ImgLabelUriHans       string `json:"img_label_uri_hans"`
	ImgLabelUriHant       string `json:"img_label_uri_hant"`
	ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"`
	ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"`
}
type LoginWallet struct {
	Mid           int64 `json:"mid"`
	BcoinBalance  int64 `json:"bcoin_balance"`
	CouponBalance int64 `json:"coupon_balance"`
	CouponDueTime int64 `json:"coupon_due_time"`
}
type LoginWbiImg struct {
	ImgUrl string `json:"img_url"`
	SubUrl string `json:"sub_url"`
}
type LoginAvatarIcon struct {
	IconResource map[string]interface{} `json:"icon_resource"`
}
type VideoStream struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		From              string   `json:"from"`
		Result            string   `json:"result"`
		Message           string   `json:"message"`
		Quality           int      `json:"quality"`
		Format            string   `json:"format"`
		Timelength        int      `json:"timelength"`
		AcceptFormat      string   `json:"accept_format"`
		AcceptDescription []string `json:"accept_description"`
		AcceptQuality     []int    `json:"accept_quality"`
		VideoCodecid      int      `json:"video_codecid"`
		SeekParam         string   `json:"seek_param"`
		SeekType          string   `json:"seek_type"`
		Dash              struct {
			Duration       int     `json:"duration"`
			MinBufferTime  float64 `json:"minBufferTime"`
			MinBufferTime0 float64 `json:"min_buffer_time"`
			Video          []struct {
				ID            int      `json:"id"`
				BaseURL       string   `json:"baseUrl"`
				BaseURL0      string   `json:"base_url"`
				BackupURL     []string `json:"backupUrl"`
				BackupURL0    []string `json:"backup_url"`
				Bandwidth     int      `json:"bandwidth"`
				MimeType      string   `json:"mimeType"`
				MimeType0     string   `json:"mime_type"`
				Codecs        string   `json:"codecs"`
				Width         int      `json:"width"`
				Height        int      `json:"height"`
				FrameRate     string   `json:"frameRate"`
				FrameRate0    string   `json:"frame_rate"`
				Sar           string   `json:"sar"`
				StartWithSap  int      `json:"startWithSap"`
				StartWithSap0 int      `json:"start_with_sap"`
				SegmentBase   struct {
					Initialization string `json:"Initialization"`
					IndexRange     string `json:"indexRange"`
				} `json:"SegmentBase"`
				SegmentBase0 struct {
					Initialization string `json:"initialization"`
					IndexRange     string `json:"index_range"`
				} `json:"segment_base"`
				Codecid int `json:"codecid"`
			} `json:"video"`
			Audio []struct {
				ID            int      `json:"id"`
				BaseURL       string   `json:"baseUrl"`
				BaseURL0      string   `json:"base_url"`
				BackupURL     []string `json:"backupUrl"`
				BackupURL0    []string `json:"backup_url"`
				Bandwidth     int      `json:"bandwidth"`
				MimeType      string   `json:"mimeType"`
				MimeType0     string   `json:"mime_type"`
				Codecs        string   `json:"codecs"`
				Width         int      `json:"width"`
				Height        int      `json:"height"`
				FrameRate     string   `json:"frameRate"`
				FrameRate0    string   `json:"frame_rate"`
				Sar           string   `json:"sar"`
				StartWithSap  int      `json:"startWithSap"`
				StartWithSap0 int      `json:"start_with_sap"`
				SegmentBase   struct {
					Initialization string `json:"Initialization"`
					IndexRange     string `json:"indexRange"`
				} `json:"SegmentBase"`
				SegmentBase0 struct {
					Initialization string `json:"initialization"`
					IndexRange     string `json:"index_range"`
				} `json:"segment_base"`
				Codecid int `json:"codecid"`
			} `json:"audio"`
			Dolby struct {
				Type  int         `json:"type"`
				Audio interface{} `json:"audio"`
			} `json:"dolby"`
			Flac interface{} `json:"flac"`
		} `json:"dash"`
		SupportFormats []struct {
			Quality        int      `json:"quality"`
			Format         string   `json:"format"`
			NewDescription string   `json:"new_description"`
			DisplayDesc    string   `json:"display_desc"`
			Superscript    string   `json:"superscript"`
			Codecs         []string `json:"codecs"`
		} `json:"support_formats"`
		HighFormat   interface{} `json:"high_format"`
		LastPlayTime int         `json:"last_play_time"`
		LastPlayCid  int64       `json:"last_play_cid"`
		ViewInfo     interface{} `json:"view_info"`
	} `json:"data"`
}
type VideoInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Bvid      string `json:"bvid"`
		Aid       int    `json:"aid"`
		Videos    int    `json:"videos"`
		Tid       int    `json:"tid"`
		Tname     string `json:"tname"`
		Copyright int    `json:"copyright"`
		Pic       string `json:"pic"`
		Title     string `json:"title"`
		Pubdate   int    `json:"pubdate"`
		Ctime     int    `json:"ctime"`
		Desc      string `json:"desc"`
		DescV2    []struct {
			RawText string `json:"raw_text"`
			Type    int    `json:"type"`
			BizID   int    `json:"biz_id"`
		} `json:"desc_v2"`
		State     int `json:"state"`
		Duration  int `json:"duration"`
		MissionID int `json:"mission_id"`
		Rights    struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"`
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UgcPay        int `json:"ugc_pay"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
			CleanMode     int `json:"clean_mode"`
			IsSteinGate   int `json:"is_stein_gate"`
			Is360         int `json:"is_360"`
			NoShare       int `json:"no_share"`
			ArcPay        int `json:"arc_pay"`
			FreeWatch     int `json:"free_watch"`
		} `json:"rights"`
		Owner struct {
			Mid  int    `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct {
			Aid        int    `json:"aid"`
			View       int    `json:"view"`
			Danmaku    int    `json:"danmaku"`
			Reply      int    `json:"reply"`
			Favorite   int    `json:"favorite"`
			Coin       int    `json:"coin"`
			Share      int    `json:"share"`
			NowRank    int    `json:"now_rank"`
			HisRank    int    `json:"his_rank"`
			Like       int    `json:"like"`
			Dislike    int    `json:"dislike"`
			Evaluation string `json:"evaluation"`
			Vt         int    `json:"vt"`
		} `json:"stat"`
		ArgueInfo struct {
			ArgueMsg  string `json:"argue_msg"`
			ArgueType int    `json:"argue_type"`
			ArgueLink string `json:"argue_link"`
		} `json:"argue_info"`
		Dynamic   string `json:"dynamic"`
		Cid       int    `json:"cid"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		Premiere           interface{} `json:"premiere"`
		TeenageMode        int         `json:"teenage_mode"`
		IsChargeableSeason bool        `json:"is_chargeable_season"`
		IsStory            bool        `json:"is_story"`
		IsUpowerExclusive  bool        `json:"is_upower_exclusive"`
		IsUpowerPlay       bool        `json:"is_upower_play"`
		IsUpowerPreview    bool        `json:"is_upower_preview"`
		EnableVt           int         `json:"enable_vt"`
		VtDisplay          string      `json:"vt_display"`
		NoCache            bool        `json:"no_cache"`
		Pages              []struct {
			Cid       int    `json:"cid"`
			Page      int    `json:"page"`
			From      string `json:"from"`
			Part      string `json:"part"`
			Duration  int    `json:"duration"`
			Vid       string `json:"vid"`
			Weblink   string `json:"weblink"`
			Dimension struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				Rotate int `json:"rotate"`
			} `json:"dimension"`
		} `json:"pages"`
		Subtitle struct {
			AllowSubmit bool          `json:"allow_submit"`
			List        []interface{} `json:"list"`
		} `json:"subtitle"`
		Staff []struct {
			Mid   int    `json:"mid"`
			Title string `json:"title"`
			Name  string `json:"name"`
			Face  string `json:"face"`
			Vip   struct {
				Type       int   `json:"type"`
				Status     int   `json:"status"`
				DueDate    int64 `json:"due_date"`
				VipPayType int   `json:"vip_pay_type"`
				ThemeType  int   `json:"theme_type"`
				Label      struct {
					Path                  string `json:"path"`
					Text                  string `json:"text"`
					LabelTheme            string `json:"label_theme"`
					TextColor             string `json:"text_color"`
					BgStyle               int    `json:"bg_style"`
					BgColor               string `json:"bg_color"`
					BorderColor           string `json:"border_color"`
					UseImgLabel           bool   `json:"use_img_label"`
					ImgLabelURIHans       string `json:"img_label_uri_hans"`
					ImgLabelURIHant       string `json:"img_label_uri_hant"`
					ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
					ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
				} `json:"label"`
				AvatarSubscript    int    `json:"avatar_subscript"`
				NicknameColor      string `json:"nickname_color"`
				Role               int    `json:"role"`
				AvatarSubscriptURL string `json:"avatar_subscript_url"`
				TvVipStatus        int    `json:"tv_vip_status"`
				TvVipPayType       int    `json:"tv_vip_pay_type"`
				TvDueDate          int    `json:"tv_due_date"`
				AvatarIcon         struct {
					IconType     int `json:"icon_type"`
					IconResource struct {
					} `json:"icon_resource"`
				} `json:"avatar_icon"`
			} `json:"vip"`
			Official struct {
				Role  int    `json:"role"`
				Title string `json:"title"`
				Desc  string `json:"desc"`
				Type  int    `json:"type"`
			} `json:"official"`
			Follower   int `json:"follower"`
			LabelStyle int `json:"label_style"`
		} `json:"staff"`
		IsSeasonDisplay bool `json:"is_season_display"`
		UserGarb        struct {
			URLImageAniCut string `json:"url_image_ani_cut"`
		} `json:"user_garb"`
		HonorReply struct {
			Honor []struct {
				Aid                int    `json:"aid"`
				Type               int    `json:"type"`
				Desc               string `json:"desc"`
				WeeklyRecommendNum int    `json:"weekly_recommend_num"`
			} `json:"honor"`
		} `json:"honor_reply"`
		LikeIcon          string `json:"like_icon"`
		NeedJumpBv        bool   `json:"need_jump_bv"`
		DisableShowUpInfo bool   `json:"disable_show_up_info"`
		IsStoryPlay       int    `json:"is_story_play"`
		IsViewSelf        bool   `json:"is_view_self"`
	} `json:"data"`
}
