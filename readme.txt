A simple chatops client poc written in go for slack and Otter.

How to set it up:

1. have a otter server running
2. in otter, generate an apikey
3. in otter, create a job template
4. build and run this app with the flags -otterServer "AAA" -otterApiKey "BBBB". Where AAA is your otterhost(http://localhost:82), BBB your generated api-key
5. in slack, add a custom slash command
	3.1 command /otter
	3.2 url = the url to this go application(must be https - use aws apigateway if you do not have a ssl cert)
	3.3 Method = POST
	3.4 Token = doesn't matter right - the application doesn't check for it yet
6. in slack, write
	/otter deploy [templatename] prod
	the environment name doesn't matter yet