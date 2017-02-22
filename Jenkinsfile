#!/usr/bin/env groovy
node {
	
	def goroot = tool "Go"
	env.PATH =  "${goroot}/bin:${env.PATH}"
	gopath = pwd()
	env.GOPATH = gopath
	env.PATH =  "${gopath}/bin:${env.PATH}"
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

				env.PATH =  "/usr/local/bin:${env.PATH}"
				echo "Build RPM package..."
				sh 'go run build.go pkg-rpm'
		
		stage 'RPM Dist Copy'
			  VERSION = readFile 'VERSION'
			  step([$class: 'UCDeployPublisher',
        		siteName: 'IBM UCD',
        		component: [
            		$class: 'com.urbancode.jenkins.plugins.ucdeploy.VersionHelper$VersionBlock',
            		componentName: 'cdsmon:alertmanager',
            	delivery: [
					$class: 'com.urbancode.jenkins.plugins.ucdeploy.DeliveryHelper$Push',
					pushVersion: '${VERSION}.${BUILD_NUMBER}',
					baseDir: '\\var\\lib\\jenkins\\workspace\\go\\src\\github.com\\prometheus\\alertmanager\\dist',
					
					fileIncludePatterns: 'alertmanager-*.rpm',
					fileExcludePatterns: '',
					pushProperties: 'jenkins.server=Local\njenkins.reviewed=false',
					pushDescription: 'Pushed from Jenkins',
					pushIncremental: false
        	    ]
        	]
   		 ])
	}
}