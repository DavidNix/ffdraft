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

src = c(
  "FantasySharks",
  "NumberFire",
  "NFL",
  "WalterFootball",
  "CBS", 
  "ESPN", 
  "FantasyData", 
  "FantasyPros", 
  "FFToday",
  "FleaFlicker", 
  "Yahoo", 
  "FantasyFootballNerd", 
  "RTSports")

# Temporarily remove sources because they are causing problems
if (week > 0) {
  print("WARNING REMOVING SOURCES")
  src = src[! src %in% c(
    "FantasySharks",
    "NumberFire",
    "NFL",
    "WalterFootball"
    )]   
}

pos = c("QB", "RB", "WR", "TE", "K", "DST")
scraped = scrape_data(src = src, pos = pos, season = as.integer(year), week = week)

projections = NULL
filename = NULL

if (ppr > 0) {
  ppr_scoring = custom_scoring(pass_yds = 0.04, pass_tds = 4,
                               rush_yds = 0.1, rush_tds = 6,
                               rec = 0.5, rec_yds = 0.1, rec_tds = 6)
  print("Calculating PPR projections")
  filename = "ppr_projections.csv"
  projections = projections_table(scraped, scoring_rules = ppr_scoring)
  
} else {
  print("Calculating standard projections")
  filename = "standard_projections.csv"
  projections = projections_table(scraped)
}

final = projections %>% add_adp() %>% add_player_info() %>% add_ecr() %>% add_risk()
fname = paste("week", week, "_", filename, sep="")
# avg_types: average, weighted, robust
sprintf("file saved as %s", fname)
final %>% filter(avg_type == "weighted") %>% write.csv(file = fname, row.names=FALSE)
  
if (ppr > 0) {
  warning("WARNING PPR HARDCODED TO 0.5")
}