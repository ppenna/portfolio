Portfolio Assistant
===================

This repository hosts the source code of a portfolio assistant that I
wrote for helping me on my investment decisions. It relies on a [genetic
algorithm](https://en.wikipedia.org/wiki/Genetic_algorithm) and is
narrowed for [Brazilian
REITs](https://pt.wikipedia.org/wiki/Fundo_de_Investimento_Imobili%C3%A1rio).


*DISCLAIMER: you should not rely on the output of this program to guide
your investment decisions and strategies. This program should be used
for studying purposes only. Use this program at your own risk.*


Built-In Asset Evaluation Methodology
-------------------------------------

By default, this portfolio assistant relies on the following equation to
evaluate a portfolio `x`:

```
val(x) = perf(x) + risk(x) + cost(x)

```
Where:
- `perf()` expresses the performance of the portfolio
- `risk()` quantifies the risk of the portfolio
- `cost()` gives the cost of the portfolio

To compute these three components, the following metrics are considered:

- Book Value Price per Share (BVPS)
- Dividend Yield (DY)
- Equity Value
- Gross Leasable Area (GLA)
- Share Price (P)

Usage
------

```
assistant [options]

Options:

  -output string  Name of the wallet file (default "new.wallet")
  -print          Print wallet?
  -save           Save wallet to a file?
  -stats          Print statistics? (default true)
```
