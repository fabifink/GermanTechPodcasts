package cmd

import (
	"time"
)

type PodcastInformation struct {
	Name                   string   `yaml:"name" json:"name"`
	Slug                   string   `json:"slug"`
	Website                string   `yaml:"website" json:"website"`
	PodcastIndexID         int64    `yaml:"podcastIndexID" json:"podcastIndexID"`
	RSSFeed                string   `yaml:"rssFeed" json:"rssFeed"` // TODO Should be better a url.URL
	Spotify                string   `yaml:"spotify" json:"spotify"` // TODO Should be better a url.URL
	Description            string   `yaml:"description" json:"description"`
	EpisodeCount           int      `json:"episodeCount"`
	LatestEpisodePublished int64    `json:"latestEpisodePublished"`
	ItunesID               int64    `json:"itunesID"`
	Categories             []string `json:"categories"`
	Image                  string   `json:"image"`
}

func (p PodcastInformation) GetHumanReadableDate() string {
	t := time.Unix(p.LatestEpisodePublished, 0)
	s := t.Format("Monday, 02 January 2006")

	return s
}