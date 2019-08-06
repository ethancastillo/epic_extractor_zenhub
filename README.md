# Epics Extractor for Zenhub

This Go code is designed to work with both the public and enterprise API, just replace the endpoint url in the `cmd` and `pull_epics` variables.

# What Exactly Does It Do?
It pulls all of the issues tied to a Github repo in Zenhub, and parses the data into a readable format by pipeline. You don't even need your Repo ID from Github, just the name it. It outputs the Repo ID, the number of pipelines in the Zenhub board, the issues in each pipeline with their position, and the IDs for each epic.
