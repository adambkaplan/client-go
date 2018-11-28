package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_DelegatedAuthentication = map[string]string{
	"":         "DelegatedAuthentication allows authentication to be disabled.",
	"disabled": "disabled indicates that authentication should be disabled.  By default it will use delegated authentication.",
}

func (DelegatedAuthentication) SwaggerDoc() map[string]string {
	return map_DelegatedAuthentication
}

var map_DelegatedAuthorization = map[string]string{
	"":         "DelegatedAuthorization allows authorization to be disabled.",
	"disabled": "disabled indicates that authorization should be disabled.  By default it will use delegated authorization.",
}

func (DelegatedAuthorization) SwaggerDoc() map[string]string {
	return map_DelegatedAuthorization
}

var map_GenerationHistory = map[string]string{
	"":               "GenerationHistory keeps track of the generation for a given resource so that decisions about forced updated can be made.",
	"group":          "group is the group of the thing you're tracking",
	"resource":       "resource is the resource type of the thing you're tracking",
	"namespace":      "namespace is where the thing you're tracking is",
	"name":           "name is the name of the thing you're tracking",
	"lastGeneration": "lastGeneration is the last generation of the workload controller involved",
}

func (GenerationHistory) SwaggerDoc() map[string]string {
	return map_GenerationHistory
}

var map_GenericOperatorConfig = map[string]string{
	"":               "GenericOperatorConfig provides information to configure an operator",
	"servingInfo":    "ServingInfo is the HTTP serving information for the controller's endpoints",
	"leaderElection": "leaderElection provides information to elect a leader. Only override this if you have a specific need",
	"authentication": "authentication allows configuration of authentication for the endpoints",
	"authorization":  "authorization allows configuration of authentication for the endpoints",
}

func (GenericOperatorConfig) SwaggerDoc() map[string]string {
	return map_GenericOperatorConfig
}

var map_LoggingConfig = map[string]string{
	"":        "LoggingConfig holds information about configuring logging",
	"level":   "level is passed to glog.",
	"vmodule": "vmodule is passed to glog.",
}

func (LoggingConfig) SwaggerDoc() map[string]string {
	return map_LoggingConfig
}

var map_NodeStatus = map[string]string{
	"":                               "NodeStatus provides information about the current state of a particular node managed by this operator.",
	"nodeName":                       "nodeName is the name of the node",
	"currentDeploymentGeneration":    "currentDeploymentGeneration is the generation of the most recently successful deployment",
	"targetDeploymentGeneration":     "targetDeploymentGeneration is the generation of the deployment we're trying to apply",
	"lastFailedDeploymentGeneration": "lastFailedDeploymentGeneration is the generation of the deployment we tried and failed to deploy.",
	"lastFailedDeploymentErrors":     "lastFailedDeploymentGenerationErrors is a list of the errors during the failed deployment referenced in lastFailedDeploymentGeneration",
}

func (NodeStatus) SwaggerDoc() map[string]string {
	return map_NodeStatus
}

var map_OperatorCondition = map[string]string{
	"": "OperatorCondition is just the standard condition fields.",
}

func (OperatorCondition) SwaggerDoc() map[string]string {
	return map_OperatorCondition
}

var map_OperatorSpec = map[string]string{
	"":                "OperatorSpec contains common fields for an operator to need.  It is intended to be anonymous included inside of the Spec struct for you particular operator.",
	"managementState": "managementState indicates whether and how the operator should manage the component",
	"imagePullSpec":   "imagePullSpec is the image to use for the component.",
	"imagePullPolicy": "imagePullPolicy specifies the image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.",
	"version":         "version is the desired state in major.minor.micro-patch.  Usually patch is ignored.",
	"logging":         "logging contains glog parameters for the component pods.  It's always a command line arg for the moment",
}

func (OperatorSpec) SwaggerDoc() map[string]string {
	return map_OperatorSpec
}

var map_OperatorStatus = map[string]string{
	"":                           "OperatorStatus contains common fields for an operator to need.  It is intended to be anonymous included inside of the Status struct for you particular operator.",
	"observedGeneration":         "observedGeneration is the last generation change you've dealt with",
	"conditions":                 "conditions is a list of conditions and their status",
	"state":                      "state indicates what the operator has observed to be its current operational status.",
	"taskSummary":                "taskSummary is a high level summary of what the controller is currently attempting to do.  It is high-level, human-readable and not guaranteed in any way. (I needed this for debugging and realized it made a great summary).",
	"currentVersionAvailability": "currentVersionAvailability is availability information for the current version.  If it is unmanged or removed, this doesn't exist.",
	"targetVersionAvailability":  "targetVersionAvailability is availability information for the target version if we are migrating",
}

func (OperatorStatus) SwaggerDoc() map[string]string {
	return map_OperatorStatus
}

var map_StaticPodOperatorStatus = map[string]string{
	"":                                    "StaticPodOperatorStatus is status for controllers that manage static pods.  There are different needs because individual node status must be tracked.",
	"latestAvailableDeploymentGeneration": "latestAvailableDeploymentGeneration is the deploymentID of the most recent deployment",
	"nodeStatuses":                        "nodeStatuses track the deployment values and errors across individual nodes",
}

func (StaticPodOperatorStatus) SwaggerDoc() map[string]string {
	return map_StaticPodOperatorStatus
}

var map_VersionAvailability = map[string]string{
	"":                "VersionAvailability gives information about the synchronization and operational status of a particular version of the component",
	"version":         "version is the level this availability applies to",
	"updatedReplicas": "updatedReplicas indicates how many replicas are at the desired state",
	"readyReplicas":   "readyReplicas indicates how many replicas are ready and at the desired state",
	"errors":          "errors indicates what failures are associated with the operator trying to manage this version",
	"generations":     "generations allows an operator to track what the generation of \"important\" resources was the last time we updated them",
}

func (VersionAvailability) SwaggerDoc() map[string]string {
	return map_VersionAvailability
}

// AUTO-GENERATED FUNCTIONS END HERE
