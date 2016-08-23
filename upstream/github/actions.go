// This file has been created by "go generate" as initial code. go generate will never update it, EXCEPT if you remove it.

// So, update it for your need.
package main

// You can remove following comments.
// It has been designed fo you, to implement the core of your plugin task.
//
// You can use use it to write your own plugin handler for additional functionnality
// Like Index which currently return a basic code.

import (
    "net/http"
    "github.hpe.com/christophe-larsonneur/goforjj"
    "log"
    "fmt"
    "path"
    "os"
)

const github_file = "github.yaml"

// Do creating plugin task
// req_data contains the request data posted by forjj. Structure generated from 'github.yaml'.
// ret_data contains the response structure to return back to forjj.
//
func DoCreate(w http.ResponseWriter, r *http.Request, req *CreateReq, ret *goforjj.PluginData) {

    gws := GitHubStruct{
        source_mount: req.ForjjSourceMount,
        token: req.GithubToken,
    }
    check := make(map[string]bool)
    check["token"] = true
    log.Printf("Checking parameters : %#v", gws)

    //ensure source path is writeable
    if gws.verify_req_fails(ret, check) {
        return
    }

    // A create won't be possible if source files already exist. The Update is the only possible option.
    log.Printf("Checking Infrastructure code existence.")
    if _, err := os.Stat(path.Join(req.ForjjSourceMount, github_file)) ; err == nil {
        ret.Errorf("Unable to create the github configuration which already exist.\nUse update to update it (or update %s), and maintain to update github according to his configuration.", github_file)
        return
    }

    log.Printf("Checking github connection : %#v", gws)

    if gws.github_connect(req.GithubServer, ret) == nil {
        return
    }

    // Build gws.github_source yaml structure.
    if err := gws.create_yaml_data(req) ; err != nil {
        ret.Errorf("%s", err)
        return
    }

    // A create won't be possible if repo requested already exist. The Update is the only possible option.
    if err := gws.repos_exists() ; err != nil {
        ret.Errorf("%s\nUnable to create the github configuration when github already has repositories created. Use 'update' instead.", err)
        return
    }

    ret.StatusAdd("Environment checked. Ready to be created.")

    // Save gws.github_source.
    if err := gws.save_yaml(path.Join(req.ForjjSourceMount, github_file)) ; err != nil {
        ret.Errorf("%s", err)
        return
    }
    log.Printf(ret.StatusAdd("Configuration saved in '%s'.", github_file))

    // Building final Post answer
    // We assume ssh is used and forjj can push with appropriate credential.
    infra_repo := gws.github_source.Repos[req.ForjjInfra]
    ret.Repos[req.ForjjInfra] = goforjj.PluginRepo{ infra_repo.Name, infra_repo.Upstream, infra_repo.Exist }
    for k, v := range gws.github_source.Urls {
        ret.Services.Urls[k] = v
    }

    ret.CommitMessage = fmt.Sprintf("Create github configuration")
    ret.Files = append(ret.Files, github_file)

}

// Do updating plugin task
// req_data contains the request data posted by forjj. Structure generated from 'github.yaml'.
// ret_data contains the response structure to return back to forjj.
//
func DoUpdate(w http.ResponseWriter, r *http.Request, req *UpdateReq, ret *goforjj.PluginData) {

    gws := GitHubStruct{
        source_mount: req.ForjjSourceMount,
    }
    check := make(map[string]bool)

    if gws.verify_req_fails(ret, check) {
        return
    }

    // TODO: Add code to update github service source files
}

// Do maintaining plugin task
// req_data contains the request data posted by forjj. Structure generated from 'github.yaml'.
// ret_data contains the response structure to return back to forjj.
//
func DoMaintain(w http.ResponseWriter, r *http.Request, req *MaintainReq, ret *goforjj.PluginData) {

    gws := GitHubStruct{
        source_mount: req.ForjjSourceMount,
        workspace_mount: req.ForjjWorkspaceMount,
        token: req.GithubToken,
    }
    check := make(map[string]bool)
    check["token"] = true
    check["workspace"] = true

    if gws.verify_req_fails(ret, check) { // true => include workspace testing.
        return
    }

    // Read the github.yaml file.
    gws.load_yaml(path.Join(req.ForjjSourceMount, github_file))

    if gws.github_connect(gws.github_source.Urls["github-base-url"], ret) == nil {
        return
    }

    // ensure organization exist
    if ! gws.ensure_organization_exists(ret) {
        return
    }
    log.Printf(ret.StatusAdd("Organization maintained."))

    // loop on list of repos, and ensure they exist with minimal config and rights
    for name, repo_data := range  gws.github_source.Repos {
        if err := repo_data.ensure_exists(&gws, ret) ; err != nil {
           return
        }
        log.Printf(ret.StatusAdd("Repo maintained: %s", name))
    }
}
