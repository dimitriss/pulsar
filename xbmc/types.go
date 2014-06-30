package xbmc

type View struct {
	ContentType string    `json:"content_type"`
	Items       ListItems `json:"items"`
}

type GUIIconOverlay int

const (
	IconOverlayNone GUIIconOverlay = iota
	IconOverlayRAR
	IconOverlayZip
	IconOverlayLocked
	IconOverlayHasTrainer
	IconOverlayTrained
	IconOverlayWatched
	IconOverlayHD
)

type ListItems []*ListItem

type ListItem struct {
	Label       string             `json:"label"`
	Icon        string             `json:"icon"`
	Thumbnail   string             `json:"thumbnail"`
	IsPlayable  bool               `json:"is_playable"`
	Path        string             `json:"path"`
	Info        *ListItemInfo      `json:"info,omitempty"`
	Properties  map[string]string  `json:"properties,omitempty"`
	Art         *ListItemArt       `json:"art,omitempty"`
	StreamInfo  *StreamInfo        `json:"stream_info,omitempty"`
	ContextMenu []*ContextMenuItem `json:"context_menu,omitempty"`
}

type ListItemInfo struct {
	// General Values that apply to all types
	Count string `json:"count,omitempty"`
	Size  int    `json:"size,omitempty"`
	Date  string `json:"date,omitempty"`

	// Video Values
	Genre         string         `json:"genre,omitempty"`
	Year          string         `json:"year,omitempty"`
	Episode       int            `json:"episode,omitempty"`
	Season        int            `json:"season,omitempty"`
	Top250        int            `json:"top250,omitempty"`
	TrackNumber   int            `json:"tracknumber,omitempty"`
	Rating        float32        `json:"rating,omitempty"`
	PlayCount     int            `json:"playcount,omitempty"`
	Overlay       GUIIconOverlay `json:"overlay,omitempty"`
	Cast          []string       `json:"cast,omitempty"`
	CastAndRole   []string       `json:"castandrole,omitempty"`
	Director      string         `json:"director,omitempty"`
	MPAA          string         `json:"mpaa,omitempty"`
	Plot          string         `json:"plot,omitempty"`
	PlotOutline   string         `json:"plotoutline,omitempty"`
	Title         string         `json:"title,omitempty"`
	OriginalTitle string         `json:"originaltitle,omitempty"`
	SortTitle     string         `json:"sorttitle,omitempty"`
	Duration      int            `json:"duration,omitempty"`
	Studio        string         `json:"studio,omitempty"`
	TagLine       string         `json:"tagline,omitempty"`
	Writer        string         `json:"writer,omitempty"`
	TVShowTitle   string         `json:"tvshowtitle,omitempty"`
	Premiered     string         `json:"premiered,omitempty"`
	Status        string         `json:"status,omitempty"`
	Code          string         `json:"code,omitempty"`
	Aired         string         `json:"aired,omitempty"`
	Credits       string         `json:"credits,omitempty"`
	LastPlayed    string         `json:"lastplayed,omitempty"`
	Album         string         `json:"album,omitempty"`
	Artist        []string       `json:"artist,omitempty"`
	Votes         string         `json:"votes,omitempty"`
	Trailer       string         `json:"trailer,omitempty"`
	DateAdded     string         `json:"dateadded,omitempty"`

	// Music Values
	Lyrics string `json:"lyrics,omitempty"`

	// Picture Values
	PicturePath string `json:"picturepath,omitempty"`
	Exif        string `json:"exif,omitempty"`
}

type ListItemArt struct {
	Thumbnail string `json:"thumb,omitempty"`
	Poster    string `json:"poster,omitempty"`
	Banner    string `json:"banner,omitempty"`
	FanArt    string `json:"fanart,omitempty"`
	ClearArt  string `json:"clearart,omitempty"`
	ClearLogo string `json:"clearlogo,omitempty"`
	Landscape string `json:"landscape,omitempty"`
}

type ContextMenuItem struct {
	Label  string
	Action string
}

type StreamInfo struct {
	Video    *StreamInfoEntry `json:"video,omitempty"`
	Audio    *StreamInfoEntry `json:"audio,omitempty"`
	Subtitle *StreamInfoEntry `json:"subtitle,omitempty"`
}

type StreamInfoEntry struct {
	Codec    string  `json:"codec,omitempty"`
	Aspect   float32 `json:"aspect,omitempty"`
	Width    int     `json:"width,omitempty"`
	Height   int     `json:"height,omitempty"`
	Duration int     `json:"duration,omitempty"`
	Language string  `json:"language,omitempty"`
	Channels int     `json:"channels,omitempty"`
}

func NewView(contentType string, items ListItems) *View {
	return &View{
		ContentType: contentType,
		Items:       items,
	}
}