/*
 * Copyright 2022 XXIV
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package audiodb

type baseArtist struct {
	Artists []Artist `json:"artists"`
}

type Artist struct {
	IDArtist           string      `json:"idArtist"`
	IDLabel            string      `json:"idLabel"`
	IntBornYear        string      `json:"intBornYear"`
	IntCharted         string      `json:"intCharted"`
	IntDiedYear        string      `json:"intDiedYear"`
	IntFormedYear      string      `json:"intFormedYear"`
	IntMembers         string      `json:"intMembers"`
	StrArtist          string      `json:"strArtist"`
	StrArtistAlternate string      `json:"strArtistAlternate"`
	StrArtistBanner    string      `json:"strArtistBanner"`
	StrArtistClearart  string      `json:"strArtistClearart"`
	StrArtistCutout    string      `json:"strArtistCutout"`
	StrArtistFanart    string      `json:"strArtistFanart"`
	StrArtistFanart2   string      `json:"strArtistFanart2"`
	StrArtistFanart3   string      `json:"strArtistFanart3"`
	StrArtistFanart4   string      `json:"strArtistFanart4"`
	StrArtistLogo      string      `json:"strArtistLogo"`
	StrArtistStripped  string      `json:"strArtistStripped"`
	StrArtistThumb     string      `json:"strArtistThumb"`
	StrArtistWideThumb string      `json:"strArtistWideThumb"`
	StrBiographyCN     string      `json:"strBiographyCN"`
	StrBiographyDE     string      `json:"strBiographyDE"`
	StrBiographyEN     string      `json:"strBiographyEN"`
	StrBiographyES     string      `json:"strBiographyES"`
	StrBiographyFR     string      `json:"strBiographyFR"`
	StrBiographyHU     string      `json:"strBiographyHU"`
	StrBiographyIL     string      `json:"strBiographyIL"`
	StrBiographyIT     string      `json:"strBiographyIT"`
	StrBiographyJP     string      `json:"strBiographyJP"`
	StrBiographyNL     string      `json:"strBiographyNL"`
	StrBiographyNO     string      `json:"strBiographyNO"`
	StrBiographyPL     string      `json:"strBiographyPL"`
	StrBiographyPT     string      `json:"strBiographyPT"`
	StrBiographyRU     string      `json:"strBiographyRU"`
	StrBiographySE     string      `json:"strBiographySE"`
	StrCountry         string      `json:"strCountry"`
	StrCountryCode     string      `json:"strCountryCode"`
	StrDisbanded       string      `json:"strDisbanded"`
	StrFacebook        string      `json:"strFacebook"`
	StrGender          string      `json:"strGender"`
	StrGenre           string      `json:"strGenre"`
	StrISNIcode        string      `json:"strISNIcode"`
	StrLabel           string      `json:"strLabel"`
	StrLastFMChart     string      `json:"strLastFMChart"`
	StrLocked          string      `json:"strLocked"`
	StrMood            string      `json:"strMood"`
	StrMusicBrainzID   string      `json:"strMusicBrainzID"`
	StrStyle           string      `json:"strStyle"`
	StrTwitter         string      `json:"strTwitter"`
	StrWebsite         string      `json:"strWebsite"`
}

type baseDiscography struct {
	Album []DiscographyResponse `json:"album"`
}

type DiscographyResponse struct {
	IntYearReleased string `json:"intYearReleased"`
	StrAlbum        string `json:"strAlbum"`
}

type baseAlbum struct {
	Album []Album `json:"album"`
}

type Album struct {
	IDAlbum                string      `json:"idAlbum"`
	IDArtist               string      `json:"idArtist"`
	IDLabel                string      `json:"idLabel"`
	IntLoved               string      `json:"intLoved"`
	IntSales               string      `json:"intSales"`
	IntScore               string      `json:"intScore"`
	IntScoreVotes          string      `json:"intScoreVotes"`
	IntYearReleased        string      `json:"intYearReleased"`
	StrAlbum               string      `json:"strAlbum"`
	StrAlbum3DCase         string      `json:"strAlbum3DCase"`
	StrAlbum3DFace         string      `json:"strAlbum3DFace"`
	StrAlbum3DFlat         string      `json:"strAlbum3DFlat"`
	StrAlbum3DThumb        string      `json:"strAlbum3DThumb"`
	StrAlbumCDart          string      `json:"strAlbumCDart"`
	StrAlbumSpine          string      `json:"strAlbumSpine"`
	StrAlbumStripped       string      `json:"strAlbumStripped"`
	StrAlbumThumb          string      `json:"strAlbumThumb"`
	StrAlbumThumbBack      string      `json:"strAlbumThumbBack"`
	StrAlbumThumbHQ        string      `json:"strAlbumThumbHQ"`
	StrAllMusicID          string      `json:"strAllMusicID"`
	StrAmazonID            string      `json:"strAmazonID"`
	StrArtist              string      `json:"strArtist"`
	StrArtistStripped      string      `json:"strArtistStripped"`
	StrBBCReviewID         string      `json:"strBBCReviewID"`
	StrDescription         string      `json:"strDescription"`
	StrDescriptionCN       string      `json:"strDescriptionCN"`
	StrDescriptionDE       string      `json:"strDescriptionDE"`
	StrDescriptionEN       string      `json:"strDescriptionEN"`
	StrDescriptionES       string      `json:"strDescriptionES"`
	StrDescriptionFR       string      `json:"strDescriptionFR"`
	StrDescriptionHU       string      `json:"strDescriptionHU"`
	StrDescriptionIL       string      `json:"strDescriptionIL"`
	StrDescriptionIT       string      `json:"strDescriptionIT"`
	StrDescriptionJP       string      `json:"strDescriptionJP"`
	StrDescriptionNL       string      `json:"strDescriptionNL"`
	StrDescriptionNO       string      `json:"strDescriptionNO"`
	StrDescriptionPL       string      `json:"strDescriptionPL"`
	StrDescriptionPT       string      `json:"strDescriptionPT"`
	StrDescriptionRU       string      `json:"strDescriptionRU"`
	StrDescriptionSE       string      `json:"strDescriptionSE"`
	StrDiscogsID           string      `json:"strDiscogsID"`
	StrGeniusID            string      `json:"strGeniusID"`
	StrGenre               string      `json:"strGenre"`
	StrItunesID            string      `json:"strItunesID"`
	StrLabel               string      `json:"strLabel"`
	StrLocation            string      `json:"strLocation"`
	StrLocked              string      `json:"strLocked"`
	StrLyricWikiID         string      `json:"strLyricWikiID"`
	StrMood                string      `json:"strMood"`
	StrMusicBrainzArtistID string      `json:"strMusicBrainzArtistID"`
	StrMusicBrainzID       string      `json:"strMusicBrainzID"`
	StrMusicMozID          string      `json:"strMusicMozID"`
	StrRateYourMusicID     string      `json:"strRateYourMusicID"`
	StrReleaseFormat       string      `json:"strReleaseFormat"`
	StrReview              string      `json:"strReview"`
	StrSpeed               string      `json:"strSpeed"`
	StrStyle               string      `json:"strStyle"`
	StrTheme               string      `json:"strTheme"`
	StrWikidataID          string      `json:"strWikidataID"`
	StrWikipediaID         string      `json:"strWikipediaID"`
}

type baseTrack struct {
	Track []Track `json:"track"`
}

type Track struct {
	IDAlbum                string      `json:"idAlbum"`
	IDArtist               string      `json:"idArtist"`
	IDIMVDB                string      `json:"idIMVDB"`
	IDLyric                string      `json:"idLyric"`
	IDTrack                string      `json:"idTrack"`
	IntCD                  string      `json:"intCD"`
	IntDuration            string      `json:"intDuration"`
	IntLoved               string      `json:"intLoved"`
	IntMusicVidComments    string      `json:"intMusicVidComments"`
	IntMusicVidDislikes    string      `json:"intMusicVidDislikes"`
	IntMusicVidFavorites   string      `json:"intMusicVidFavorites"`
	IntMusicVidLikes       string      `json:"intMusicVidLikes"`
	IntMusicVidViews       string      `json:"intMusicVidViews"`
	IntScore               string      `json:"intScore"`
	IntScoreVotes          string      `json:"intScoreVotes"`
	IntTotalListeners      string      `json:"intTotalListeners"`
	IntTotalPlays          string      `json:"intTotalPlays"`
	IntTrackNumber         string      `json:"intTrackNumber"`
	StrAlbum               string      `json:"strAlbum"`
	StrArtist              string      `json:"strArtist"`
	StrArtistAlternate     string      `json:"strArtistAlternate"`
	StrDescriptionCN       string      `json:"strDescriptionCN"`
	StrDescriptionDE       string      `json:"strDescriptionDE"`
	StrDescriptionEN       string      `json:"strDescriptionEN"`
	StrDescriptionES       string      `json:"strDescriptionES"`
	StrDescriptionFR       string      `json:"strDescriptionFR"`
	StrDescriptionHU       string      `json:"strDescriptionHU"`
	StrDescriptionIL       string      `json:"strDescriptionIL"`
	StrDescriptionIT       string      `json:"strDescriptionIT"`
	StrDescriptionJP       string      `json:"strDescriptionJP"`
	StrDescriptionNL       string      `json:"strDescriptionNL"`
	StrDescriptionNO       string      `json:"strDescriptionNO"`
	StrDescriptionPL       string      `json:"strDescriptionPL"`
	StrDescriptionPT       string      `json:"strDescriptionPT"`
	StrDescriptionRU       string      `json:"strDescriptionRU"`
	StrDescriptionSE       string      `json:"strDescriptionSE"`
	StrGenre               string      `json:"strGenre"`
	StrLocked              string      `json:"strLocked"`
	StrMood                string      `json:"strMood"`
	StrMusicBrainzAlbumID  string      `json:"strMusicBrainzAlbumID"`
	StrMusicBrainzArtistID string      `json:"strMusicBrainzArtistID"`
	StrMusicBrainzID       string      `json:"strMusicBrainzID"`
	StrMusicVid            string      `json:"strMusicVid"`
	StrMusicVidCompany     string      `json:"strMusicVidCompany"`
	StrMusicVidDirector    string      `json:"strMusicVidDirector"`
	StrMusicVidScreen1     string      `json:"strMusicVidScreen1"`
	StrMusicVidScreen2     string      `json:"strMusicVidScreen2"`
	StrMusicVidScreen3     string      `json:"strMusicVidScreen3"`
	StrStyle               string      `json:"strStyle"`
	StrTheme               string      `json:"strTheme"`
	StrTrack               string      `json:"strTrack"`
	StrTrack3DCase         string      `json:"strTrack3DCase"`
	StrTrackLyrics         string      `json:"strTrackLyrics"`
	StrTrackThumb          string      `json:"strTrackThumb"`
}

type baseMusicVideo struct {
	Mvids []MusicVideo `json:"mvids"`
}

type MusicVideo struct {
	IDAlbum          string `json:"idAlbum"`
	IDArtist         string `json:"idArtist"`
	IDTrack          string `json:"idTrack"`
	StrDescriptionEN string `json:"strDescriptionEN"`
	StrMusicVid      string `json:"strMusicVid"`
	StrTrack         string `json:"strTrack"`
	StrTrackThumb    string `json:"strTrackThumb"`
}