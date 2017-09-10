// This file has been created by "go generate" as initial code. go generate will never update it, EXCEPT if you remove it.
package main

import (
	"github.com/forj-oss/goforjj"
	"log"
	"net/http"
)

// Do creating plugin task
// req_data contains the request data posted by forjj. Structure generated from 'jenkins.yaml'.
// ret_data contains the response structure to return back to forjj.
//
func DoCreate(w http.ResponseWriter, r *http.Request, req *CreateReq, ret *goforjj.PluginData) (httpCode int) {
	var p *JenkinsPlugin

	if pr, code := req.check_source_existence(ret); pr == nil {
		return code
	} else {
		p = pr
	}

	if !p.initialize_from(req, ret) {
		return
	}

	if !p.create_jenkins_sources(req.Forj.ForjjInstanceName, ret) {
		return
	}

	if !p.save_yaml(ret) {
		return
	}
	return
}

// DoUpdate is the update plugin task
// req_data contains the request data posted by forjj. Structure generated from 'jenkins.yaml'.
// ret_data contains the response structure to return back to forjj.
// forjj-jenkins.yaml is loaded by default.
//
func DoUpdate(w http.ResponseWriter, r *http.Request, req *UpdateReq, ret *goforjj.PluginData) (httpCode int) {
	var p *JenkinsPlugin

	if pr, ok := req.check_source_existence(ret); !ok {
		return
	} else {
		p = pr
	}

	if !p.load_yaml(ret) {
		return
	}

	// TODO: Use the GithubStruct.UpdateFrom(...)
	instance := req.Forj.ForjjInstanceName
	p.yaml.Forjj.InstanceName = instance
	p.yaml.Forjj.OrganizationName = req.Forj.ForjjOrganization
	p.yaml.Forjj.InfraUpstream = req.Forj.ForjjInfraUpstream

	status := p.update_from(req, ret)
	status = p.update_projects(req, ret) || status
	status = p.update_jenkins_sources(req.Forj.ForjjInstanceName, ret) || status
	status = p.save_yaml(ret) || status

	if !status {
		log.Print(ret.StatusAdd("No update detected. Jenkins source files hasn't been updated."))
		return
	}
	return
}

// Do maintaining plugin task
// req_data contains the request data posted by forjj. Structure generated from 'jenkins.yaml'.
// ret_data contains the response structure to return back to forjj.
//
func DoMaintain(w http.ResponseWriter, r *http.Request, req *MaintainReq, ret *goforjj.PluginData) (httpCode int) {
	if !req.check_source_existence(ret) {
		return
	}

	// loop on list of jenkins instances defined by a collection of */jenkins.yaml
	if !req.Instantiate(req, ret) {
		return
	}
	return
}
