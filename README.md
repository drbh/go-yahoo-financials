# Welcome to go-yahoo-financials

![Screenshot](gyf.png?raw=true "Application Screenshot")

## What I do

This is a simple application that does a few things.

1. Download quarterly statement data for +3600 companies from `https://finance.yahoo.com/quote/SYMBOL/financials?p=SYMBOL`
2. Download technical data from `https://query1.finance.yahoo.com/v7/finance/download/`
3. Calculate ReturnOnAssets, ReturnOnEquity, ProfitMargin, QuickRatio, CurrentRatio, DebtToEquity for each company
4. Calculate RelativeStrengthIndex and MoneyFlowIndex for each company
5. Calculate 50th percentile for each ratio per sector
6. Check if Ratios are above 50th percentile and that RSI and MFI are sufficently low
7. Return Companies that meet Evaluation criteria

## How to use

1. Click Statements to download all +3600 company quarterlys - wait about an hour
2. Click Technicals to download all +3600 company pricing data - wait about an hour
3. Click Evaluate to compare the data to the evaluation criteria

## When to run

*Technicals*
You must run the Technical data after 4PM EST in order to have correct data for the next day. This is when the data is in for the current day in the market.

*Statements*
This data get cached in the BoltDB and doesnt download unless the whole BoltDB is deleted. This will remove all the pricing data as well. (need to come up with a better way to remove statements from the cache individually)