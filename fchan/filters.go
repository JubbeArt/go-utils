package fchan

func Images(posts []Post) []string {
	urls := make([]string, 0)

	for _, post := range posts {
		if post.HasImage() {
			urls = append(urls, post.ImageUrl())
		}
	}

	return urls
}
