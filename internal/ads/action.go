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
	for i := len(in) - 1; i >= 0; i-- {
		if len(in[i].Targeting.Browser) == 0 {
			// Empty browser means no restrictions, so don't remove the campaign
			continue
		}

		if in[i].Targeting.Browser == u.Browser {
			// Browser matches – campaign passes
			continue
		}

		// At this point we know there's a browser filter, and the user's browser doesn't match it
		// Remove the campaign
		in[i] = in[0]
		in = in[1:]
	}

	return in
}

func filterByCountry(in []*Campaign, u *User) []*Campaign {
	for i := len(in) - 1; i >= 0; i-- {
		if len(in[i].Targeting.Country) == 0 {
			// Empty country means no restrictions, so don't remove the campaign
			continue
		}

		if in[i].Targeting.Country == u.Country {
			// Country matches – campaign passes
			continue
		}

		// At this point we know there's a country filter, and the user's country doesn't match it
		// Remove the campaign
		in[i] = in[0]
		in = in[1:]
	}

	return in
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
	}
}
