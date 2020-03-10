package thumbnail

//Image Reading an image from INFILE,and write its thumbnail to the same directory
func ImageFile(infile string) (string err)

func makeThumbnails(filenames string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		/*
			if _, err := ImageFile(f); err != nil {
				log.Println(err)
			}
		*/
		go ImageFile(f)
	}
}
