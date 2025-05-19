import jenkins.model.Jenkins;

pm = Jenkins.instance.pluginManager
uc = Jenkins.instance.updateCenter

pm.doCheckUpdatesServer()

["github", "junit", "docker-build-publish", "workflow-aggregator"].each {
    if (! pm.getPlugin(it)) {
        deployment = uc.getPlugin(it).deploy(true)
        deployment.get()
    }
}

Jenkins.instance.restart()