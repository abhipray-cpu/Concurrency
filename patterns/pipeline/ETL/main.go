package main

import (
	"ETL/data"
)

/*
ETL stands for Extract, Transform, Load. It's a process used in database usage and especially in data
 warehousing. The ETL process involves three steps:

1)Extract: The first step is to extract data from various sources. These sources could be databases,
CRM systems, flat files, web services, etc. The goal is to collect the necessary data for analysis
and storage. During extraction, the focus is on accessing the data and ensuring it's accurately
and efficiently retrieved, without much concern for the data's actual content or structure.

2)Transform: Once the data is extracted, it undergoes transformation.
This step is crucial for preparing the data for its intended use.
Transformation can involve a variety of processes, including:

a)Cleaning: Removing inaccuracies, duplicates, or irrelevant data.
b)Standardization: Converting data to a common format to ensure consistency across all data.
c)Enrichment: Enhancing data by adding additional information from other sources.
d)Filtering: Selecting only the parts of the data that are needed for the analysis.
e)Aggregation: Summarizing detailed data for a higher-level view.
f)Splitting: Dividing data into more granular parts for detailed analysis.
g)Merging: Combining data from different sources to provide a unified view.
The transformation step is tailored to the specific requirements of the target system and
the business objectives.

Load: The final step is to load the transformed data into a target system, which could be a database,
a data warehouse, a data mart, or any other storage system.
The loading process can be done in batches (batch loading) at specific intervals
(e.g., nightly, weekly) or in real-time (streaming load) where data is loaded as soon as it is
extracted and transformed.
*/

func init() {
	data.GenerateMockJsonData()
	data.GenerateCSVData()
}
func main() {
	var productChannel chan Product
	go ProductGenerator(productChannel)
	pipeline := Extract(productChannel)
	pipeline = Transform(pipeline)
	loadToCSV(pipeline)
}
