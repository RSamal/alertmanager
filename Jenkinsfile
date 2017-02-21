#!/usr/bin/env groovy
node {
	
	def goroot = tool "Go"
	env.PATH =  "${goroot}/bin:${env.PATH}"
	gopath = pwd()
	env.GOPATH = gopath
	workspace = gopath + "/src/github.com/prometheus/alertmanager"

	dir(workspace) {
		stage 'Checkout'	
			   echo "checking out the code.."			
	    	   checkout scm
		stage 'Test'
			   sh 'make test'
		stage 'Build'
				echo "Build  binary..."
				sh 'promu crossbuild'

				echo "Build RPM package..."
				sh 'go run build.go pkg-rpm'
	}
}