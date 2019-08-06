package main

import (
"bytes"
"fmt"
"log"
"os/exec"
"github.com/tidwall/gjson"
"strings"
"strconv"
)

const api_key = "" // your Zenhub API key
const git_token = "" // your Github token
const repo = "" // input as `user/repo_name`
var names []string // slice of pipeline names
var num_issues []int // slice of number of issues in each pipeline
var issues []string // slice of all issues
var positions []string // slice of all positions
var epic_ids []string // slice of epic IDs

func main() {
    // Get repo id from Github API
    cmd := exec.Command("curl", fmt.Sprintf("https://github.com/api/v3/repos/%s?access_token=%s", repo, git_token))
    var out1 bytes.Buffer
    cmd.Stdout = &out1
    err1 := cmd.Run()

    if err1 != nil {
        log.Fatal(err1)
    }

    // Extracts repo id from JSON output
    repo_id := gjson.Get(out1.String(), "id")
    fmt.Printf("Repo id: %s\n", repo_id)

    // Get epics for a repo
    pull_epics := exec.Command("curl", fmt.Sprintf("https://zenhub.com/p1/repositories/%s/board?access_token=%s", repo_id.String(), api_key))
    var out2 bytes.Buffer
    pull_epics.Stdout = &out2
    err2 := pull_epics.Run()

    if err2 != nil {
        log.Fatal(err2)
    }

    // Get number of pipelines
    count_pipelines := gjson.Get(out2.String(), `pipelines.#`).String()
    fmt.Printf("Number of pipelines: %s\n\n", count_pipelines)

    // Create map of [pipeline] = number_of_issues
    p, errr := strconv.Atoi(count_pipelines)
    if errr != nil {
      log.Fatal(errr)
    }

    // Loop that grabs names of each pipeline, number of issues in each one, the epic id, an array of the issues, and their positions
    for i := 0; i < p; i++ {
      names := append(names, gjson.Get(out2.String(), strings.Join([]string{"pipelines.",strconv.Itoa(i),".name"}, "")).String())
      num_issues := append(num_issues, int(gjson.Get(out2.String(), strings.Join([]string{"pipelines.",strconv.Itoa(i),".issues.#"}, "")).Int()))
      epic_ids := append(issues, gjson.Get(out2.String(), strings.Join([]string{"pipelines.",strconv.Itoa(i),".id"}, "")).String())
      issues := append(issues, gjson.Get(out2.String(), strings.Join([]string{"pipelines.",strconv.Itoa(i),".issues.#.issue_number"}, "")).String())
      positions := append(positions, gjson.Get(out2.String(), strings.Join([]string{"pipelines.",strconv.Itoa(i),".issues.#.position"}, "")).String())

      // Print issues and positions
      defer fmt.Println(issues, positions, epic_ids, "\n")

      // Print pipeline name and number of issues
      defer fmt.Println(names, num_issues)
    }
}
