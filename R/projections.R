
# Instructions: https://fantasyfootballanalytics.net/2016/06/ffanalytics-r-package-fantasy-football-data-analysis.html

library("ffanalytics")

year = format(Sys.Date(), "%Y")
print("Season")
print(year)

pos = c("QB", "RB", "WR", "TE", "K", "DST")
scraped = scrape_data(pos = pos, season = as.integer(year), week = 0)

# 0.5 PPR, leave default for standard
ppr_scoring = custom_scoring(pass_yds = 0.04, pass_tds = 4,
                             rush_yds = 0.1, rush_tds = 6,
                             rec = 0.5, rec_yds = 0.1, rec_tds = 6)

projections = projections_table(scraped, scoring_rules = ppr_scoring)

final = projections %>% add_adp() %>% add_player_info() %>% add_ecr() %>% add_risk()

# types: average, weighted, robust
final %>% filter(avg_type == "weighted") %>% write.csv(file = "projections.csv", row.names=FALSE)
  