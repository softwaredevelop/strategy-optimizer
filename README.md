# strategy-optimizer: Python Trading Strategy Research

This repository is a quantitative lab for developing, backtesting, and optimizing trading strategies in Python. The primary goal is to leverage Python's powerful data science ecosystem to create robust strategies intended for final implementation as Expert Advisors (EAs) on the MetaTrader 5 (MQL5) platform.

## Notebook Naming Convention

To maintain a clean and scalable project structure, all Jupyter notebooks follow a thematic prefix-based naming convention. The format is `prefix_descriptive-name.ipynb`.

---

### **`setup_`**

* **Theme:** Environment & Connection Setup
* **Purpose:** Notebooks for initial project configuration, dependency checks, and testing API connections (e.g., MetaTrader 5, data providers). These are typically run once at the beginning of the project or on a new machine.
* **Examples:** `setup_environment_check.ipynb`, `setup_mt5_connection.ipynb`

### **`data_`**

* **Theme:** Data Acquisition & Exploratory Data Analysis (EDA)
* **Purpose:** Scripts for downloading, cleaning, formatting, and performing initial statistical analysis on historical data. This is where we get to know our raw materials.
* **Examples:** `data_download_forex_pairs.ipynb`, `data_exploratory_analysis_dax.ipynb`

### **`feat_`**

* **Theme:** Feature Engineering
* **Purpose:** Focused analysis and visualization of individual technical indicators or custom-derived features. The goal is to understand the behavior of a single feature before incorporating it into a full strategy.
* **Examples:** `feat_rsi_distribution_analysis.ipynb`, `feat_bollinger_bands_behavior.ipynb`

### **`strat_`**

* **Theme:** Strategy Implementation & Backtesting
* **Purpose:** Implementation and backtesting of a complete trading logic (entry, exit, SL/TP) using a *fixed* set of parameters. This step is for validating the core concept of a strategy.
* **Examples:** `strat_sma_cross_backtest.ipynb`, `strat_backtestingpy_sma_comparison.ipynb`

### **`optim_`**

* **Theme:** Strategy Optimization
* **Purpose:** Fine-tuning the parameters of a validated strategy to find the most robust and profitable combinations. This is where we leverage the power of `vectorbt`'s grid search or `backtesting.py`'s optimization engine.
* **Examples:** `optim_sma_grid_search_vectorbt.ipynb`, `optim_rsi_walk_forward.ipynb`

### **`analysis_`**

* **Theme:** Results Analysis & Reporting
* **Purpose:** In-depth analysis, comparison, and visualization of backtesting and optimization results. This is where final conclusions about a strategy's performance are drawn.
* **Examples:** `analysis_compare_sma_vs_ema.ipynb`, `analysis_optimization_heatmap.ipynb`

### **`sandbox_`**

* **Theme:** Sandbox & Experiments
* **Purpose:** A catch-all category for quick experiments, prototyping new ideas, or testing third-party libraries that don't yet fit into the structured workflow. This helps keep the other directories clean.
* **Examples:** `sandbox_quick_plot_test.ipynb`, `sandbox_new_library_check.ipynb`
