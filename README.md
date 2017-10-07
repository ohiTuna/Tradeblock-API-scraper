# Tradeblock-API-scraper
## This is still very much beta!  

  
simple tool for querying Tradeblock's Bitcoin tx history (and mempool history, see below) into .csv   
sum block/tx size, sum txs, sum fees, and total blocks are available to be outputted at 1h, 2h, 6h, and 24h intervals with three column format:  
timestamp(unix), value, excel converted date ("=(A1/24/60/60)+DATE(1970,1,1)") 

Requires Go prog language (golang): https://golang.org/dl/   


## version history notes/update notes       
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

version 0.20alpha
 - available via tbkcustNEW.go. 
 - adding ability to pull more than 1000 records at once (currently not working, DO NOT enter value >1 for 'instances back')
 - fixing the date column so that it correctly outputs Excel formatted datestamp
 - code cleanup so that not everything is in func main()


## Usage:  
$ go run $GOPATH/Tradeblock-API-scraper/tbkcustom.go  
  
alternatively you can use:
$ go run $GOPATH/Tradeblock-API-scraper/tbkcustNEW.go
to run the most current version (which may or may not be stable)

## available commands (self-explanatory usage):  
txsize_new, txes_new, txfee_new, blocks_new  
1h, 2h, 6h, 1d  
unix timestamp val (i.e. enter 1471327200 for 8/16/16 at 0600)
positive integer (1,2,3,4...) (only in tbkcustNEW.go, specifies how many additional thousand of interval records to retrieve)

example usage (entered when prompted):
txfee_new
6h
1507278520
2

will retrieve the sum transaction fees (in satoshis) for the 2000 six hour intervals before October 6 2017 04:30 (500 days of tx fee vals ending on Oct. 6 2017)


# Tradeblock mempool scraper
Gets the last week of mempool data, giving snapshots at 10 min intervals(code can be modified to do 1 min intervals if desired). Data is spit out into 8 columns and (typically) just under 1,000 rows. The columns aren't currently labeled, but they are: date/time (unix), tx bytes added to MP since last block found, tx bytes in MP since BEFORE last block, txs added to MP since last block, txs in MP since BEFORE last block, fees added to MP since last block found (in satoshis), fees added to MP since BEFORE last block, tx per min.     
So if you divide the sixth column (F in excel), by the second column (B in excel) you can get the fee per byte for all txs in MP that have been sent AFTER the last block was found. If you add columns F+G and B+C and do (F+G)/(B+C) the result is the fee per byte for all txs in the mempool at that moment. 
Basicaly this scraper is the data that is visualized here: https://tradeblock.com/bitcoin/. It can easily be modified to run API queries at weekly intervals if you want to just let it run in the background. 
Two points of caution: one, there are sometimes gaps in the data and the time will jump from 8/30 12:10 to 8/30 12:40 (or similar), it is rare but it does happen. Two, remember that Tradeblock is like any other node and will 'count' thier mempool differently. Some will hold on to txs for 72 hours with no min fee, others will only count txs as in the mempool fo 48 hours and only if the fee is >0. 

## Usage
$ go run $GOPATH/SRC/github.com/test1/tbkmempool.go (your location may vary of course)
Data is then automatically parsed out and, unlike the above tool, will actually display the data in your terminal rather than the "looks good" message.
