package parser

import (
	"reflect"
	"testing"

	// "../db"
	"github.com/likipiki/goSerialLoader/db"
)

const (
	SAMPLE_TEXT = `<?xml version="1.0" encoding="utf-8" ?>
        <rss version="0.91">
        <channel>
        <title>LostFilm.TV</title>
        <description>Свежачок от LostFilm.TV</description>
        <link>https://www.lostfilm.tv/</link>
        <lastBuildDate>Tue, 07 Aug 2018 08:27:07 +0000</lastBuildDate>
        <language>ru</language><item>
        	<title>Грезы (Reverie). Чёрная мандала (S01E07) [MP4]</title>
        	<category>[MP4]</category>
        	<pubDate>Tue, 07 Aug 2018 08:26:05 +0000</pubDate>
        	<link>http://tracktor.in/rssdownloader.php?id=29285</link>
        </item>
        <item>
        	<title>Грезы (Reverie). Чёрная мандала (S01E07) [1080p]</title>
        	<category>[1080p]</category>
        	<pubDate>Tue, 07 Aug 2018 08:26:05 +0000</pubDate>
        	<link>http://tracktor.in/rssdownloader.php?id=29284</link>
        </item>
        <item>
        	<title>Грезы (Reverie). Чёрная мандала (S01E07) [SD]</title>
        	<category>[SD]</category>
        	<pubDate>Tue, 07 Aug 2018 08:26:05 +0000</pubDate>
        	<link>http://tracktor.in/rssdownloader.php?id=29283</link>
        </item>
        <item>
        	<title>Сирена (Siren). Противостояние (S01E06) [MP4]</title>
        	<category>[MP4]</category>
        	<pubDate>Mon, 06 Aug 2018 08:07:31 +0000</pubDate>
        	<link>http://tracktor.in/rssdownloader.php?id=29282</link>
        </item>
        <item>
        	<title>Сирена (Siren). Противостояние (S01E06) [1080p]</title>
        	<category>[1080p]</category>
        	<pubDate>Mon, 06 Aug 2018 08:07:31 +0000</pubDate>
        	<link>http://tracktor.in/rssdownloader.php?id=29281</link>
        </item>
        <item>
        	<title>Сирена (Siren). Противостояние (S01E06) [SD]</title>
        	<category>[SD]</category>
        	<pubDate>Mon, 06 Aug 2018 08:07:31 +0000</pubDate>
        	<link>http://tracktor.in/rssdownloader.php?id=29280</link>
        </item>
        <item>
        	<title>Киллджойс (Killjoys). Не по пути (S04E03) [MP4]</title>
        	<category>[MP4]</category>
        	<pubDate>Sun, 05 Aug 2018 17:22:17 +0000</pubDate>
        	<link>http://tracktor.in/rssdownloader.php?id=29279</link>
        </item>
        <item>
        	<title>Киллджойс (Killjoys). Не по пути (S04E03) [1080p]</title>
        	<category>[1080p]</category>
        	<pubDate>Sun, 05 Aug 2018 17:22:17 +0000</pubDate>
        	<link>http://tracktor.in/rssdownloader.php?id=29278</link>
        </item>
        <item>
        	<title>Киллджойс (Killjoys). Не по пути (S04E03) [SD]</title>
        	<category>[SD]</category>
        	<pubDate>Sun, 05 Aug 2018 17:22:17 +0000</pubDate>
        	<link>http://tracktor.in/rssdownloader.php?id=29277</link>
        </item>
        <item>
        	<title>Королева юга (Queen of the South). Королева мечей (S03E07) [MP4]</title>
        	<category>[MP4]</category>
        	<pubDate>Sun, 05 Aug 2018 16:46:43 +0000</pubDate>
        	<link>http://tracktor.in/rssdownloader.php?id=29276</link>
        </item>
        <item>
        	<title>Королева юга (Queen of the South). Королева мечей (S03E07) [1080p]</title>
        	<category>[1080p]</category>
        	<pubDate>Sun, 05 Aug 2018 16:46:43 +0000</pubDate>
        	<link>http://tracktor.in/rssdownloader.php?id=29275</link>
        </item>
        <item>
        	<title>Королева юга (Queen of the South). Королева мечей (S03E07) [SD]</title>
        	<category>[SD]</category>
        	<pubDate>Sun, 05 Aug 2018 16:46:43 +0000</pubDate>
        	<link>http://tracktor.in/rssdownloader.php?id=29274</link>
        </item>
        <item>
        	<title>Пастырь (Preacher). В гробу (S03E05) [MP4]</title>
        	<category>[MP4]</category>
        	<pubDate>Wed, 25 Jul 2018 17:43:09 +0000</pubDate>
        	<link>http://tracktor.in/rssdownloader.php?id=29273</link>
        </item>
        <item>
        	<title>Пастырь (Preacher). В гробу (S03E05) [1080p]</title>
        	<category>[1080p]</category>
        	<pubDate>Wed, 25 Jul 2018 17:43:09 +0000</pubDate>
        	<link>http://tracktor.in/rssdownloader.php?id=29272</link>
        </item>
        <item>
        	<title>Пастырь (Preacher). В гробу (S03E05) [SD]</title>
        	<category>[SD]</category>
        	<pubDate>Wed, 25 Jul 2018 17:43:09 +0000</pubDate>
        	<link>http://tracktor.in/rssdownloader.php?id=29271</link>
        </item>
        </channel>
        </rss>`
)

var (
	SAMPLE_RESULT = []Serial{
		Serial{
			Serial: db.Serial{
				Name:    "Reverie",
				Season:  1,
				Episode: 7,
			},
			SeasonData: "S01E07",
			Resolutions: []Resolution{
				Resolution{
					Link:   "http://tracktor.in/rssdownloader.php?id=29285",
					Format: "MP4",
				},
				Resolution{
					Link:   "http://tracktor.in/rssdownloader.php?id=29284",
					Format: "1080p",
				},
				Resolution{
					Link:   "http://tracktor.in/rssdownloader.php?id=29283",
					Format: "SD",
				},
			},
		},
		Serial{
			Serial: db.Serial{
				Name:    "Siren",
				Season:  1,
				Episode: 6,
			},
			SeasonData: "S01E06",
			Resolutions: []Resolution{
				Resolution{
					Link:   "http://tracktor.in/rssdownloader.php?id=29282",
					Format: "MP4",
				},
				Resolution{
					Link:   "http://tracktor.in/rssdownloader.php?id=29281",
					Format: "1080p",
				},
				Resolution{
					Link:   "http://tracktor.in/rssdownloader.php?id=29280",
					Format: "SD",
				},
			},
		},
		Serial{
			Serial: db.Serial{
				Name:    "Killjoys",
				Season:  4,
				Episode: 3,
			},
			SeasonData: "S04E03",
			Resolutions: []Resolution{
				Resolution{
					Link:   "http://tracktor.in/rssdownloader.php?id=29279",
					Format: "MP4",
				},
				Resolution{
					Link:   "http://tracktor.in/rssdownloader.php?id=29278",
					Format: "1080p",
				},
				Resolution{
					Link:   "http://tracktor.in/rssdownloader.php?id=29277",
					Format: "SD",
				},
			},
		},
		Serial{
			Serial: db.Serial{
				Name:    "Queen of the South",
				Season:  3,
				Episode: 7,
			},
			SeasonData: "S03E07",
			Resolutions: []Resolution{
				Resolution{
					Link:   "http://tracktor.in/rssdownloader.php?id=29276",
					Format: "MP4",
				},
				Resolution{
					Link:   "http://tracktor.in/rssdownloader.php?id=29275",
					Format: "1080p",
				},
				Resolution{
					Link:   "http://tracktor.in/rssdownloader.php?id=29274",
					Format: "SD",
				},
			},
		},
		Serial{
			Serial: db.Serial{
				Name:    "Preacher",
				Season:  3,
				Episode: 5,
			},
			SeasonData: "S03E05",
			Resolutions: []Resolution{
				Resolution{
					Link:   "http://tracktor.in/rssdownloader.php?id=29273",
					Format: "MP4",
				},
				Resolution{
					Link:   "http://tracktor.in/rssdownloader.php?id=29272",
					Format: "1080p",
				},
				Resolution{
					Link:   "http://tracktor.in/rssdownloader.php?id=29271",
					Format: "SD",
				},
			},
		},
	}
)

func TestParse(t *testing.T) {

	serials, err := Parse(SAMPLE_TEXT)
	if err != nil {
		t.Fatal("parse not completed")
	}
	if !reflect.DeepEqual(serials, SAMPLE_RESULT) {
		t.Error("expected", SAMPLE_RESULT, "got", serials)
	}

}
