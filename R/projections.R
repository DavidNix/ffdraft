# Instructions: https://fantasyfootballanalytics.net/2016/06/ffanalytics-r-package-fantasy-football-data-analysis.html

library("ffanalytics")

pos = c("QB", "RB", "WR", "TE", "K", "DST")
year = as.integer(format(Sys.Date(), "%Y"))
scraped <-scrape_data(pos = pos, season = year, week = 0)

projections = projections_table(scraped)

final = projections %>% add_adp() %>% add_player_info()

# types: average, weighted, robust
final %>% filter(avg_type == "robust") %>% write.csv(file = "projections.csv")
  
