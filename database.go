package ghost

type (
	// Export holds the entire export data structure.
	Export struct {
		DB []Database `json:"db"`
	}

	// Database holds the meta/data for a publication.
	Database struct {
		Meta DatabaseMeta `json:"meta"`
		Data DatabaseData `json:"data"`
	}

	// DatabaseMeta holds metadata about the exported data in the Database.
	DatabaseMeta struct {
		ExportedOn int64  `json:"exported_on"`
		Version    string `json:"version"`
	}

	// DatabaseData holds posts and other data that make up a publication.
	DatabaseData struct {
		Posts []Post `json:"posts"`
	}
)
