package one_to_one

func CleanCreateWeeklyReportRequest(req *CreateWeeklyReportRequest) {
	req.GoneWell = FilterEmptyLabels(req.GoneWell, func(g GoneWell) string {
		return g.Label
	})
	req.Challenges = FilterEmptyLabels(req.Challenges, func(c Challenges) string {
		return c.Label
	})
}

func CleanUpdateWeeklyReportRequest(req *UpdateWeeklyReportRequest) {
	req.GoneWell = FilterEmptyLabels(req.GoneWell, func(g GoneWell) string {
		return g.Label
	})
	req.Challenges = FilterEmptyLabels(req.Challenges, func(c Challenges) string {
		return c.Label
	})
}
