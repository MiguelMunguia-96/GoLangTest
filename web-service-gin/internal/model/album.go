package model

type Album struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	MetaReady   bool   `json:"metaReady"`
	BackupReady bool   `json:"backupReady"`
	Mp3Ready    bool   `json:"mp3Ready"`
}

func PrintAlbum(a Album) string {
	return "Title : " + a.Title + "\n" +
		"Artist : " + a.Artist + "\n"

}
