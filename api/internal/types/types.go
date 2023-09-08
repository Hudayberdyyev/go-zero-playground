// Code generated by goctl. DO NOT EDIT.
package types

type GetChannelInfoReq struct {
	YoutubeId string `path:"youtubeId"`
	BaseReqHeaders
}

type GetChannelInfoResp struct {
	YoutubeId        string `json:"youtubeId"`
	Title            string `json:"title"`
	AvatarUrl        string `json:"avatarUrl"`
	BannerUrl        string `json:"bannerUrl"`
	Description      string `json:"description"`
	SubscribersCount int    `json:"subscribersCount"`
	UploadedDate     string `json:"uploadedDate"`
	IsSubscribed     bool   `json:"isSubscribed"`
}

type GetChannelVideosReq struct {
	YoutubeId string `path:"youtubeId"`
	Pagination
	BaseReqHeaders
}

type GetChannelVideosResp struct {
	Content  []ChannelVideo `json:"content"`
	NextPage int            `json:"nextPage"`
}

type GetChannelPlaylistsReq struct {
	YoutubeId string `path:"youtubeId"`
	Pagination
	BaseReqHeaders
}

type GetChannelPlaylistsResp struct {
	Content  []ChannelPlaylist `json:"content"`
	NextPage int               `json:"nextPage"`
}

type GetPlaylistInfoReq struct {
	PlaylistId string `path:"playlistId"`
	BaseReqHeaders
}

type GetPlaylistInfoResp struct {
	PlaylistId       string `json:"playlistId"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	Thumbnail        string `json:"thumbnail"`
	ChannelName      string `json:"channelName"`
	ChannelYoutubeId string `json:"channelYoutubeId"`
	IsReadOnly       bool   `json:"isReadOnly"`
	VideosCount      int    `json:"videosCount"`
	UpdatedDate      string `json:"updatedAt"`
}

type GetPlaylistVideosReq struct {
	PlaylistId string `path:"playlistId"`
	Pagination
	BaseReqHeaders
}

type GetPlaylistVideosResp struct {
	Content  []PlaylistVideo `json:"content"`
	NextPage int             `json:"nextPage"`
}

type GetUserPlaylistsReq struct {
	Pagination
	BaseReqHeaders
}

type GetUserPlaylistsResp struct {
	Content  []ChannelPlaylist `json:"content"`
	NextPage int               `json:"nextPage"`
}

type GetUserPlaylistsByVideoReq struct {
	VideoYoutubeId string `path:"videoYoutubeId"`
	Pagination
	BaseReqHeaders
}

type GetUserPlaylistsByVideoResp struct {
	Content  []UserPlaylistForVideo `json:"content"`
	NextPage int                    `json:"nextPage"`
}

type GetVideoInfoReq struct {
	YoutubeId string `path:"youtubeId"`
	BaseReqHeaders
}

type GetVideoInfoResp struct {
	YoutubeId              string         `json:"youtubeId"`
	Title                  string         `json:"title"`
	Description            string         `json:"description"`
	Thumbnail              string         `json:"thumbnail"`
	Duration               int            `json:"duration"`
	AdUri                  string         `json:"adUri"`
	Hls                    string         `json:"hls"`
	IsSubscribed           bool           `json:"isSubscribed"`
	Likes                  int            `json:"likes"`
	VideoReaction          string         `json:"videoReaction"`
	WatchedTime            int            `json:"watchedTime"`
	UploadedDate           string         `json:"uploadedDate"`
	ChannelName            string         `json:"channelName"`
	ChannelAvatarUrl       string         `json:"channelAvatarUrl"`
	ChannelYoutubeId       string         `json:"channelYoutubeId"`
	ChannelSubscriberCount int            `json:"channelSubscriberCount"`
	RelatedVideos          []RelatedVideo `json:"relatedVideos"`
}

type SearchReq struct {
	SearchQuery string `form:"searchQuery"`
	Type        string `form:"type,options=channel|video|all,default=all"`
	BaseReqHeaders
}

type SearchResp struct {
	Content  []SearchResult `json:"content"`
	NextPage int            `json:"nextPage"`
}

type SuggestionReq struct {
	SearchQuery string `form:"searchQuery"`
	BaseReqHeaders
}

type SuggestionResp struct {
	Content []string `json:"content"`
}

type GetTrendingListReq struct {
	Genre int `form:"genreId,optional"`
	Pagination
	BaseReqHeaders
}

type GetTrendingListResp struct {
	Content  []TrendingVideo `json:"content"`
	NextPage int             `json:"nextPage"`
}

type GetGenresReq struct {
	Hl string `form:"hl,options=ru|en|tm,default=ru""`
	BaseReqHeaders
}

type GetGenresResp struct {
	Content []Genre `json:"content"`
}

type GetReportTypesReq struct {
	Hl string `form:"hl,options=ru|en|tm,default=ru"`
	BaseReqHeaders
}

type GetReportTypesResp struct {
	Content []ReportType `json:"content"`
}

type GetReportsReq struct {
	Pagination
	BaseReqHeaders
}

type GetReportsResp struct {
	Content  []Report `json:"content"`
	NextPage int      `json:"nextPage"`
}

type SubscriptionsReq struct {
	BaseReqHeaders
	Pagination
}

type SubscriptionsResp struct {
	Content  []SubscriptionVideo `json:"content"`
	NextPage int                 `json:"nextPage"`
}

type FeedChannelsReq struct {
	BaseReqHeaders
	Pagination
}

type FeedChannelsResp struct {
	Content  []SubscribedChannel `json:"content"`
	NextPage int                 `json:"nextPage"`
}

type LikedVideosReq struct {
	BaseReqHeaders
	Pagination
}

type LikedVideosResp struct {
	Content    []FeedVideo `json:"content"`
	TotalCount int         `json:"totalCount"`
	NextPage   int         `json:"nextPage"`
}

type HistoryReq struct {
	BaseReqHeaders
	Pagination
}

type HistoryResp struct {
	Content  []FeedVideo `json:"content"`
	NextPage int         `json:"nextPage"`
}

type LibraryResp struct {
	History     []FeedVideo       `json:"history"`
	Playlists   []ChannelPlaylist `json:"playlists"`
	LikedVideos []FeedVideo       `json:"liked"`
}

type ChannelSubscriptionReq struct {
	Action    string `form:"action,options=subscribe|unsubscribe"`
	YoutubeId string `path:"youtubeId"`
	BaseReqHeaders
}

type PlaylistVideoActionReq struct {
	VideoYoutubeId string `json:"videoYoutubeId"`
	PlaylistId     string `path:"playlistId"`
	BaseReqHeaders
}

type UpdateUserPlaylistReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PlaylistId  string `path:"playlistId"`
	BaseReqHeaders
}

type DeleteUserPlaylistReq struct {
	PlaylistId string `path:"playlistId"`
	BaseReqHeaders
}

type ImportYoutubePlaylistReq struct {
	ChannelPlaylistId string `path:"channelPlaylistId"`
	BaseReqHeaders
}

type VideoReactionReq struct {
	Action    string `form:"action,options=like|dislike|revertLike|revertDislike"`
	YoutubeId string `path:"youtubeId"`
	BaseReqHeaders
}

type VideoIncViewReq struct {
	YoutubeId string `path:"youtubeId"`
	BaseReqHeaders
}

type VideoSetWatchTimeReq struct {
	WatchedTime int    `json:"watchedTime"`
	YoutubeId   string `path:"youtubeId"`
	BaseReqHeaders
}

type CreateUserPlaylistReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	BaseReqHeaders
}

type SendFeedbackReq struct {
	Image       string `form:"image"`
	Description string `form:"description"`
	BaseReqHeaders
}

type SendReportReq struct {
	ContentType  string `json:"content,options=channel|video"`
	ReportTypeId int    `json:"reportTypeId"`
	YoutubeId    string `json:"youtube"`
	BaseReqHeaders
}

type MessageResp struct {
	Message string `json:"message"`
}

type Pagination struct {
	Page  int `form:"page,optional,range=[1:1000],default=1"`
	Limit int `form:"limit,optional,range=[1:1000],default=10"`
}

type BaseReqHeaders struct {
	Token        string `header:"authorization,optional"`
	PlatformType string `header:"platformType"`
	VersionCode  int    `header:"versionCode"`
}

type ChannelVideo struct {
	YoutubeId    string `json:"youtubeId"`
	Title        string `json:"title"`
	Duration     int    `json:"duration"`
	Thumbnail    string `json:"thumbnail"`
	ViewsCount   int    `json:"viewsCount"`
	WatchedTime  int    `json:"watchedTime"`
	UploadedDate string `json:"uploadedDate"`
}

type TrendingVideo struct {
	YoutubeId        string `json:"youtubeId"`
	Title            string `json:"title"`
	Duration         int    `json:"duration"`
	Thumbnail        string `json:"thumbnail"`
	ViewsCount       int    `json:"viewsCount"`
	Hls              string `json:"hls"`
	WatchedTime      int    `json:"watchedTime"`
	UploadedDate     string `json:"uploadedDate"`
	ChannelName      string `json:"channelName"`
	ChannelAvatarUrl string `json:"channelAvatarUrl"`
	ChannelYoutubeId string `json:"channelYoutubeId"`
}

type RelatedVideo struct {
	YoutubeId        string `json:"youtubeId"`
	Title            string `json:"title"`
	Duration         int    `json:"duration"`
	Thumbnail        string `json:"thumbnail"`
	ViewsCount       int    `json:"viewsCount"`
	WatchedTime      int    `json:"watchedTime"`
	UploadedDate     string `json:"uploadedDate"`
	ChannelName      string `json:"channelName"`
	ChannelAvatarUrl string `json:"channelAvatarUrl"`
	ChannelYoutubeId string `json:"channelYoutubeId"`
}

type PlaylistVideo struct {
	YoutubeId        string `json:"youtubeId"`
	Title            string `json:"title"`
	Thumbnail        string `json:"thumbnail"`
	Duration         int    `json:"duration"`
	WatchedTime      int    `json:"watchedTime"`
	ChannelName      string `json:"channelName"`
	ChannelYoutubeId string `json:"channelYoutubeId"`
	ViewsCount       int    `json:"viewsCount"`
	UploadedDate     string `json:"uploadedDate"`
}

type ChannelPlaylist struct {
	YoutubeId      string `json:"youtubeId"`
	Title          string `json:"title"`
	Thumbnail      string `json:"thumbnail"`
	VideoYoutubeId string `json:"videoYoutubeId"`
	VideosCount    int    `json:"videosCount"`
}

type UserPlaylistForVideo struct {
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	YoutubeId string `json:"youtubeId"`
	IsExist   bool   `json:"isExist"`
}

type SearchPlaylistVideos struct {
	Thumbnail string `json:"thumbnail"`
	Title     string `json:"title"`
	Duration  int    `json:"duration"`
}

type SearchResult struct {
	Type             string                 `json:"type"`
	YoutubeId        string                 `json:"youtubeId"`
	Title            string                 `json:"title"`
	Description      string                 `json:"description"`
	Thumbnail        string                 `json:"thumbnail"`
	UploadedDate     string                 `json:"uploadedDate"`
	Duration         int                    `json:"duration"`
	ViewsCount       int                    `json:"viewsCount"`
	WatchedTime      int                    `json:"watchedTime"`
	VideoCount       int                    `json:"videoCount"`
	PlaylistVideos   []SearchPlaylistVideos `json:"playlistVideos"`
	ChannelYoutubeId string                 `json:"uploaderYoutubeId"`
	ChannelName      string                 `json:"channelName"`
	ChannelAvatarUrl string                 `json:"channelAvatarUrl"`
	IsSubscribed     bool                   `json:"isSubscribed"`
	SubscriberCount  int                    `json:"subscriberCount"`
}

type SubscriptionVideo struct {
	YoutubeId        string `json:"youtubeId"`
	Title            string `json:"title"`
	Duration         int    `json:"duration"`
	Thumbnail        string `json:"thumbnail"`
	ViewsCount       int    `json:"viewsCount"`
	Hls              string `json:"hls"`
	WatchedTime      int    `json:"watchedTime"`
	UploadedDate     string `json:"uploadedDate"`
	ChannelName      string `json:"channelName"`
	ChannelAvatarUrl string `json:"channelAvatarUrl"`
	ChannelYoutubeId string `json:"channelYoutubeId"`
}

type SubscribedChannel struct {
	YoutubeId       string `json:"youtubeId"`
	Title           string `json:"title"`
	SubscriberCount int    `json:"subscriberCount"`
	VideoCount      int    `json:"videoCount"`
	Description     string `json:"description"`
	AvatarUrl       string `json:"avatarUrl"`
}

type FeedVideo struct {
	YoutubeId   string `json:"youtubeId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	ChannelName string `json:"channelName"`
	ViewsCount  int    `json:"viewsCount"`
	Duration    int    `json:"duration"`
	WatchedTime int    `json:"watchedTime"`
}

type Genre struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type ReportType struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type Report struct {
	Type         string `json:"type"`
	YoutubeId    string `json:"youtubeId"`
	Title        string `json:"title"`
	ChannelName  string `json:"channelName"`
	ReportReason string `json:"reportReason"`
	Status       string `json:"status"` //moderator reaction
	ReportedDate string `json:"reportedDate"`
}
