# TheAudioDB

[![](https://img.shields.io/github/v/tag/thechampagne/audiodb-go?label=version)](https://github.com/thechampagne/audiodb-go/releases/latest) [![](https://img.shields.io/github/license/thechampagne/audiodb-go)](https://github.com/thechampagne/audiodb-go/blob/main/LICENSE)

TheAudioDB API client for **Go**.

### Download

```
go get github.com/thechampagne/audiodb-go
```

### Example

```go
func main() {
	albums, err := audiodb.SearchAlbumsByArtistId(111674)
	if err != nil {
		log.Fatal(err)
	}
	for _,i := range albums {
		fmt.Println(i.StrAlbum)
	}
}
```

### License

TheAudioDB is released under the [Apache License 2.0](https://github.com/thechampagne/audiodb-go/blob/main/LICENSE).

```
 Copyright 2022 XXIV

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
```