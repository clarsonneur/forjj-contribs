// This file is autogenerated by "go generate". Do not modify it.
// It has been generated from your 'github.yaml' file.
// To update those structure, update the 'github.yaml' and run 'go generate'
package main

// Object app groups structure

// Groups structure


// Object Instance structures

type AppInstanceStruct struct {
	ForjjInfra string `json:"forjj-infra"` // Name of the Infra repository to use in github if requested.
	ForjjOrganization string `json:"forjj-organization"` // Default FORJJ Organization. Used by default as github organization. If you want different one, use --github-organization
	Organization string `json:"organization"` // Github Organization name. By default, it uses the FORJJ organization name
	Server string `json:"server"` // Github Entreprise Server name. By default, public 'github.com' API is used.
	Token string `json:"token"` // github token to access. This token must authorize organization level access.
}

// Object group groups structure

// Groups structure


// Object Instance structures

type GroupInstanceStruct struct {
	Members []string `json:"members"` // List of users to attach to the new group.
	Name string `json:"name"` // group name
	Role string `json:"role"` // List of roles to apply to the new group.
}

// Object repo groups structure

// Groups structure


// Object Instance structures

type RepoInstanceStruct struct {
	Flow string `json:"flow"` // Flow activated on this repository
	ForjjWorkspaceMount string `json:"forjj-workspace-mount"` // Where the workspace dir is located in the github plugin container.
	Groups string `json:"groups"` // List of groups to attach to the repository, separated by comma.
	Issue_tracker string `json:"issue_tracker"` // To activate the Issue tracker to the Repository
	Name string `json:"name"` // Repository name
	Title string `json:"title"` // Github Repository title
	Users string `json:"users"` // List of users to attach to the repository, separated by comma.
}

// Object user groups structure

// Groups structure


// Object Instance structures

type UserInstanceStruct struct {
	Name string `json:"name"` // 
	Role string `json:"role"` // 
}


// ************************
// Create request structure
// ************************

type CreateReq struct {
	Forj struct {
		Debug string `json:"debug"`
		ForjjInstanceName string `json:"forjj-instance-name"`
		ForjjSourceMount string `json:"forjj-source-mount"`
	}
	Objects CreateArgReq
}

type CreateArgReq struct {
	App map[string]AppInstanceStruct `json:"app"` // Object details
	Group map[string]GroupInstanceStruct `json:"group"` // Object details
	Repo map[string]RepoInstanceStruct `json:"repo"` // Object details
	User map[string]UserInstanceStruct `json:"user"` // Object details
}

// ************************
// Update request structure
// ************************

type UpdateReq struct {
	Forj struct {
		Debug string `json:"debug"`
		ForjjInstanceName string `json:"forjj-instance-name"`
		ForjjSourceMount string `json:"forjj-source-mount"`
	}
	Objects UpdateArgReq
}

type UpdateArgReq struct {
	App map[string]AppInstanceStruct `json:"app"` // Object details
	Group map[string]GroupInstanceStruct `json:"group"` // Object details
	Repo map[string]RepoInstanceStruct `json:"repo"` // Object details
	User map[string]UserInstanceStruct `json:"user"` // Object details
}

// **************************
// Maintain request structure
// **************************

type MaintainReq struct {
	Forj struct {
		Debug string `json:"debug"`
		ForjjInstanceName string `json:"forjj-instance-name"`
		ForjjSourceMount string `json:"forjj-source-mount"`
		ForjjWorkspaceMount string `json:"forjj-workspace-mount"`
	}
	Objects MaintainArgReq
}

type MaintainArgReq struct {
	App map[string]AppMaintainStruct `json:"app"` // Object details
}

type AppMaintainStruct struct {
	Token string `json:"token"` // github token to access. This token must authorize organization level access.
}


// YamlDesc has been created from your 'github.yaml' file.
const YamlDesc = "---\n" +
   "plugin: \"github\"\n" +
   "version: \"0.1\"\n" +
   "description: \"Upstream github plugin for FORJJ. It properly configure github.com or entreprise with organisation/repos\"\n" +
   "runtime:\n" +
   "  docker:\n" +
   "    image: \"forjdevops/forjj-github\"\n" +
   "  service_type: \"REST API\"\n" +
   "  service:\n" +
   "    parameters: [ \"service\", \"start\" ]\n" +
   "created_flag_file: \"{{ .InstanceName }}/{{.Name}}.yaml\"\n" +
   "task_flags: # All task flags will be delivered by forjj to the plugin under forj/\n" +
   "  common:\n" +
   "    debug:\n" +
   "      help: \"To activate github debug information\"\n" +
   "    forjj-source-mount:\n" +
   "      help: \"Where the source dir is located for github plugin container.\"\n" +
   "    forjj-instance-name:\n" +
   "       help: \"Name of the jenkins instance given by forjj.\"\n" +
   "  maintain:\n" +
   "    forjj-workspace-mount:\n" +
   "      help: \"Where the workspace dir is located in the github plugin container.\"\n" +
   "objects: # All objects will be delivered by forjj except workspace/infra under objects/<type>/<instance>/<action>/key=value\n" +
   "  # Define infra object special flag for github\n" +
   "  app: # already defined by Forjj\n" +
   "    # Default is : actions: [\"add\", \"change\", \"remove\"] No need to define it.\n" +
   "    flags:\n" +
   "      server:\n" +
   "        help: \"Github Entreprise Server name. By default, public 'github.com' API is used.\"\n" +
   "      forjj-organization:\n" +
   "        only-for-actions: [\"add\"]\n" +
   "        help: \"Default FORJJ Organization. Used by default as github organization. If you want different one, use --github-organization\"\n" +
   "      organization:\n" +
   "        only-for-actions: [\"add\"]\n" +
   "        help: \"Github Organization name. By default, it uses the FORJJ organization name\"\n" +
   "      forjj-infra:\n" +
   "        only-for-actions: [\"add\", \"change\"]\n" +
   "        help: \"Name of the Infra repository to use in github if requested.\"\n" +
   "      token:\n" +
   "        only-for-actions: [\"add\", \"change\"]\n" +
   "        help: \"github token to access. This token must authorize organization level access.\"\n" +
   "        required: true\n" +
   "        secure: true\n" +
   "        envar: \"TOKEN\"\n" +
   "  # Define github group exposure to forjj\n" +
   "  group: # New object type in forjj\n" +
   "    # Default is : actions: [\"add\", \"change\", \"remove\", \"list\", \"rename\"]\n" +
   "    help: \"Manage teams in github\"\n" +
   "    identified_by_flag: name\n" +
   "    flags:\n" +
   "      members:\n" +
   "        only-for-actions: [\"add\"]\n" +
   "        help: \"List of users to attach to the new group.\"\n" +
   "        of-type: \"[]string\"\n" +
   "      name:\n" +
   "        help: \"group name\"\n" +
   "        required: true\n" +
   "      role:\n" +
   "        only-for-actions: [\"add\"]\n" +
   "        help: \"List of roles to apply to the new group.\"\n" +
   "  # Define github users exposure to forjj\n" +
   "  user: # New object type in forjj\n" +
   "    # Default is : actions: [\"add\", \"change\", \"remove\", \"list\", \"rename\"]\n" +
   "    help: \"Manage organization list of users\"\n" +
   "    identified_by_flag: name\n" +
   "    flags:\n" +
   "      name:\n" +
   "        options:\n" +
   "          help: \"user name\"\n" +
   "          required: true\n" +
   "      role:\n" +
   "        only-for-actions: [\"add\"]\n" +
   "        options:\n" +
   "          help: \"List of roles to apply to the new user.\"\n" +
   "  repo: # Enhance Forjj repo object\n" +
   "    actions: [\"add\", \"change\"]\n" +
   "    flags:\n" +
   "      name:\n" +
   "        help: \"Repository name\"\n" +
   "      title:\n" +
   "        help: \"Github Repository title\"\n" +
   "      issue_tracker:\n" +
   "        help: \"To activate the Issue tracker to the Repository\"\n" +
   "        default: \"true\"\n" +
   "      users:\n" +
   "        only-for-actions: [\"add\"]\n" +
   "        help: \"List of users to attach to the repository, separated by comma.\"\n" +
   "        format-regexp: \"[+-]?[a-zA-Z0-9]([a-zA-Z0-9-]*[a-zA-Z0-9])?([a-zA-Z0-9]([a-zA-Z0-9-]*[a-zA-Z0-9]))*,\"\n" +
   "      groups:\n" +
   "        only-for-actions: [\"add\"]\n" +
   "        help: \"List of groups to attach to the repository, separated by comma.\"\n" +
   "        format-regexp: \"[+-]?[a-zA-Z0-9]([a-zA-Z0-9-]*[a-zA-Z0-9])?([a-zA-Z0-9]([a-zA-Z0-9-]*[a-zA-Z0-9]))*,\"\n" +
   "      flow:\n" +
   "        help: \"Flow activated on this repository\"\n" +
   "      forjj-workspace-mount:\n" +
   "        help: \"Where the workspace dir is located in the github plugin container.\"\n" +
   ""
