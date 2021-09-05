# Instructions: https://fantasyfootballanalytics.net/2016/06/ffanalytics-r-package-fantasy-football-data-analysis.html

library("ffanalytics")
library("optparse")

parser = OptionParser()
parser <- add_option(parser, c("--ppr"), type="double", default=0,
                     help="PPR point value [default %default]. If 0 uses standard scoring.")
parser <- add_option(parser, c("--week"), type="integer", default=0,
                     help="Week of season. Use 0 for the entire season [default %default].",
                     metavar="number")
opt = parse_args(parser)

year = format(Sys.Date(), "%Y")
ppr = opt$ppr
week = opt$week 
sprintf("Getting projections for %s Week %d with PPR %s", year, week, ppr)

pos = c("QB", "RB", "WR", "TE", "K", "DST")
scraped = scrape_data(pos = pos, season = as.integer(year), week = week)


save_projections <- function(projections, filename) {
  final = projections %>% add_adp() %>% add_player_info() %>% add_ecr() %>% add_risk()
  # avg_types: average, weighted, robust
  final %>% filter(avg_type == "weighted") %>% write.csv(file = filename, row.names=FALSE)
}

if (ppr > 0) {
  ppr_scoring = custom_scoring(pass_yds = 0.04, pass_tds = 4,
                               rush_yds = 0.1, rush_tds = 6,
                               rec = ppr, rec_yds = 0.1, rec_tds = 6)
  print("Calculating PPR projections")
  projections = projections_table(scraped, scoring_rules = ppr_scoring)
  save_projections(projections, "ppr_projections.csv")
} else {
  print("Calculating standard projections")
  projections = projections_table(scraped)
  save_projections(projections, "standard_projections.csv")
}
