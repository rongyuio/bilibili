package bilibili

import "encoding/json"

type DynamicItem struct {
	Basic struct { // 见 https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/dynamic/all.md#data%E5%AF%B9%E8%B1%A1---items%E6%95%B0%E7%BB%84%E4%B8%AD%E7%9A%84%E5%AF%B9%E8%B1%A1---basic%E5%AF%B9%E8%B1%A1
		CommentIdStr string `json:"comment_id_str"`
		CommentType  int    `json:"comment_type"`
		LikeIcon     struct {
			ActionUrl string      `json:"action_url"`
			EndUrl    string      `json:"end_url"`
			Id        json.Number `json:"id"`
			StartUrl  string      `json:"start_url"`
		} `json:"like_icon"`
		RidStr string `json:"rid_str"`
	} `json:"basic"`
	IdStr   json.Number `json:"id_str"` // 这个字段，B站返回的数据有时是number，有时是string
	Modules struct {
		ModuleAuthor struct {
			Avatar struct {
				ContainerSize struct {
					Height float64 `json:"height"`
					Width  float64 `json:"width"`
				} `json:"container_size"`
				FallbackLayers struct {
					IsCriticalGroup bool `json:"is_critical_group"`
					Layers          []struct {
						GeneralSpec struct {
							PosSpec struct {
								AxisX         float64 `json:"axis_x"`
								AxisY         float64 `json:"axis_y"`
								CoordinatePos int     `json:"coordinate_pos"`
							} `json:"pos_spec"`
							RenderSpec struct {
								Opacity int `json:"opacity"`
							} `json:"render_spec"`
							SizeSpec struct {
								Height float64 `json:"height"`
								Width  float64 `json:"width"`
							} `json:"size_spec"`
						} `json:"general_spec"`
						LayerConfig struct {
							IsCritical bool `json:"is_critical,omitempty"`
							Tags       struct {
								AvatarLayer struct {
								} `json:"AVATAR_LAYER,omitempty"`
								GeneralCfg struct {
									ConfigType    int `json:"config_type"`
									GeneralConfig struct {
										WebCssStyle struct {
											BorderRadius    string `json:"borderRadius"`
											BackgroundColor string `json:"background-color,omitempty"`
											Border          string `json:"border,omitempty"`
											BoxSizing       string `json:"boxSizing,omitempty"`
										} `json:"web_css_style"`
									} `json:"general_config"`
								} `json:"GENERAL_CFG"`
								IconLayer struct{} `json:"ICON_LAYER,omitempty"`
							} `json:"tags"`
						} `json:"layer_config"`
						Resource struct {
							ResAnimation struct {
								WebpSrc struct {
									Placeholder int `json:"placeholder"`
									Remote      struct {
										BfsStyle string `json:"bfs_style"`
										Url      string `json:"url"`
									} `json:"remote"`
									SrcType int `json:"src_type"`
								} `json:"webp_src"`
							} `json:"res_animation,omitempty"`
							ResType  int `json:"res_type"`
							ResImage struct {
								ImageSrc struct {
									Local   int `json:"local"`
									SrcType int `json:"src_type"`
								} `json:"image_src"`
							} `json:"res_image,omitempty"`
						} `json:"resource"`
						Visible bool `json:"visible"`
					} `json:"layers"`
				} `json:"fallback_layers"`
				Mid string `json:"mid"`
			} `json:"avatar"`
			Face           string      `json:"face"`
			FaceNft        bool        `json:"face_nft"`
			Following      bool        `json:"following"`
			JumpUrl        string      `json:"jump_url"`
			Label          string      `json:"label"`
			Mid            json.Number `json:"mid"`
			Name           string      `json:"name"`
			OfficialVerify struct {
				Desc string `json:"desc"`
				Type int    `json:"type"`
			} `json:"official_verify"`
			Pendant struct {
				Expire            json.Number `json:"expire"`
				Image             string      `json:"image"`
				ImageEnhance      string      `json:"image_enhance"`
				ImageEnhanceFrame string      `json:"image_enhance_frame"`
				NPid              json.Number `json:"n_pid"`
				Name              string      `json:"name"`
				Pid               json.Number `json:"pid"`
			} `json:"pendant"`
			PubAction       string      `json:"pub_action"`
			PubLocationText string      `json:"pub_location_text"`
			PubTime         string      `json:"pub_time"`
			PubTs           json.Number `json:"pub_ts"`
			Type            string      `json:"type"`
			Vip             struct {
				AvatarSubscript    int         `json:"avatar_subscript"`
				AvatarSubscriptUrl string      `json:"avatar_subscript_url"`
				DueDate            json.Number `json:"due_date"`
				Label              struct {
					BgColor               string `json:"bg_color"`
					BgStyle               int    `json:"bg_style"`
					BorderColor           string `json:"border_color"`
					ImgLabelUriHans       string `json:"img_label_uri_hans"`
					ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"`
					ImgLabelUriHant       string `json:"img_label_uri_hant"`
					ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"`
					LabelTheme            string `json:"label_theme"`
					Path                  string `json:"path"`
					Text                  string `json:"text"`
					TextColor             string `json:"text_color"`
					UseImgLabel           bool   `json:"use_img_label"`
				} `json:"label"`
				NicknameColor string `json:"nickname_color"`
				Status        int    `json:"status"`
				ThemeType     int    `json:"theme_type"`
				Type          int    `json:"type"`
			} `json:"vip"`
		} `json:"module_author"`
		ModuleDynamic struct {
			Additional any `json:"additional"`
			Desc       *struct {
				RichTextNodes []struct {
					OrigText string `json:"orig_text"`
					Text     string `json:"text"`
					Type     string `json:"type"`
					JumpUrl  string `json:"jump_url,omitempty"`
					Style    any    `json:"style"`
					Emoji    struct {
						IconUrl string `json:"icon_url"`
						Size    int    `json:"size"`
						Text    string `json:"text"`
						Type    int    `json:"type"`
					} `json:"emoji,omitempty"`
					Rid string `json:"rid,omitempty"`
				} `json:"rich_text_nodes"`
				Text string `json:"text"`
			} `json:"desc"`
			Major *struct {
				Draw struct {
					Id    json.Number `json:"id"`
					Items []struct {
						Height json.Number `json:"height"`
						Size   json.Number `json:"size"`
						Src    string      `json:"src"`
						Tags   []any       `json:"tags"`
						Width  json.Number `json:"width"`
					} `json:"items"`
				} `json:"draw,omitempty"`
				Type    string `json:"type"`
				Archive struct {
					Aid   string `json:"aid"`
					Badge struct {
						BgColor string `json:"bg_color"`
						Color   string `json:"color"`
						IconUrl any    `json:"icon_url"`
						Text    string `json:"text"`
					} `json:"badge"`
					Bvid           string `json:"bvid"`
					Cover          string `json:"cover"`
					Desc           string `json:"desc"`
					DisablePreview int    `json:"disable_preview"`
					DurationText   string `json:"duration_text"`
					JumpUrl        string `json:"jump_url"`
					Stat           struct {
						Danmaku string `json:"danmaku"`
						Play    string `json:"play"`
					} `json:"stat"`
					Title string `json:"title"`
					Type  int    `json:"type"`
				} `json:"archive,omitempty"`
			} `json:"major"`
			Topic any `json:"topic"`
		} `json:"module_dynamic"`
		ModuleMore struct {
			ThreePointItems []struct {
				Label string `json:"label"`
				Type  string `json:"type"`
			} `json:"three_point_items"`
		} `json:"module_more"`
		ModuleStat struct {
			Comment struct {
				Count     json.Number `json:"count"`
				Forbidden bool        `json:"forbidden"`
			} `json:"comment"`
			Forward struct {
				Count     json.Number `json:"count"`
				Forbidden bool        `json:"forbidden"`
			} `json:"forward"`
			Like struct {
				Count     json.Number `json:"count"`
				Forbidden bool        `json:"forbidden"`
				Status    bool        `json:"status"`
			} `json:"like"`
		} `json:"module_stat"`
	} `json:"modules"`
	Orig struct {
		Basic struct {
			CommentIdStr string `json:"comment_id_str"`
			CommentType  int    `json:"comment_type"`
			LikeIcon     struct {
				ActionUrl string `json:"action_url"`
				EndUrl    string `json:"end_url"`
				Id        int    `json:"id"`
				StartUrl  string `json:"start_url"`
			} `json:"like_icon"`
			RidStr string `json:"rid_str"`
		} `json:"basic"`
		IdStr   json.Number `json:"id_str"`
		Modules struct {
			ModuleAuthor struct {
				Avatar struct {
					ContainerSize struct {
						Height float64 `json:"height"`
						Width  float64 `json:"width"`
					} `json:"container_size"`
					FallbackLayers struct {
						IsCriticalGroup bool `json:"is_critical_group"`
						Layers          []struct {
							GeneralSpec struct {
								PosSpec struct {
									AxisX         float64 `json:"axis_x"`
									AxisY         float64 `json:"axis_y"`
									CoordinatePos int     `json:"coordinate_pos"`
								} `json:"pos_spec"`
								RenderSpec struct {
									Opacity int `json:"opacity"`
								} `json:"render_spec"`
								SizeSpec struct {
									Height float64 `json:"height"`
									Width  float64 `json:"width"`
								} `json:"size_spec"`
							} `json:"general_spec"`
							LayerConfig struct {
								IsCritical bool `json:"is_critical,omitempty"`
								Tags       struct {
									AVATARLAYER struct {
									} `json:"AVATAR_LAYER,omitempty"`
									GENERALCFG struct {
										ConfigType    int `json:"config_type"`
										GeneralConfig struct {
											WebCssStyle struct {
												BorderRadius    string `json:"borderRadius"`
												BackgroundColor string `json:"background-color,omitempty"`
												Border          string `json:"border,omitempty"`
												BoxSizing       string `json:"boxSizing,omitempty"`
											} `json:"web_css_style"`
										} `json:"general_config"`
									} `json:"GENERAL_CFG,omitempty"`
									PENDENTLAYER struct {
									} `json:"PENDENT_LAYER,omitempty"`
									ICONLAYER struct {
									} `json:"ICON_LAYER,omitempty"`
								} `json:"tags"`
							} `json:"layer_config"`
							Resource struct {
								ResImage struct {
									ImageSrc struct {
										Placeholder int `json:"placeholder,omitempty"`
										Remote      struct {
											BfsStyle string `json:"bfs_style"`
											Url      string `json:"url"`
										} `json:"remote,omitempty"`
										SrcType int `json:"src_type"`
										Local   int `json:"local,omitempty"`
									} `json:"image_src"`
								} `json:"res_image"`
								ResType int `json:"res_type"`
							} `json:"resource"`
							Visible bool `json:"visible"`
						} `json:"layers"`
					} `json:"fallback_layers"`
					Mid string `json:"mid"`
				} `json:"avatar"`
				Decorate struct {
					CardUrl string `json:"card_url"`
					Fan     struct {
						Color  string      `json:"color"`
						IsFan  bool        `json:"is_fan"`
						NumStr string      `json:"num_str"`
						Number json.Number `json:"number"`
					} `json:"fan"`
					Id      json.Number `json:"id"`
					JumpUrl string      `json:"jump_url"`
					Name    string      `json:"name"`
					Type    int         `json:"type"`
				} `json:"decorate,omitempty"`
				Face           string      `json:"face"`
				FaceNft        bool        `json:"face_nft"`
				Following      bool        `json:"following"`
				JumpUrl        string      `json:"jump_url"`
				Label          string      `json:"label"`
				Mid            json.Number `json:"mid"`
				Name           string      `json:"name"`
				OfficialVerify struct {
					Desc string `json:"desc"`
					Type int    `json:"type"`
				} `json:"official_verify"`
				Pendant struct {
					Expire            json.Number `json:"expire"`
					Image             string      `json:"image"`
					ImageEnhance      string      `json:"image_enhance"`
					ImageEnhanceFrame string      `json:"image_enhance_frame"`
					NPid              json.Number `json:"n_pid"`
					Name              string      `json:"name"`
					Pid               int         `json:"pid"`
				} `json:"pendant"`
				PubAction string      `json:"pub_action"`
				PubTime   string      `json:"pub_time"`
				PubTs     json.Number `json:"pub_ts"`
				Type      string      `json:"type"`
				Vip       struct {
					AvatarSubscript    int         `json:"avatar_subscript"`
					AvatarSubscriptUrl string      `json:"avatar_subscript_url"`
					DueDate            json.Number `json:"due_date"`
					Label              struct {
						BgColor               string `json:"bg_color"`
						BgStyle               int    `json:"bg_style"`
						BorderColor           string `json:"border_color"`
						ImgLabelUriHans       string `json:"img_label_uri_hans"`
						ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"`
						ImgLabelUriHant       string `json:"img_label_uri_hant"`
						ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"`
						LabelTheme            string `json:"label_theme"`
						Path                  string `json:"path"`
						Text                  string `json:"text"`
						TextColor             string `json:"text_color"`
						UseImgLabel           bool   `json:"use_img_label"`
					} `json:"label"`
					NicknameColor string `json:"nickname_color"`
					Status        int    `json:"status"`
					ThemeType     int    `json:"theme_type"`
					Type          int    `json:"type"`
				} `json:"vip"`
			} `json:"module_author"`
			ModuleDynamic struct {
				Additional any `json:"additional"`
				Desc       *struct {
					RichTextNodes []struct {
						JumpUrl  string `json:"jump_url,omitempty"`
						OrigText string `json:"orig_text"`
						Text     string `json:"text"`
						Type     string `json:"type"`
						Emoji    struct {
							IconUrl string `json:"icon_url"`
							Size    int    `json:"size"`
							Text    string `json:"text"`
							Type    int    `json:"type"`
						} `json:"emoji,omitempty"`
					} `json:"rich_text_nodes"`
					Text string `json:"text"`
				} `json:"desc"`
				Major struct {
					Archive struct {
						Aid   string `json:"aid"`
						Badge struct {
							BgColor string `json:"bg_color"`
							Color   string `json:"color"`
							IconUrl any    `json:"icon_url"`
							Text    string `json:"text"`
						} `json:"badge"`
						Bvid           string `json:"bvid"`
						Cover          string `json:"cover"`
						Desc           string `json:"desc"`
						DisablePreview int    `json:"disable_preview"`
						DurationText   string `json:"duration_text"`
						JumpUrl        string `json:"jump_url"`
						Stat           struct {
							Danmaku string `json:"danmaku"`
							Play    string `json:"play"`
						} `json:"stat"`
						Title string `json:"title"`
						Type  int    `json:"type"`
					} `json:"archive,omitempty"`
					Type string `json:"type"`
					Draw struct {
						Id    json.Number `json:"id"`
						Items []struct {
							Height json.Number `json:"height"`
							Size   json.Number `json:"size"`
							Src    string      `json:"src"`
							Tags   []any       `json:"tags"`
							Width  json.Number `json:"width"`
						} `json:"items"`
					} `json:"draw,omitempty"`
				} `json:"major"`
				Topic any `json:"topic"`
			} `json:"module_dynamic"`
		} `json:"modules"`
		Type    string `json:"type"`
		Visible bool   `json:"visible"`
	} `json:"orig,omitempty"`
	Type    string `json:"type"`
	Visible bool   `json:"visible"`
}

type DynamicInfo struct {
	HasMore        bool          `json:"has_more"`        // 是否有更多数据
	Items          []DynamicItem `json:"items"`           // 数据数组
	Offset         string        `json:"offset"`          // 偏移量，等于items中最后一条记录的id，获取下一页时使用
	UpdateBaseline string        `json:"update_baseline"` // 更新基线，等于items中第一条记录的id
	UpdateNum      json.Number   `json:"update_num"`      // 本次获取获取到了多少条新动态，在更新基线以上的动态条数
}
