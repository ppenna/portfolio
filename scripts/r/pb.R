#
# Copyright(C) 2020 Pedro Henrique Penna <pedrohenriquepenna@gmail.com>
#
# All rights reserved.
#

library("ggplot2")

source("scripts/r/rplots/myPlots.R")
source("scripts/r/rplots/histogram.R")
source("scripts/r/paths.R")

#===============================================================================

args = commandArgs(trailingOnly=TRUE)

asset.ticker <- args[1]
plot.axis.y.max <- strtoi(args[2])

filename <- paste(
	asset.ticker,
	"csv",
	sep = "."
)

asset.df <- read.csv(
	file = paste(
		paths.data.dir,
		filename,
		sep = "/"
	),
	header = FALSE
)
names(asset.df) <- c(
	"date",
	"price",
	"bvps",
	"market.cap",
	"equity",
	"dividends",
	"ffo",
	"num.shares",
	"default",
	"gla",
	"num.share.holders"
)

asset.df$pb <- asset.df$price/asset.df$bvps

#===============================================================================

var <- "pb"
respvar <- "density"
plot.title <- paste(toupper(asset.ticker), "- Price-to-Book Ratio (P/B)", sep = " ")
plot.axis.x.title <- "P/B Ratio"
plot.axis.y.title <- "Density (%)"
plot.axis.y.limits <- c(0, plot.axis.y.max)
plot.axis.x.min <- 0.50
plot.axis.x.max <- 2.00

#===============================================================================

p <- plot.histogram(
		factor = var, respvar = respvar,
		title = plot.title,
		axis.x.title = plot.axis.x.title,
		axis.y.title = plot.axis.y.title,
		axis.y.limits = plot.axis.y.limits,
	) +
	xlim(plot.axis.x.min, plot.axis.x.max) +
	plot.theme.title +
	plot.theme.axis.x +
	plot.theme.axis.y +
	plot.theme.grid.major +
	plot.theme.grid.minor +
	plot.theme.grid.wall

plot.save(
	directory = paths.output.dir,
	plot = p,
	filename = asset.ticker
)
