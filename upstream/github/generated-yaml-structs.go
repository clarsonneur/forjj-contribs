// This file is autogenerated by "go generate". Do not modify it.
// It has been generated from your 'github.yaml' file.
// To update those structure, update the 'github.yaml' and run 'go generate'
package main

type CreateReq struct {
    ForjjOrganization string `yaml:"forjj-organization"` // Default FORJJ Organization. Used by default as github organization. If you want different one, use --github-organization
    GithubOrganization string `yaml:"github-organization"` // Github Organization name. By default, it uses the FORJJ organization name
    GithubServer string `yaml:"github-server"` // Github Entreprise Server name. By default, public 'github.com' API is used.
    GithubToken string `yaml:"github-token"` // github token to access. This token must authorize organization level access.

    // common flags
    ForjjInfra string `yaml:"forjj-infra"` // Name of the Infra repository to use
    GithubDebug string `yaml:"github-debug"` // To activate github debug information
}

type UpdateReq struct {

    // common flags
    ForjjInfra string `yaml:"forjj-infra"` // Name of the Infra repository to use
    GithubDebug string `yaml:"github-debug"` // To activate github debug information
}

type MaintainReq struct {
    GithubServer string `yaml:"github-server"` // Github Entreprise Server name. By default, public 'github.com' API is used.
    GithubToken string `yaml:"github-token"` // github token to access. This token must authorize organization level access.

    // common flags
    ForjjInfra string `yaml:"forjj-infra"` // Name of the Infra repository to use
    GithubDebug string `yaml:"github-debug"` // To activate github debug information
}

// YamlDesc has been created from your 'github.yaml' file.
const YamlDesc="---\n" +
   "plugin: \"github\"\n" +
   "version: \"0.1\"\n" +
   "description: \"Upstream github plugin for FORJJ. It properly configure github.com or entreprise with organisation/repos\"\n" +
   "runtime:\n" +
   "  docker_image: \"docker.hos.hpecorp.net/forjj-us/github\"\n" +
   "  service_type: \"REST API\"\n" +
   "  service:\n" +
   "    #socket: \"github.sock\"\n" +
   "    parameters: [ \"service\", \"start\" ]\n" +
   "actions:\n" +
   " common:\n" +
   "   flags:\n" +
   "     forjj-infra:\n" +
   "       help: \"Name of the Infra repository to use\"\n" +
   "     github-debug:\n" +
   "       help: \"To activate github debug information\"\n" +
   " create:\n" +
   "   help: \"Create the github environment to manage source and infra code.\"\n" +
   "   flags:\n" +
   "     github-token:\n" +
   "       help: \"github token to access. This token must authorize organization level access.\"\n" +
   "       required: true\n" +
   "     github-server:\n" +
   "       help: \"Github Entreprise Server name. By default, public 'github.com' API is used.\"\n" +
   "     forjj-organization:\n" +
   "       help: \"Default FORJJ Organization. Used by default as github organization. If you want different one, use --github-organization\"\n" +
   "     github-organization:\n" +
   "       help: \"Github Organization name. By default, it uses the FORJJ organization name\"\n" +
   " update:\n" +
   "   help: \"Update the github infrastructure in the infra repository.\"\n" +
   " maintain:\n" +
   "   help: \"Maintain github infrastructure from the infra repository\"\n" +
   "   flags:\n" +
   "     github-token:\n" +
   "       help: \"github token to access. This token must authorize organization level access.\"\n" +
   "       required: true\n" +
   "     github-server:\n" +
   "       help: \"Github Entreprise Server name. By default, public 'github.com' API is used.\"\n" +
   ""

