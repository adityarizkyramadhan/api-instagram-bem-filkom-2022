package entity

type Response struct {
	Success interface{}            `json:"success"`
	Data    map[string]interface{} `json:"data"`
}

type ResponseKedua struct {
	Body       map[string]interface{} `json:"body"`
	Status     string                 `json:"status"`
	StatusCode int                    `json:"status_code"`
}

type T struct {
	Body struct {
		MoreAvailable bool `json:"more_available"`
		User          struct {
			IsPrivate     bool   `json:"is_private"`
			Pk            int64  `json:"pk"`
			FullName      string `json:"full_name"`
			ProfilePicId  string `json:"profile_pic_id"`
			IsVerified    bool   `json:"is_verified"`
			Username      string `json:"username"`
			ProfilePicUrl string `json:"profile_pic_url"`
		} `json:"user"`
		Items []struct {
			FacepileTopLikers struct {
			} `json:"facepile_top_likers"`
			HasDelayedMetadata              bool        `json:"has_delayed_metadata"`
			TakenAt                         int         `json:"taken_at"`
			ProductType                     string      `json:"product_type"`
			CommerceIntegrityReviewDecision interface{} `json:"commerce_integrity_review_decision"`
			CaptionIsEdited                 bool        `json:"caption_is_edited"`
			IntegrityReviewDecision         string      `json:"integrity_review_decision"`
			Pk                              float64     `json:"pk"`
			ShouldRequestAds                bool        `json:"should_request_ads"`
			MediaType                       int         `json:"media_type"`
			HasMoreComments                 bool        `json:"has_more_comments"`
			IsPinned                        bool        `json:"is_pinned"`
			User                            struct {
				IsPrivate                  bool   `json:"is_private"`
				Pk                         int64  `json:"pk"`
				FullName                   string `json:"full_name"`
				ProfilePicId               string `json:"profile_pic_id"`
				LatestReelMedia            int    `json:"latest_reel_media"`
				ProfilePicUrl              string `json:"profile_pic_url"`
				HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
				FanClubInfo                struct {
					FanClubId   interface{} `json:"fan_club_id"`
					FanClubName interface{} `json:"fan_club_name"`
				} `json:"fan_club_info"`
				AccountBadges struct {
				} `json:"account_badges"`
				Username                   string `json:"username"`
				HasHighlightReels          bool   `json:"has_highlight_reels"`
				IsFavorite                 bool   `json:"is_favorite"`
				IsVerified                 bool   `json:"is_verified"`
				TransparencyProductEnabled bool   `json:"transparency_product_enabled"`
				IsUnpublished              bool   `json:"is_unpublished"`
			} `json:"user"`
			CanViewMorePreviewComments bool   `json:"can_view_more_preview_comments"`
			Code                       string `json:"code"`
			PhotoOfYou                 bool   `json:"photo_of_you"`
			PreviewComments            struct {
			} `json:"preview_comments"`
			Comments struct {
			} `json:"comments"`
			LikeCount                      int    `json:"like_count"`
			CommercialityStatus            string `json:"commerciality_status"`
			ClientCacheKey                 string `json:"client_cache_key"`
			OrganicTrackingToken           string `json:"organic_tracking_token"`
			Id                             string `json:"id"`
			InlineComposerDisplayCondition string `json:"inline_composer_display_condition"`
			IsInProfileGrid                bool   `json:"is_in_profile_grid"`
			IsUnifiedVideo                 bool   `json:"is_unified_video"`
			CanSeeInsightsAsBrand          bool   `json:"can_see_insights_as_brand"`
			LikeAndViewCountsDisabled      bool   `json:"like_and_view_counts_disabled"`
			SharingFrictionInfo            struct {
				SharingFrictionPayload    interface{} `json:"sharing_friction_payload"`
				BloksAppUrl               interface{} `json:"bloks_app_url"`
				ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
			} `json:"sharing_friction_info"`
			CommentInformTreatment struct {
				ShouldHaveInformTreatment bool        `json:"should_have_inform_treatment"`
				ActionType                interface{} `json:"action_type"`
				Url                       interface{} `json:"url"`
				Text                      string      `json:"text"`
			} `json:"comment_inform_treatment"`
			IsOrganicProductTaggingEligible bool `json:"is_organic_product_tagging_eligible"`
			OriginalHeight                  int  `json:"original_height"`
			MaxNumVisiblePreviewComments    int  `json:"max_num_visible_preview_comments"`
			MusicMetadata                   struct {
				MusicCanonicalId  string      `json:"music_canonical_id"`
				MusicInfo         interface{} `json:"music_info"`
				PinnedMediaIds    interface{} `json:"pinned_media_ids"`
				AudioType         interface{} `json:"audio_type"`
				OriginalSoundInfo interface{} `json:"original_sound_info"`
			} `json:"music_metadata"`
			TopLikers struct {
			} `json:"top_likers"`
			HasLiked                     bool    `json:"has_liked"`
			ProfileGridControlEnabled    bool    `json:"profile_grid_control_enabled"`
			IsPaidPartnership            bool    `json:"is_paid_partnership"`
			HideViewAllCommentEntrypoint bool    `json:"hide_view_all_comment_entrypoint"`
			CommentLikesEnabled          bool    `json:"comment_likes_enabled"`
			FilterType                   int     `json:"filter_type"`
			HasSharedToFb                int     `json:"has_shared_to_fb"`
			CommentCount                 int     `json:"comment_count"`
			DeviceTimestamp              float64 `json:"device_timestamp"`
			CommentThreadingEnabled      bool    `json:"comment_threading_enabled"`
			DeletedReason                int     `json:"deleted_reason"`
			Caption                      struct {
				Pk                 float64 `json:"pk"`
				PrivateReplyStatus int     `json:"private_reply_status"`
				MediaId            float64 `json:"media_id"`
				IsCovered          bool    `json:"is_covered"`
				User               struct {
					IsPrivate     bool   `json:"is_private"`
					Pk            int64  `json:"pk"`
					FullName      string `json:"full_name"`
					ProfilePicId  string `json:"profile_pic_id"`
					IsVerified    bool   `json:"is_verified"`
					Username      string `json:"username"`
					ProfilePicUrl string `json:"profile_pic_url"`
				} `json:"user"`
				Status          string `json:"status"`
				Text            string `json:"text"`
				ShareEnabled    bool   `json:"share_enabled"`
				DidReportAsSpam bool   `json:"did_report_as_spam"`
				BitFlags        int    `json:"bit_flags"`
				Type            int    `json:"type"`
				UserId          int64  `json:"user_id"`
				CreatedAt       int    `json:"created_at"`
				ContentType     string `json:"content_type"`
				CreatedAtUtc    int    `json:"created_at_utc"`
			} `json:"caption"`
			CanViewerReshare                    bool `json:"can_viewer_reshare,omitempty"`
			IsVisualReplyCommenterNoticeEnabled bool `json:"is_visual_reply_commenter_notice_enabled"`
			OriginalWidth                       int  `json:"original_width"`
			ImageVersions2                      struct {
				Candidates []struct {
					Url          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					ScansProfile string `json:"scans_profile,omitempty"`
				} `json:"candidates"`
				AdditionalCandidates struct {
					IgtvFirstFrame struct {
						Url          string `json:"url"`
						ScansProfile string `json:"scans_profile"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
					} `json:"igtv_first_frame"`
					FirstFrame struct {
						Url          string `json:"url"`
						ScansProfile string `json:"scans_profile"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
					} `json:"first_frame"`
				} `json:"additional_candidates,omitempty"`
			} `json:"image_versions2"`
			IsPostLive               bool `json:"is_post_live,omitempty"`
			IgtvExistsInViewerSeries bool `json:"igtv_exists_in_viewer_series,omitempty"`
			VideoVersions            []struct {
				Height int    `json:"height"`
				Url    string `json:"url"`
				Id     string `json:"id"`
				Width  int    `json:"width"`
				Type   int    `json:"type"`
			} `json:"video_versions,omitempty"`
			NearlyCompleteCopyrightMatch bool    `json:"nearly_complete_copyright_match,omitempty"`
			HasAudio                     bool    `json:"has_audio,omitempty"`
			VideoDuration                float64 `json:"video_duration,omitempty"`
			MediaCroppingInfo            struct {
			} `json:"media_cropping_info,omitempty"`
			Title      string `json:"title,omitempty"`
			ViewCount  int    `json:"view_count,omitempty"`
			Thumbnails struct {
				SpriteWidth                int      `json:"sprite_width"`
				ThumbnailDuration          float64  `json:"thumbnail_duration"`
				VideoLength                float64  `json:"video_length"`
				MaxThumbnailsPerSprite     int      `json:"max_thumbnails_per_sprite"`
				SpriteHeight               int      `json:"sprite_height"`
				SpriteUrls                 []string `json:"sprite_urls"`
				TotalThumbnailNumPerSprite int      `json:"total_thumbnail_num_per_sprite"`
				FileSizeKb                 int      `json:"file_size_kb"`
				RenderedWidth              int      `json:"rendered_width"`
				ThumbnailWidth             int      `json:"thumbnail_width"`
				ThumbnailsPerRow           int      `json:"thumbnails_per_row"`
				ThumbnailHeight            int      `json:"thumbnail_height"`
			} `json:"thumbnails,omitempty"`
			Usertags struct {
				In []struct {
					StartTimeInVideoInSec interface{} `json:"start_time_in_video_in_sec"`
					Position              []float64   `json:"position"`
					DurationInVideoInSec  interface{} `json:"duration_in_video_in_sec"`
					User                  struct {
						IsPrivate     bool   `json:"is_private"`
						Pk            int64  `json:"pk"`
						FullName      string `json:"full_name"`
						ProfilePicId  string `json:"profile_pic_id"`
						IsVerified    bool   `json:"is_verified"`
						Username      string `json:"username"`
						ProfilePicUrl string `json:"profile_pic_url"`
					} `json:"user"`
				} `json:"in"`
			} `json:"usertags,omitempty"`
		} `json:"items"`
		NextMaxId           string `json:"next_max_id"`
		AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
		Status              string `json:"status"`
		NumResults          int    `json:"num_results"`
	} `json:"body"`
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
}
