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

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	httpClient "net/http"
	"net/url"
)

func http(endpoint string) (string, error) {
	response, err := httpClient.Get(fmt.Sprintf("https://theaudiodb.com/api/v1/json/2/%s", endpoint))
	if err != nil {
		return "", errors.New("ERROR")
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.New("ERROR")
	}
	return string(body), nil
}

// SearchArtist Return Artist details from artist name
func SearchArtist(s string) (Artist, error) {
	response, err := http(fmt.Sprintf("search.php?s=%s",url.QueryEscape(s)))
	if err != nil {
		return Artist{}, newError("error")
	}
	if len(response) == 0 {
		return Artist{}, newError("error")
	}
	var data baseArtist
	jsonError := json.Unmarshal([]byte(response), &data)
	if jsonError != nil {
		return Artist{}, newError("error")
	}
	if len(data.Artists) == 0 {
		return Artist{}, newError("error")
	}
	return data.Artists[0], nil
}

// Discography Return Discography for an Artist with Album names and year only
func Discography(s string) ([]DiscographyResponse, error) {
	response, err := http(fmt.Sprintf("discography.php?s=%s",url.QueryEscape(s)))
	if err != nil {
		return []DiscographyResponse{}, newError("error")
	}
	if len(response) == 0 {
		return []DiscographyResponse{}, newError("error")
	}
	var data baseDiscography
	jsonError := json.Unmarshal([]byte(response), &data)
	if jsonError != nil {
		return []DiscographyResponse{}, newError("error")
	}
	if len(data.Album) == 0 {
		return []DiscographyResponse{}, newError("error")
	}
	var responseSlice []DiscographyResponse
	for _, v := range data.Album {
		responseSlice= append(responseSlice,v)
	}
	return responseSlice, nil
}

// SearchArtistById Return individual Artist details using known Artist ID
func SearchArtistById(i int64) (Artist, error) {
	response, err := http(fmt.Sprintf("artist.php?i=%d",i))
	if err != nil {
		return Artist{}, newError("error")
	}
	if len(response) == 0 {
		return Artist{}, newError("error")
	}
	var data baseArtist
	jsonError := json.Unmarshal([]byte(response), &data)
	if jsonError != nil {
		return Artist{}, newError("error")
	}
	if len(data.Artists) == 0 {
		return Artist{}, newError("error")
	}
	return data.Artists[0], nil
}

// SearchAlbumById Return individual Album info using known Album ID
func SearchAlbumById(i int64) (Album, error) {
	response, err := http(fmt.Sprintf("album.php?m=%d",i))
	if err != nil {
		return Album{}, newError("error")
	}
	if len(response) == 0 {
		return Album{}, newError("error")
	}
	var data baseAlbum
	jsonError := json.Unmarshal([]byte(response), &data)
	if jsonError != nil {
		return Album{}, newError("error")
	}
	if len(data.Album) == 0 {
		return Album{}, newError("error")
	}
	return data.Album[0], nil
}

// SearchAlbumsByArtistId Return All Albums for an Artist using known Artist ID
func SearchAlbumsByArtistId(i int64) ([]Album, error) {
	response, err := http(fmt.Sprintf("album.php?i=%d",i))
	if err != nil {
		return []Album{}, newError("error")
	}
	if len(response) == 0 {
		return []Album{}, newError("error")
	}
	var data baseAlbum
	jsonError := json.Unmarshal([]byte(response), &data)
	if jsonError != nil {
		return []Album{}, newError("error")
	}
	if len(data.Album) == 0 {
		return []Album{}, newError("error")
	}
	var responseSlice []Album
	for _, v := range data.Album {
		responseSlice= append(responseSlice,v)
	}
	return responseSlice, nil
}

// SearchTracksByAlbumId Return All Tracks for Album from known Album ID
func SearchTracksByAlbumId(i int64) ([]Track, error) {
	response, err := http(fmt.Sprintf("track.php?m=%d", i))
	if err != nil {
		return []Track{}, newError("error")
	}
	if len(response) == 0 {
		return []Track{}, newError("error")
	}
	var data baseTrack
	jsonError := json.Unmarshal([]byte(response), &data)
	if jsonError != nil {
		return []Track{}, newError("error")
	}
	if len(data.Track) == 0 {
		return []Track{}, newError("error")
	}
	var responseSlice []Track
	for _, v := range data.Track {
		responseSlice= append(responseSlice,v)
	}
	return responseSlice, nil
}

// SearchTrackById Return individual track info using a known Track ID
func SearchTrackById(i int64) (Track, error) {
	response, err := http(fmt.Sprintf("track.php?h=%d",i))
	if err != nil {
		return Track{}, newError("error")
	}
	if len(response) == 0 {
		return Track{}, newError("error")
	}
	var data baseTrack
	jsonError := json.Unmarshal([]byte(response), &data)
	if jsonError != nil {
		return Track{}, newError("error")
	}
	if len(data.Track) == 0 {
		return Track{}, newError("error")
	}
	return data.Track[0], nil
}

// SearchMusicVideosByArtistId Return all the Music videos for a known Artist ID
func SearchMusicVideosByArtistId(i int64) ([]MusicVideo, error) {
	response, err := http(fmt.Sprintf("mvid.php?i=%d", i))
	if err != nil {
		return []MusicVideo{}, newError("error")
	}
	if len(response) == 0 {
		return []MusicVideo{}, newError("error")
	}
	var data baseMusicVideo
	jsonError := json.Unmarshal([]byte(response), &data)
	if jsonError != nil {
		return []MusicVideo{}, newError("error")
	}
	if len(data.Mvids) == 0 {
		return []MusicVideo{}, newError("error")
	}
	var responseSlice []MusicVideo
	for _, v := range data.Mvids {
		responseSlice= append(responseSlice,v)
	}
	return responseSlice, nil
}