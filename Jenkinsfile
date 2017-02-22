#!/usr/bin/env groovy
node {
	
	// Make Sure GO tool installed on Jenkins Server and the fpm is avaliable in the server
	def goroot = tool "Go"
	env.PATH =  "${goroot}/bin:${env.PATH}"
	gopath = pwd()
	env.GOPATH = gopath
	FPMPATH = "/usr/local/rvm/gems/ruby-1.9.3-p551/bin"
	RUBY =    "/usr/local/rvm/rubies/ruby-1.9.3-p551/bin"
	RVM  = "/usr/local/rvm/bin"
	// if fpm tool is not avaliable then do "gem install fpm", and provide the binary path below
	env.PATH =  "${gopath}/bin:${RUBY}:${RVM}:${env.PATH}"
	env.PATH = "${FPMPATH}:${env.PATH}"
	workspace = gopath + "/src/github.com/prometheus/alertmanager"

	dir(workspace) {
		stage 'Checkout'	
			   echo "checking out the code.."			
	    	   checkout scm

		stage 'Build'
				sh 'gem env'
				sh 'rvm info'
				echo "Build Linux binary..."
				sh 'promu crossbuild'

				echo "Build RPM package..."
				sh "go run build.go pkg-rpm"
		
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