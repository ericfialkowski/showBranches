# env
Environment Variable Parsing in Go

This is a simple library to parse environment variables into various types. There is also the option for
providing a default value. This helps in writing 12 factor applications, specifically storing the config
in environment variables. See <https://12factor.net/config>

The one thing I don't like about this library is I worry the go could fall into the same problems that 
node.js had with leftpad. So make sure you vendor, fork, or even just copy pasta this code in case it
goes away. I have no plans on it going away but protect your projects! 😀
