package ads

import "sort"

type (
	User struct {
		Country string
		Browser string
	}

	Campaign struct {
		ClickUrl  string
		Price     float64
		Targeting Targeting
	}

	Targeting struct {
		Country string
		Browser string
	}

	filterFunc func(in []*Campaign, u *User) []*Campaign
)

var (
	filters = []filterFunc{
		filterByCountry,
		filterByBrowser,
	}
)

func MakeAuction(in []*Campaign, u *User) (winner *Campaign) {
	campaigns := make([]*Campaign, len(in))
	copy(campaigns, in)

	for _, f := range filters {
		campaigns = f(campaigns, u)
	}

	if len(campaigns) == 0 {
		return nil
	}

	sort.Slice(campaigns, func(i, j int) bool {
		return campaigns[i].Price < campaigns[j].Price
	})

	return campaigns[0]
}

func filterByBrowser(in []*Campaign, u *User) []*Campaign {
	var out []*Campaign
	for _, c := range in {
		if c.Targeting.Browser == "" || c.Targeting.Browser == u.Browser {
			out = append(out, c)
		}
	}
	return out
}

func filterByCountry(in []*Campaign, u *User) []*Campaign {
	var out []*Campaign
	for _, c := range in {
		if c.Targeting.Country == "" || c.Targeting.Country == u.Country {
			out = append(out, c)
		}
	}
	return out
}

func GetStaticCampaigns() []*Campaign {
	return []*Campaign{
		{
			Price: 1,
			Targeting: Targeting{
				Country: "RU",
				Browser: "Chrome",
			},
			ClickUrl: "https://yandex.ru",
		},
		{
			Price: 1,
			Targeting: Targeting{
				Country: "DE",
				Browser: "Chrome",
			},
			ClickUrl: "https://google.com",
		},
		{
			Price: 1,
			Targeting: Targeting{
				Browser: "Firefox",
			},
			ClickUrl: "https://duckduckgo.com",
		},
		{
			Price: 0.9,
			Targeting: Targeting{
				Country: "AO", // Angola
				Browser: "Chrome",
			},
			ClickUrl: "https://zap.co.ao",
		},
		{
			Price: 1.2,
			Targeting: Targeting{
				Country: "BR", // Brasil
				Browser: "Firefox",
			},
			ClickUrl: "https://globo.com",
		},
		{
			Price: 1.1,
			Targeting: Targeting{
				Country: "BR", // Brasil
				Browser: "Chrome",
			},
			ClickUrl: "https://uol.com.br",
		},
	}
}
