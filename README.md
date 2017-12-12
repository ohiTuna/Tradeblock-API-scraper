# Tradeblock-API-scraper
## This is still very much beta!  

  
simple tool for querying Tradeblock's Bitcoin tx history (and mempool history, see below) into .csv   
sum block/tx size, sum txs, sum fees, and total blocks are available to be outputted at 1h, 2h, 6h, and 24h intervals with three column format:  
timestamp(unix), value, excel converted date ("=(A1/24/60/60)+DATE(1970,1,1)") 

Requires Go prog language (golang): https://golang.org/dl/   
(simple windows executable may be available in future)


## Usage:  
$ go run $GOPATH/Tradeblock-API-scraper/tbkcustom.go  

follow prompts on console          
requires Go (golang) distribution installed and is run through terminal on Linux or whatever you like on Win systems: MinGW, Win command prompt, git-bash, etc. I use git-bash.



## version history notes/update notes       
version 0.32
  - cleanup
  - started work on win binary 


version 0.31 
  - forking mempool scraper off into its own new repo
  
version 0.30 
 - both tbkcustNEW.go and tbkcustom.go have been updated with latest stable version
 - excel counter fixed
 - debugging messages pruned
 - scanner and console instructions improved
 
version 0.21a
 - available via tbkcustNEW.go
 - ability to pull multiple record sets now working. Program will ask user for 'instances back' to run, each instance is 1000 values i.e. 4 instances back = 4000 records retrieved.
 - minor bugs remain: excel counter needs fixed, debugging status output msgs need pruned, record set output order needs fixed
 - rompt for user date now has a unix time epoch with current time
 
version 0.20a
 - available via tbkcustNEW.go. 
 - adding ability to pull more than 1000 records at once (currently not working, DO NOT enter value >1 for 'instances back')
 - fixing the date column so that it correctly outputs Excel formatted datestamp
 - code cleanup so that not everything is in func main()


