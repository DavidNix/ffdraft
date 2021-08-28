
# Instructions: https://fantasyfootballanalytics.net/2016/06/ffanalytics-r-package-fantasy-football-data-analysis.html

library("ffanalytics")

year = format(Sys.Date(), "%Y")
print("Season")
print(year)

pos = c("QB", "RB", "WR", "TE", "K", "DST")
scraped = scrape_data(pos = pos, season = as.integer(year), week = 0)

projections = projections_table(scraped)

final = projections %>% add_adp() %>% add_player_info() %>% add_ecr() %>% add_risk()

# types: average, weighted, robust
final %>% filter(avg_type == "weighted") %>% write.csv(file = "projections.csv", row.names=FALSE)
  