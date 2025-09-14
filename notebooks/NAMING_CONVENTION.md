# Notebook Naming Convention

To maintain a clean and scalable project structure, all Jupyter notebooks follow a thematic prefix-based naming convention. The format is `prefix_descriptive-name.ipynb`.

## **`setup_`**

- **Theme:** Environment & Connection Setup
- **Purpose:** Notebooks for initial project configuration, dependency checks, and testing API connections (e.g., MetaTrader 5, data providers). These are typically run once at the beginning of the project or on a new machine.
- **Examples:** `setup_environment_check.ipynb`, `setup_mt5_connection.ipynb`

## **`data_`**

- **Theme:** Data Acquisition & Persistence
- **Purpose:** Scripts dedicated to downloading, cleaning, and saving historical data to a persistent format (e.g., Parquet, CSV). The primary output of these notebooks is a clean, reusable data file.
- **Examples:** `data_download_mt5_to_parquet.ipynb`, `data_yf_download_sp500.ipynb`

## **`explore_`**

- **Theme:** Exploratory Data Analysis (EDA)
- **Purpose:** Notebooks for initial, high-level exploration of data and the trading environment. This includes analyzing account details, visualizing the basic properties of downloaded datasets, or getting a first look at an indicator's behavior.
- **Examples:** `explore_account_info.ipynb`, `explore_eurusd_data_properties.ipynb`, `explore_supertrend_visuals.ipynb`

## **`feat_`**

- **Theme:** Feature Engineering
- **Purpose:** Focused, in-depth numerical analysis of individual technical indicators or custom-derived features. The goal is to understand the statistical properties of a feature before incorporating it into a strategy.
- **Examples:** `feat_rsi_distribution_analysis.ipynb`, `feat_bollinger_bands_volatility.ipynb`

## **`strat_`**

- **Theme:** Strategy Implementation & Backtesting
- **Purpose:** Implementation and backtesting of a complete trading logic (entry, exit, SL/TP) using a _fixed_ set of parameters. This step is for validating the core concept of a strategy.
- **Examples:** `strat_sma_cross_backtest.ipynb`, `strat_backtestingpy_sma_comparison.ipynb`

## **`optim_`**

- **Theme:** Strategy Optimization
- **Purpose:** Fine-tuning the parameters of a validated strategy to find the most robust and profitable combinations. This is where we leverage `vectorbt`'s grid search or `backtesting.py`'s optimization engine.
- **Examples:** `optim_sma_grid_search_vectorbt.ipynb`, `optim_rsi_walk_forward.ipynb`

## **`analysis_`**

- **Theme:** Results Analysis & Reporting
- **Purpose:** In-depth analysis, comparison, and visualization of backtesting and optimization results. This is where final conclusions about a strategy's performance are drawn.
- **Examples:** `analysis_compare_sma_vs_ema.ipynb`, `analysis_optimization_heatmap.ipynb`

## **`sandbox_`**

- **Theme:** Sandbox & Experiments
- **Purpose:** A catch-all category for quick experiments, prototyping new ideas, or testing third-party libraries that don't yet fit into the structured workflow. This helps keep the other directories clean.
- **Examples:** `sandbox_quick_plot_test.ipynb`, `sandbox_new_library_check.ipynb`
