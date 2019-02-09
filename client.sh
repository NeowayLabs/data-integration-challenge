awk 'BEGIN {FS=";"}{print "{\"name\": \""$1"\", \"zip\": \""$2"\", \"website\": \""$3"\"}"}' q2_clientData.csv |
    xargs -i -d'\n' curl -d '{}' -X PUT localhost:3000/companies/update
