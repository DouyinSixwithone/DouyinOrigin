package publish

import (
	"Douyin/config"
	"Douyin/middleware/ffmpeg"
	"Douyin/repository"
	"fmt"
)

func GetVideoCover(path string, name string, playSuffix string) (string, error) {
	coverSuffix := ".jpg"
	videoPath := path + name + playSuffix
	outputPath := path + name + coverSuffix
	err := ffmpeg.Do(videoPath, outputPath)
	if err != nil {
		return "", err
	}
	return coverSuffix, nil
}

func UploadVideo(authorId uint, title string, playName string, coverName string) error {
	ip := config.Conf.Ip
	urlPrefix := "http://" + ip + "/data/" + fmt.Sprintf("%v/", authorId)
	newVideo := repository.Video{
		AuthorId: authorId,
		PlayUrl:  urlPrefix + playName,
		CoverUrl: urlPrefix + coverName,
		Title:    title,
	}
	err := repository.InsertNewVideo(&newVideo)
	if err != nil {
		return err
	}
	return nil
}
